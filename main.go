package main

import (
  "os"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
  "github.com/gorilla/handlers"
  "github.com/alif-faishol/devchallenge-alifaishol-backend/handler"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  router := mux.NewRouter()

  router.HandleFunc("/cors", handler.CorsProxy).Methods("POST")

  headers := handlers.AllowedHeaders([]string{"Content-Type"})
  methods := handlers.AllowedHeaders([]string{"GET", "POST"})
  origins := handlers.AllowedOrigins([]string{"*"})

  log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), handlers.CORS(headers, methods, origins)(router)))
}
