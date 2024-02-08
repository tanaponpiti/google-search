package connector

import (
	"bytes"
	"context"
	"github.com/goccy/go-json"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	"io"
	"net/http"
)

type ICloudRunConnector interface {
	GetRenderedHTMLFromCloudRun(url string) (*string, error)
}

var CloudRunConnectorInstance ICloudRunConnector

type CloudRunConnector struct {
	CloudRunUrl string
	KeyPath     string
	Client      *http.Client
}

type Payload struct {
	URL string `json:"url"`
}

func InitCloudRunConnector(cloudRunUrl string, keyPath string) error {
	instance, err := NewCloudRunConnector(cloudRunUrl, keyPath)
	if err != nil {
		return err
	}
	CloudRunConnectorInstance = instance
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

func (c *CloudRunConnector) GetRenderedHTMLFromCloudRun(url string) (*string, error) {
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
