package service

// https://stackoverflow.com/questions/64693710/parse-json-file-in-golang
type Services_slice struct {
	Services []Services `json:"Services"`
}

type Services struct {
	Ip        string `json:"Ip"`
	Service   string `json:"Service"`
	Port      string `json:"Port"`
}
