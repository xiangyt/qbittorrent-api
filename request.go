package qbittorrent_api

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/xiangyt/qbittorrent-api/definition"
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Request struct {
	Jar      http.CookieJar
	http     *http.Client
	host     string
	basePath string
}

func (r *Request) initialize() {
	logrus.Debug("initializing Request...")
	// base path for all API endpoints
	r.basePath = "api/v2"
	r.Jar, _ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	r.http = &http.Client{
		Jar: r.Jar,
	}
}

func (r *Request) get(name definition.APIName, path string) (*http.Response, error) {
	urlStr, err := url.JoinPath(r.host, r.basePath, name, path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate url")
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}

	// add user-agent header to allow qbittorrent to identify us
	req.Header.Set("User-Agent", "go-qbittorrent v0.1")

	// add optional parameters that the user wants
	//if opts != nil {
	//	query := req.URL.Query()
	//	for k, v := range opts {
	//		query.Add(k, v)
	//	}
	//	req.URL.RawQuery = query.Encode()
	//}

	resp, err := r.http.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	return resp, nil

}

func (r *Request) post(name definition.APIName, path string, params map[string]string) (*http.Response, error) {
	urlStr, err := url.JoinPath(r.host, r.basePath, name, path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate url")
	}
	req, err := http.NewRequest("POST", urlStr, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}

	// add the content-type so qbittorrent knows what to expect
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// add user-agent header to allow qbittorrent to identify us
	req.Header.Set("User-Agent", "go-qbittorrent v0.1")

	// add optional parameters that the user wants
	if params != nil {
		form := url.Values{}
		for k, v := range params {
			form.Add(k, v)
		}
		req.PostForm = form
	}

	resp, err := r.http.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	return resp, nil
}

func (r *Request) request(method, name, path string) (*http.Response, error) {

	return nil, nil
}
