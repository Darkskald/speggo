package models

type ResponseWrapper struct {
	Data   interface{} `json:"data"`
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
}

func Wrap(data interface{}) ResponseWrapper {
	return ResponseWrapper{
		Data:   data,
		Limit:  0,
		Offset: 0,
	}
}
