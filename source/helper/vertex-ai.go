package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2/google"
)

type Credentials struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

// VertexAIAuth handles authentication for Vertex AI services
type VertexAIAuth struct {
	ProjectID   string
	Location    string
	AccessToken string
}

// NewVertexAIAuth creates a new VertexAIAuth instance
func NewVertexAIAuth(credentialsPath string) (*VertexAIAuth, error) {
	// Read and parse credentials file
	data, err := os.ReadFile(credentialsPath)
	if err != nil {
		return nil, fmt.Errorf("error reading credentials file: %v", err)
	}

	// Parse project ID from credentials
	var creds map[string]interface{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, fmt.Errorf("error parsing credentials: %v", err)
	}

	projectID, ok := creds["project_id"].(string)
	if !ok {
		return nil, fmt.Errorf("project_id not found in credentials")
	}

	// Vertex AI specific scopes
	scopes := []string{
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/cloud-language",
		"https://www.googleapis.com/auth/dialogflow",
	}

	// Create configuration from JSON credentials
	config, err := google.JWTConfigFromJSON(data, scopes...)
	if err != nil {
		return nil, fmt.Errorf("error creating JWT config: %v", err)
	}

	// Get token
	tokenSource := config.TokenSource(context.Background())
	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("error getting token: %v", err)
	}

	return &VertexAIAuth{
		ProjectID:   projectID,
		Location:    "us-central1", // Default location, can be modified
		AccessToken: token.AccessToken,
	}, nil
}

func GetAccessToken() (string, error) {
	credentialsPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	scopes := []string{
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/cloud-language",
		"https://www.googleapis.com/auth/dialogflow",
	}
	// Read the credentials file
	data, err := os.ReadFile(credentialsPath)
	if err != nil {
		return "", fmt.Errorf("error reading credentials file: %v", err)
	}

	// Parse the credentials
	var creds Credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return "", fmt.Errorf("error parsing credentials: %v", err)
	}

	// Create configuration from JSON credentials
	config, err := google.JWTConfigFromJSON(data, scopes...)
	if err != nil {
		return "", fmt.Errorf("error creating JWT config: %v", err)
	}

	// Create a token source
	tokenSource := config.TokenSource(context.Background())

	// Get the token
	token, err := tokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("error getting token: %v", err)
	}

	return token.AccessToken, nil
}
