package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	HttpClient *http.Client // nolint:golint
	Host       string
	Endpoint   string
	Port       string
	Token      string

	Headers http.Header
	Method  string
}

func NewClient(host, endpoint, port, token string) (*Client, error) {
	if isURL(host) {
		return &Client{
			HttpClient: &http.Client{Timeout: 30 * time.Second},
			Host:       host,
			Endpoint:   endpoint,
			Port:       port,
			Token:      token,
			Headers:    http.Header{},
		}, nil
	}

	return nil, fmt.Errorf("host (%v) is not a valid url", host)
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (c *Client) SetHeaders(headers map[string]string) *Client {
	for k, v := range headers {
		c.Headers.Set(k, v)
	}

	return c
}

func (c *Client) SetMethod(method string) *Client {
	c.Method = method

	return c
}

func buildURL(host, port, endpoint, token string) string {
	u := strings.TrimSuffix(host, "/")
	if port != "" {
		u = u + "/" + port
	}
	return u + "/" + endpoint + "?circle-token=" + token
}

func (c *Client) prepareRequest() *http.Request {
	rawURL := buildURL(c.Host, c.Port, c.Endpoint, c.Token)
	u, err := url.Parse(rawURL)
	if err != nil {
		log.Fatalf("invalid URL (%v): %v", c.Host+"/"+c.Endpoint, err)
	}

	return &http.Request{
		Method: c.Method,
		URL:    u,
		Header: c.Headers,
	}
}

func (c *Client) Run(obj interface{}) error {
	resp, err := c.HttpClient.Do(c.prepareRequest())
	if err != nil {
		u := buildURL(c.Host, c.Port, c.Endpoint, "")
		return fmt.Errorf("error completing request (url: %v): %v", u, err)
	}

	defer func() {
		responseBodyCloseErr := resp.Body.Close()
		if responseBodyCloseErr != nil {
			log.Printf(responseBodyCloseErr.Error())
		}
	}()

	if err := json.NewDecoder(resp.Body).Decode(obj); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	return nil
}
