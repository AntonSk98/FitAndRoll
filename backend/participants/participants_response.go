package participants

// DTO that encapsulates participant details
type ParticipantDto struct {
	ID      uint   `json:"id,omitempty"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Group   string `json:"group"`
}

// Creates a new participant DTO from the Participant entity
func NewParticipantDto(participant Participant) *ParticipantDto {
	return &ParticipantDto{
		ID:      participant.ID,
		Name:    participant.Name,
		Surname: participant.Surname,
		Group:   participant.Group,
	}
}
