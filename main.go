package main

import(
    "ct-feedback-manager/mongo"
    "log"
    "net/http"
    "fmt"
    "os"
)

func main() {
    defer mongo.Connection.Close()

    fmt.Println("Router initialization")

    router := NewRouter()

    fmt.Println("Server is listening on port " + os.Getenv("SERVER_PORT"))

    log.Fatal(http.ListenAndServe(":" + os.Getenv("SERVER_PORT"), router))
}
