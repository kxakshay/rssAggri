package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

func main(){
	godotenv.Load(".env")
	portSting := os.Getenv("PORTS")
	if portSting == "" {
		log.Fatal("PORT env variable not found")
	}
	fmt.Println("Port  : " + portSting)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1_router := chi.NewRouter()
	v1_router.Get("/healthz", handlerReadiness)
	v1_router.Get("/err", handlerErr)

	router.Mount("/v1",v1_router)
	srv := &http.Server{
		Addr:    ":" + portSting,
		Handler: router,
	}


	log.Printf("Serving on port: %s\n", portSting)
	log.Fatal(srv.ListenAndServe())
}
