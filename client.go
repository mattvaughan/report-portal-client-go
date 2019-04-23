package reportportal

import "net/http"
import "encoding/json"
import "bytes"
import "io/ioutil"

type Client struct {
	ApiKey string
	ApiUrl string
}

type CreateProjectRequest struct {
	ProjectName string
	EntryType string
}

type GetProjectResponse struct {
	ProjectId string
	Configuration ProjectConfiguration
}

type ProjectConfiguration struct {
	ExternalSystem []string
	EntryType string
	StatisticCalculationStrategy string
	ProjectSpecific string
	InterruptedJob string
	KeepLogs string
	KeepScreenshots string
	EmailConfiguration ProjectEmailConfiguration
	AnalyzerConfiguration ProjectAnalyzerConfiguration
}

type ProjectEmailConfiguration struct {
	EmailEnabled bool
	FromAddress string
	EmailCases EmailCase
}

type EmailCase struct {
	Recipients []string
	SendCase string
	LaunchNames []string
	Tags []string
}

type ProjectAnalyzerConfiguration struct {
	MinDocFreq int
	MinTermFreq int
	MinShouldMatch int
	NumberOfLogLines int
	IsAutoAnalyzerEnabled bool
	AnalyzerMode string `json:"analyzer_mode"`
	IndexingRunning string `json:"indexing_running"`
}

func (c *Client) CreateProject(projectName string) string{
	requestBody := &CreateProjectRequest{
		ProjectName: projectName,
		EntryType: "INTERNAL",
	}

	bodyStream, _ := json.Marshal(requestBody)

	client := &http.Client{}

	req, _ := http.NewRequest("POST", c.ApiUrl + "/project", bytes.NewReader(bodyStream))
	req.Header.Set("Authorization", "bearer " + c.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func (c *Client) DeleteProject(projectName string) string{
	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", c.ApiUrl + "/project/" + projectName, nil)
	req.Header.Set("Authorization", "bearer " + c.ApiKey)
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func (c *Client) GetProject(projectName string) *GetProjectResponse{
	client := &http.Client{}

	req, _ := http.NewRequest("GET", c.ApiUrl + "/project/" + projectName, nil)
	req.Header.Set("Authorization", "bearer " + c.ApiKey)
	resp, _ := client.Do(req)

	getProjectResponse := GetProjectResponse{}
	json.NewDecoder(resp.Body).Decode(&getProjectResponse)

	defer resp.Body.Close()
	return &getProjectResponse
}
