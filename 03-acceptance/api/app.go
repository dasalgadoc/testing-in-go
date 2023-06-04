package api

import (
	"dasalgadoc.com/go-testing/03-acceptance/application"
	"dasalgadoc.com/go-testing/03-acceptance/configs"
	"dasalgadoc.com/go-testing/03-acceptance/infrastructure/database"
	"dasalgadoc.com/go-testing/03-acceptance/infrastructure/entrypoints"
	"dasalgadoc.com/go-testing/03-acceptance/repository"
)

type Application struct {
	Config configs.Config

	Repositories struct {
		studentRepository repository.StudentRepository
	}

	UseCases struct {
		useSearchStudent application.StudentSearcher
	}

	Controllers struct {
		Ping       entrypoints.Ping
		GetStudent entrypoints.StudentGetter
	}
}

func BuildApplication() *Application {
	app := Application{}

	app.getConfiguration()
	app.buildRepositories()
	app.buildUseCases()
	app.buildControllers()

	return &app
}

func (app Application) getConfiguration() {
	appConfig, err := configs.LoadConfig("./03-acceptance/configs/config.yaml")
	if err != nil {
		panic(err)
	}

	app.Config = appConfig
}

func (app Application) buildRepositories() {
	app.Repositories.studentRepository = database.NewMysqlStudentRepository(
		app.Config.DB.User,
		app.Config.DB.Password,
		app.Config.DB.Host,
		app.Config.DB.Port,
		app.Config.DB.Database)
}

func (app Application) buildUseCases() {
	app.UseCases.useSearchStudent = application.NewStudentSearcher(app.Repositories.studentRepository)
}

func (app Application) buildControllers() {
	app.Controllers.Ping = entrypoints.NewPing()
	app.Controllers.GetStudent = entrypoints.NewStudentGetter(app.UseCases.useSearchStudent)
}
