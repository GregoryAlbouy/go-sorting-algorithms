build:
	@make tests && \
	rm ./bin/* && \
	go build -o ./bin/gsa && \
	GOOS=windows GOARCH=amd64 go build -o ./bin/gsa.exe

tests:
	@cd algorithms && go test -run TestAll

benchmarks:
	@cd algorithms && go test -run=" " -bench BenchmarkAll

runtest:
	@cd algorithms && go test -run $(t)

runbench:
	@cd algorithms && go test -run=" " -bench $(b)

quickcompare:


example:
	@go run . -s "100 1000 10000 100000" -output "_example/results.csv _example/results.json"

servedoc:
	@godoc -http=:8080