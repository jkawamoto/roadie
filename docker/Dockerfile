#
# Dockerfile
#
# Copyright (c) 2015 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
FROM ubuntu:latest
MAINTAINER Junpei Kawamoto <kawamoto.junpei@gmail.com>

# Install packages
RUN apt-get update && apt-get install -y curl python-pymongo
RUN curl https://bootstrap.pypa.io/get-pip.py | python
RUN pip install --upgrade google-api-python-client

# Copy entrypoint
COPY bin /root/

# Change working directory
VOLUME /data
WORKDIR /data

ENTRYPOINT ["/root/entrypoint.py"]
