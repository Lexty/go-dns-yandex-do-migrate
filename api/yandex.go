package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// YandexClient foo
type YandexClient struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

func (c *YandexClient) setDefaults() {
	if c.UserAgent == "" {
		c.UserAgent = "DNS Yandex-DO Migrate Tool"
	}
}

// ListRecords foo
func (c *YandexClient) ListRecords(token string, domain string) ([]DNSRecord, error) {
	c.setDefaults()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://pddimp.yandex.ru/api2/admin/dns/list?domain=%s", domain), nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("PddToken", token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var v YandexListRecordsResponse
	err = json.NewDecoder(resp.Body).Decode(&v)
	return v.Records, err

	// req, err := c.newRequest("GET", fmt.Sprintf("https://pddimp.yandex.ru/api2/admin/dns/list?domain=%s", domain), nil, token)
	// if err != nil {
	// 	return nil, err
	// }
	// var response YandexListRecordsResponse
	// _, err = c.do(req, &response)
	// return response.Records, err
}
func (c *YandexClient) newRequest(method, path string, body interface{}, token string) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("PddToken", token)
	return req, nil
}
func (c *YandexClient) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

// YandexListRecordsResponse foo
type YandexListRecordsResponse struct {
	Domain  string      `json:"domain"`
	Records []DNSRecord `json:"records"`
	Success string      `json:"success"`
}
