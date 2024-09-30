package main

import (
    "fmt"
    "log"
    "net/http"
    "user-service/config"
    "user-service/routes"
)

func main() {
    config.ConnectDB()

    routes.InitializeRoutes()

    fmt.Println("Starting server on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
