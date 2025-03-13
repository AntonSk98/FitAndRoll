package participants

import (
	"fit_and_roll/backend/common"
	"fit_and_roll/backend/config"
	"fmt"

	"gorm.io/gorm"
)

// Provides CRUD operations to manage participants
type ParticipantsHandler struct {
	dbManager *config.DatabaseManager
}

// Creates a new participant handler
func NewParticipantsHandler(dbManager *config.DatabaseManager) *ParticipantsHandler {
	return &ParticipantsHandler{dbManager: dbManager}
}

// Finds Participants
// The method supports filtering and returns the page of found participants
func (controller *ParticipantsHandler) FindParticipants(filter FindParticipantsParams, pageParams common.PageParams) (*common.Page[ParticipantDto], error) {
	findAllParticipantsTemplateQuery := controller.findParticipantsFilterQuery(filter)

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

// Finds details of one participant by participant identifier.
func (controller *ParticipantsHandler) FindParticipantDetails(id uint) (*ParticipantDto, error) {
	var participant Participant
	if err := controller.dbManager.DB.First(&participant, id).Error; err != nil {
		return nil, err
	}
	return NewParticipantDto(participant), nil
}

// Creates or updates a participant
func (controller *ParticipantsHandler) CreateUpdateParticipant(command ParticipantCommand) error {
	if err := command.Validate(); err != nil {
		return err
	}

	return controller.dbManager.Transactional(func(tx *gorm.DB) error {
		if command.ID != nil {
			var participant Participant
			if err := tx.First(&participant, *command.ID).Error; err != nil {
				return fmt.Errorf("failed to fetch a participant with id %d. Error: %v", *command.ID, err)
			}
			participant.Update(command)
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

// Archives a participant.
// The participant will be marked as removed in the persistent storage.
func (controller *ParticipantsHandler) ArchiveParticipant(id uint) error {
	if err := controller.dbManager.DB.Delete(&Participant{}, id).Error; err != nil {
		return fmt.Errorf("failed to archive a participant. Error: %v", err)
	}
	return nil
}

func (controller *ParticipantsHandler) findParticipantsFilterQuery(filter FindParticipantsParams) *gorm.DB {
	query := controller.dbManager.DB

	if filter.Name != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+filter.Name+"%")
	}

	if filter.Surname != "" {
		query = query.Where("LOWER(surname) LIKE LOWER(?)", "%"+filter.Surname+"%")
	}

	if filter.Group != "" {
		query = query.Where("LOWER(`group`) LIKE LOWER(?)", "%"+filter.Group+"%")
	}

	if filter.WithActiveCard {
		query = query.Where("participants.id IN (?)",
			controller.dbManager.DB.Model(&MemberCard{}).
				Select("participant_id").
				Where("deleted IS NULL"),
		)
	}

	return query
}
