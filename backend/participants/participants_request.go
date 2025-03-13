package participants

import (
	"fmt"
	"strings"
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
	ID      *uint  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Group   string `json:"group"`
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
