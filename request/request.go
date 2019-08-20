package request

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//
const Timeout = 10 * time.Second

//
const DefaultContentType = "application/x-www-form-urlencoded"

// Get
func Get(url string) (body []byte, resp *http.Response, err error) {
	httpClient := &http.Client{
		Timeout:   Timeout,
		Transport: &ExportRedirectURL{},
	}
	resp, err = httpClient.Get(url)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

// Post
func Post(url, contentType, data string) (body []byte, resp *http.Response, err error) {
	if contentType == "" {
		contentType = DefaultContentType
	}
	httpClient := &http.Client{
		Timeout: Timeout,
	}
	resp, err = httpClient.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	return
}

//
type ExportRedirectURL struct {
	Transport   http.RoundTripper
	RedirectURL string
}

func (l *ExportRedirectURL) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t := l.Transport
	if t == nil {
		t = http.DefaultTransport
	}
	resp, err = t.RoundTrip(req)
	if err != nil {
		return
	}
	switch resp.StatusCode {
	case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther, http.StatusTemporaryRedirect:
		// log.Println("Request for", req.URL, "redirected with status", resp.StatusCode, resp.Header["Location"][0])
		resp.StatusCode = http.StatusOK
		return resp, nil
	}
	return
}
