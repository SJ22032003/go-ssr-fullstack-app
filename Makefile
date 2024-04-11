run:
	@air -c .air.toml

templ:
	@templ generate -watch -proxy=http://localhost:8080

tailwind:
	@tailwindcss -i ./static/css/styles.css -o ./static/css/output.css --watch


install:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod vendor
	@npm install -D tailwindcss
