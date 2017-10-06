#!/bin/bash
apt-get update && apt-get install -y wget
wget -O roadie.tar.gz "https://www.dropbox.com/s/julstcq5764oaq5/roadie-azure_linux_amd64.tar.gz?dl=1"
tar -zxvf roadie.tar.gz --strip=1

cat <<EOS > config.yml
{{ .Config }}
EOS

cat <<EOS > script.yml
{{ .Script }}
EOS

./roadie-azure config.yml script.yml
