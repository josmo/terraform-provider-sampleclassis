package classis

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/hashicorp/go-uuid"
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

func (c *Client) CreateSpotGroup(spotGroup SpotGroup) (string, error) {
	var a [2]interface{}
	generatedUID, err := uuid.GenerateUUID()
	a[0] = generatedUID
	a[1] = spotGroup
	spotBytes, _ := json.Marshal(a)
	spotReader := bytes.NewReader(spotBytes)

	req, err := http.NewRequest("POST", c.base+"/methods/sgUpsert", spotReader)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)
	response, err := c.client.Do(req)


	defer response.Body.Close()

	return generatedUID, err
}


func(c *Client) DeleteSpotGroup(groupId string) error {
	req, err := http.NewRequest("DELETE", c.base+"/spot-groups/"+groupId, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)
	response, deleteError := c.client.Do(req)
	if response.StatusCode == 400 {
		return errors.New("Sorry but the group must be stopped before it's removed")
	}
	if response.StatusCode == 401 {
		return errors.New("Sorry this is not your spot group")
	}
	return deleteError
}


