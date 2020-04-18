package api

type Nameservers []string

type ZoneFile string

type Domain struct {
	DomainHref         string `json:"domain_href"`
	DomainRecordsHref  string `json:"domain__records_href"`
	AutomaticSnapshots bool   `json:"automatic_snapshots"`
	Fqdn               string `json:"fqdn"`
}

type Record struct {
	RrsetHref   string   `json:"rrset_href"`
	RrsetName   string   `json:"rrset_name"`
	RrsetType   string   `json:"rrset_type"`
	RrsetValues []string `json:"rrset_values"`
	RrsetTtl    int      `json:"rrset_ttl"`
}
