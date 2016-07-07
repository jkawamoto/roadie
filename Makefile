default: build

.PHONY: asset
asset:
	go-bindata -pkg util -o command/util/assets.go assets/startup.sh

.PHONY: build
build: asset
	gox --output pkg/{{.Dir}}_{{.OS}}_{{.Arch}} -arch="amd64" -os="darwin linux windows"
