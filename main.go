package main

import (
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/yantology/simple-http-server-golang/docs" // Import generated docs
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// formHandler godoc
// @Summary Handle form submission
// @Description Handle form submission and parse form data
// @Tags form
// @Accept  json
// @Produce  json
// @Param name query string false "Name"
// @Param address query string false "Address"
// @Success 200 {string} string "POST request successful"
// @Failure 400 {string} string "ParseForm() err"
// @Router /form [post]
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() err: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// helloHandler godoc
// @Summary Say hello
// @Description Respond with a hello message
// @Tags hello
// @Accept  json
// @Produce  json
// @Success 200 {string} string "hello!"
// @Failure 404 {string} string "404 not found"
// @Router /hello [get]
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func serveForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/form.html")
}

func setupServer() *http.ServeMux {
	fileServer := http.FileServer(http.Dir("./static"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/form", serveForm)

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/form", formHandler)
	apiMux.HandleFunc("/hello", helloHandler)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiMux))

	// Swagger endpoint
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	return mux
}

func main() {
	mux := setupServer()
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
