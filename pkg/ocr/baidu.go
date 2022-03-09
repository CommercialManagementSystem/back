package ocr

import (
	"encoding/json"
	"fmt"
	"github.com/CommercialManagementSystem/back/internal/config"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	POST = "POST"
	GET  = "GET"
)

type BaiDu struct {
	Token string
}

func GetToken() (*BaiDu, error) {
	c := config.C.OCR
	url := c.TokenUrl()

	client := &http.Client{}
	req, err := http.NewRequest(POST, url, nil)

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var token *TokenResponse
	err = json.Unmarshal(body, token)
	if err != nil {
		return nil, err
	}
	return &BaiDu{Token: token.AccessToken}, err
}

func (b *BaiDu) AccurateBasic(file string) (*AccurateBasicResponse, error) {
	url := fmt.Sprintf(
		"https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic?access_token=%s",
		b.Token,
	)

	payload := strings.NewReader(
		fmt.Sprintf(
			"url=%s",
			file,
		),
	)

	client := &http.Client{}
	req, err := http.NewRequest(POST, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var results *AccurateBasicResponse
	err = json.Unmarshal(body, results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
