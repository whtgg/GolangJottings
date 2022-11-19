package main

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var url string

func main() {
	discoverWithConsul()
	log.Println("discover starting")
	var client = &http.Client{
		Timeout: time.Second * 30,
	}
	callServerEvery(time.Second*10, client)
}

func discoverWithConsul() {
	config := consulApi.DefaultConfig()
	consul, err := consulApi.NewClient(config)
	if err != nil {
		log.Println("client error", err)
	}
	services, err := consul.Agent().Services()
	if err != nil {
		log.Println("fetch services error", err)
	}
	service := services["helloWorld-server"]
	log.Printf("%v", service)
	url = fmt.Sprintf("http://%s:%d/hello_world", service.Address, service.Port)
}

func hello(t time.Time, client *http.Client) {
	res, err := client.Get(url)
	if err != nil {
		log.Println("Get url error", err)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Read body error", err)
		return
	}
	fmt.Printf("%s time is %v \n", body, t)
}

func callServerEvery(duration time.Duration, client *http.Client) {
	for x := range time.Tick(duration) {
		hello(x, client)
	}
}
