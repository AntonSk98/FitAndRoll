package participants

type ParticipantDto struct {
	ID      uint   `json:"id,omitempty"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Group   string `json:"group"`
}

func NewParticipantDto(participant Participant) *ParticipantDto {
	return &ParticipantDto{
		ID:      participant.ID,
		Name:    participant.Name,
		Surname: participant.Surname,
		Group:   participant.Group,
	}
}
