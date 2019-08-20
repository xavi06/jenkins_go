package gojenkins

// APIInfo struct
type APIInfo struct {
	Class           interface{} `json:"_class"`
	AssignedLabels  interface{} `json:"assignedLabels"`
	Mode            string
	NodeDescription string `json:"nodeDescription"`
	NodeName        string `json:"nodeName"`
	NumExecutors    int    `json:"numExecutors"`
	Description     interface{}
	Jobs            []Job
	OverallLoad     interface{}
	PrimaryView     interface{} `json:"primaryView"`
	QuietingDown    bool        `json:"quietingDown"`
	SlaveAgentPort  int         `json:"slaveAgentPort"`
	UnlabeledLoad   interface{} `json:"unlabeledLoad"`
	UseCrumbs       bool        `json:"useCrumbs"`
	UseSecurity     bool        `json:"useSecurity"`
	Views           []View
}

// Job struct
type Job struct {
	Class interface{} `json:"_class"`
	Name  string
	URL   string `json:"url"`
	Color string
}

// View struct
type View struct {
	Class interface{} `json:"_class"`
	Name  string
	URL   string `json:"url"`
}

// JobInfo struct
type JobInfo struct {
	Class                 interface{} `json:"_class"`
	Actions               interface{} `json:"actions"`
	Description           string
	DisplayName           string      `json:"displayName"`
	DisplayNameOrNull     interface{} `json:"displayNameOrNull"`
	FullDisplayName       string      `json:"fullDisplayName"`
	FullName              string      `json:"fullName"`
	Name                  string
	URL                   string `json:"url"`
	Buildable             bool
	Builds                []Build
	Color                 string
	FirstBuild            Build       `json:"firstBuild"`
	HealthReport          interface{} `json:"healthReport"`
	InQueue               bool        `json:"inQueue"`
	KeepDependencies      bool        `json:"keepDependencies"`
	LastBuild             Build       `json:"lastBuild"`
	LastCompletedBuild    Build       `json:"lastCompletedBuild"`
	LastFailedBuild       Build       `json:"lastFailedBuild"`
	LastStableBuild       Build       `json:"lastStableBuild"`
	LastSuccessfulBuild   Build       `json:"lastSuccessfulBuild"`
	LastUnstableBuild     Build       `json:"lastUnstableBuild"`
	LastUnsuccessfulBuild Build       `json:"lastUnsuccessfulBuild"`
	NextBuildNumber       int         `json:"nextBuildNumber"`
	Property              []Property
	QueueItem             interface{} `json:"queueItem"`
	ConcurrentBuild       bool        `json:"concurrentBuild"`
	ResumeBlocked         bool        `json:"resumeBlocked"`
}

// Build struct
type Build struct {
	Class  interface{} `json:"_class"`
	Number int
	URL    string `json:"url"`
}

// Property struct
type Property struct {
	Class                interface{}           `json:"_class"`
	ParameterDefinitions []ParameterDefinition `json:"parameterDefinitions"`
}

// ParameterDefinition struct
type ParameterDefinition struct {
	Class                 interface{}           `json:"_class"`
	DefaultParameterValue DefaultParameterValue `json:"defaultParameterValue"`
	Description           string                `json:"description"`
	Name                  string                `json:"name"`
	Type                  string                `json:"type"`
}

// DefaultParameterValue struct
type DefaultParameterValue struct {
	Class interface{} `json:"_class"`
	Name  string
	Value string
}

// JobBuild struct
type JobBuild struct {
	Class             interface{} `json:"_class"`
	Actions           interface{}
	Artifacts         interface{}
	Building          bool
	Description       interface{}
	DisplayName       string `json:"displayName"`
	Duration          int
	EstimatedDuration int `json:"estimatedDuration"`
	Executor          interface{}
	FullDisplayName   string `json:"fullDisplayName"`
	ID                string `json:"id"`
	KeepLog           bool   `json:"keepLog"`
	Number            int
	QueueID           int `json:"queueId"`
	Result            string
	Timestamp         int64       `json:"timestamp"`
	URL               string      `json:"url"`
	ChangeSets        interface{} `json:"changeSets"`
	Culprits          interface{} `json:"culprits"`
	NextBuild         interface{} `json:"nextBuild"`
	PreviousBuild     interface{} `json:"previousBuild"`
}

