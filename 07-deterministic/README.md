# 👨🏻‍🌾 Good Practices

## 📋 Table of contents
1. [Test who randomly fails](#Test who randomly fails)
2. [Deterministic tests](#Deterministic tests)
3. [SRP](#SRP in test)
4. [Given-When-Then](#Given-When-Then)
5. [TCR](#TCR)

7. [Githooks](#Githooks)
8. [Guidelines](#Guidelines)

## 😵‍💫 Test who randomly fails

In some suites, test can fail randomly. This is a problem because it can make the team lose confidence in the test suite.

The usage of global variables and singletons is one of the most frequent sources of this issue, but another one is when a test depends on the order in which it is executed. For instance, one test may write information to a database while another might retrieve it.

A smart technique to prevent this issue is to manage isolation between tests, manage the state before and after each test, and run the tests in random order.

### 💦 Clean state before and after each test

If you're familiar with @Before and @After annotations in Java, we can emulate this behavior in Go using functions.

```go
func TestSome(m *testing.M) {
    before()
    defer after()
        // ... test code
}
```

### 🤷🏻‍♀️ Run tests in random order

One more time, in Java we count with @TestMethodOrder(MethodOrdered.Random.class) annotation, but in Go we rely on command to run tests in random order.

```bash
go test -shufle
```

## 🕐 Deterministic tests

When we run a test, we expect the same result every time. If we don't get the same result, we can't trust the test.
If a feature depends on current date or time, we must to isolates the test from the system clock and mock it.

The production code might be:
```go
var CurrentTime = func() time.Time {
    return time.Now()
}
```

But test code should be:
```go
CurrentTime = func() int {
    // mock the system clock
}
```

And CurrentTime will be replaced by the mock function.

## 👁️‍🗨️ SRP in test

The Single Responsibility Principle (SRP) states that a class should have only one reason to change. In other words, a class should have only one job.

In test, we can apply the same principle. [A test should have only one reason to fail](https://twitter.com/unclebobmartin/status/1079396026996441088). If a test fails for more than one reason, it will be difficult to identify the root cause of the failure.

## 🔋 Given-When-Then

Also known as, Arrange-Act-Assert, is a pattern that helps us to write tests in a clear and understandable way.

This consists of three steps:
1. Given: We prepare the data for the test.
2. When: We execute the code we want to test.
3. Then: We verify the result.

If we isolate these three steps in functions, we will gain some benefits.

- Writing tests that are easy to read and understand
- Reusing the same code in different tests
- Encapsulating the test logic

## 🪓 TCR

TCR stands for `Test && Commit || Revert`. It is radical a technique that helps us to commit reliable code to the repo.

The idea is simple: If the tests pass, the code is committed. If the tests fail, the code is reverted.

TCR can be implemented through aliases in the terminal.

```bash
alias test='go test'
alias commit='git commit -a'
alias revert='git reset --hard'

alias tcr='test && commit || revert'
```

## 🧠 Githooks

[Githooks](https://git-scm.com/docs/githooks) are scripts that Git executes before or after events such as: commit, push, and receive. Git hooks are a built-in feature - no need to download anything. Git hooks are run locally.

See the .pre-commit file in this repo.

## 🎲 Guidelines

Some guidelines in test are:

- Do not do refactor with broken test
- Production code and test have to be separated.
- Do not commit broken test or features that not pass the test.
