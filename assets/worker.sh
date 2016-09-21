#!/bin/bash
#
# startup script for queue worker
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
cd /root

# Start logging.
for i in `seq 5`
do
  if [ -n "`docker ps -a | grep fluentd`" ]; then
    docker rm -f fluentd
  fi
  docker run -d --name fluentd -e "INSTANCE={{.Name}}" -e "USERNAME=roadie" \
    -v /var/lib/docker:/var/lib/docker jkawamoto/docker-google-fluentd \
    || break
done
sleep 30s

# Prepare Roadie Queue Manager.
FILENAME=roadie-queue-manager_{{.Version}}_linux_amd64
wget https://github.com/jkawamoto/roadie-queue-manager/releases/download/v0.1.1/${FILENAME}.tar.gz
tar -zxvf ${FILENAME}.tar.gz
cd ${FILENAME}

# Run the script
./roadie-queue-manager {{.ProjectID}} {{.Name}}

# Shutdown
docker run -it jkawamoto/roadie-gcp shutdown
