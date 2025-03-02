package participants

import (
	"fit_and_roll/backend/common"
	"fit_and_roll/backend/config"
	"fmt"

	"gorm.io/gorm"
)

type ParticipantsController struct {
	dbManager *config.DatabaseManager
}

func NewParticipantsController(dbManager *config.DatabaseManager) *ParticipantsController {
	return &ParticipantsController{dbManager: dbManager}
}

func (controller *ParticipantsController) FindParticipants(filter FindParticipantsParams, pageParams common.PageParams) (*common.Page[ParticipantDto], error) {

	findAllParticipantsTemplateQuery := controller.dbManager.DB

	if filter.Name != "" {
		findAllParticipantsTemplateQuery = findAllParticipantsTemplateQuery.Where("LOWER(name) LIKE LOWER(?)", "%"+filter.Name+"%")
	}

	if filter.Surname != "" {
		findAllParticipantsTemplateQuery = findAllParticipantsTemplateQuery.Where("LOWER(surname) LIKE LOWER(?)", "%"+filter.Surname+"%")
	}

	if filter.Group != "" {
		findAllParticipantsTemplateQuery = findAllParticipantsTemplateQuery.Where("LOWER(`group`) LIKE LOWER(?)", "%"+filter.Group+"%")
	}

	if filter.WithActiveCard {
		findAllParticipantsTemplateQuery = findAllParticipantsTemplateQuery.InnerJoins("MemberCards", "member_cards.deleted IS NULL")
	}

	var total int64

	if err := findAllParticipantsTemplateQuery.Model(&Participant{}).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("error counting the total number of participants. Error: %v", err)
	}

	var participants []Participant
	if err := findAllParticipantsTemplateQuery.Scopes(controller.dbManager.Paginate(pageParams.Page, pageParams.Size)).Order("`group` desc").Find(&participants).Error; err != nil {
		return nil, fmt.Errorf("error fetching participants: %v", err)
	}

	var participantsDto []ParticipantDto
	for _, participant := range participants {
		participantsDto = append(participantsDto, *NewParticipantDto(participant))
	}

	return &common.Page[ParticipantDto]{
		Data:  participantsDto,
		Total: int(total),
		Page:  pageParams.Page,
		Size:  pageParams.Size,
	}, nil
}

func (controller *ParticipantsController) FindParticipantDetails(id uint) (*ParticipantDto, error) {
	var participant Participant
	if err := controller.dbManager.DB.First(&participant, id).Error; err != nil {
		return nil, err
	}
	return NewParticipantDto(participant), nil
}

func (controller *ParticipantsController) CreateUpdateParticipant(command ParticipantCommand) error {
	if err := command.Validate(); err != nil {
		return err
	}

	return controller.dbManager.Transactional(func(tx *gorm.DB) error {
		if command.ID != nil {
			var participant Participant
			if err := tx.First(&participant, *command.ID).Error; err != nil {
				return fmt.Errorf("failed to fetch a participant with id %d. Error: %v", *command.ID, err)
			}
			participant.Name = command.Name
			participant.Surname = command.Surname
			participant.Group = command.Group
			if err := tx.Save(&participant).Error; err != nil {
				return fmt.Errorf("failed to update a participant. Command: %v. Error: %v", command, err)
			}
			return nil
		}

		participant := NewParticipant(command)
		if err := tx.Create(participant).Error; err != nil {
			return fmt.Errorf("failed to create a participant. Command: %v. Error: %v", command, err)
		}
		return nil
	})
}

func (controller *ParticipantsController) ArchiveParticipant(id uint) error {
	if err := controller.dbManager.DB.Delete(&Participant{}, id).Error; err != nil {
		return fmt.Errorf("Failed to archive a participant. Error: %v", err)
	}
	return nil
}
