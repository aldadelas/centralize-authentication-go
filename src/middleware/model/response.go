package model

type MetaVersion struct {
	Version string `json:"version"`
}
type ResponseError struct {
	Meta   MetaVersion `json:"meta"`
	Reason string      `json:"reason"`
}

type ResponseSuccess struct {
	Meta MetaVersion `json:"meta"`
	Data interface{} `json:"data"`
}
