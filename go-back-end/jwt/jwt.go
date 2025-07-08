package main

import (
 "fmt"
 "github.com/golang-jwt/jwt/v5"
 "time"
 "net/http"
 "encoding/json"
)

var secretKey = []byte("secret-key")

func createToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
    return "", err
    }

 return tokenString, nil
}

func verifyToken(tokenString string) error {
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return secretKey, nil
   })
  
   if err != nil {
      return err
   }
  
   if !token.Valid {
      return fmt.Errorf("invalid token")
   }
  
   return nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
 
  var u User
  json.NewDecoder(r.Body).Decode(&u)
  fmt.Printf("The user request value %v", u)
  
  if u.Username == "Chek" && u.Password == "123456" {
    tokenString, err := CreateToken(u.Username)
    if err != nil {
       w.WriteHeader(http.StatusInternalServerError)
       fmt.Errorf("No username found")
     }
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, tokenString)
    return
  } else {
    w.WriteHeader(http.StatusUnauthorized)
    fmt.Fprint(w, "Invalid credentials")
  }
}
