OS := $(shell uname)

TAGS = ""
BUILD_FLAGS = "-v"

SERVICE_NAME = char7
BUILD_NAME = char7
RELEASE_ROOT = release

COVERAGE_FOLDER = coverage

# DON'T EDIT THE BELOW
RELEASE_SERVICE = $(RELEASE_ROOT)/$(SERVICE_NAME)
NOW = $(shell date -u '+%Y%m%d%I%M%S')
GOVET = go tool vet -composites=false -methods=false -structtags=false

all: build

check: test

dist: release

web: build
	./$(BUILD_NAME)

govet:
	$(GOVET) main.go
	$(GOVET) cmd modules

build: $(GENERATED)
	go install $(BUILD_FLAGS) -ldflags '-s -w $(LDFLAGS)' -tags '$(TAGS)'
	cp '$(GOPATH)/bin/$(BUILD_NAME)' .

build-dev: $(GENERATED) govet
	go install $(BUILD_FLAGS) -ldflags '$(LDFLAGS)' -tags '$(TAGS)'
	cp '$(GOPATH)/bin/$(BUILD_NAME)' .

build-dev-race: $(GENERATED) govet
	go install $(BUILD_FLAGS) -race -tags '$(TAGS)'
	cp '$(GOPATH)/bin/$(BUILD_NAME)' .

pack:
	rm -rf $(RELEASE_SERVICE)
	mkdir -p $(RELEASE_SERVICE)
	cp -r $(BUILD_NAME) $(RELEASE_SERVICE)
	strip -x $(RELEASE_SERVICE)/$(BUILD_NAME) -o $(RELEASE_SERVICE)/$(BUILD_NAME)
	cd $(RELEASE_ROOT) && zip -r $(SERVICE_NAME).$(NOW).zip $(SERVICE_NAME)

clean-coverage:
	rm -rf $(COVERAGE_FOLDER)
	find . -name cover.out -delete
	find . -name coverage.xml -delete

coverage: clean-coverage
	mkdir -p $(COVERAGE_FOLDER)
	./coverage.sh

release: clean clean-mac fixme todo build pack

clean:
	go clean -i ./...

clean-mac: clean
	find . -name ".DS_Store" -print0 | xargs -0 rm -f

test:
	go test -cover -race ./...

sonarscaner: coverage lint
	sonar-scanner

clean-test-coverage:
	rm -rf $(TEST_FOLDER) | true

clean-lint:
	rm -rf $(LINT_FOLDER) | true

lint: clean-lint
	mkdir -p $(LINT_FOLDER)
	$(GOPATH)/bin/gometalinter.v1 --vendor --checkstyle --deadline=5m ./... > $(LINT_FOLDER)/report.xml | true

fixme:
	grep -rnw "FIXME" transponder &

todo:
	grep -rnw "TODO" transponder &
