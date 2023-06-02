package application

import (
	"dasalgadoc.com/go-testing/02-integration/infrastructure"
	"database/sql"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

const (
	user     = "root"
	password = "password"
	host     = "localhost"
	database = "test"

	// I use a unix socket to connect to docker from colima. A default docker installation will use empty string.
	poolEndpoint = "unix:///Users/diesalgado/.colima/docker.sock"
)

type studentCreatorIntegrationTest struct {
	studentName     string
	studentAge      int
	err             error
	mysqlRepository *infrastructure.MysqlStudentRepository
	test            *testing.T
}

func TestStudentCreator(t *testing.T) {
	s := startServiceIntegrationTest(t)

	pool, resource, _ := setupMySQLContainer(t)
	defer teardown(t, resource, pool)

	s.givenAStudent("John Doe", 31)
	s.andMySQLRepository(resource.GetPort("3306/tcp"))
	s.whenServiceIsInvoked()

	assert.NoError(t, s.err, "Error should be nil")
}

func setupMySQLContainer(t *testing.T) (pool *dockertest.Pool, resource *dockertest.Resource, db *sql.DB) {
	pool, err := dockertest.NewPool(poolEndpoint)
	assert.NoError(t, err, "Could not create the docker pool")

	resource, err = pool.Run("mysql", "5.7",
		[]string{
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", password),
			fmt.Sprintf("MYSQL_DATABASE=%s", database),
		})
	assert.NoError(t, err, "Could not start MySQL container")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, resource.GetPort("3306/tcp"), database)
	err = pool.Retry(func() error {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			return err
		}
		return db.Ping()
	})
	assert.NoError(t, err, "Could not connect to docker")

	script, err := os.ReadFile("../../sql/init.sql")
	assert.NoError(t, err, "Could not read init.sql file")

	queries := strings.Split(string(script), "\n")
	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Error executing query: %v", err)
		}
	}

	assert.NoError(t, err, "Could not execute init.sql file")

	return pool, resource, db
}

func teardown(t *testing.T, resource *dockertest.Resource, pool *dockertest.Pool) {
	err := pool.Purge(resource)
	assert.NoError(t, err, "Could not purge resource")
}

func startServiceIntegrationTest(t *testing.T) *studentCreatorIntegrationTest {
	t.Parallel()
	return &studentCreatorIntegrationTest{
		test: t,
	}
}

func (s *studentCreatorIntegrationTest) givenAStudent(name string, age int) {
	s.studentName = name
	s.studentAge = age
}

func (s *studentCreatorIntegrationTest) andMySQLRepository(port string) {
	s.mysqlRepository = infrastructure.NewMysqlStudentRepository(user, password, host, port, database)
}

func (s *studentCreatorIntegrationTest) whenServiceIsInvoked() {
	target := NewStudentCreator(s.mysqlRepository)
	s.err = target.CreateStudent(s.studentName, s.studentAge)
}
