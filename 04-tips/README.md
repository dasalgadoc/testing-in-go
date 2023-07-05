# 🛠️ Tips

## 1️⃣ Exclude testing

To exclude some files from testing, you can use the follow comments at the top of the file.

```go
//go:build exclude
// +build exclude
```

This will exclude the file from testing when running `go test ./...`.

But in some cases, its necessary to exclude ONE test from file and no the entire file. To do so, you can use the `t.SkipNow()` function.

```go
func TestSomething(t *testing.T) {
    t.SkipNow()
}
```

## 2️⃣ Parallel testing

Run test in parallel it is a good practice to improve the performance of your test suite. 
To do so, you can use the `-p` flag when running `go test` or add the `t.Parallel()` function in your test.

```go 
func TestSome(t *testing.T) {
    t.Parallel()
}
```

## 3️⃣ Assert Equals and assert same

When testing, you may need to compare two values. To do so, you can use the `assert` package from `testify`.

- Assert equals will compare the values and return an error if they are not equal.
- Assert same will compare the pointers and return an error if they are not the same including its reference.


## 4️⃣ Makefile for testing

You can create a Makefile to run your tests. This will help you to run your tests faster, with less typing and enables the integration with CI/CD tools.

```makefile
run-test:
    @echo "Running tests..."
    @go test -v ./... -tags '!exclude'
    @echo "Done."
```