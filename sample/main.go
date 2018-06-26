package main

import (
	//"bufio"
	//"bytes"
	//"crypto/tls"
	//"fmt"
	//"io"
	//"io/ioutil"
	//"log"
	"net/http"
	//"os"
	//"strconv"
	//"time"
	//"google.golang.org/api/googleapi"
	//"encoding/json"
	//"os/exec"
	//"os/user"
	//log "github.com/Sirupsen/logrus"

	//yaml "gopkg.in/yaml.v2"
	//"github.com/gorilla/mux"

	"encoding/base64"


)
func basicAuth(username, password string) string {
	auth := username + ":" + password
	 return base64.StdEncoding.EncodeToString([]byte(auth))
  }
  
  func redirectPolicyFunc(req *http.Request, via []*http.Request) error{
   req.Header.Add("Authorization","Basic " + basicAuth("yamugcp9@gmail.com","9pcgumay"))
   return nil
  }

  // init function
  func init() {
		  utilities.Configuration()
  }

  func CoresMiddleware(h http.Handler) http.Handler {
	      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				  log.Println("middleware", r.URL)
				  
				  allowedHeaders :=handlers.AllowedHeaders([]string{"Acess-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "token"})
				  allowedOrigins :=handlers.AllowedOrigins([]string{"*"})
				  allowedMethods :=handlers.AllowedMethods([]string{"GET", "HEAD","POST", "PUT", "DELETE", "OPTIONS"})
				  handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)
				  h.ServeHTTP(w, r)
		  })
  }
  
  func main() {

	router :=mux.NewRouter()
	router.Use(CoresMiddleware)

	//create project
	router.HandleFunc("/project", utilities.CreateProject).Methods("POST", "OPTIONS")
	router.HandleFunc("/projects", utilities.GetProjects).Methods("GET")

	//delete project
	router.HandleFunc("/project/{projectID}", utilities.ProjectDel).Methods("DELETE")

	// auth
	router.Handlefunc("/login", users.Login).Methods("POST")
	router.HandleFunc("/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/resetpassword", utilities.Resetpassword).Methods("PUT", "OPTIONS")
	router.HandleFunc("/resetpassword_link/{user}", utilities.SendResetPassword).Methods("GET")

	//status
	router.HandleFunc("/getDepStatus", controllers.GetDepStatus).Methods("POST")
	router.HandleFunc("/getPodStatus", controllers.GetPodStatus).Methods("POST")
	router.HandleFunc("/getAllDep", controllers.GetAllKubeDep).Methods("POST")


	//utilities.GetInstance()
  }
