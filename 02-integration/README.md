# ✌🏻 Integration Testing 

_"Integration tests determine if independently developed units of software work correctly when they are connected to each other"_
[Article](https://martinfowler.com/bliki/IntegrationTest.html)

Since it calls for the system to be executed as a whole, this kind of test is more complicated than unit testing. Additionally, it takes longer and costs more to maintain.

In clean architecture, the use cases and domain models are combined with a concrete infrastructure implementation.

## 📝 How to write integration tests in Go:

1. Create a file with the suffix `_test.go` in the same directory as the file you want to test.
2. Import the `testing` package.
   
   - [Testing package](https://golang.org/pkg/testing/)
   - [Testify](https://github.com/stretchr/testify)
3. Create a strategy to run concrete infrastructure implementations.
    - [Docker test](https://github.com/ory/dockertest)
4. With Docker you can build concrete and ephemere infrastructure to control the test.

## 🏃🏻‍♀️ Considerations and Good practices

- Integration test would be fast.
- Integration test should be isolated.
- Integration test should be repeatable.
- Integration test should be self-validating.
- Integration test should be timely.
- Integration test are not necessarily exhaustive.

# 🔺 Test Pyramid

[Article](https://martinfowler.com/articles/practical-test-pyramid.html)

![PyramidTest](../imgs/PyramidTest.jpg)

Unit test are faster and cheaper to maintain than integration test and integration test are faster and cheaper to maintain than end-to-end test.

Unit test has minor coverage than integration test and integration test has minor coverage than end-to-end test.
