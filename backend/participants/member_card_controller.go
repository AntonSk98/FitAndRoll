package participants

import (
	"fit_and_roll/backend/common"
	"fit_and_roll/backend/config"

	"gorm.io/gorm"
)

type MemberCardController struct {
	dbManager *config.DatabaseManager
}

func NewMemberCardController(dbManager *config.DatabaseManager) *MemberCardController {
	return &MemberCardController{dbManager: dbManager}
}

func (controller *MemberCardController) FindAllMemberCards(participantId uint, pageParams common.PageParams) (*common.Page[MemberCardInfo], error) {
	var total int64
	var participantMemberCards []MemberCard
	var participantMemberCardInfos []MemberCardInfo

	baseQuery := controller.dbManager.DB.Where("participant_id = ?", participantId)

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := baseQuery.
		Scopes(controller.dbManager.Paginate(pageParams.Page, pageParams.Size)).
		Order("created_at desc").
		Find(&participantMemberCards).Error; err != nil {
		return nil, err
	}

	for _, memparticipantMemberCard := range participantMemberCards {
		participantMemberCardInfos = append(participantMemberCardInfos, *NewMemberCardInfo(memparticipantMemberCard))
	}

	return &common.Page[MemberCardInfo]{
		Data:  participantMemberCardInfos,
		Total: int(total),
		Page:  pageParams.Page,
		Size:  pageParams.Size,
	}, nil
}

func (controller *MemberCardController) IssueNewMemberCard(participantId uint) error {
	return controller.dbManager.Transactional(func(tx *gorm.DB) error {
		participantPointer, err := requireParticipant(participantId, tx)
		if err != nil {
			return err
		}

		memberCard := NewMemberCard()

		if err := tx.Model(participantPointer).Association("MemberCards").Append(&memberCard); err != nil {
			return err
		}

		return nil
	})
}

func (controller *MemberCardController) UndoIssueNewMemberCard(participantId uint, memberCardId uint) error {
	if err := controller.dbManager.DB.
		Unscoped().
		Where("id = ? AND participant_id = ?", memberCardId, participantId).
		Delete(&MemberCard{}).Error; err != nil {
		return err
	}

	return nil
}

func requireParticipant(participantId uint, db *gorm.DB) (*Participant, error) {
	var participant Participant

	if err := db.First(&participant, participantId).Error; err != nil {
		return nil, err
	}

	return &participant, nil
}
