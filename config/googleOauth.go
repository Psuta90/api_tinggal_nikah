package config

import (
	"encoding/json"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Source struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Metadata struct {
	Primary       bool   `json:"primary"`
	Source        Source `json:"source"`
	SourcePrimary bool   `json:"sourcePrimary"`
	Verified      bool   `json:"verified"`
}

type EmailAddress struct {
	Metadata Metadata `json:"metadata"`
	Value    string   `json:"value"`
}

type Name struct {
	DisplayName          string   `json:"displayName"`
	DisplayNameLastFirst string   `json:"displayNameLastFirst"`
	GivenName            string   `json:"givenName"`
	Metadata             Metadata `json:"metadata"`
	UnstructuredName     string   `json:"unstructuredName"`
}

type UserInfo struct {
	EmailAddresses []EmailAddress `json:"emailAddresses"`
	Etag           string         `json:"etag"`
	Names          []Name         `json:"names"`
	ResourceName   string         `json:"resourceName"`
}

var GoogleOauthConfig = oauth2.Config{
	ClientID:     os.Getenv("ClientID"),
	ClientSecret: os.Getenv("ClientSecret"),
	RedirectURL:  os.Getenv("RedirectURL"),
	Scopes:       []string{"openid", "profile", "email"},
	Endpoint:     google.Endpoint,
}

func GetUserInfo(client *http.Client) (*UserInfo, error) {
	// Make a GET request to Google People API
	resp, err := client.Get("https://people.googleapis.com/v1/people/me?personFields=names,emailAddresses")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
