package participants

import (
	"time"

	"gorm.io/gorm"
)

// Entity that represent a participant in the application
type Participant struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"not null"`
	Surname     string       `gorm:"not null"`
	Group       string       `gorm:"not null"`
	MemberCards []MemberCard `gorm:"foreignKey:ParticipantID;onDelete:CASCADE"`
	CreatedAt   time.Time
	Deleted     gorm.DeletedAt
}

// Creates a new participant instance
func NewParticipant(command ParticipantCommand) *Participant {
	return &Participant{
		Name:    command.Name,
		Surname: command.Surname,
		Group:   command.Group,
	}
}

// Updates participant details
func (participant *Participant) Update(command ParticipantCommand) {
	participant.Name = command.Name
	participant.Surname = command.Surname
	participant.Group = command.Surname
}
