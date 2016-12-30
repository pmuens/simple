package alexa

func NewAlexaSkillEvent() *alexaSkillEvent {
	alexaSkill := new(alexaSkillEvent)

	alexaSkill.Type = "AlexaSkill"

	return alexaSkill
}

type alexaSkillEvent struct {
	Type string `json:"Type" yaml:"Type"`
}
