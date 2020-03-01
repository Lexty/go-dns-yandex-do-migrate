package api

// DNSRecord foo
type DNSRecord struct {
	RecordID  int `json:"record_id"`
	Type      string `json:"type"`
	Domain    string `json:"domain"`
	Fqdn      string `json:"fqdn"`
	TTL       int `json:"ttl"`
	Subdomain string `json:"subdomain"`
	Content   string `json:"content"`
	Priority  interface{} `json:"priority"`
}