# 🖐🏻 End-To-End Test

Also known as, acceptance test, functional test, black-box test, UI test, or browser test.

End-To-End test tries to simulate the real user behavior, it is the most expensive and slowest test to run. This will use a concrete infrastructure implementation and users interfaces.

## 📝 How to write end-to-end tests in Go:

1. Create a file with the suffix `_test.go` in the same directory as the file you want to test.
2. Import the `testing` package.
   
   - [Testing package](https://golang.org/pkg/testing/)
   - [Testify](https://github.com/stretchr/testify)
3. Create a strategy to run concrete infrastructure implementations.
    - [Docker test](https://github.com/ory/dockertest)
4. With Docker you can build concrete and ephemere infrastructure to control the test.
5. Use some BDD framework to write the test.
    - [Cucumber](https://cucumber.io/)
    - [Godog - Go Cucumber framework](https://github.com/cucumber/godog)
6. Write an .feature file with the test scenarios.
7. Structure the test in the following way:

```go
func TestFeatures(t *testing.T) {
    suite := godog.TestSuite{
            ScenarioInitializer: FeatureGetStudentOk, // The function that will initialize the scenario
            Options: &godog.Options{
            Format:   "pretty",
            Paths:    []string{"../../features"}, // The path where the .feature files are located
            TestingT: t,
        },
    }
	
    if suite.Run() != 0 {
        t.Fatal("non-zero status returned, failed to run feature tests")
    }
}

func FeatureGetStudentOk(ctx *godog.ScenarioContext) {
    s := startStudentTestScenario()
    if err := s.setup(); err != nil {
        fmt.Println("Error setting up the scenario")
        panic(err)
    }
    //defer s.teardown()
    
    // The steps that will be executed in the scenario
    ctx.Step(`^I send a GET request to '([^']*)'`, s.iSendAGETRequestTo)
    ctx.Step(`^the response status should be (\d+)$`, s.theResponseStatusShouldBe)
}
```

## 🏃🏻‍♀️ Considerations and Good practices

### 🤓 General

- End-To-End test would be slow.
- End-To-End test should be isolated.
- End-To-End test should be repeatable.
- End-To-End test should be self-validating.
- End-To-End test should be timely.

### 😎 Design

- End-To-End test are not necessarily exhaustive.
- End-To-End test should be written in a declarative way.
- End-To-End test should be written in a business language.
- End-To-End might include a happy path and some edge cases. 
