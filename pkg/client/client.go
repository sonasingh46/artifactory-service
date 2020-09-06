package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sonasingh46/artifactory-service/pkg/aql"
	"github.com/sonasingh46/artifactory-service/pkg/types"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Artifactory interface abstracts the operations on artifacts API
// The artifactory clients should implement the methods.
type Artifactory interface {
	GetTopDownloaded() (*types.Artifacts, error)
}

// GetTopDownloaded returns top 2 doownloaded artifacts based on AQL.
// Note : For more info : pkg/aql/aql.go
// Here ArtifactoryClient is the Jfrog artifactory client implementation.
func (ac *ArtifactoryClient) GetTopDownloaded() (*types.Artifacts, error) {
	payload := strings.NewReader(aql.PayLoad)
	newReq, err := http.NewRequest("POST",
		ac.buildURL("artifactory/api/search/aql"),
		payload,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build artifactory client:{%s}", err.Error())
	}
	ac.request = newReq
	resp, err := ac.MakeRequest()
	artifactory := &types.Artifacts{}
	err = json.Unmarshal(resp, artifactory)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal:{%s}", err.Error())
	}
	return artifactory, nil
}

// SetArtifactoryClient sets the client that interacts with the artifactory API
func SetArtifactoryClient() error {
	artifactoryIP := os.Getenv("ART_IP")
	if artifactoryIP == "" {
		return errors.New("Empty value for ART_IP")
	}
	artifactoryPort := os.Getenv("ART_PORT")
	if artifactoryPort == "" {
		return errors.New("Empty value for ART_PORT")
	}

	client = NewArtifactoryClient().
		SetIP(artifactoryIP).
		SetPort(artifactoryPort).
		SetHttpScheme()

	return nil
}

// GetArtifactoryClient return the artifactory API client
func GetArtifactoryClient() Artifactory {
	return client
}

// client holds the artifactory API client.
// This is here as a variable so that the client is set up only once when this
// proxy service warms up.
// This can avoid repeated access of ENVs for details like artifactory API IP and port.
var client *ArtifactoryClient

// ArtifactoryClient is the Jfrog artifactory client
type ArtifactoryClient struct {
	request   *http.Request
	hostIP    string
	port      string
	scheme    string
	basicAuth bool
}

func NewArtifactoryClient() *ArtifactoryClient {
	return &ArtifactoryClient{
		basicAuth: true, // set basic auth true by default
	}
}

func (ac *ArtifactoryClient) SetIP(ip string) *ArtifactoryClient {
	ac.hostIP = ip
	return ac
}

func (ac *ArtifactoryClient) SetPort(port string) *ArtifactoryClient {
	ac.port = port
	return ac
}

func (ac *ArtifactoryClient) SetHttpScheme() *ArtifactoryClient {
	ac.scheme = "http"
	return ac
}

// buildURL builds the API url to execute AQL
func (ac *ArtifactoryClient) buildURL(uri string) string {
	return ac.scheme + "://" + ac.hostIP + ":" + ac.port + "/" + uri
}

// MakeRequest makes an http request bases on the request type set on client.
func (ac *ArtifactoryClient) MakeRequest() ([]byte, error) {
	secret := os.Getenv("ART_SECRET")
	if secret == "" {
		return nil, errors.New("failed to get secret.")
	}
	if ac.basicAuth {
		ac.request.Header.Add("authorization", "Basic "+secret)
	}
	res, _ := http.DefaultClient.Do(ac.request)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, errors.New("failed to make request")
	}
	return body, nil
}
