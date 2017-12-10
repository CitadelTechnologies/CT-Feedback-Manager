package main

import(
    "ct-feedback-manager/manager"
    "log"
    "net/http"
    "fmt"
    "os"
)

func main() {
    fmt.Println("MongoDB server initialization")

    manager.InitMongo()
    defer manager.MongoDBConnection.Close()

    fmt.Println("MongoDB is ready")
    fmt.Println("Router initialization")

    router := NewRouter()

    fmt.Println("Server is listening on port " + os.Getenv("SERVER_PORT"))

    log.Fatal(http.ListenAndServe(":" + os.Getenv("SERVER_PORT"), router))
}
