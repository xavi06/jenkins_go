package gojks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// API struct
type API struct {
	baseurl  string
	username string
	password string
	client   *http.Client
}

type param struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// NewAPI func
func NewAPI(ur string, username string, password string) *API {
	return &API{ur, username, password, &http.Client{}}
}

// HTTPDo func
func (api *API) HTTPDo(url string, method string) (body []byte, err error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}
	req.SetBasicAuth(api.username, api.password)
	response, err := api.client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	return

}

// formatParams func
func formatParams(params map[string]string) url.Values {
	data := url.Values{}
	for k, v := range params {
		data.Set(k, v)
	}
	return data
}

// HTTPDoHeader func
func (api *API) HTTPDoHeader(ur string, method string, header map[string]string, bd io.Reader) (body []byte, err error) {
	req, err := http.NewRequest(method, ur, bd)
	if err != nil {
		return
	}
	req.SetBasicAuth(api.username, api.password)

	for k, v := range header {
		req.Header.Add(k, v)
	}
	response, err := api.client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	return

}

// ListJobs func
func (api *API) ListJobs() (data []Job, err error) {
	ur := api.baseurl + "/api/json"
	body, err := api.HTTPDo(ur, "GET")
	if err != nil {
		return
	}
	var apiinfo APIInfo
	err = json.Unmarshal(body, &apiinfo)
	if err != nil {
		return
	}
	data = apiinfo.Jobs
	return
}

// ListViews func
func (api *API) ListViews() (data []View, err error) {
	ur := api.baseurl + "/api/json"
	body, err := api.HTTPDo(ur, "GET")
	if err != nil {
		return
	}
	var apiinfo APIInfo
	err = json.Unmarshal(body, &apiinfo)
	if err != nil {
		return
	}
	data = apiinfo.Views
	return
}

// GetJob func
func (api *API) GetJob(j string) (data JobInfo, err error) {
	ur := fmt.Sprintf("%s/job/%s/api/json", api.baseurl, j)
	body, err := api.HTTPDo(ur, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return
}

// GetJobParams func
func (api *API) GetJobParams(j string) (data []ParameterDefinition, err error) {
	ur := fmt.Sprintf("%s/job/%s/api/json", api.baseurl, j)
	body, err := api.HTTPDo(ur, "GET")
	if err != nil {
		return
	}
	var jobInfo JobInfo
	json.Unmarshal(body, &jobInfo)
	for _, property := range jobInfo.Property {
		if property.Class == "hudson.model.ParametersDefinitionProperty" {
			data = property.ParameterDefinitions
		}
	}
	return
}

// GetJobBuild func
func (api *API) GetJobBuild(j string, b int) (data JobBuild, err error) {
	ur := fmt.Sprintf("%s/job/%s/%d/api/json", api.baseurl, j, b)
	body, err := api.HTTPDo(ur, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return
}

// GetBuildConsoleOutput func
func (api *API) GetBuildConsoleOutput(j string, b int) (data string, err error) {
	ur := fmt.Sprintf("%s/job/%s/%d/consoleText", api.baseurl, j, b)
	body, err := api.HTTPDo(ur, "GET")
	if err != nil {
		return
	}
	return string(body), nil

}

// Build func
func (api *API) Build(j string) error {
	ur := fmt.Sprintf("%s/job/%s/build", api.baseurl, j)
	body, err := api.HTTPDo(ur, "POST")

	if err != nil {
		return err
	}
	data := string(body)
	fmt.Println(data)
	return nil

}

// BuildWithParams func
func (api *API) BuildWithParams(j string, d map[string]string) error {
	ur := fmt.Sprintf("%s/job/%s/buildWithParameters", api.baseurl, j)

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	data := formatParams(d)

	_, err := api.HTTPDoHeader(ur, "POST", header, bytes.NewBufferString(data.Encode()))

	if err != nil {
		return err
	}
	//da := string(body)
	//fmt.Println(da)
	return nil
}

// CreateJob func
func (api *API) CreateJob(j string, config string) error {
	ur := fmt.Sprintf("%s/createItem?name=%s", api.baseurl, j)
	header := map[string]string{
		"Content-Type": "text/xml",
	}
	_, err := api.HTTPDoHeader(ur, "POST", header, bytes.NewBufferString(config))
	if err != nil {
		return err
	}
	//fmt.Println(string(body))
	return nil

}

// DeleteJob func
func (api *API) DeleteJob(j string) error {
	ur := fmt.Sprintf("%s/job/%s/doDelete", api.baseurl, j)
	_, err := api.HTTPDo(ur, "POST")
	if err != nil {
		return err
	}
	//fmt.Println(string(body))
	return nil
}

// UpdateJob func
func (api *API) UpdateJob(j, config string) error {
	ur := fmt.Sprintf("%s/job/%s/config.xml", api.baseurl, j)
	body, err := api.HTTPDoHeader(ur, "POST", nil, bytes.NewBufferString(config))
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
