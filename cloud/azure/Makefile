#
# cloud/azure/Makefile
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
.PHONY: build go-deps

build: storage resource compute subscriptions network disk batchmanagement batch
	go get -d -t -v .

go-deps:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

resource: go-deps
	swagger generate client -f https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-resources/resources/2016-09-01/swagger/resources.json -t resource

compute: go-deps
	swagger generate client -f https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-compute/2016-04-30-preview/swagger/compute.json -t compute

subscriptions: go-deps
	swagger generate client -f https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-resources/subscriptions/2016-06-01/swagger/subscriptions.json -t subscriptions

batch: go-deps
	swagger generate client -f https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/batch/2016-07-01.3.1/swagger/BatchService.json -t batch
