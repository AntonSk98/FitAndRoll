package participants

import (
	"fit_and_roll/backend/mappers"
)

type MemberCardInfo struct {
	Capacity  uint   `json:"capacity"`
	IssuedAt  string `json:issuedAt`
	ExpiredAt string `json:expiredAt,omitempty`
}

func NewMemberCardInfo(memberCard MemberCard) *MemberCardInfo {
	memberCardInfo := &MemberCardInfo{
		Capacity: memberCard.Capacity,
		IssuedAt: mappers.ToDateString(memberCard.CreatedAt),
	}

	if memberCard.Deleted.Valid {
		memberCardInfo.ExpiredAt = mappers.ToDateString(memberCard.Deleted.Time)
	}

	return memberCardInfo
}
