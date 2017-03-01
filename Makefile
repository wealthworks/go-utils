.SILENT :
.PHONY : vet clean dns_query


ROOF:=github.com/wealthworks/go-utils

all: vet

vet:
	echo "Checking ."
	go tool vet -atomic -bool -copylocks -nilfunc -printf -shadow -rangeloops -unreachable -unsafeptr -unusedresult .

clean:
	echo "Cleaning dist"
	rm -rf dist

dns_query:
	echo "Building $@"
	mkdir -p dist/linux_amd64 && GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/linux_amd64/$@ $(ROOF)/cmd/$@
	mkdir -p dist/darwin_amd64 && GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/darwin_amd64/$@ $(ROOF)/cmd/$@
.PHONY: dns_query
