package mental_models

const (
	ReqDataForm   = "form"
	ReqDataRAW    = "raw"
	ReqDataHeader = "header"
)

const (
	RespDataBody   = "body"
	RespDataHeader = "header"
)

const (
	ReqPayloadURLCheckForm          = "form"
	ReqPayloadURLCheckRAWWithHeader = "raw_with_header"
	ReqPayloadURLCheckHeader        = "header"
)

const (
	RespPayloadURLCheckHeader = "header"
	RespPayloadURLCheckBody   = "body"
)

type C2Profile struct {
	//ReqProfile ReqProfile `yaml:"req_profile"`
	CoreConfig        CoreConfig        `json:"core_config" yaml:"core_config"`
	RequestConfig     RequestConfig     `json:"request_config" yaml:"request_config"`
	ResponseConfig    ResponseConfig    `json:"response_config" yaml:"response_config"`
	OverrideFormParam OverrideFormParam `json:"override_form_param" yaml:"override_form_param"`
}

type OverrideFormParam struct {
	Enable         bool   `json:"enable" yaml:"enable"`
	ConnPassword   string `json:"conn_password" yaml:"conn_password"`
	ReqHeader      string `json:"req_header" yaml:"req_header"`
	RespStatusCode string `json:"resp_status_code" yaml:"resp_status_code"`
	//RespCharset     string `json:"resp_charset" yaml:"resp_charset"`
	RespHeader     string `json:"resp_header" yaml:"resp_header"`
	RespBodyFormat string `json:"resp_body_format" yaml:"resp_body_format"`
}

type CoreConfig struct {
	SupportCodeType      string `json:"support_code_type" yaml:"support_code_type"`
	SupportEncryptMethod string `json:"support_encrypt_method" yaml:"support_encrypt_method"`

	ReqEncode  string `json:"req_encode" yaml:"req_encode"`
	RespEncode string `json:"resp_encode" yaml:"resp_encode"`

	ConnMaxRetry            int `json:"conn_max_retry" yaml:"conn_max_retry"`
	UploadFileSectionSize   int `json:"upload_file_section_size" yaml:"upload_file_section_size"`     // B
	DownloadFileSectionSize int `json:"download_file_section_size" yaml:"download_file_section_size"` // B
	FileTransMaxRetry       int `json:"file_trans_max_retry" yaml:"file_trans_max_retry"`
	PluginLoadMaxRetry      int `json:"plugin_load_max_retry" yaml:"plugin_load_max_retry"`

	HeartBeatEnabled   bool    `json:"heart_beat_enabled" yaml:"heart_beat_enabled"`
	HeartBeatMaxRetry  int     `json:"heart_beat_max_retry" yaml:"heart_beat_max_retry"`
	HeartBeatJitter    float32 `json:"heart_beat_jitter" yaml:"heart_beat_jitter"`
	HeartBeatSleepTime int     `json:"heart_beat_sleep_time" yaml:"heart_beat_sleep_time"`

	SessionId string `json:"session_id" yaml:"session_id"`
	Debug     bool   `json:"debug" yaml:"debug"`
}

type RequestConfig struct {
	UserAgents []string `json:"user_agents" yaml:"user_agents"`

	RequestPayloadURLConfig RequestPayloadURLConfig `json:"request_payload_url_config" yaml:"request_payload_url_config"`
	RequestDataConfig       RequestDataConfig       `json:"request_data_config" yaml:"request_data_config"`

	SectionEnabled       bool `json:"section_enabled" yaml:"section_enabled"`
	SectionTransMaxRetry int  `json:"section_trans_max_retry" yaml:"section_trans_max_retry"`
	SectionPerDataLength int  `json:"section_per_data_length" yaml:"section_per_data_length"`

	RequestAuth RequestAuth `json:"request_auth" yaml:"request_auth"`
	Chunked     bool        `json:"chunked" yaml:"chunked"`
}

type RequestAuth struct {
	AuthEnable bool   `json:"auth_enable" yaml:"auth_enable"`
	AuthType   string `json:"auth_type" yaml:"auth_type"` // Basic/NTLMv1/NTLMv2
	Username   string `json:"username" yaml:"username"`
	Password   string `json:"password" yaml:"password"`
	Domain     string `json:"domain" yaml:"domain"`
}

type RequestPayloadURLConfig struct {
	Method string `json:"method" yaml:"method"`

	PayloadURLContentType   string `json:"payload_url_content_type" yaml:"payload_url_content_type"`
	PayloadURLCheckPosition string `json:"payload_url_check_position" yaml:"payload_url_check_position"`
	PayloadURLCheckParam    string `json:"payload_url_check_param" yaml:"payload_url_check_param"`
}

