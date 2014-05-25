package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/menu"
	"io/ioutil"
	"net/http"
)

func (c *Client) MenuCreate(mn *menu.Menu) error {
	if mn == nil {
		return errors.New("menu == nil")
	}
	jsonData, err := json.Marshal(mn)
	if err != nil {
		return err
	}

	token, err := c.Token()
	if err != nil {
		return err
	}

	url := fmt.Sprintf(menuCreateUrlFormat, token)
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result Error
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return &result
	}
	return nil
}

func (c *Client) MenuDelete() error {
	token, err := c.Token()
	if err != nil {
		return err
	}

	url := fmt.Sprintf(menuDeleteUrlFormat, token)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result Error
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return &result
	}
	return nil
}

func (c *Client) MenuGet() (*menu.GetMenuResponse, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(menuGetUrlFormat, token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		menu.GetMenuResponse
		Error
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.GetMenuResponse, nil
}
