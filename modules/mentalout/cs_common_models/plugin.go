package cs_common_models

type MentalPlugin struct {
	ID          int    `yaml:"-" json:"id"`
	CodeType    string `yaml:"code_type" json:"code_type"` // JAVA/PHP/CSHARP
	PluginName  string `yaml:"plugin_name" json:"plugin_name"`
	ScriptName  string `yaml:"script_name" json:"script_name"`
	PluginAlias string `yaml:"plugin_alias" json:"plugin_alias"`
	Description string `yaml:"plugin_description" json:"description"`
	Version     string `yaml:"version" json:"version"`
	OS          string `yaml:"os" json:"os"`
	Middleware  string `yaml:"middleware" json:"middleware"`
	Author      string `yaml:"author" json:"author"`
	Repository  string `yaml:"repository" json:"repository"`

	UpdateDate string `yaml:"update_date" json:"update_date"`
	CreateDate string `yaml:"create_date" json:"create_date"`

	CacheEnable bool `yaml:"cache_enable,omitempty" json:"cache_enable"`
	AutoRun     bool `yaml:"auto_run,omitempty" json:"auto_run"`

	PluginPath string `yaml:"-" json:"plugin_path"`

	//Params PluginParams `yaml:"params" json:"params"`

	Layout MentalUILayout `yaml:"layout,omitempty" json:"layout"`
	Rule   ExtendInitRule `yaml:"rule" json:"rule"`
}
type PluginParams struct {
	Format string       `yaml:"format" json:"format"`
	Flags  []PluginFlag `yaml:"flags" json:"flags"`
}

type PluginFlag struct {
	Name        string        `yaml:"name" json:"name"`
	Shorthand   string        `yaml:"shorthand" json:"shorthand"`
	Value       interface{}   `yaml:"value" json:"value"`
	ValueBase64 bool          `yaml:"value_base64" json:"value_base64"`
	Ref         string        `yaml:"ref" json:"ref"`
	Usage       string        `yaml:"usage" json:"usage"`
	Required    bool          `yaml:"required" json:"required"`
	UIParam     PluginUIParam `yaml:"ui_param,omitempty" json:"ui_param"`
}
type PluginUIParam struct {
	Title        string            `json:"title" yaml:"title"`
	ParamType    string            `json:"param_type" yaml:"param_type,omitempty"`
	Tips         string            `json:"tips" yaml:"tips,omitempty"`
	DefaultValue interface{}       `json:"default_value" yaml:"default_value"`
	Required     bool              `yaml:"required"`
	AliasMap     map[string]string `json:"alias_map" yaml:"alias_map,omitempty"`
}

type MentalUILayout struct {
	LabelSpan   int `yaml:"label_span" json:"label_span"`
	ResultWidth int `yaml:"result_width" json:"result_width"`
}

func (this *MentalPlugin) GetPureName() string {
	return ""
}

func (this *MentalPlugin) GetFullName() string {
	return ""
}

type ExtendInitRule struct {
	Params         []ExtendParam               `json:"params" yaml:"params"`
	FilterCheckKey string                      `json:"filter_check_key,omitempty" yaml:"filter_check_key,omitempty"` // 过滤检查的key
	FilterRules    map[string]ExtendFilterRule `json:"filter_rules,omitempty" yaml:"filter_rules,omitempty"`         // 过滤规则
	PreLoad        ExtendPreLoad               `json:"pre_load,omitempty" yaml:"pre_load,omitempty"`
}
type ExtendFilterRule struct {
	Includes []string `json:"includes" yaml:"includes,omitempty"`
	Excludes []string `json:"excludes" yaml:"excludes,omitempty"`
}

type ExtendPreLoad struct {
	CheckName string `yaml:"check_name" json:"check_name"`
	LoadFile  string `yaml:"load_file" json:"load_file"`
}

type ExtendParam struct {
	Title        string            `json:"title" yaml:"title"`
	Key          string            `json:"key" yaml:"key"`
	ParamType    int               `json:"param_type" yaml:"param_type,omitempty"`
	Tips         string            `json:"tips" yaml:"tips,omitempty"`
	DefaultValue interface{}       `json:"default_value" yaml:"default_value"`
	Required     bool              `json:"required" yaml:"required,omitempty"`
	AliasMap     map[string]string `json:"alias_map" yaml:"alias_map,omitempty"`
}
