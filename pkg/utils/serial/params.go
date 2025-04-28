package serial

type ParamMeta map[string]interface{}

type ControlParams struct {
}

func NewControlParams() *ControlParams {
	params := &ControlParams{}
	return params
}

func (this *ControlParams) GetMeta() ParamMeta {
	return nil
}

func (this *ControlParams) SetMeta(meta ParamMeta) {
}
func (this *ControlParams) RemoveParam(key string) {
}
func (this *ControlParams) SetEncoding(charset string) {
}

func (this *ControlParams) SetParam(key string, value interface{}) {
}

func (this *ControlParams) GetParam(key string) (value interface{}, err error) {
	return
}
