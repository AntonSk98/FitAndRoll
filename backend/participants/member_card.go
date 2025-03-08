package participants

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

const capacityLimit = 10

type MemberCard struct {
	ID            uint `gorm:"primaryKey"`
	Capacity      uint
	CreatedAt     time.Time
	Deleted       gorm.DeletedAt
	ParticipantID uint
}

func NewMemberCard(participant uint) *MemberCard {
	return &MemberCard{Capacity: capacityLimit, ParticipantID: participant}
}

func (memberCard *MemberCard) isValid() bool {
	return memberCard.Capacity > 0
}

func (memberCard *MemberCard) visitCourse() bool {
	if !memberCard.isValid() {
		fmt.Println("Participant must buy a new card!")
		return false
	}
	memberCard.Capacity -= 1
	return memberCard.isValid()
}

func (memberCard *MemberCard) markAsUsed() {
	if !memberCard.isValid() {
		memberCard.Deleted = gorm.DeletedAt{Time: time.Now(), Valid: true}
	}
}

func (memberCard *MemberCard) restoreVisit() bool {
	if memberCard.Capacity == capacityLimit {
		fmt.Println("Cannot restore a visit for that card as it has never been used!")
		return false
	}
	memberCard.Capacity += 1
	return true
}

func (memberCard *MemberCard) isUnused() bool {
	return memberCard.Capacity == capacityLimit
}
