package archive

import (
	"fit_and_roll/backend/common"
	"fit_and_roll/backend/courses"
	"fit_and_roll/backend/participants"
	"fmt"
)

// Parameters to find an archived entry
type FindArchivedEntryParams struct {
	NameLike *string `json:"name,omitempty"`
}

// Represents an archived entry dto
type ArchivedEntryDto struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	ArchivedAt string `json:"archivedAt"`
}

// Maps archived participant to archived entry
func mapArchivedParticipant(participant participants.Participant) (*ArchivedEntryDto, error) {
	if !participant.Deleted.Valid {
		return nil, fmt.Errorf("cannot map a non-archived participant")
	}

	return &ArchivedEntryDto{
		ID:         participant.ID,
		Name:       participant.Name + " " + participant.Surname,
		ArchivedAt: common.ToDateTime(participant.Deleted.Time),
	}, nil
}

// Maps archived course to archived entry
func mapArchivedCourse(course courses.Course) (*ArchivedEntryDto, error) {
	if !course.Deleted.Valid {
		return nil, fmt.Errorf("cannot map a non-archived course")
	}
	return &ArchivedEntryDto{
		ID:         course.ID,
		Name:       course.Name,
		ArchivedAt: common.ToDateTime(course.Deleted.Time),
	}, nil
}
