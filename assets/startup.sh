#!/bin/bash
#
# startup script
#
cd /root

# Load configurations.
SCRIPT=$(curl http://metadata/computeMetadata/v1/instance/attributes/script -H "Metadata-Flavor: Google")
NAME=$(curl http://metadata/computeMetadata/v1/instance/hostname -H "Metadata-Flavor: Google")
NAME=${NAME%%.*}

# Start logging.
docker run -d --name fluentd -e "INSTANCE=${NAME}" -e "USERNAME=roadie" \
  -v /var/lib/docker:/var/lib/docker jkawamoto/docker-google-fluentd

# Run the script
echo $SCRIPT > run.yml
docker run -i --name "$NAME" jkawamoto/roadie-gcp < run.yml
