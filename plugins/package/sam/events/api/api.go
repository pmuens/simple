package api

func NewAPIEvent() *apiEvent {
	api := new(apiEvent)

	api.Type = "Api"

	return api
}

type apiEvent struct {
	Type       string     `json:"Type" yaml:"Type"`
	Properties properties `json:"Properties" yaml:"Properties"`
}

type properties struct {
	Path   string `json:"Path" yaml:"Path"`
	Method string `json:"Method" yaml:"Method"`
}
