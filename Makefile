format:
	@find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/pinpointdev90/go-crud-api/tree/main"
verify:
	@staticcheck ./... && go vet ./...
unit-test:
	@go test ./... -coverprofile=./test_results/cover.out && go tool cover -html=./test_results/cover.out -o ./test_results/cover.html
serve:
	@docker-compose -f build/docker-compose.yml up