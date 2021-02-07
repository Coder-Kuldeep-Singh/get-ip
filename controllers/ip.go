package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// IP helps to generate the json data
type IP struct {
	Ip string `json:"ip"`
}

func formatter(req *http.Request) {
	log.Printf("%s %s", req.Method, req.RequestURI)
}

// GetIP returns the ip of the user
func GetIP(res http.ResponseWriter, req *http.Request) {
	formatter(req)
	// ip := getIPAdress(req)
	// log.Println("one", ip)
	// log.Println(req.RemoteAddr)
	found := IP{Ip: req.RemoteAddr}
	res.Header().Set("Content-Type", "application/json")
	build, err := json.Marshal(found)
	if err != nil {
		log.Println("error to marshall the ip:", err)
		res.Write([]byte("error:Internal Error"))
		return
	}
	// res.Write(build)
	fmt.Fprintf(res, "%v", string(build))
}
