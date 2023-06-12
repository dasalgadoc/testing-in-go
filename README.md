<h1 align="center">
  🚀 🐹 Testing techniques in Go 🐹 🚀 
</h1>

<p align="center">
    <a href="#"><img src="https://img.shields.io/badge/technology-go-blue.svg" alt="Go"/></a>
</p>

<p align="center">
  This repository contains a project to explore basic test techniques in Go.
</p>

__Notes__: 
- All test are passing, even the bad ones.
- To explore similar techniques in Java, please check [this repository](https://github.com/dasalgadoc/testing-in-java)

## 🧲 Environment Setup

### 🛠️ Needed tools

1. Go 1.18 or higher
2. Docker and docker compose are recommended but not required

### 🏃🏻 Application execution

1. Make sure to download all Needed tools
2. Clone the repository
```bash
git clone https://github.com/dasalgadoc/testing-in-go.git
```
3. Build up go project
```bash
go mod download
go get .
```
4. Run tests
```bash
go test ./...
```
5. Enjoy! 😎

## 📚 References

- [Testify](https://github.com/stretchr/testify)
- [Docker test](https://github.com/ory/dockertest)
- [Test Containers](https://github.com/testcontainers/testcontainers-go)
- [Test with Cucumber](https://github.com/cucumber/godog)
- [Gin](https://github.com/gin-gonic/gin)

## 📝 Useful commands

- Go test commands

Run all
```bash
go test -v ./...
```

Get coverage
```bash
go test -cover
```

Profiling tests
```bash
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

More profiling
```bash
go test -cpuprofile=cpu.out
go tool pprof cpu.out
```

- Want to check your dockers?
```bash
docker ps

docker exec -it <DOCKER_ID/NAME> bash
mysql -u root -p

docker ps | awk '{print $1}' | xargs docker stop
```

- Makefile for CI pipeline flows
```bash
make
make run-tests
```

