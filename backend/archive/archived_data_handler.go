package archive

import (
	"fit_and_roll/backend/common"
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/courses"
	"fit_and_roll/backend/participants"
	"fmt"

	"gorm.io/gorm"
)

// ArchivedDataHandler provides functionality to view archived data and unarchive them
type ArchivedDataHandler struct {
	dbManager *config.DatabaseManager
}

// NewArchivedDataHandler creates a new instance of ArchivedDataHandler
func NewArchivedDataHandler(dbManager *config.DatabaseManager) *ArchivedDataHandler {
	return &ArchivedDataHandler{dbManager: dbManager}
}

// Finds archived participants
func (handler *ArchivedDataHandler) FindArchivedParticipants(params FindArchivedEntryParams, pageParams common.PageParams) (*common.Page[ArchivedEntryDto], error) {
	var archivedParticipantsDto []ArchivedEntryDto
	var archivedParticipants []participants.Participant
	var total int64

	// Build the query with conditions
	query := handler.dbManager.DB.Model(participants.Participant{}).Unscoped().Where("participants.deleted IS NOT NULL")

	// Filter by NameLike if provided
	if params.NameLike != nil {
		query = query.Where("LOWER(CONCAT(participants.name, ' ', participants.surname)) LIKE ?", "%"+*params.NameLike+"%")
	}

	// Get total count of participants
	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count archived participants: %w", err)
	}

	// Fetch paginated results
	if err := query.
		Scopes(handler.dbManager.Paginate(pageParams.Page, pageParams.Size)).
		Order("participants.deleted DESC").
		Find(&archivedParticipants).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch archived participants: %w", err)
	}

	// Map the fetched participants to DTOs
	for _, archivedParticipant := range archivedParticipants {
		archivedParticipantDto, err := mapArchivedParticipant(archivedParticipant)
		if err != nil {
			return nil, err
		}
		archivedParticipantsDto = append(archivedParticipantsDto, *archivedParticipantDto)
	}

	// Return the paginated result
	return &common.Page[ArchivedEntryDto]{
		Data:  archivedParticipantsDto,
		Total: int(total),
		Page:  pageParams.Page,
		Size:  pageParams.Size,
	}, nil
}

// Finds archived courses
func (handler *ArchivedDataHandler) FindArchivedCourses(params FindArchivedEntryParams, pageParams common.PageParams) (*common.Page[ArchivedEntryDto], error) {
	var archivedCoursesDto []ArchivedEntryDto
	var archivedCourses []courses.Course
	var total int64

	// Build the query with conditions
	query := handler.dbManager.DB.Model(courses.Course{}).Unscoped().Where("courses.deleted IS NOT NULL")

	// Filter by NameLike if provided
	if params.NameLike != nil {
		query = query.Where("LOWER(courses.name) LIKE ?", "%"+*params.NameLike+"%")
	}

	// Get total count of courses
	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count archived courses: %w", err)
	}

	// Fetch paginated results
	if err := query.
		Scopes(handler.dbManager.Paginate(pageParams.Page, pageParams.Size)).
		Order("courses.deleted DESC").
		Find(&archivedCourses).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch archived courses: %w", err)
	}

	// Map the fetched courses to DTOs
	for _, archivedCourse := range archivedCourses {
		archivedCourseDto, err := mapArchivedCourse(archivedCourse)
		if err != nil {
			return nil, err
		}
		archivedCoursesDto = append(archivedCoursesDto, *archivedCourseDto)
	}

	// Return the paginated result
	return &common.Page[ArchivedEntryDto]{
		Data:  archivedCoursesDto,
		Total: int(total),
		Page:  pageParams.Page,
		Size:  pageParams.Size,
	}, nil
}

// Unarchives a participant by identifier
func (handler *ArchivedDataHandler) UnarchiveParticipant(id *uint) error {
	if err := assertIdNotNull(id); err != nil {
		return err
	}

	return handler.dbManager.Transactional(func(tx *gorm.DB) error {
		var archivedParticipant participants.Participant
		if err := tx.Unscoped().First(&archivedParticipant, id).Error; err != nil {
			return err
		}

		if err := archivedParticipant.Unarchive(); err != nil {
			return err
		}

		tx.Save(&archivedParticipant)

		return nil
	})
}

// Unarchives a course by identifier
func (handler *ArchivedDataHandler) UnarchiveCourse(id *uint) error {
	if err := assertIdNotNull(id); err != nil {
		return err
	}

	return handler.dbManager.Transactional(func(tx *gorm.DB) error {
		var archivedCourse courses.Course
		if err := tx.Unscoped().First(&archivedCourse, id).Error; err != nil {
			return err
		}

		if err := archivedCourse.Unarchive(); err != nil {
			return err
		}

		tx.Save(&archivedCourse)

		return nil
	})
}

func assertIdNotNull(id *uint) error {
	if id == nil {
		return fmt.Errorf("id cannot be null")
	}

	return nil
}