/**/
// Pipeline struct
type Pipeline struct {
	Class                     interface{} `json:"_class"`
	Links                     interface{} `json:"_links"`
	Actions                   interface{} `json:"actions"`
	DisplayName               string      `json:"displayName"`
	EstimatedDurationInMillis int         `json:"estimatedDurationInMillis"`
	FullDisplayName           string      `json:"fullDisplayName"`
	FullName                  string      `json:"fullName"`
	LatestRun                 interface{} `json:"latestRun"`
	Name                      string      `json:"name"`
	Oganization               string      `json:"organization"`
	Parameters                []Parameter `json:"parameters"`
	Permissions               interface{} `json:"permissions"`
	WeatherScore              int         `json:"weatherScore"`
}

// Parameter struct
type Parameter struct {
	Class                 interface{}           `json:"_class"`
	DefaultParameterValue DefaultParameterValue `json:"defaultParameterValue"`
	Description           string                `json:"description"`
	Name                  string                `json:"name"`
	Type                  string                `json:"type"`
}

// PipelineRun struct
type PipelineRun struct {
	Class                     interface{} `json:"_class"`
	Links                     interface{} `json:"_links"`
	Actions                   interface{} `json:"actions"`
	ArtifactsZipFile          interface{} `json:"artifactsZipFile"`
	CauseOfBlockage           interface{} `json:"causeOfBlockage"`
	Causes                    []Cause     `json:"causes"`
	ChangeSet                 []ChangeSet `json:"changeSet"`
	Description               string      `json:"description"`
	DurationInMillis          int         `json:"durationInMillis"`
	EnQueueTime               int         `json:"enQueueTime"`
	EndTime                   string      `json:"endTime"`
	EstimatedDurationInMillis int         `json:"estimatedDurationInMillis"`
	ID                        string      `json:"id"`
	Name                      string      `json:"name"`
	Organization              string      `json:"organization"`
	Pipeline                  string      `json:"pipeline"`
	Replayable                bool        `json:"replayable"`
	Result                    string      `json:"result"`
	RunSummary                string      `json:"runSummary"`
	StartTime                 string      `json:"startTime"`
	State                     string      `json:"state"`
	Type                      string      `json:"type"`
	Branch                    string      `json:"branch"`
	CommitID                  string      `json:"commitId"`
	CommitURL                 string      `json:"commitUrl"`
	PullRequest               string      `json:"pullRequest"`
}

// Cause struct
type Cause struct {
	Class            interface{} `json:"_class"`
	ShortDescription string      `json:"shortDescription"`
	UserID           string      `json:"userId"`
	UserName         string      `json:"userName"`
}

// ChangeSet struct
type ChangeSet struct {
	Class         interface{} `json:"_class"`
	Links         interface{} `json:"_links"`
	AffectedPaths []string    `json:"affectedPaths"`
	Author        Author      `json:"author"`
	CommitID      string      `json:"commitId"`
	Issues        interface{} `json:"issues"`
	Msg           string      `json:"msg"`
	Timestamp     string      `json:"timestamp"`
	URL           interface{} `json:"url"`
}

// Author struct
type Author struct {
	Class      interface{} `json:"_class"`
	Links      interface{} `json:"_links"`
	Avatar     interface{} `json:"avator"`
	Email      string      `json:"email"`
	FullName   string      `json:"fullName"`
	ID         string      `json:"id"`
	Permission interface{} `json:"permission"`
}

// Node struct
type Node struct {
	Class              interface{} `json:"_class"`
	Links              interface{} `json:"_links"`
	Actions            interface{} `json:"actions"`
	DisplayDescription string      `json:"displayDescription"`
	DisplayName        string      `json:"displayName"`
	DurationInMillis   int         `json:"durationInMillis"`
	ID                 string      `json:"id"`
	Input              string      `json:"input"`
	Result             string      `json:"result"`
	StartTime          string      `json:"startTime"`
	State              string      `json:"state"`
	Type               string      `json:"type"`
	CauseOfBlockage    string      `json:"causeOfBlockage"`
	Edges              interface{} `json:"edges"`
	FirstParent        interface{} `json:"firstParent"`
	Restartable        bool        `json:"restartable"`
}

// Step struct
type Step struct {
	Class              interface{} `json:"_class"`
	Links              interface{} `json:"_links"`
	Actions            interface{} `json:"actions"`
	DisplayDescription string      `json:"displayDescription"`
	DisplayName        string      `json:"displayName"`
	DurationInMillis   int         `json:"durationInMillis"`
	ID                 string      `json:"id"`
	Input              string      `json:"input"`
	Result             string      `json:"result"`
	StartTime          string      `json:"startTime"`
	State              string      `json:"state"`
	Type               string      `json:"type"`
}

// Favorite struct
type Favorite struct {
	Class interface{} `json:"_class"`
	Links interface{} `json:"_links"`
	Item  Pipeline    `json:"item"`
}
