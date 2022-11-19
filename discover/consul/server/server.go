package main

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	registerWithConsul()
	log.Println("Starting Hello World Server...")
	http.HandleFunc("/hello_world", helloWorld)
	http.HandleFunc("/check", check)
	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		log.Println(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("hello world service is called.")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Hello world.")
	if err != nil {
		log.Println("hello service error ", err)
	}
}

func check(w http.ResponseWriter, r *http.Request) {
	log.Println("check service is called.")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Consul Check")
	if err != nil {
		log.Println("check service error ", err)
	}
}

func registerWithConsul() {
	config := consulApi.DefaultConfig()
	consul, err := consulApi.NewClient(config)
	if err != nil {
		log.Println("client-error", err)
	}

	serviceId := "helloWorld-server"
	port, err := strconv.Atoi(getPort()[1:len(getPort())])
	if err != nil {
		log.Println("strconv-error", err)
	}
	address := getHostName()

	regConf := &consulApi.AgentServiceRegistration{
		ID:      serviceId,
		Name:    serviceId,
		Port:    port,
		Address: address,
		Check: &consulApi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/check", address, port),
			Interval: "10s",
			Timeout:  "30s",
		},
	}
	regErr := consul.Agent().ServiceRegister(regConf)
	if regErr != nil {
		log.Printf("Failed to register service: %s:%v ", address, port)
	} else {
		log.Printf("successfully register service: %s:%v", address, port)
	}

}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return fmt.Sprintf(":%s", port)
}

func getHostName() string {
	hostName, err := os.Hostname()
	if err != nil {
		log.Println("hostname get error", err.Error())
	}
	return hostName
}
