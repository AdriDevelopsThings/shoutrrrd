package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AdriDevelopsThings/shoutrrrd/src"
)

func main() {
	errLoadConfig := src.LoadConfig()
	if errLoadConfig != nil {
		fmt.Printf("Error while loading config: %v\n", errLoadConfig)
		os.Exit(1)
	}
	http.HandleFunc("/send", src.Send)
	host := os.Getenv("SHOUTRRRD_HOST")
	if host == "" {
		host = ":8000"
	}
	fmt.Printf("Started http server: %q\n", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Printf("Error while starting http server %v\n", err)
	}
}
