
test:
	go test ./...

cover-html: cover
	go tool cover -html=c.out

cover:
	go test -v -cover -coverprofile=c.out ./...

