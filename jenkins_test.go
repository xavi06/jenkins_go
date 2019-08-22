package gojks

import (
	"fmt"
	"testing"
)

var (
	baseurl = "http://172.16.151.114:8080"
	user    = "root"
	passwd  = "jks2018"
)

func TestListJobs(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	data, err := api.ListJobs()
	if err != nil {
		panic(err)
	}
	fmt.Println(data[0].Name)
}

/*
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
*/

/*
func TestBuildWithParams(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	d := make(map[string]string)
	d["aa"] = "xx"
	err := api.BuildWithParams("test11", d)
	fmt.Println(err)
}
*/

/*
func TestGetBuildConsoleOutput(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	data, err := api.GetBuildConsoleOutput("app-microservice-web", 39)
	fmt.Println(data, err)
}
*/

/*
func TestCreateJob(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	config := `<?xml version='1.1' encoding='UTF-8'?>
	<flow-definition plugin="workflow-job@2.29">
	  <actions/>
	  <description></description>
	  <keepDependencies>false</keepDependencies>
	  <properties>
		<hudson.model.ParametersDefinitionProperty>
		  <parameterDefinitions>
			<hudson.model.StringParameterDefinition>
			  <name>aa</name>
			  <description></description>
			  <defaultValue>aa</defaultValue>
			  <trim>false</trim>
			</hudson.model.StringParameterDefinition>
		  </parameterDefinitions>
		</hudson.model.ParametersDefinitionProperty>
	  </properties>
	  <definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps@2.62">
		<script>node {
	   echo &quot;${params.aa}&quot;;
	}</script>
		<sandbox>true</sandbox>
	  </definition>
	  <triggers/>
	  <disabled>false</disabled>
	</flow-definition>
	`
	err := api.CreateJob("test22", config)
	fmt.Println(err)
}
*/
/*
func TestDeleteJob(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	api.DeleteJob("test22")
}
*/

/*
func TestUpdateJob(t *testing.T) {
	api := NewAPI(baseurl, user, passwd)
	config := `<?xml version='1.1' encoding='UTF-8'?>
	<flow-definition plugin="workflow-job@2.29">
	  <actions/>
	  <description></description>
	  <keepDependencies>false</keepDependencies>
	  <properties>
		<hudson.model.ParametersDefinitionProperty>
		  <parameterDefinitions>
			<hudson.model.StringParameterDefinition>
			  <name>aa</name>
			  <description></description>
			  <defaultValue>bb</defaultValue>
			  <trim>false</trim>
			</hudson.model.StringParameterDefinition>
		  </parameterDefinitions>
		</hudson.model.ParametersDefinitionProperty>
	  </properties>
	  <definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps@2.62">
		<script>node {
	   echo &quot;${params.aa}&quot;;
	}</script>
		<sandbox>true</sandbox>
	  </definition>
	  <triggers/>
	  <disabled>false</disabled>
	</flow-definition>
	`
	api.UpdateJob("test22", config)
}
*/
