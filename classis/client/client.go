package classis

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
	base   string
	token  string
}

func NewClientWith(url string, username string, password string) (*Client, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	login := Login{
		EmailAddress: username,
		Password:     password,
	}
	loginBytes, err := json.Marshal(login)
	if err != nil {
		return nil, err
	}
	loginReader := bytes.NewReader(loginBytes)

	response, err := netClient.Post(url+"/users/login", "application/json", loginReader)
	if response.StatusCode == 400 {
		return nil, errors.New("Sorry but the log in info is incorrect")
	}
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	loginResponse := LoginResponse{}
	err = json.Unmarshal(contents, &loginResponse)
	if err != nil {
		return nil, err
	}
	return &Client{netClient, url, loginResponse.Token}, nil
}
