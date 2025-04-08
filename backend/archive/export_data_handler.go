package archive

import (
	"context"
	"fit_and_roll/backend/common"
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/courses"
	"fit_and_roll/backend/membercardattendance"
	"fit_and_roll/backend/participants"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// ExportDataHandler provides functionality to export data into excel file.
type ExportDataHandler struct {
	dbManager *config.DatabaseManager
	ctx       context.Context
}

// Createas a new instance of ExportDataHandler
func NewExportDataHandler(dbManager *config.DatabaseManager) *ExportDataHandler {
	return &ExportDataHandler{dbManager: dbManager}
}

// Sets the application context to allow calling the runtime fuctions
func (controller *ExportDataHandler) SetContext(ctx context.Context) {
	controller.ctx = ctx
}

// Exports all data from the database into an excel file
func (handler *ExportDataHandler) ExportData() error {
	exportDataPath, err := handler.selectExportPath()

	fmt.Println(exportDataPath, err)
	// Check if the user clicked "Cancel" or provided an empty path
	if exportDataPath == "" {
		return nil
	}

	if err != nil {
		return err
	}

	file := excelize.NewFile()
	handler.withCoursesSheet(file)
	handler.withParticipantsSheet(file)
	handler.withMemberCardsSheet(file)
	handler.withParticipationHistorySheet(file)
	file.DeleteSheet("Sheet1")
	withHeaderStyle(file)

	if err := file.SaveAs(exportDataPath); err != nil {
		return err
	}

	return nil
}

func (handler *ExportDataHandler) withCoursesSheet(file *excelize.File) error {
	var courses []courses.Course
	if err := handler.dbManager.DB.Unscoped().Order("created_at desc").Find(&courses).Error; err != nil {
		return err
	}

	headers := []string{"Name", "Description", "Created_At", "Archived_At"}
	sheetName := "Courses"

	createSheet(headers, sheetName, file)

	for i, course := range courses {
		file.SetSheetRow(
			sheetName,
			toCell(i),
			&[]any{course.Name, course.Description, common.ToDateTime(course.CreatedAt), common.FormatDeletedAt(course.Deleted)},
		)
	}
	return nil
}

func (handler *ExportDataHandler) withParticipantsSheet(file *excelize.File) error {
	var participants []participants.Participant
	if err := handler.dbManager.DB.Unscoped().Order("created_at desc").Find(&participants).Error; err != nil {
		return err
	}
	sheetName := "Participants"
	headers := []string{"Name", "Surname", "Group", "Joined_At", "Archived_At"}
	createSheet(headers, sheetName, file)

	for i, participant := range participants {
		file.SetSheetRow(
			sheetName,
			toCell(i),
			&[]any{participant.Name,
				participant.Surname,
				participant.Group,
				common.ToDateTime(participant.CreatedAt),
				common.FormatDeletedAt(participant.Deleted)},
		)
	}

	return nil
}

func (handler *ExportDataHandler) withMemberCardsSheet(file *excelize.File) error {
	var memberCards []struct {
		ID        uint
		Capacity  uint
		Deleted   gorm.DeletedAt
		CreatedAt time.Time
		Name      string
		Surname   string
	}

	err := handler.dbManager.DB.Model(&participants.MemberCard{}).
		Select("member_cards.id, member_cards.capacity, member_cards.created_at, member_cards.deleted, participants.name, participants.surname").
		Joins("JOIN participants ON participants.id = member_cards.participant_id").
		Order("member_cards.created_at desc").
		Unscoped().
		Scan(&memberCards).
		Error

	if err != nil {
		return err
	}

	sheetName := "Member_Cards"
	headers := []string{"ID", "Name", "Surname", "Capacity", "Issued_At", "Depleted_At"}
	createSheet(headers, sheetName, file)

	for i, memberCard := range memberCards {
		file.SetSheetRow(
			sheetName,
			toCell(i),
			&[]any{
				memberCard.ID,
				memberCard.Name,
				memberCard.Surname,
				memberCard.Capacity,
				common.ToDateTime(memberCard.CreatedAt),
				common.FormatDeletedAt(memberCard.Deleted)},
		)
	}
	return nil
}

func (handler *ExportDataHandler) withParticipationHistorySheet(file *excelize.File) error {
	var attendanceHistory []struct {
		Name           string
		Surname        string
		Course         string
		MemberCard     string
		AttendedAt     time.Time
		AttendanceType string
	}

	err := handler.dbManager.DB.Model(&membercardattendance.MemberCardAttendance{}).
		Select("participants.name, participants.surname, courses.name as course, member_card_attendances.member_card_id as member_card, member_card_attendances.created_at as attended_at, member_card_attendances.attendance_type").
		Joins("JOIN participants on participants.id = member_card_attendances.participant_id").
		Joins("JOIN courses on courses.id = member_card_attendances.course_id").
		Order("member_card_attendances.created_at desc").
		Scan(&attendanceHistory).
		Error

	if err != nil {
		return err
	}

	sheetName := "Attendance_History"
	headers := []string{"Name", "Surname", "Course", "Card_ID", "Attended_At", "Attendance_Type"}
	createSheet(headers, sheetName, file)

	for i, attendanceHistoryEntry := range attendanceHistory {
		file.SetSheetRow(
			sheetName,
			toCell(i),
			&[]any{
				attendanceHistoryEntry.Name,
				attendanceHistoryEntry.Surname,
				attendanceHistoryEntry.Course,
				attendanceHistoryEntry.MemberCard,
				common.ToDateTime(attendanceHistoryEntry.AttendedAt),
				attendanceHistoryEntry.AttendanceType,
			},
		)
	}
	return nil
}

func (handler *ExportDataHandler) selectExportPath() (string, error) {
	timestamp := time.Now().Format(time.DateTime) // YYYY-MM-DD_HH-MM
	filename := fmt.Sprintf("protect_yourself_export_%s.xlsx", timestamp)

	return runtime.SaveFileDialog(handler.ctx, runtime.SaveDialogOptions{
		Title:           "Save File",
		DefaultFilename: filename,
	})
}

func createSheet(headers []string, sheetName string, file *excelize.File) {
	file.NewSheet(sheetName)
	file.SetSheetRow(sheetName, "A1", &headers)
}

func withHeaderStyle(file *excelize.File) error {
	style, err := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "#FFFFFF",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#10B981"},
		},
	})

	if err != nil {
		return err
	}

	for _, sheet := range file.GetSheetList() {
		cols, err := file.GetCols(sheet)
		if err != nil {
			return err
		}
		if len(cols) == 0 {
			continue // Skip empty sheets
		}

		// Apply style to the first row (headers)
		for colIdx := range cols {
			columnName, err := excelize.ColumnNumberToName(colIdx + 1) // Convert index to column name
			if err != nil {
				return err
			}
			// Apply style to header row: A1,B1,C1,etc
			err = file.SetCellStyle(sheet, columnName+"1", columnName+"1", style)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func toCell(index int) string {
	return fmt.Sprintf("A%d", index+2)
}
