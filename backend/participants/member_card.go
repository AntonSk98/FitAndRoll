package participants

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

const capacityLimit = 10

// Entity that represents a member card
// A member card is used to attend courses
type MemberCard struct {
	ID            uint `gorm:"primaryKey"`
	Capacity      uint
	CreatedAt     time.Time
	Deleted       gorm.DeletedAt
	ParticipantID uint
}

// Creates a new MemberCard instance
func NewMemberCard(participant uint) *MemberCard {
	return &MemberCard{Capacity: capacityLimit, ParticipantID: participant}
}

// Check is a member card is valid
// A valid card is considered valid if there is at least one free slot that could be used to attend a course.
func (memberCard *MemberCard) isValid() bool {
	return memberCard.Capacity > 0
}

// Use the member card to visit a course
// It reduces the capacity of the card
// If the member card is used then it will be soft deleted
func (memberCard *MemberCard) VisitCourse() error {
	if !memberCard.isValid() {
		return fmt.Errorf("Participant %v must buy a new card", memberCard.ParticipantID)
	}
	memberCard.Capacity -= 1
	if !memberCard.isValid() {
		memberCard.markAsUsed()
	}
	return nil
}

// UndoAttendance restores attendance for a member card by performing the following actions:
// - If the member card was fully used, it will be restored.
// - If the member card has available capacity below the defined limit, it will increment the capacity by 1.
// - If the member card's capacity has already reached the limit, a message will be printed indicating that the capacity cannot be increased while undoing the course attendance.
func (memberCard *MemberCard) UndoAttendance() error {
	if memberCard.Deleted.Valid {
		memberCard.Deleted = gorm.DeletedAt{}
	}

	if memberCard.Capacity < capacityLimit {
		memberCard.Capacity += 1
		return nil
	}

	return fmt.Errorf("Cannot increase capacity while undoing the course attendance: Member card is already full")
}

func (memberCard *MemberCard) markAsUsed() {
	if !memberCard.isValid() {
		memberCard.Deleted = gorm.DeletedAt{Time: time.Now(), Valid: true}
	}
}

func (memberCard *MemberCard) isUnused() bool {
	return memberCard.Capacity == capacityLimit
}
