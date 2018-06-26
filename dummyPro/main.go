package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func test() {
	fmt.Println("I am runnning task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println(" Oops ! Not received proper arguments, list of " +
			"arguments are {topic Name, Number of Iterations}")
		os.Exit(1)
	} else {

		var NoofIterationsInString = os.Args[2]
		fmt.Println(NoofIterationsInString)
		var NoofIterations, err = strconv.Atoi(NoofIterationsInString)

		if NoofIterations == -1 {
			NoofIterations = 900000
		}

		fmt.Println(NoofIterations)
		fmt.Println(err)
		for i := 0; i < NoofIterations; i++ {

			fmt.Println(time.Now())

			var username string = "admin"
			var passwd string = "XOIknQf3gqXrX0rr"
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}

			client := &http.Client{Transport: tr}

			f, err := os.Open(os.Args[3])
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			r := bufio.NewReaderSize(f, 4*1024)
			line, isPrefix, err := r.ReadLine()
			var req *http.Request

			for err == nil && !isPrefix {
				s := string(line)
				req, _ = http.NewRequest("GET", s, nil)
				fmt.Println(s)
				line, isPrefix, err = r.ReadLine()
				req.SetBasicAuth(username, passwd)
				resp, err := client.Do(req)
				if err != nil {
					log.Fatal(err)
				}
				bodyText, err := ioutil.ReadAll(resp.Body)

				byteR := bytes.NewReader(bodyText)

				s = string(bodyText)
				fmt.Println(s)
				fmt.Println(time.Now())
				fmt.Println(resp)

				req1, err1 := http.NewRequest("POST", "http://localhost/api/v1/ingest/analytics", byteR)

				req1.Header.Set("content-type", "application/json")
				req1.Header.Set("topic-name", os.Args[1])
				req1.Header.Set("x-auth-header", "abc")
				res, _ := client.Do(req1)

				fmt.Println(res)
				fmt.Println(err1)
			}
			if isPrefix {
				fmt.Println("buffer size to small")
				return
			}
			if err != io.EOF {
				fmt.Println(err)
				return
			}
			//for i:=0;i<len(file);i++{
			//req, err := http.NewRequest("GET",file[i] , nil)
			//}
			//"https://35.184.170.175/apis/metrics.k8s.io/v1beta1/nodes/"

			//resp, err = http.Post("http://localhost/api/v1/ingest/analytics", "application/json", byteR)

		}

	}
}
