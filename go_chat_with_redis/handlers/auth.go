
package handlers

import (
    "encoding/json"
    "net/http"
    "strings"

    "golang.org/x/crypto/bcrypt"
    "github.com/go-redis/redis/v9"
    "go_chat/redis"
)

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    json.NewDecoder(r.Body).Decode(&creds)
    creds.Username = strings.ToLower(creds.Username)

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    key := "user:" + creds.Username
    err = redis.Rdb.HSet(redis.Ctx, key, "password", hashedPassword).Err()
    if err != nil {
        http.Error(w, "Error saving user", http.StatusInternalServerError)
        return
    }

    w.Write([]byte("User registered successfully"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    json.NewDecoder(r.Body).Decode(&creds)
    creds.Username = strings.ToLower(creds.Username)

    key := "user:" + creds.Username
    storedHash, err := redis.Rdb.HGet(redis.Ctx, key, "password").Result()
    if err == redis.Nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    } else if err != nil {
        http.Error(w, "Error fetching user", http.StatusInternalServerError)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(creds.Password))
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    w.Write([]byte("Login successful"))
}
