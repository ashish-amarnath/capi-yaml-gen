test:
	go test ./...

update-golden-files:
	go test ./... -update
