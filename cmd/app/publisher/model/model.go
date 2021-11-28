package model

type Report struct {
	EngagementType       string `json:"source_engagement_type"`
	SourceSite           string `json:"source_site"`
	SourceId             int    `json:"source_id"`
	AttributedOnSite     string `json:"attributed_on_site"`
	TriggerData          int    `json:"trigger_data"`
	Version              int    `json:"version"`
	SecretToken          string `json:"source_secret_token"`
	SecretTokenSignature string `json:"source_secret_token_signature"`
}

type Sign struct {
	EngagementType string `json:"source_engagement_type"`
	Nonce          string `json:"source_nonce"`
	SourceToken    string `json:"source_unlinkable_token"`
	Version        int    `json:"version"`
}
