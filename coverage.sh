#!/usr/bin/env bash

for pkg in $(go list ./...); do
    name=$(echo "$pkg" | rev | cut -d'/' -f1 | rev)
    go test -coverprofile=$(pwd)/coverage/${name}.cover.out ${pkg}
done

echo "mode: set" > coverage/coverage.out && cat $(pwd)/coverage/*.cover.out | grep -v mode: | sort -r | awk '{if($1 != last) {print $0;last=$1}}' >> $(pwd)/coverage/coverage.out
go tool cover -html=$(pwd)/coverage/coverage.out -o $(pwd)/coverage/coverage.html

# If exist gocov application
if type "$GOPATH/bin/gocov" > /dev/null; then
  $GOPATH/bin/gocov convert coverage/coverage.out | $GOPATH/bin/gocov-xml > coverage/coverage.xml
fi