function showGlobalAlert(message, isError) {
  let alertElem = $('#global-alert').clone();

  alertElem.find("span.content").text(message);
  if (isError) {
    alertElem.addClass('bg-red-500 border-red-600');
    alertElem.find("span.title").text("Error: ");
  } else {
    alertElem.addClass('bg-green-500 border-green-600');
    alertElem.find("span.title").text("Info: ");
  }
  alertElem.removeAttr('id');
  alertElem.removeClass('hidden');

  $('#global-alerts').prepend(alertElem);

  var progressBarInner = $('#progress-bar-inner')[0];
  // Update the progress bar every 100ms
  var progress = 0;
  setInterval(function() {
    progress = Math.min(100, progress + 1);
    progressBarInner.style.width = progress + "%";
    if (progress >= 100) {
      clearInterval(this);
      alertElem.fadeOut(200, () => {
        $(this).remove()
      });
    }
  }, 25);
}

function addUrlParamAndRefresh(key, value) {
  var currentUrl = new URL(window.location.href);
  currentUrl.searchParams.set(key, value);
  window.location = currentUrl.href;
};
