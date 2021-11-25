package model

type Report struct {
	EngagementType       string `json:"source_engagement_type"`
	SourceSite           string `json:"source_site"`
	SourceId             int    `json:"source_id"`
	AttributedOnSite     string `json:"attributed_on_site"`
	TriggerData          int    `json:"trigger_data"`
	PcmVersion           int    `json:"version"`
	SecretToken          string `json:"source_secret_token"`
	SecretTokenSignature string `json:"source_secret_token_signature"`
}
