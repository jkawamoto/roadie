#
# environment.py
#
# Copyright (c) 2015 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
from collections import namedtuple
from os import environ
from urlparse import urlparse

_MYSQL_ID = namedtuple("_MYSQL_ID", "user, passwd, db")

def _environ(key, alt):
    return environ[key] if key in environ else alt


MONGO = urlparse(_environ("MONGO_PORT", "tcp://127.0.0.1:27017"))
MYSQL = urlparse(_environ("MYSQL_PORT", "tcp://127.0.0.1:3308"))

MYSQL_ID = _MYSQL_ID(
    _environ("MYSQL_ENV_MYSQL_USER", None),      # user
    _environ("MYSQL_ENV_MYSQL_PASSWORD", None),  # passwd
    _environ("MYSQL_ENV_MYSQL_DATABASE", None)   # db
)
