package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"

    _ "github.com/go-sql-driver/mysql"
)

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Hello World")
}

func env(res http.ResponseWriter, req *http.Request) {
    env := os.Environ()
    fmt.Fprintln(res, "List of Environtment variables : \n")
    for index, value := range env {
        name := strings.Split(value, "=")
        fmt.Fprintf(res, "[%d] %s : %v\n", index, name[0], name[1])
    }
}

func connect(host, port, user, password, database string) error {
    uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

    db, err := sql.Open("mysql", uri)
    if  err != nil {
        return err
    }

    return db.Ping()
}

func main() {
    address := os.Getenv("MYSQL_ADDRESS")
    port := os.Getenv("MYSQL_PORT")
    user := os.Getenv("MYSQL_USER")
    passwd := os.Getenv("MYSQL_PASSWORD")
    db := os.Getenv("MYSQL_DATABASE")

    if len(address) == 0 || len(port) == 0 || len(user) == 0 || len(passwd) == 0 || len(db) == 0 {
        log.Fatal("Invalid mysql env")
    }

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/env", env)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
