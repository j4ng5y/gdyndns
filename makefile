all: test build

.PHONY: test
test:
	@go test -v ./... -race

.PHONY: build
build:
	@go build -a -ldflags='-extldflags "-static"' gdyndns.go

.PHONY: install
install: ./gdyndns
	@mv ./gdyndns /usr/bin/gdyndns
	@chmod +x /usr/bin/gdyndns

.PHONY: install-service
install-service: gdyndns.service gdyndns.timer
	@mv ./gdyndns.* /etc/systemd/system/
	@sudo systemctl daemod-reload
	@sudo systemctl enable gdyndns.timer
	@sudo systemctl start gdyndns.timer