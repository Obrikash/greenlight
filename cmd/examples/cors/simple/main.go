package main

import (
	"flag"
	"log"
	"net/http"
)

const html = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
</head>
<body>
<h1>Simple CORS</h1>
<div id="output"></div>
<script>
    document.addEventListener("DOMContentLoaded", () => {
  fetch("http://localhost:4000/v1/healthcheck")
    .then((response) => {
      response.text().then((text) => {
        document.getElementById("output").innerHTML = text;
      });
    })
    .catch((err) => {
      document.getElementById("output").innerHTML = err;
    });
});

</script>
</body>
</html>
`

func main() {
	addr := flag.String("addr", "localhost:9002", "Server address")
	flag.Parse()

	log.Printf("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	}))
	log.Fatal(err)
}
