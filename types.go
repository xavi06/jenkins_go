package gojks

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
