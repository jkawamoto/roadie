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
	goxc -arch="amd64" -os="darwin linux windows" -d=pkg -pv=$(VERSION)