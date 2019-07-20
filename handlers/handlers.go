package handlers


import (
	"encoding/json"
	"fmt"
	"net/http"
	"oauth_demo/config"
	"golang.org/x/oauth2"
	"io/ioutil"
)
func Index (w http.ResponseWriter, r *http.Request) {
	var page = `
	<html>
		<body>
			<button onclick="window.location='/login';"> Ugly button to login with Gitlab </button>
		</body>
	</html>
	`
	fmt.Fprintf(w, page)
}

func Login (w http.ResponseWriter, r *http.Request) {
	url := config.GitlabOauthConfig.AuthCodeURL (config.State)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func Error (w http.ResponseWriter, r *http.Request) {
	var page = `
	<html>
		<body>
			<h1>WHOOOPS!</h1><button onclick="window.location='/';"> Volver al inicio </button>
		</body>
	</html>	
	`
	fmt.Fprintf(w, page)
}

func Callback (w http.ResponseWriter, r *http.Request) {
	encodedInfo, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
	}
	var userInfo map[string]interface{}
	json.Unmarshal(encodedInfo, &userInfo)
	var page = 
	`<html>
		<body>
			<h1> User information presented in a ugly way </h1>
			<img style="max-width:300px;" src="%s"> <br>
			<h2>Username: %s <br><h2>
			Name: %s <br>
			email: %s <br>
		<body>
	</html>`
	fmt.Fprintf(w, page, userInfo["avatar_url"], userInfo["username"], userInfo["name"],
	userInfo["email"])
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != config.State {
		return nil, fmt.Errorf("Bad State")
	}
	token, err := config.GitlabOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("Code failed: %s", err)
	}
	response, err := http.Get("https://gitlab.com/api/v4/user?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Error getting user info:")
	}
	defer response.Body.Close()
	userInfo, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("REsponse error")
	}
	return userInfo, nil
}