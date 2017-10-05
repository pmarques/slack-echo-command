package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// Bind HTTP server port
	viper.SetDefault("port", 8080)
	viper.BindEnv("PORT")
	// Bind HTTP server address (empty will bind to all interfaces)
	viper.SetDefault("bind", "")
	viper.BindEnv("BIND_ADDRESS")

	// Start HTTP server
	listenAddress := fmt.Sprintf("%s:%d", "", viper.GetInt("port"))
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(listenAddress, nil)
	log.Fatal(err)
}

func handle(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form.", http.StatusBadRequest)
		return
	}

	text := r.Form.Get("text")

	fmt.Fprintf(w, text)
}
