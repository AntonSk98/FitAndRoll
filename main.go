package main

import (
	"context"
	"embed"
	"fit_and_roll/backend/archive"
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/courseattendance"
	"fit_and_roll/backend/courses"
	"fit_and_roll/backend/import_"
	"fit_and_roll/backend/membercardattendance"
	"fit_and_roll/backend/participants"
	"fit_and_roll/backend/scaler"
	"fit_and_roll/backend/statistics"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	dbManager := config.NewDatabaseManager()
	dbManager.Migrate(
		&courses.Course{},
		&courses.ScheduleEntry{},
		&participants.Participant{},
		&participants.MemberCard{},
		&membercardattendance.MemberCardAttendance{},
		&scaler.Scaler{},
	)

	courseHandler := courses.NewCourseHandler(dbManager)
	participantsHandler := participants.NewParticipantsHandler(dbManager)
	importParticipantsHandler := import_.NewImportParticipantsHandler(dbManager, participantsHandler)
	memberCardHandler := participants.NewMemberCardHandler(dbManager)
	memberCardAttendanceHandler := membercardattendance.NewMemberCardAttendanceHandler(dbManager)
	courseAttendanceHandler := courseattendance.NewCourseAttendanceHandler(dbManager)
	exportDataHandler := archive.NewExportDataHandler(dbManager)
	archiveDataHandler := archive.NewArchivedDataHandler(dbManager)
	statisticsHandler := statistics.NewStatisticsHandler(dbManager)
	scaler := scaler.NewScalerHandler(dbManager)

	fileLogs := logger.NewFileLogger("logs.log")

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "PROTECT YOURSELF",
		MinWidth:         769,
		MinHeight:        500,
		WindowStartState: options.Maximised,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			exportDataHandler.SetContext(ctx)
			importParticipantsHandler.SetContext(ctx)
		},
		Logger: fileLogs,
		Bind: []interface{}{
			courseHandler,
			participantsHandler,
			memberCardHandler,
			memberCardAttendanceHandler,
			courseAttendanceHandler,
			exportDataHandler,
			archiveDataHandler,
			statisticsHandler,
			importParticipantsHandler,
			scaler,
		},
	})

	if err != nil {
		panic(err)
	}
}
