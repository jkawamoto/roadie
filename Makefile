default: build


.PHONY: build
build:
	go-bindata -pkg util -o util/assets.go assets/startup.sh
	gox --output pkg/{{.Dir}}_{{.OS}}_{{.Arch}} -arch="amd64" -os="darwin linux windows"
