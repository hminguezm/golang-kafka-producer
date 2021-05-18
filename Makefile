config:
	@echo "Environment file generating"
	@cp example.env .env
	@echo "Init..."
	@find . -type f -print0 | xargs -0 sed -i 's/wrk-connector/$(project-name)/gI'
	@rm -rf .idea .git buildspec.yml # delete .idea file, for resolve issue in editor
	@echo "Downloading modules..."
	@go mod download
	@go mod vendor
	@echo "Done"

start:
	@echo "Starting..."
	@go run src/main.go


