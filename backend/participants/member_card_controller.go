package participants

import (
	"fit_and_roll/backend/config"
	"fmt"

	"gorm.io/gorm"
)

type MemberCardController struct {
	dbManager *config.DatabaseManager
}

func NewMemberCardController(dbManager *config.DatabaseManager) *MemberCardController {
	return &MemberCardController{dbManager: dbManager}
}

func (controller *MemberCardController) FindAllMemberCards(participantId uint) ([]MemberCardInfo, error) {
	var participantMemberCards []MemberCard
	var participantMemberCardInfos []MemberCardInfo

	baseQuery := controller.dbManager.DB.Model(&MemberCard{}).Where("participant_id = ?", participantId)

	if err := baseQuery.
		Order("created_at desc").
		Find(&participantMemberCards).Error; err != nil {
		return nil, err
	}

	for _, memparticipantMemberCard := range participantMemberCards {
		participantMemberCardInfos = append(participantMemberCardInfos, *NewMemberCardInfo(memparticipantMemberCard))
	}

	return participantMemberCardInfos, nil
}

func (controller *MemberCardController) IssueNewMemberCard(participantId uint) error {
	return controller.dbManager.Transactional(func(tx *gorm.DB) error {
		participantPointer, err := requireParticipant(participantId, tx)
		if err != nil {
			return err
		}

		memberCard := NewMemberCard(participantPointer.ID)

		if err := tx.Create(&memberCard).Error; err != nil {
			return err
		}

		return nil
	})
}

func (controller *MemberCardController) UndoIssueNewMemberCard(participantId uint, memberCardId uint) error {
	result := controller.dbManager.DB.
		Unscoped().
		Where("id = ? AND participant_id = ?", memberCardId, participantId).
		Delete(&MemberCard{})

	if result.RowsAffected == 0 {
		return fmt.Errorf("could not remove the member card. Either participant or member card do not exist")
	}

	if err := result.Error; err != nil {
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
