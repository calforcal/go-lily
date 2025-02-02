package google

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/calforcal/can-lily-eat-it/config"
	"github.com/calforcal/can-lily-eat-it/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleService struct {
	conf *oauth2.Config
}

func NewGoogleService() *GoogleService {
	return &GoogleService{
		conf: &oauth2.Config{
			ClientID:     config.GOOGLE_CLIENT_ID,
			ClientSecret: config.GOOGLE_CLIENT_SECRET,
			RedirectURL:  config.GOOGLE_REDIRECT_URL,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}

func (s *GoogleService) GetAuthURL() string {
	url := s.conf.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return url
}

func (s *GoogleService) GetToken(code string) (*oauth2.Token, error) {
	token, err := s.conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %v", err)
	}

	return token, nil
}

func (s *GoogleService) GetUserInfo(token *oauth2.Token) (*models.GoogleUserInfo, error) {
	client := s.conf.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %v", err)
	}

	var userInfo models.GoogleUserInfo
	if err := json.Unmarshal(data, &userInfo); err != nil {
		return nil, fmt.Errorf("failed parsing user info: %v", err)
	}

	return &userInfo, nil
}
