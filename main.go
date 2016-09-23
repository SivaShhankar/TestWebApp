package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
	<head>
		<title>Hello Gopher</title>
	</head>
	<body>
		Hello Gopher </br>
		It is really awesome that both Docker and Kubernetes are written in Go!
	</body>
</html>`,
	)
}
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go web app powered by Docker")
}

// GetPort -- get the Port from the Dynamic environment
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	fmt.Println("Running Port is" + port)
	return ":" + port
}
func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
