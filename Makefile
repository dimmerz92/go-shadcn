.SILENT:

server:
	air

css:
	watchexec -e css -i pkg/themes/output.css lightningcss \
		--bundle pkg/themes/input.css -o pkg/themes/output.css

deps:
	go install github.com/air-verse/air@latest && \
		go install github.com/a-h/templ/cmd/templ@latest && \
		go mod tidy

dev:
	make -j2 css server
