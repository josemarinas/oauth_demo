package config

import (
	"os"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/gitlab"
	
)
var GitlabOauthConfig *oauth2.Config
var State = "Totally_Random_And_CSRF_safe_String"
func Init() {
	GitlabOauthConfig = &oauth2.Config{
		RedirectURL: 	"http://localhost:3000/callback",
		ClientID: 		os.Getenv("GITLAB_CLIENT_ID"),
		ClientSecret: 	os.Getenv("GITLAB_CLIENT_SECRET"),
		Scopes:			[]string{"read_user"},
		Endpoint:		gitlab.Endpoint,
	}
}