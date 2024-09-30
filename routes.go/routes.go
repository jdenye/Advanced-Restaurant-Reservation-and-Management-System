package routes

import (
    "net/http"
    "user-service/controllers"
)

func InitializeRoutes() {
    http.HandleFunc("/register", controllers.Register)
    http.HandleFunc("/login", controllers.Login)
}
