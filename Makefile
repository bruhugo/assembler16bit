
.PHONY: test

test:
	go test ./...

cover-html: cover
	go tool cover -html=c.out

cover:
	go test -v -cover -coverprofile=c.out ./...

bench:
	TEST_PROGRAM=../programs/test_all go test -count=5 -v -bench=. ./...
