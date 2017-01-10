#!/bin/bash
#
# startup script.
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
# along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
#

# This script starts both fluentd and a given program.
#
cd /root

# Start logging.
if [[ -n $(docker ps -a | grep fluentd) ]]; then
  docker rm -f fluentd
fi
for i in $(seq 10); do
  docker run -d --name fluentd -e 'INSTANCE={{.Name}}' -e 'USERNAME=roadie' \
    -v /var/lib/docker:/var/lib/docker jkawamoto/docker-google-fluentd
  sleep 30s
  if [[ -n $(docker ps -a | grep fluentd) ]]; then
    break
  fi
done
sleep 30s

# Run the script
cat <<EOF > run.yml
{{.Script}}
EOF

for i in $(seq {{.Retry}}); do
  if [[ -n $(docker ps -a | grep {{.Name}}) ]]; then
    docker rm -f {{.Name}}
  fi
  docker run -i --name {{.Name}} {{.Image}} {{.Options}} < run.yml || break
done
