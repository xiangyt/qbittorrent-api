package qbittorrent_api

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/publicsuffix"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
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

func (r *Request) get(name APIName, path string, params map[string]string) (*http.Response, error) {
	urlStr, err := url.JoinPath(r.host, r.basePath, name, path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate url")
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}

	// add user-agent header to allow qbittorrent to identify us
	req.Header.Set("User-Agent", "go-qbittorrent "+GoQBitVersion)

	// add optional parameters that the user wants
	if params != nil {
		query := req.URL.Query()
		for k, v := range params {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	resp, err := r.http.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	return resp, nil

}

func (r *Request) post(name APIName, path string, params map[string]string) (*http.Response, error) {
	urlStr, err := url.JoinPath(r.host, r.basePath, name, path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate url")
	}

	form := url.Values{}
	if params != nil {
		for k, v := range params {
			form.Add(k, v)
		}
	}

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}

	// add the content-type so qbittorrent knows what to expect
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// add user-agent header to allow qbittorrent to identify us
	req.Header.Set("User-Agent", "go-qbittorrent "+GoQBitVersion)
	//req.Header.Set("X-Requested-With", "XMLHttpRequest")
	//req.Header.Set("Host", "nas.zerotier.xyt:8999")
	//req.Header.Set("Origin", "http://nas.zerotier.xyt:8999")

	// add optional parameters that the user wants

	resp, err := r.http.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	return resp, nil
}

func (r *Request) postMultipart(urlStr string, buffer bytes.Buffer, contentType string) (*http.Response, error) {

	req, err := http.NewRequest("POST", urlStr, &buffer)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}

	// add the content-type so qbittorrent knows what to expect
	req.Header.Set("Content-Type", contentType)
	// add user-agent header to allow qbittorrent to identify us
	req.Header.Set("User-Agent", "go-qbittorrent "+GoQBitVersion)

	resp, err := r.http.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	return resp, nil
}

func (r *Request) postMultipartData(name APIName, path string, params map[string]string) (*http.Response, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	for key, val := range params {
		writer.WriteField(key, val)
	}

	if err := writer.Close(); err != nil {
		return nil, errors.Wrap(err, "failed to close writer")
	}

	urlStr, err := url.JoinPath(r.host, r.basePath, name, path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate url")
	}

	return r.postMultipart(urlStr, buffer, writer.FormDataContentType())
}

func (r *Request) postMultipartFile(name APIName, urlPath, fileName string, params map[string]string) (*http.Response, error) {
	//var buffer bytes.Buffer
	//writer := multipart.NewWriter(&buffer)
	//
	//// open the file for reading
	//file, err := os.Open(fileName)
	//if err != nil {
	//	return nil, errors.Wrap(err, "error opening file")
	//}
	//
	//// create form for writing the file to and give it the filename
	//formWriter, err := writer.CreateFormFile("torrents", path.Base(fileName))
	//if err != nil {
	//	return nil, errors.Wrap(err, "error adding file")
	//}
	//
	//for key, val := range params {
	//	writer.WriteField(key, val)
	//}
	//
	//// copy the file contents into the form
	//if _, err = io.Copy(formWriter, file); err != nil {
	//	return nil, errors.Wrap(err, "error copying file")
	//}
	//
	//if err := writer.Close(); err != nil {
	//	return nil, errors.Wrap(err, "failed to close writer")
	//}
	//
	//urlStr, err := url.JoinPath(r.host, r.basePath, name, urlPath)
	//if err != nil {
	//	return nil, errors.Wrap(err, "failed to generate url")
	//}
	//
	//return r.postMultipart(urlStr, buffer, writer.FormDataContentType())
	return nil, errors.New("not implement")
}
