package import_

import (
	"fit_and_roll/backend/common"
	"time"
)

// ImportParticipantCommand is a command to import a participant from an Excel file
type ImportParticipantCommand struct {
	Name                    string
	Surname                 string
	Group                   string
	TrainingStart           *time.Time
	Birthday                *time.Time
	Phone                   *string
	Email                   *string
	Address                 *string
	Parents                 *string
	PrivacyPolicy           bool
	PrivacyPolicyAcceptedAt *time.Time
	Notes                   *string
}

func Parse(row []string) (*ImportParticipantCommand, error) {
	if len(row) != 12 {
		return nil, &ParticipantsImportError{
			Code: REQUIRED_COLUMNS_ERROR,
		}
	}
	name := row[0]
	surname := row[1]
	group := row[2]
	phone := row[3]
	email := row[4]
	birthdayString := row[5]
	trainingStartString := row[6]
	privacyPolicyString := row[7]
	privacyPolicyAcceptedAtString := row[8]
	address := row[9]
	parents := row[10]
	notes := row[11]

	id := name + surname

	if name == "" || surname == "" || group == "" {
		return nil, &ParticipantsImportError{
			Id:   id,
			Code: REQUIRED_PARAMETER_MISSING_ERROR,
		}
	}

	birthday, err := toDateOnlyOrNil(birthdayString)
	if err != nil {
		return nil, &ParticipantsImportError{
			Id:   id,
			Code: BIRTHDAY_PARSE_ERROR,
		}

	}

	trainingStart, err := toDateOnlyOrNil(trainingStartString)
	if err != nil {
		return nil, &ParticipantsImportError{
			Id:   id,
			Code: TRAINING_START_PARSE_ERROR,
		}
	}

	privacyPolicy := toBool(privacyPolicyString)
	privacyPolicyAcceptedAt, err := toDateOnlyOrNil(privacyPolicyAcceptedAtString)
	if err != nil {
		return nil, &ParticipantsImportError{
			Id:   id,
			Code: PRIVACY_POLICY_ACCEPTED_AT_PARSE_ERROR,
		}
	}

	if privacyPolicy && privacyPolicyAcceptedAt == nil {
		return nil, &ParticipantsImportError{
			Id:   id,
			Code: PRIVACY_POLICY_ACCEPTED_AT_REQUIRED_ERROR,
		}

	}

	return &ImportParticipantCommand{
		Name:                    name,
		Surname:                 surname,
		Group:                   group,
		TrainingStart:           trainingStart,
		Birthday:                birthday,
		Phone:                   nilIfEmpty(phone),
		Email:                   nilIfEmpty(email),
		Parents:                 nilIfEmpty(parents),
		Address:                 nilIfEmpty(address),
		PrivacyPolicy:           privacyPolicy,
		PrivacyPolicyAcceptedAt: privacyPolicyAcceptedAt,
		Notes:                   nilIfEmpty(notes),
	}, nil
}

func nilIfEmpty(inputString string) *string {
	if inputString != "" {
		return &inputString
	}
	return nil
}

func toBool(inputString string) bool {
	return inputString == "1"
}

func toDateOnlyOrNil(inputString string) (*time.Time, error) {
	if inputString != "" {
		trainingStart, err := common.ToDateOnly(inputString)
		if err != nil {
			return nil, err
		}
		return &trainingStart, nil
	}
	return nil, nil
}
