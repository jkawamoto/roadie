#!/bin/bash
#
# startup script for queue worker.
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

# This script starts fluentd for logging, and then starts queue-manager.
# After the manager finishes tasks, this script shutdowns the VM.
#
cd /root

# Start logging.
if [[ -n $(docker ps -a | grep fluentd) ]]; then
  docker rm -f fluentd
fi
for i in $(seq 10); do
  docker run -d --name fluentd \
    -e 'INSTANCE={{.InstanceName}}' -e 'USERNAME=roadie' \
    -v /var/lib/docker:/var/lib/docker jkawamoto/docker-google-fluentd
  sleep 30s
  if [[ -n $(docker ps -a | grep fluentd) ]]; then
    break
  fi
done
sleep 30s

# Prepare Roadie Queue Manager.
readonly FILENAME='roadie-queue-manager_{{.Version}}_linux_amd64'
wget https://github.com/jkawamoto/roadie-queue-manager/releases/download/v{{.Version}}/${FILENAME}.tar.gz
tar -zxvf ${FILENAME}.tar.gz
cd ${FILENAME}

# Run the script
./roadie-queue-manager {{.ProjectID}} {{.Name}}

# Shutdown
docker run -i jkawamoto/roadie-gcp shutdown
