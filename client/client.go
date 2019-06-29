package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	HttpClient *http.Client // nolint:golint
	Host       string
	Endpoint   string
	Port       string
	Token      string

	Headers http.Header
	Method  string
	Query   map[string]interface{}
}

func NewClient(host, port, endpoint, token string) (*Client, error) {
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

func (c *Client) AddQuery(key string, value interface{}) *Client {
	if c.Query == nil {
		c.Query = make(map[string]interface{})
	}
	c.Query[key] = value

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
		log.WithField("method", "prepareRequest()").
			Fatalf("invalid URL (%v): %v", rawURL, err)
	}

	request := &http.Request{
		Method: c.Method,
		URL:    u,
		Header: c.Headers,
	}

	if c.Query != nil {
		query := request.URL.Query()
		for k, v := range c.Query {
			query.Add(k, fmt.Sprintf("%v", v))
		}
		request.URL.RawQuery = query.Encode()
	}

	return request
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