type RequestDataConfig struct {
	Method      string `json:"method" yaml:"method"`
	ContentType string `json:"content_type" yaml:"content_type"`

	DataMode   string `json:"data_mode" yaml:"data_mode"`
	DataFormat string `json:"data_format" yaml:"data_format"`

	DataPrefix string `json:"data_prefix" yaml:"data_prefix"`
	DataSuffix string `json:"data_suffix" yaml:"data_suffix"`

	DataSplitLeft  string `json:"data_split_left" yaml:"data_split_left"`
	DataSplitRight string `json:"data_split_right" yaml:"data_split_right"`

	PayloadPrefix string `json:"payload_prefix" yaml:"payload_prefix"`
	PayloadSuffix string `json:"payload_suffix" yaml:"payload_suffix"`

	ChildDataPrefix string `json:"child_data_prefix" yaml:"child_data_prefix"`
	ChildDataSuffix string `json:"child_data_suffix" yaml:"child_data_suffix"`

	ChildPayloadPrefix string `json:"child_payload_prefix" yaml:"child_payload_prefix"`
	ChildPayloadSuffix string `json:"child_payload_suffix" yaml:"child_payload_suffix"`

	RandomParamCount string `json:"random_param_count" yaml:"random_param_count"`
	RandomParamName  string `json:"random_param_name" yaml:"random_param_name"`
	RandomParamValue string `json:"random_param_value" yaml:"random_param_value"`
}

type ResponseConfig struct {
	ResponsePayloadURLConfig ResponsePayloadURLConfig `json:"response_payload_url_config" yaml:"response_payload_url_config"`
	ResponseDataConfig       ResponseDataConfig       `json:"response_data_config" yaml:"response_data_config"`

	SectionEnabled       bool `json:"section_enabled" yaml:"section_enabled"`
	SectionTransMaxRetry int  `json:"section_trans_max_retry" yaml:"section_trans_max_retry"`
	SectionPerDataLength int  `json:"section_per_data_length" yaml:"section_per_data_length"`
}

type ResponsePayloadURLConfig struct {
	PayloadURLCheckPosition string `json:"payload_url_check_position" yaml:"payload_url_check_position"`
	PayloadURLCheckParam    string `json:"payload_url_check_param" yaml:"payload_url_check_param"`
}
type ResponseDataConfig struct {
	DataMode string `json:"data_mode" yaml:"data_mode"`

	DataHeaderFormat string `json:"data_header_format" yaml:"data_header_format"` // {{data}}

	DataPrefixRange string `json:"data_prefix_range" yaml:"data_prefix_range"`
	DataSuffixRange string `json:"data_suffix_range" yaml:"data_suffix_range"`

	//DataSplitLeft  string `json:"data_split_left" yaml:"data_split_left"`
	//DataSplitRight string `json:"data_split_right" yaml:"data_split_right"`

	PayloadPrefixRange string `json:"payload_prefix_range" yaml:"payload_prefix_range"`
	PayloadSuffixRange string `json:"payload_suffix_range" yaml:"payload_suffix_range"`
}

type ReqProfile struct {
	ProfileName     string        `json:"profile_name" yaml:"profile_name"`
	ReqHeader       string        `json:"req_header" yaml:"req_header"`
	RespStatusCode  string        `json:"resp_status_code" yaml:"resp_status_code"`
	RespCharset     string        `json:"resp_charset" yaml:"resp_charset"`
	RespHeader      string        `json:"resp_header" yaml:"resp_header"`
	RespBodyFormat  string        `json:"resp_body_format" yaml:"resp_body_format"`
	RespBodyDynamic func() string `json:"-" yaml:"-"`
}

type ReqProfileNamesResp struct {
	Profiles []string `json:"profiles"`
	Status   bool     `json:"status"`
	Err      string   `json:"err"`
}

type ReqProfileResp struct {
	Profile ReqProfile `json:"profile"`
	Status  bool       `json:"status"`
	Err     string     `json:"err"`
}

type C2ProfileNamesResp struct {
	Profiles []string `json:"profiles"`
	Status   bool     `json:"status"`
	Err      string   `json:"err"`
}

type C2ProfileSaveReq struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type C2ProfileContentResp struct {
	Content string `json:"content"`
	Status  bool   `json:"status"`
	Err     string `json:"err"`
}

type C2ProfileGenerateReq struct {
	RequestData  string `json:"request_data"`
	ResponseData string `json:"response_data"`
}
