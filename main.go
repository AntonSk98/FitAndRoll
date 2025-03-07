package main

import (
	"embed"
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/courses"
	"fit_and_roll/backend/participants"

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
		&participants.CourseAttendance{},
	)

	app := NewApp()
	coursesController := courses.NewCoursesController(dbManager)
	participantsController := participants.NewParticipantsController(dbManager)
	memberCardController := participants.NewMemberCardController(dbManager)
	courseParticipationController := participants.NewMemberCardCourseParticipationController(dbManager)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "fit_and_roll",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			coursesController,
			participantsController,
			memberCardController,
			courseParticipationController,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
