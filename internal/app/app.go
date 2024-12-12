package app

import (
	"context"
	"encoding/gob"
	"net/http"
	"os"

	"generic.com/internal/labels"
	"generic.com/internal/models"
	"generic.com/internal/repo"
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo *repo.UserRepository
}

func NewRepositories(db *gorm.DB) *Repositories {

	gob.Register(models.User{})

	db.AutoMigrate(&models.User{})

	return &Repositories{
		UserRepo: repo.NewUserRepository(db),
	}
}

type Application struct {
	Ctx            context.Context
	SessionManager *scs.SessionManager
	DB             *gorm.DB
	LabelMapper    *labels.LabelMapper
	Repos          Repositories
}

func NewApplication(
	sessionManager *scs.SessionManager,
	appDB *gorm.DB,
) *Application {
	labelMapper, err := labels.NewLabelMapper("config/ui", &models.Language{Value: "en"})
	if err != nil {
		ELOG.Fatalf("Error loading labels: %v", err)
	}

	a := &Application{
		SessionManager: sessionManager,
		DB:             appDB,
		LabelMapper:    labelMapper,
		Repos:          *NewRepositories(appDB),
	}

	return a
}

func (app *Application) Start() {

	if err := app.CreateAdminUser(); err != nil {
		ELOG.Fatalf("Error creating admin user: %v", err)
	}

	srv := &http.Server{
		Addr:     os.Getenv("SERVER_ADDR"),
		ErrorLog: ELOG,
		Handler:  app.Routes(),
	}

	LOG.Printf("Starting server on %s\n", os.Getenv("SERVER_ADDR"))
	err := srv.ListenAndServe()
	ELOG.Fatal(err)
}

func (app *Application) CreateAdminUser() error {

	if _, err := app.Repos.UserRepo.FindByUsername("Admin"); err != nil {
		if err == gorm.ErrRecordNotFound {
			admin := models.NewSuperAdminUser()
			return app.Repos.UserRepo.CreateUser(admin)
		} else {
			return err
		}

	} else {
		return err
	}
}

func (app *Application) AddDefaultData(r *http.Request, data map[string]interface{}) map[string]interface{} {

	user := app.SessionManager.Get(r.Context(), "user").(models.User)

	lang := r.URL.Query().Get("lang")
	if lang != "" {
		if l, ok := models.LanguageFromValue(lang); ok {
			user.Settings.Language = &l
			app.SessionManager.Put(r.Context(), "user", user)
		}
	}

	data["user"] = user
	data["labelMapper"] = app.LabelMapper.WithLanguage(user.Settings.Language)
	data["CSRFToken"] = nosurf.Token(r)
	data["flashMessage"] = app.SessionManager.Pop(r.Context(), "flash")
	return data
}
