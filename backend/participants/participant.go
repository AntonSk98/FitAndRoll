package participants

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Entity that represent a participant in the application
type Participant struct {
	ID                      uint         `gorm:"primaryKey"`
	Name                    string       `gorm:"not null"`
	Surname                 string       `gorm:"not null"`
	Group                   string       `gorm:"not null"`
	MemberCards             []MemberCard `gorm:"foreignKey:ParticipantID;onDelete:CASCADE"`
	CreatedAt               time.Time
	Deleted                 gorm.DeletedAt
	Phone                   *string
	Email                   *string
	Address                 *string
	Parents                 *string
	PrivacyPolicy           bool
	PrivacyPolicyAcceptedAt *time.Time
	Notes                   *string
	Birthday                *time.Time
}

// Creates a new participant instance
func NewParticipant(command ParticipantCommand) *Participant {
	return &Participant{
		Name:          command.Name,
		Surname:       command.Surname,
		Group:         command.Group,
		Phone:         command.Phone,
		Email:         command.Email,
		Address:       command.Address,
		Parents:       command.Parents,
		PrivacyPolicy: command.PrivacyPolicy,
		Notes:         command.Notes,
		Birthday:      command.Birthday,
	}
}

// Updates participant details
func (participant *Participant) Update(command ParticipantCommand) {
	participant.Name = command.Name
	participant.Surname = command.Surname
	participant.Group = command.Group
	participant.Phone = command.Phone
	participant.Email = command.Email
	participant.Notes = command.Notes
	participant.Birthday = command.Birthday
	participant.Address = command.Address
	participant.Parents = command.Parents
	participant.PrivacyPolicy = command.PrivacyPolicy
	if command.PrivacyPolicy {
		now := time.Now()
		participant.PrivacyPolicyAcceptedAt = &now
	} else {
		participant.PrivacyPolicyAcceptedAt = nil
	}
}

// Unarchives an archived participant
func (participant *Participant) Unarchive() error {
	if !participant.Deleted.Valid {
		return fmt.Errorf("unable to unarchive an active participant")
	}

	participant.Deleted = gorm.DeletedAt{}
	return nil
}
