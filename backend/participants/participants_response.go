package participants

type ParticipantDto struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Group   string `json:"group"`
}

func NewParticipantDto(participant Participant) *ParticipantDto {
	return &ParticipantDto{
		Name:    participant.Name,
		Surname: participant.Surname,
		Group:   participant.Group,
	}
}
