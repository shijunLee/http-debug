package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var port int = 9001
var f *flag.FlagSet

func init() {
	f = flag.NewFlagSet("http-debug", flag.ContinueOnError)
	f.IntVar(&port, "port", 9001, "the server start port")
}

func main() {
	f.Parse(os.Args[1:])
	r := mux.NewRouter()
	envPort := os.Getenv("HTTP_DEBUG_PORT")
	if envPort != "" {
		if envPortValue, err := strconv.Atoi(envPort); err == nil {
			port = envPortValue
		}
	}
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header
		query := r.URL.Query()
		path := r.URL.Path
		host := r.Host
		proto := r.Proto
		fragment := r.URL.Fragment
		bodyData, err := ioutil.ReadAll(r.Body)
		body := ""
		if err == nil && len(bodyData) > 0 {
			body = string(bodyData)
		}
		resultMap := map[string]interface{}{}
		resultMap["header"] = header
		resultMap["query"] = query
		resultMap["path"] = path
		resultMap["host"] = host
		resultMap["fragment"] = fragment
		resultMap["proto"] = proto
		resultMap["body"] = body
		responseMap := map[string]interface{}{}
		responseMap["requestId"] = uuid.NewString()
		responseMap["result"] = resultMap
		responseData, _ := json.Marshal(responseMap)
		w.Header().Set("content-type", "application/json")
		w.Write(responseData)
		w.WriteHeader(200)
	}).Methods("GET", "POST", "DELETE", "PUT", "PATCH", "HEAD")
	address := fmt.Sprintf("0.0.0.0:%d", port)
	log.Fatal(http.ListenAndServe(address, r))
}
