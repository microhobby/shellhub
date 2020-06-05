package models

type FirewallRule struct {
	Action   string `json:"action"`
	Active   bool   `json:"active"`
	SourceIP string `json:"source_ip"`
	User     string `json:"user"`
	Hostname string `json:"hostname"`
	TenantID string `json:"tenant_id"`
}
