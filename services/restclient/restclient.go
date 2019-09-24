package restclient

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type RestClient struct {

}

func (r *RestClient) Get(url string, bodyRequest io.Reader) map[string]string {
	req, err := http.NewRequest(http.MethodGet, url, bodyRequest)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	m := make(map[string]string)
	err = json.Unmarshal(body, &m)

	return m
}
