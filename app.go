package qbittorrent_api

import (
	"io"
)

type applicationApi struct {
	client *Client
}
type applicationAction = string

const (
	actionVersion applicationAction = "version"
)

func (a *authorization) Version() (string, error) {
	resp, err := a.get(apiNameApplication, actionVersion, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(body))
	//fmt.Println(resp.StatusCode)
	if err := handleResponsesErr(resp.StatusCode); err != nil {
		return "", err
	}
	return string(body), nil
}
