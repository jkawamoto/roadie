#
# Makefile.go
#
# Copyright (c) 2016-2017 Junpei Kawamoto
#
# This file is part of Roadie.
#
# Roadie is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Roadie is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Roadie.  If not, see <http://www.gnu.org/licenses/>.
#
VERSION = snapshot
GHRFLAGS =

.PHONY: asset build release get-deps test
default: build

asset: get-deps
	rm assets/assets.go
	go-bindata -pkg assets -o assets/assets.go -nometadata assets/*

build: asset
	goxc -os="darwin linux windows" -d=pkg -pv=$(VERSION)

release:
	ghr  -u jkawamoto $(GHRFLAGS) v$(VERSION) pkg/$(VERSION)

get-deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get -d -t -v .

test: asset
	go test -v ./... -tags=dummy

local-test: asset
	go run *.go -c gcp.cfg.yml --verbose run sss.script.yml --name sss --follow

queue-test: asset
	go run *.go -c gcp.cfg.yml --verbose queue add sssq sss.script.yml --name sssq-1
	go run *.go -c gcp.cfg.yml --verbose queue add sssq sss.script.yml --name sssq-2
	go run *.go -c gcp.cfg.yml --verbose queue add sssq sss.script.yml --name sssq-3
	go run *.go -c gcp.cfg.yml --verbose queue add sssq sss.script.yml --name sssq-4
