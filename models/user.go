package models

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Name      string             `bson:"name"`
    Email     string             `bson:"email"`
    Password  string             `bson:"password"`
    Role      string             `bson:"role"`
    CreatedAt time.Time          `bson:"created_at"`
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func VerifyPassword(hashedPwd string, plainPwd string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
    return err == nil
}

func (u *User) SaveUser(db *mongo.Collection) (*mongo.InsertOneResult, error) {
    u.CreatedAt = time.Now()
    u.Role = "customer"  // Default role
    u.Password, _ = HashPassword(u.Password)
    return db.InsertOne(context.Background(), u)
}

func FindUserByEmail(db *mongo.Collection, email string) (*User, error) {
    var user User
    err := db.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
    return &user, err
}
