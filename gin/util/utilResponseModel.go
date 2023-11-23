package util

type FDModelResult struct {
	Result []int  `json:"result"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type TSModelResult struct {
	Result []TSModelData `json:"result"`
	Status string        `json:"status"`
	Msg    string        `json:"msg"`
}

type TSModelData struct {
	Amplitude float64 `json:"amplitude"`
	Freq      float64 `json:"freq"`
	Phase     float64 `json:"phase"`
	ACF       float64 `json:"acf"`
}
