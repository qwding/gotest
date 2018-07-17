package main

import (
	"net/http"
)

func RequestWithFormData(code string) (*GitlabAccessToken, error) {
	u, err := url.Parse("/oauth/token")
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	query := u.Query()
	query.Add("client_id", GetGitlabClientId())
	query.Add("client_secret", GetGitlabSercetId())
	query.Add("code", code)
	query.Add("grant_type", "authorization_code")
	query.Add("redirect_uri", GetCallBackUrl())

	q := query.Encode()
	uri := u.Path + "?" + q

	targetUrl := beego.AppConfig.String("Gitlab") + uri
	beego.Debug("GetGitlabAccessToken url:", targetUrl, "gitlab:", beego.AppConfig.String("Gitlab"))

	resp, err := http.Post(targetUrl, "application/json", nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	accToken := &GitlabAccessToken{}
	err = json.Unmarshal(body, &accToken)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	//fmt.Printf("body is %#v\n", string(body))

	return accToken, nil
}
 