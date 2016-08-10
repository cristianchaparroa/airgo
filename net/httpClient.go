package net

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

type IHttpClient interface {
	Get(url string) ([]byte, error)
}

type HttpClient struct {

}

func (h*HttpClient) Get(url string) ([]byte, error) {
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

	fmt.Println("###-----###----###-####")
	s := string(b[:])
	fmt.Println(s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

