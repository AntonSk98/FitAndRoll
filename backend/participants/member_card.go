package participants

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

const capacityLimit = 10

type MemberCard struct {
	Capacity  uint `json:"capacity"`
	CreatedAt time.Time
	Deleted   gorm.DeletedAt
}

func NewMemberCard() *MemberCard {
	return &MemberCard{Capacity: capacityLimit}
}

func (memberCard *MemberCard) visitCourse() bool {
	if memberCard.Capacity == 0 {
		fmt.Println("Participant must buy a new card!")
		return false
	}
	memberCard.Capacity -= 1
	return true
}

func (memberCard *MemberCard) restoreVisit() bool {
	if memberCard.Capacity == capacityLimit {
		fmt.Println("Cannot restore a visit for that card as it has never been used!")
		return false
	}
	memberCard.Capacity += 1
	return true
}
