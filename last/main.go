package main

import (
	"net/http"
	"log"
    "io/ioutil"
    "fmt"
)

func main() {
    data := basicAuth()
    fmt.Println("response:", data)
}

func basicAuth() string {
    var username string = "yamugcp9@gmail.com"
    var passwd string = "9pcgumay"
    client := &http.Client{}
    req, err := http.NewRequest("GET", "http://yamugcp9@gmail.com", nil)
    req.SetBasicAuth(username, passwd)
    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
    s := string(bodyText)
    return s
}