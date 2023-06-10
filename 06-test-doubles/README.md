# 👯‍♀️ Test Doubles

Unit tests run without external dependencies, only taking effort for uses cases and domain. However, sometimes you need to test code that depends on external resources, such as a database or a web service. In these cases, you can use test doubles to replace the external resources. Test doubles are objects that replace real objects in your code while simulating the behavior of the real objects.

The test doubles make our test faster, less coupling and more isolated.

We can simply change the implementation of the interface to use a test double instead of the real implementation, applying __DIP__.

There are several types of test doubles:

## 🧙🏻‍♀️ Fakes

Fakes are objects that have working implementations, but not the same as production one. Usually, they take some shortcut and have simplified versions of production code. For example, an in-memory database is a fake database because it has the same interface as a real database, but it doesn’t persist data between tests.

```go
type FakeDatabase struct {
    data map[string]string
}
```

## 🚬 Stubs

Stubs provide canned answers to calls made during the test, usually not responding at all to anything outside what’s programmed in for the test. Stubs may also record information about calls.

```go
func (s *stubDatabase) Get(key string) (string, error) {
    return "fixed value", nil
}
```

## 🙃 Dummies

Dummies like stubs has a fixed behaviour but no implementation. It's just a empty object.

```go
func (s *dummyDatabase) Get(key string) (string, error) {
    return "", nil
}
```

## 🕵🏻‍♀️ Spies

Spies are stubs that also record some information based on how or how many times they were called.

```go
type SpyDatabase struct {
    Called      bool
    CalledTimes int
}

func (s *spyDatabase) Get(key string) (string, error) {
    s.Called = true
    s.CalledTimes++	
	// ... Stub implementation
}
```

## 🧑🏻‍🔬 Mocks

Mocks are pre-programmed with expectations which form a specification of the calls they are expected to receive. They can throw an exception if they receive a call they don’t expect and are checked during verification to ensure they got all the calls they were expecting.

In `testify`, we can use the `mock` package to create mocks.

[Documentation](https://pkg.go.dev/github.com/stretchr/testify/mock)