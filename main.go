package main

import (
    "encoding/json"
    "database/sql"
    "fmt"
    "io/ioutil"
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

type Config struct {
    Address  string `json:"mysql_address"`
    Port     string `json:"mysql_port"`
    User     string `json:"mysql_user"`
    Password string `json:"mysql_password"`
    DB       string `json:"mysql_db"`
}

func main() {
    address := os.Getenv("MYSQL_ADDRESS")
    port := os.Getenv("MYSQL_PORT")
    user := os.Getenv("MYSQL_USER")
    passwd := os.Getenv("MYSQL_PASSWORD")
    db := os.Getenv("MYSQL_DATABASE")

    if bytes, err := ioutil.ReadFile("/tmp/web.json"); err == nil {
        conf := new(Config)
        if json.Unmarshal(bytes, conf) == nil {
            address = conf.Address
            port = conf.Port
            user = conf.User
            passwd = conf.Password
            db = conf.DB
        }
    }

    if len(address) == 0 || len(port) == 0 || len(user) == 0 || len(passwd) == 0 || len(db) == 0 {
        log.Fatalln("Invalid mysql connection args")
    }

    if err := connect(address, port, user, passwd, db); err != nil {
        log.Fatalln(err)
    }

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/env", env)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
