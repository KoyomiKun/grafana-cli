package client

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/KoyomiKun/grafana-cli/utils/log"
)

type Client struct {
	baseUrl    string
	client     *http.Client
	authString string
}

func NewClient(
	timeout time.Duration,
	baseUrl string,
	apiKey string) *Client {

	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
		baseUrl:    baseUrl,
		authString: "Bearer " + apiKey,
	}
}

func (c *Client) GetApi(
	api string,
	headers map[string]string,
	querys map[string]string) ([]byte, error) {

	url := c.baseUrl + api

	// init request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("Fail creating request %s: %v\n", url, err)
		return nil, err
	}

	// init header
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Authorization", c.authString)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// init querys
	query := req.URL.Query()
	for k, v := range querys {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()

	// do request
	resp, err := c.client.Do(req)
	if err != nil {
		log.Errorf("Fail doing request %s: %v\n", url, err)
		return nil, err
	}
	defer resp.Body.Close()
	log.Infof("Request %s\n", url)

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Fail read response %s: %v\n", url, err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Errorf("Fail requset %s: code = %d, msg = %s", url, resp.StatusCode, ret)
	}
	return ret, nil
}

func (c *Client) PostApi(
	api string,
	headers map[string]string,
	body io.Reader) ([]byte, error) {

	url := c.baseUrl + api

	// init request
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		log.Errorf("Fail creating request %s: %v\n", url, err)
		return nil, err
	}

	// init header
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Authorization", c.authString)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// do request
	resp, err := c.client.Do(req)
	if err != nil {
		log.Errorf("Fail doing request %s: %v\n", url, err)
		return nil, err
	}
	defer resp.Body.Close()
	log.Infof("Request %s\n", url)

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Fail read response %s: %v\n", url, err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Errorf("Fail requset %s: code = %d, msg = %s", url, resp.StatusCode, ret)
	}
	return ret, nil
}
