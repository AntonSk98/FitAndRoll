package participants

import (
	"fmt"
	"strings"
	"time"
)

// Request with parameters to find participants
type FindParticipantsParams struct {
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Group          string `json:"group"`
	WithActiveCard bool   `json:"withActiveCard"`
}

// Command to create or update a participant.
// If ID is not provided it means that a participant should be created.
type ParticipantCommand struct {
	ID                      *uint      `json:"id"`
	Name                    string     `json:"name"`
	Surname                 string     `json:"surname"`
	Group                   string     `json:"group"`
	Phone                   *string    `json:"phone"`
	Email                   *string    `json:"email"`
	Address                 *string    `json:"address"`
	Parents                 *string    `json:"parents"`
	PrivacyPolicy           bool       `json:"privacyPolicy"`
	Notes                   *string    `json:"notes"`
	Birthday                *time.Time `json:"birthday"`
	CreatedAt               *time.Time `json:"-"`
	PrivacyPolicyAcceptedAt *time.Time `json:"-"`
}

// Validates the ParticipantCommand
func (command *ParticipantCommand) Validate() error {
	if strings.TrimSpace(command.Name) == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if strings.TrimSpace(command.Surname) == "" {
		return fmt.Errorf("surname cannot be empty")
	}
	if strings.TrimSpace(command.Group) == "" {
		return fmt.Errorf("group cannot be empty")
	}
	return nil
}
