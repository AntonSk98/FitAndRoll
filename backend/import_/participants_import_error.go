package import_

import "encoding/json"

const (
	SELECT_FILE_ERROR                         = "SELECT_FILE_ERROR"
	OPEN_FILE_ERROR                           = "OPEN_FILE_ERROR"
	NO_SHEET_ERROR                            = "NO_SHEET_ERROR"
	GET_ROWS_ERROR                            = "GET_ROWS_ERROR"
	REQUIRED_COLUMNS_ERROR                    = "REQUIRED_COLUMNS_ERROR"
	REQUIRED_PARAMETER_MISSING_ERROR          = "REQUIRED_PARAMETER_MISSING_ERROR"
	TRAINING_START_PARSE_ERROR                = "TRAINING_START_PARSE_ERROR"
	BIRTHDAY_PARSE_ERROR                      = "BIRTHDAY_PARSE_ERROR"
	PRIVACY_POLICY_ACCEPTED_AT_PARSE_ERROR    = "PRIVACY_POLICY_ACCEPTED_AT_PARSE_ERROR"
	PRIVACY_POLICY_ACCEPTED_AT_REQUIRED_ERROR = "PRIVACY_POLICY_ACCEPTED_AT_REQUIRED_ERROR"
	NO_DATA_ERROR                             = "NO_DATA_ERROR"
)

// ParticipantsImportError represents an error that occurred during the import of participants.
type ParticipantsImportError struct {
	Id   string `json:"id"`
	Code string `json:"code"`
}

// Implements the error interface.
func (e *ParticipantsImportError) Error() string {
	json, _ := json.Marshal(e)
	return string(json)
}
