# auth.py
#
# Copyright (c) 2015 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
import json
import urllib2

_METADATA_SERVER = "http://169.254.169.254/computeMetadata/v1/instance/service-accounts"
_SERVICE_ACCOUNT = "default"


class Auth(object):

    def __init__(self):
        self.execute()

    def execute(self):

        req = urllib2.Request("{0}/{1}/token".format(_METADATA_SERVER, _SERVICE_ACCOUNT))
        req.add_header("Metadata-Flavor", "Google")

        data = json.load(urllib2.urlopen(req))

        self._token = data["access_token"]
        self._type = data["token_type"]

    def header_str(self):
        return "{0} {1}".format(self.type, self.token)

    @property
    def token(self):
        return self._token

    @property
    def type(self):
        return self._type
