package entrypoints

import (
	"dasalgadoc.com/go-testing/03-acceptance/application"
	"dasalgadoc.com/go-testing/03-acceptance/infrastructure/database"
	"database/sql"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"github.com/ory/dockertest/v3"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

const (
	user     = "root"
	password = "password"
	host     = "localhost"
	dbName   = "test"

	// I use a unix socket to connect to docker from colima. A default docker installation will use empty string.
	poolEndpoint = "unix:///Users/diesalgado/.colima/docker.sock"
)

type studentGetterTestScenario struct {
	router   *gin.Engine
	resource *dockertest.Resource
	pool     *dockertest.Pool
	writer   *httptest.ResponseRecorder
}

func FeatureContext(ctx *godog.ScenarioContext) {
	s := startStudentTestScenario()
	if err := s.setup(); err != nil {
		fmt.Println("Error setting up the scenario")
		panic(err)
	}
	defer s.teardown()

	ctx.Step(`^I send a GET request to "([^"]*)"`, s.iSendAGETRequestTo)
	ctx.Step(`^the response status should be "([^"]*)"`, s.theResponseStatusShouldBe)
	ctx.Step(`^the response body should be a "([^"]*)" object`, s.theResponseBodyShouldBeA)
}

func startStudentTestScenario() *studentGetterTestScenario {
	return &studentGetterTestScenario{}
}

func (s *studentGetterTestScenario) setup() error {
	s.pool, s.resource, _ = setupMySQLContainer()

	repository := database.NewMysqlStudentRepository(user, password, host, s.resource.GetPort("3306/tcp"), dbName)
	useCase := application.NewStudentSearcher(repository)
	studentGet := NewStudentGetter(useCase)

	s.router = gin.Default()
	s.router.GET("/students/:student_id", studentGet.Get)

	return nil
}

func (s *studentGetterTestScenario) teardown() {
	if err := teardownDatabaseContainer(s.resource, s.pool); err != nil {
		fmt.Println("Error tearing down the scenario")
		panic(err)
	}
}

func setupMySQLContainer() (pool *dockertest.Pool, resource *dockertest.Resource, db *sql.DB) {
	pool, err := dockertest.NewPool(poolEndpoint)
	if err != nil {
		log.Fatalf("Could not create the docker pool: %s", err)
		panic(err)
	}

	resource, err = pool.Run("mysql", "5.7",
		[]string{
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", password),
			fmt.Sprintf("MYSQL_DATABASE=%s", dbName),
		})
	if err != nil {
		log.Fatalf("Could not start MySQL container: %s", err)
		panic(err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, resource.GetPort("3306/tcp"), dbName)
	err = pool.Retry(func() error {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			return err
		}
		return db.Ping()
	})
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
		panic(err)
	}

	script, err := os.ReadFile("../../sql/init.sql")
	if err != nil {
		log.Fatalf("Could not read init.sql file: %s", err)
		panic(err)
	}

	queries := strings.Split(string(script), "\n")
	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Error executing query: %v", err)
		}
	}

	script, err = os.ReadFile("../../sql/insert.sql")
	if err != nil {
		log.Fatalf("Could not read insert.sql file: %s", err)
		panic(err)
	}

	insertion := strings.Split(string(script), "\n")
	for _, query := range insertion {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Error executing query: %v", err)
		}
	}

	if err != nil {
		log.Fatalf("Could not execute init.sql file: %s", err)
		panic(err)
	}

	return pool, resource, db
}

func teardownDatabaseContainer(resource *dockertest.Resource, pool *dockertest.Pool) error {
	err := pool.Purge(resource)
	if err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	return err
}

func (s *studentGetterTestScenario) iSendAGETRequestTo(endpoint string) error {
	s.writer = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", endpoint, nil)
	s.router.ServeHTTP(s.writer, req)

	return nil
}

func (s *studentGetterTestScenario) theResponseStatusShouldBe(status int) error {
	if s.writer.Code != status {
		return fmt.Errorf("expected response status to be %d but was %d", status, s.writer.Code)
	}
	return nil
}

func (s *studentGetterTestScenario) theResponseBodyShouldBeA(responseType string) error {
	if s.writer.Body.String() != responseType {
		return fmt.Errorf("expected response body to be %s but was %s", responseType, s.writer.Body.String())
	}
	return nil
}
