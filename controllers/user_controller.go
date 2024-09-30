package controllers

import (
    "net/http"
    "encoding/json"
    "user-service/config"
    "user-service/models"
    "user-service/utils"
    
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = config.GetCollection("users")

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    _, err := models.FindUserByEmail(userCollection, user.Email)
    if err == nil {
        http.Error(w, "User already exists", http.StatusBadRequest)
        return
    }

    _, err = user.SaveUser(userCollection)
    if err != nil {
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
    var credentials models.User
    _ = json.NewDecoder(r.Body).Decode(&credentials)

    user, err := models.FindUserByEmail(userCollection, credentials.Email)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    isValid := models.VerifyPassword(user.Password, credentials.Password)
    if !isValid {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateJWT(user.Email)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
