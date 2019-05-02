package main

import (
        "fmt"
        "log"
        "net/http"
        "golang.org/x/crypto/bcrypt"
)



func HashPassword(password string) (string, error) {
        bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
        return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
        err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
        return err == nil
}

func HomeEndPoint(w http.ResponseWriter, r *http.Request) {
        password := "secret"
        hash, _ := HashPassword(password) // ignore error for the sake of simplicity
        match := CheckPasswordHash(password, hash)

        fmt.Fprintln(w, "Password:", password)
        // fmt.Fprintln(w,password)
        fmt.Fprintln(w,"Hash:    ", hash)
        fmt.Fprintln(w,"Match:   ", match)

}

func main() {
        http.HandleFunc("/", HomeEndPoint)

        if err := http.ListenAndServe(":3000", nil); err != nil {
                log.Fatal(err)
        }

        password := "secret"
        hash, _ := HashPassword(password) // ignore error for the sake of simplicity




        fmt.Println("Password:", password)
        fmt.Println("Hash:    ", hash)

        match := CheckPasswordHash(password, hash)
        fmt.Println("Match:   ", match)
}

