build:
	@make tests && \
	rm ./bin/* && \
	go build -o ./bin/gsa && \
	GOOS=windows GOARCH=amd64 go build -o ./bin/gsa.exe

tests:
	@cd sorting && go test -run TestAll

benchmarks:
	@cd sorting && go test -run=" " -benchmem -bench=./...

runtest:
	@cd sorting && go test -run $(t)

runbench:
	@go test -run=" " -benchmem -bench=$(b)

clitest:
	@go run . -s "100 1000 10000" -o "test.csv test.json"

example:
	@go run . -s "100 1000 10000 100000" -o "_example/results.csv _example/results.json"

servedoc:
	@godoc -http=:8080