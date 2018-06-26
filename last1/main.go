package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/base64"
	//"fmt"
	//"time"
)

func main() {
	client := &http.Client{}
	//basicAuth()
	req, err := http.NewRequest("GET", "http://localhost/", nil)
    req.Header.Add("Authorization","Basic " + basicAuth("yamugcp9@gmail.com","9pcgumay"))

	resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
    s := string(bodyText)
	return s
	
}
func basicAuth(username, password string) string {
	auth := username + ":" + password
	 return base64.StdEncoding.EncodeToString([]byte(auth))
  }
  
  func redirectPolicyFunc(req *http.Request, via []*http.Request) error{
   req.Header.Add("Authorization","Basic " + basicAuth("yamugcp9@gmail.com","9pcgumay"))
   return nil
  }