package gojks

import (
	"bytes"
	"encoding/json"
	"fmt"
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
func NewAPI(url string, username string, password string) *API {
	return &API{url, username, password, &http.Client{}}
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

// getParams func
func getParams(d map[string]interface{}) (res []byte, err error) {
	/*
		var res string
		var resArr []string
		for k, v := range d {
			resArr = append(resArr, fmt.Sprintf("%s=%s", k, v))
		}
		res = strings.Join(resArr, "&")
		return res
	*/
	var params map[string][]param
	params = make(map[string][]param)
	var resArr []param
	for k, v := range d {
		resArr = append(resArr, param{Name: k, Value: v})
	}
	params["parameter"] = resArr
	res, err = json.Marshal(params)
	fmt.Println(string(res))
	return

}

// HTTPDoHeader func
func (api *API) HTTPDoHeader(ur string, method string, header map[string]string, params map[string]string) (body []byte, err error) {
	/*
		data, err := getParams(d)
		if err != nil {
			return
		}
	*/
	//fmt.Println(bytes.NewBuffer(data))
	data := url.Values{}
	for k, v := range params {
		data.Set(k, v)
	}
	req, err := http.NewRequest(method, ur, bytes.NewBufferString(data.Encode()))
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
	url := api.baseurl + "/api/json"
	body, err := api.HTTPDo(url, "GET")
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
	url := api.baseurl + "/api/json"
	body, err := api.HTTPDo(url, "GET")
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
	url := fmt.Sprintf("%s/job/%s/api/json", api.baseurl, j)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return
}

// GetJobParams func
func (api *API) GetJobParams(j string) (data []ParameterDefinition, err error) {
	url := fmt.Sprintf("%s/job/%s/api/json", api.baseurl, j)
	body, err := api.HTTPDo(url, "GET")
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
	url := fmt.Sprintf("%s/job/%s/%d/api/json", api.baseurl, j, b)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return
}

// GetBuildConsoleOutput func
func (api *API) GetBuildConsoleOutput(j string, b int) (data string, err error) {
	url := fmt.Sprintf("%s/job/%s/%d/consoleText", api.baseurl, j, b)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	return string(body), nil

}

// Build func
func (api *API) Build(j string) error {
	url := fmt.Sprintf("%s/job/%s/build", api.baseurl, j)
	body, err := api.HTTPDo(url, "POST")
	/*
		header := map[string]string{
			"Content-Type": "application/json",
		}
		data, err := getParams(d)
		if err != nil {
			return err
		}
		_, err = api.HTTPDoHeader(url, "POST", header, data)
	*/
	if err != nil {
		return err
	}
	data := string(body)
	fmt.Println(data)
	return nil

}

// BuildWithParams func
func (api *API) BuildWithParams(j string, d map[string]string) error {
	url := fmt.Sprintf("%s/job/%s/buildWithParameters", api.baseurl, j)

	header := map[string]string{
		//"Content-Type": "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	_, err := api.HTTPDoHeader(url, "POST", header, d)

	if err != nil {
		return err
	}
	//da := string(body)
	//fmt.Println(da)
	return nil
}

// ListRunNodes func
func (api *API) ListRunNodes(p string, r string) (data []Node, err error) {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/%s/nodes/", api.baseurl, p, r)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return

}

// ListNodeSteps func
func (api *API) ListNodeSteps(p string, r string, n string) (data []Step, err error) {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/%s/nodes/%s/steps/", api.baseurl, p, r, n)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return
}

// GetStepLog func
func (api *API) GetStepLog(p string, r string, n string, s string) (data string, err error) {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/%s/nodes/%s/steps/%s/log/", api.baseurl, p, r, n, s)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	data = string(body)
	return
}

// GetRunLog func
func (api *API) GetRunLog(p string, r string) (data string, err error) {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/%s/log/", api.baseurl, p, r)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	data = string(body)
	return
}

// Stop func
// Stop a build
func (api *API) Stop(p string, r string) error {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/%s/stop", api.baseurl, p, r)
	header := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := api.HTTPDoHeader(url, "PUT", header, nil)
	if err != nil {
		return err
	}
	return err

}
