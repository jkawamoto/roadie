#!/bin/bash
#
# startup script
#
cd /root

# Start logging.
docker run -d --name fluentd -e "INSTANCE={{.Name}}" -e "USERNAME=roadie" \
  -v /var/lib/docker:/var/lib/docker jkawamoto/docker-google-fluentd

sleep 30s

# Run the script
cat <<EOF > run.yml
{{.Script}}
EOF

docker run -i --name {{.Name}} jkawamoto/roadie-gcp {{.Options}} < run.yml
