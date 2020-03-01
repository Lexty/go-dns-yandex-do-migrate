package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// DOAPICreateRecordRequest foo
type DOAPICreateRecordRequest struct {
	ID       *int        `json:"id"`
	Type     string      `json:"type"`
	Name     string      `json:"name"`
	Data     string      `json:"data"`
	Priority interface{} `json:"priority"`
	Port     *int        `json:"port"`
	TTL      int         `json:"ttl"`
	Weight   *int        `json:"weight"`
	Flags    *string     `json:"flags"`
	Tag      *string     `json:"tag"`
}

// DOAPICreateRecordResponse foo
type DOAPICreateRecordResponse struct {
	DomainRecord DOAPICreateRecordRequest `json:"domain_record"`
}

// {"domain_record":{"id":91349395,"type":"TXT","name":"foo","data":"162.10.66.0","priority":null,"port":null,"ttl":21600,"weight":null,"flags":null,"tag":null}}

// DOClient foo
type DOClient struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

func (c *DOClient) setDefaults() {
	if c.UserAgent == "" {
		c.UserAgent = "DNS Yandex-DO Migrate Tool"
	}
}

// CreateRecord Create a new Domain Record
func (c *DOClient) CreateRecord(data DOAPICreateRecordRequest, token, domain string) (*DOAPICreateRecordRequest, error) {
	c.setDefaults()
	client := &http.Client{}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s\n", jsonStr)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://api.digitalocean.com/v2/domains/%s/records", domain), bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
		return nil, nil
	} else {
		defer resp.Body.Close()
		var v DOAPICreateRecordResponse
		err = json.NewDecoder(resp.Body).Decode(&v)
		return &v.DomainRecord, err
	}
}
