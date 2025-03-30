package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Kawilki-M/receipt-processor/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Starting Receipt Processing Service...")

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println(`
   __               _       _       ___                                        
  /__\ ___  ___ ___(_)_ __ | |_    / _ \_ __ ___   ___ ___  ___ ___  ___  _ __ 
 / \/// _ \/ __/ _ \ | '_ \| __|  / /_)/ '__/ _ \ / __/ _ \/ __/ __|/ _ \| '__|
/ _  \  __/ (_|  __/ | |_) | |_  / ___/| | | (_) | (_|  __/\__ \__ \ (_) | |   
\/ \_/\___|\___\___|_| .__/ \__| \/    |_|  \___/ \___\___||___/___/\___/|_|   
                     |_|                                                       `)

	port, exists := os.LookupEnv("PORT") // Default to 8000 if PORT is not set
	if !exists {
		port = "8000"
	}

	fmt.Printf("\nServer running on localhost:%s\n", port)

	err := http.ListenAndServe("localhost:"+port, r)
	if err != nil {
		log.Error(err)
	}
}
