package participants

import (
	"fit_and_roll/backend/common"
)

type MemberCardInfo struct {
	ID        uint   `json:"id"`
	Capacity  uint   `json:"capacity"`
	IssuedAt  string `json:"issuedAt"`
	ExpiredAt string `json:"expiredAt,omitempty"`
}

func NewMemberCardInfo(memberCard MemberCard) *MemberCardInfo {
	memberCardInfo := &MemberCardInfo{
		ID:       memberCard.ID,
		Capacity: memberCard.Capacity,
		IssuedAt: common.ToDateString(memberCard.CreatedAt),
	}

	if memberCard.Deleted.Valid {
		memberCardInfo.ExpiredAt = common.ToDateString(memberCard.Deleted.Time)
	}

	return memberCardInfo
}
