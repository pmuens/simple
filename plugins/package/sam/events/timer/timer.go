package timer

func NewTimerEvent() *timerEvent {
	timer := new(timerEvent)

	timer.Type = "Schedule"
	timer.Properties.Schedule = "rate(2 minutes)"

	return timer
}

type timerEvent struct {
	Type       string     `json:"Type" yaml:"Type"`
	Properties properties `json:"Properties" yaml:"Properties"`
}

type properties struct {
	Schedule string `json:"Scheudle" yaml:"Schedule"`
}
