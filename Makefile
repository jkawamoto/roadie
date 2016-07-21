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
	go-bindata -pkg util -o command/util/assets.go assets/startup.sh


.PHONY: build
build: asset
	goxc -os="darwin linux windows" -d=pkg -pv=$(VERSION)


.PHONY: release
release:
	ghr  -u jkawamoto  v$(VERSION) pkg/$(VERSION)


.PHONY: get-deps
get-deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get -d github.com/ttacon/chalk
	go get -d github.com/urfave/cli
	go get -d github.com/gosuri/uitable
	go get -d github.com/deiwin/interact
	go get -d github.com/mitchellh/mapstructure
	go get -d github.com/briandowns/spinner
	go get -d github.com/naoina/toml
	go get -d golang.org/x/net/context
	go get -d golang.org/x/oauth2/google
	go get -d google.golang.org/api/compute/v1
	go get -d google.golang.org/api/storage/v1
	go get -d google.golang.org/api/logging/v2beta1
	go get -d gopkg.in/yaml.v2
	go get -d github.com/jkawamoto/pb # Use `public_pool_add` branch.
	cd $(GOPATH)/src/github.com/jkawamoto/pb && \
		git checkout -b public_pool_add origin/public_pool_add && cat pool.go

.PHONY: test
test: asset
	go test -v ./...
