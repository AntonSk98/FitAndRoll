package participants

import (
	"time"

	"gorm.io/gorm"
)

type Participant struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"not null"`
	Surname     string       `gorm:"not null"`
	Group       string       `gorm:"not null"`
	MemberCards []MemberCard `gorm:"foreignKey:ParticipantID;onDelete:CASCADE"`
	CreatedAt   time.Time
	Deleted     gorm.DeletedAt
}

func NewParticipant(command ParticipantCommand) *Participant {
	return &Participant{
		Name:    command.Name,
		Surname: command.Surname,
		Group:   command.Group,
	}
}

func (participant *Participant) update(command ParticipantCommand) {
	participant.Name = command.Name
	participant.Surname = command.Surname
	participant.Group = command.Surname
}
