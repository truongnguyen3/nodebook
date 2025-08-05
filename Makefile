.PHONY: deps build install deps-frontend build-frontend embed-frontend

deps:
	# No longer needed - embed is built into Go

embed-frontend:
	# No longer needed - embed happens at compile time

build:
	go build -o dist/nodebook .

install:
	go install .
	@echo "nodebook built and installed."

deps-frontend:
	cd src/frontend && npm i

build-frontend:
	cd src/frontend && npm run build
	rm -Rf dist/frontend/*.map
