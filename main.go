package main

import (
	"context"
	"embed"
	"fit_and_roll/backend/archive"
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/courseattendance"
	"fit_and_roll/backend/courses"
	"fit_and_roll/backend/membercardattendance"
	"fit_and_roll/backend/participants"
	"fit_and_roll/backend/statistics"

	"github.com/wailsapp/wails/v2"
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
	)

	courseHandler := courses.NewCourseHandler(dbManager)
	participantsHandler := participants.NewParticipantsHandler(dbManager)
	memberCardHandler := participants.NewMemberCardHandler(dbManager)
	memberCardAttendanceHandler := membercardattendance.NewMemberCardAttendanceHandler(dbManager)
	courseAttendanceHandler := courseattendance.NewCourseAttendanceHandler(dbManager)
	exportDataHandler := archive.NewExportDataHandler(dbManager)
	archiveDataHandler := archive.NewArchivedDataHandler(dbManager)
	statisticsHandler := statistics.NewStatisticsHandler(dbManager)

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "PROTECT YOURSELF",
		Width:     1024,
		MinWidth:  769,
		Height:    768,
		MinHeight: 500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			exportDataHandler.SetContext(ctx)
		},
		Bind: []interface{}{
			courseHandler,
			participantsHandler,
			memberCardHandler,
			memberCardAttendanceHandler,
			courseAttendanceHandler,
			exportDataHandler,
			archiveDataHandler,
			statisticsHandler,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
