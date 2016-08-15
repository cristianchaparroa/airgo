package net

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"net/url"
	"strings"
)

type IHttpClient interface {
	Get(url string) ([]byte, error)
}

type HttpClient struct {

}

func (h *HttpClient) Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	//fmt.Println("###-----###----###-####")
	//s := string(b[:])
	//fmt.Println(s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (h *HttpClient) Post(url string, params url.Values, headers http.Header) ([]byte, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	fmt.Println("Currently headers ", req.Header)
	req.Header = headers

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println("###-----###----###-####")
	s := string(b[:])
	fmt.Println(s)

	if err != nil {
		return nil, err
	}
	return b, nil
}

