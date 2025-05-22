package import_

import (
	"context"
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/participants"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// Handler to import participants
type ImportParticipantsHandler struct {
	ctx                 *context.Context
	dbManager           *config.DatabaseManager
	participantsHandler *participants.ParticipantsHandler
}

// Creates a new import participants handler
func NewImportParticipantsHandler(dbManager *config.DatabaseManager, participantsHandler *participants.ParticipantsHandler) *ImportParticipantsHandler {
	return &ImportParticipantsHandler{
		dbManager:           dbManager,
		participantsHandler: participantsHandler,
	}
}

func (handler *ImportParticipantsHandler) ImportParticipants() error {
	importParticipantsFilePath, err := handler.selectImportParticipantsFile()
	if importParticipantsFilePath == "" {
		return nil
	}

	if err != nil {
		return err

	}

	participantRows, err := handler.extractParticipantRows(importParticipantsFilePath)
	if err != nil {
		return err
	}

	return handler.dbManager.Transactional(func(tx *gorm.DB) error {

		for i := 1; i < len(participantRows); i++ {
			participantRow := participantRows[i]
			importParticipantCommand, err := Parse(participantRow)
			if err != nil {
				return err
			}

			participantCommand := &participants.ParticipantCommand{
				Name:                    importParticipantCommand.Name,
				Surname:                 importParticipantCommand.Surname,
				Group:                   importParticipantCommand.Group,
				CreatedAt:               importParticipantCommand.TrainingStart,
				Birthday:                importParticipantCommand.Birthday,
				Phone:                   importParticipantCommand.Phone,
				Email:                   importParticipantCommand.Email,
				Address:                 importParticipantCommand.Address,
				Parents:                 importParticipantCommand.Parents,
				PrivacyPolicy:           importParticipantCommand.PrivacyPolicy,
				PrivacyPolicyAcceptedAt: importParticipantCommand.PrivacyPolicyAcceptedAt,
				Notes:                   importParticipantCommand.Notes,
			}
			participant, err := handler.participantsHandler.FindIdentifierByRequiredAttributes(
				importParticipantCommand.Name,
				importParticipantCommand.Surname)

			if err != nil {
				return err
			}

			if participant != nil {
				participantCommand.ID = participant.ID
			}

			err = handler.participantsHandler.CreateUpdateParticipant(*participantCommand)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (handler *ImportParticipantsHandler) SetContext(ctx context.Context) {
	handler.ctx = &ctx
}

func (handler *ImportParticipantsHandler) selectImportParticipantsFile() (string, error) {
	file, err := runtime.OpenFileDialog(*handler.ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Excel File (*.xlsx)",
				Pattern:     "*.xlsx",
			},
		},
	})

	if err != nil {
		return "", &ParticipantsImportError{
			Code: SELECT_FILE_ERROR,
		}
	}

	return file, nil
}

func (handler *ImportParticipantsHandler) extractParticipantRows(importParticipantsFilePath string) ([][]string, error) {
	importParticipantsFile, err := excelize.OpenFile(importParticipantsFilePath)
	if err != nil {
		return nil, &ParticipantsImportError{Code: OPEN_FILE_ERROR}
	}
	defer importParticipantsFile.Close()

	sheetName := importParticipantsFile.GetSheetName(0)
	if sheetName == "" {
		return nil, &ParticipantsImportError{Code: NO_SHEET_ERROR}
	}

	rowsRaw, err := importParticipantsFile.GetRows(sheetName)
	if err != nil {
		return nil, &ParticipantsImportError{Code: GET_ROWS_ERROR}
	}

	if len(rowsRaw) == 0 {
		return nil, &ParticipantsImportError{Code: NO_DATA_ERROR}
	}

	// Expected number of columns from header row
	expectedColumns := len(rowsRaw[0])
	var normalizedRows [][]string

	// Iterate starting from the first row (including header)
	for rowIndex := 1; rowIndex <= len(rowsRaw); rowIndex++ {
		normalizedRow := make([]string, expectedColumns)

		for colIndex := 1; colIndex <= expectedColumns; colIndex++ {
			cellName, err := excelize.CoordinatesToCellName(colIndex, rowIndex)
			if err != nil {
				return nil, &ParticipantsImportError{Code: GET_ROWS_ERROR}
			}
			val, err := importParticipantsFile.GetCellValue(sheetName, cellName)
			if err != nil {
				return nil, &ParticipantsImportError{Code: GET_ROWS_ERROR}
			}
			normalizedRow[colIndex-1] = val
		}

		normalizedRows = append(normalizedRows, normalizedRow)
	}

	return normalizedRows, nil
}
