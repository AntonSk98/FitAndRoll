package participants

import (
	"fit_and_roll/backend/common"
	"time"
)

// DTO that encapsulates participant details
type ParticipantDto struct {
	ID                      uint       `json:"id,omitempty"`
	Name                    string     `json:"name"`
	Surname                 string     `json:"surname"`
	Group                   string     `json:"group"`
	Phone                   *string    `json:"phone,omitempty"`
	Email                   *string    `json:"email,omitempty"`
	Address                 *string    `json:"address,omitempty"`
	Parents                 *string    `json:"parents,omitempty"`
	PrivacyPolicy           bool       `json:"privacyPolicy"`
	Notes                   *string    `json:"notes,omitempty"`
	Birthday                *time.Time `json:"birthday,omitempty"`
	CreatedAt               string     `json:"createdAt"`
	PrivacyPolicyAcceptedAt *string    `json:"privacyPolicyAcceptedAt,omitempty"`
}

// Creates a new participant DTO from the Participant entity
func NewParticipantDto(participant Participant) *ParticipantDto {
	var privacyPolicyAcceptedAt *string

	if participant.PrivacyPolicyAcceptedAt != nil {
		formatted := common.ToDateString(*participant.PrivacyPolicyAcceptedAt)
		privacyPolicyAcceptedAt = &formatted
	}

	createdAt := common.ToDateString(participant.CreatedAt)

	return &ParticipantDto{
		ID:                      participant.ID,
		Name:                    participant.Name,
		Surname:                 participant.Surname,
		Group:                   participant.Group,
		Phone:                   participant.Phone,
		Email:                   participant.Email,
		Address:                 participant.Address,
		Parents:                 participant.Parents,
		PrivacyPolicy:           participant.PrivacyPolicy,
		Notes:                   participant.Notes,
		Birthday:                participant.Birthday,
		CreatedAt:               createdAt,
		PrivacyPolicyAcceptedAt: privacyPolicyAcceptedAt,
	}
}
