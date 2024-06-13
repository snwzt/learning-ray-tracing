build:
	@go build -o bin/rtg .

run:
	@./bin/rtg

clean:
	@rm -rf bin