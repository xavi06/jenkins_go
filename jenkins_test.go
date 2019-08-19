package gojenkins

import (
	"fmt"
	"testing"
)

var (
	baseurl = "http://172.16.151.114:8080"
	user    = "root"
	passwd  = "jks2018"
)

/*
func TestListJobs(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	data, err := api.ListJobs()
	if err != nil {
		panic(err)
	}
	fmt.Println(data[0].Name)
}

func TestListViews(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	data, err := api.ListViews()
	if err != nil {
		panic(err)
	}
	fmt.Println(data[0].Name)
}

func TestGetJob(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	data, err := api.GetJob("app-microservice-web")
	if err != nil {
		panic(err)
	}
	fmt.Println(data.Name)
}

func TestGetParams(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	data, err := api.GetJobParams("app-microservice-web")
	//data, err := api.GetJobParams("static")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func TestGetJobBuild(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	data, err := api.GetJobBuild("app-microservice-web", 39)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func TestBuildWithParams(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	d := make(map[string]string)
	d["aa"] = "cc"
	err := api.BuildWithParams("test11", d)
	fmt.Println(err)
}
*/
func TestGetBuildConsoleOutput(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	data, err := api.GetBuildConsoleOutput("app-microservice-web", 39)
	fmt.Println(data, err)
}
