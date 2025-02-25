document.addEventListener("DOMContentLoaded", () => {
  fetch("http://localhost:4000/v1/healthcheck")
    .then((response) => {
      response
        .text()
        .then((text) => {
          document.getElementById("output").innerHTML = text;
        })
        .catch((err) => {
          document.getElementById("output").innerHTML = err;
        });
    })
    .catch((err) => {
      document.getElementById("output").innerHTML = err;
    });
});
