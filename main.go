package main

import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
    "database/sql"
   _ "github.com/lib/pq"
     "time"
)

const (
     DB_USER     = "postgres"
     DB_PASSWORD = ""
     DB_NAME     = "ab_log_db"
)


func main() {
        temp, hum :=getLight("http://192.168.71.74/sec/?pt=32&cmd=get")
        //fmt.Println(light)
        inserter(temp, hum)
}

func inserter(temp_val string,hum_val string) {
        dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
        db, err := sql.Open("postgres", dbinfo)
        checkErr(err)
        defer db.Close()

        fmt.Println("# Inserting values")
        dt := time.Now()
        var lastInsertId int
        err = db.QueryRow("INSERT INTO weather (temp_val,hum_val,w_date) VALUES($1,$2,$3) returning w_id;", temp_val,hum_val, dt).Scan(&lastInsertId)
        checkErr(err)
        //fmt.Println("last inserted id =", lastInsertId)

}

func getLight(host string )(temp_val string, hum_val string){
    resp, err := http.Get(host) 
    if err != nil { 
        fmt.Println(err) 
        return
    } 
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
          fmt.Println(err)
          return
    }
    //fmt.Println(string(body))
    temphum := strings.Replace(string(body),"/",":",-1)
    temphum22 := strings.Split(temphum,":")
    temp_val = temphum22[1]
    hum_val = temphum22[3]
    fmt.Println(temp_val, hum_val)
    return temp_val, hum_val 
}

func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }

