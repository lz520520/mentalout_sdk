package mental_models

type CommonResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Err    string `json:"err"`
}
