#
# Makefile.go
#
# Copyright (c) 2016 Junpei Kawamoto
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
# along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
#
VERSION = snapshot

default: build


.PHONY: asset
asset:
	rm assets/assets.go
	go-bindata -pkg assets -o assets/assets.go -nometadata assets/*


.PHONY: build
build: asset
	goxc -os="darwin linux windows" -d=pkg -pv=$(VERSION)


.PHONY: release
release:
	ghr  -u jkawamoto  v$(VERSION) pkg/$(VERSION)


.PHONY: get-deps
get-deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get -d -t -v .

.PHONY: test
test: asset
	go test -v ./...
