package connector

import (
	"bytes"
	"context"
	"github.com/goccy/go-json"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	"io"
	"net/http"
	"time"
)

type IHTMLRetrieverConnector interface {
	GetRenderedHTML(url string) (*string, error)
}

var HTMLRetrieverConnectorInstance IHTMLRetrieverConnector

type Payload struct {
	URL string `json:"url"`
}

type CloudRunConnector struct {
	CloudRunUrl string
	KeyPath     string
	Client      *http.Client
}

func InitCloudRunConnector(cloudRunUrl string, keyPath string) error {
	instance, err := NewCloudRunConnector(cloudRunUrl, keyPath)
	if err != nil {
		return err
	}
	HTMLRetrieverConnectorInstance = instance
	return nil
}

func NewCloudRunConnector(cloudRunUrl string, keyPath string) (*CloudRunConnector, error) {
	ctx := context.Background()
	client, err := idtoken.NewClient(ctx, cloudRunUrl, option.WithCredentialsFile(keyPath))
	if err != nil {
		return nil, err
	}
	return &CloudRunConnector{CloudRunUrl: cloudRunUrl, KeyPath: keyPath, Client: client}, nil
}

func (c *CloudRunConnector) GetRenderedHTML(url string) (*string, error) {
	data := Payload{
		URL: url,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Post(c.CloudRunUrl+"/request-html", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(body)
	return &bodyString, err
}

type StandaloneConnector struct {
	StandaloneUrl string
	Client        *http.Client
}

func InitStandaloneConnector(standaloneUrl string) error {
	instance, err := NewStandaloneConnector(standaloneUrl)
	if err != nil {
		return err
	}
	HTMLRetrieverConnectorInstance = instance
	return nil
}

func NewStandaloneConnector(standaloneUrl string) (*StandaloneConnector, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	return &StandaloneConnector{StandaloneUrl: standaloneUrl, Client: client}, nil
}

func (c *StandaloneConnector) GetRenderedHTML(url string) (*string, error) {
	data := Payload{
		URL: url,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Post(c.StandaloneUrl+"/request-html", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(body)
	return &bodyString, err
}
