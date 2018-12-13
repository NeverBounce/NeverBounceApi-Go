// Package nbModels provides the data mappings for API requests and API responses
package nbModels

// JobsDownloadRequestModel is the response model for Jobs.Download()
type JobsDownloadRequestModel struct {
	GenericRequestModel
	JobID int `json:"job_id"`

	// Segmentation options
	IncludeValids      bool `json:"valids,omitempty"`
	IncludeInvalids    bool `json:"invalids,omitempty"`
	IncludeCatchalls   bool `json:"catchalls,omitempty"`
	IncludeUnknowns    bool `json:"unknowns,omitempty"`
	IncludeDisposables bool `json:"disposables,omitempty"`
	IncludeDuplicates  bool `json:"include_duplicates,omitempty"`
	OnlyDuplicates     bool `json:"only_duplicates,omitempty"`
	OnlyBadSyntax      bool `json:"only_bad_syntax,omitempty"`

	// Data appends
	IdentifyBadSyntax      bool `json:"bad_syntax,omitempty"`
	IdentifyFreeEmailHosts bool `json:"free_email_host,omitempty"`
	IdentifyRoleAccounts   bool `json:"role_account,omitempty"`
	AppendAddr             bool `json:"addr,omitempty"`
	AppendAlias            bool `json:"alias,omitempty"`
	AppendHost             bool `json:"host,omitempty"`
	AppendSubdomain        bool `json:"subdomain,omitempty"`
	AppendDomain           bool `json:"domain,omitempty"`
	AppendTLD              bool `json:"tld,omitempty"`
	AppendFQDN             bool `json:"fqdn,omitempty"`
	AppendNetwork          bool `json:"network,omitempty"`
	HasDNSInfo             bool `json:"has_dns_info,omitempty"`
	MailServerReachable    bool `json:"mail_server_reachable,omitempty"`
	EmailStatusAsInt       bool `json:"email_status_int,omitempty"`
	EmailStatusAsString    bool `json:"email_status,omitempty"`

	// Settings for CSV data
	BinaryOperator string `json:"binary_operator_type,omitempty"`
	LineFeedType   string `json:"line_feed_type,omitempty"`
}
