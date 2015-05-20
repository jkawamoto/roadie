<<<<<<< HEAD
#! /usr/bin/env python
#
# shutdown.py
#
# Copyright (c) 2015 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
import urllib2
from apiclient import discovery
=======
import urllib2
from apiclient import discovery 
>>>>>>> pre-merge
from auth import Auth

_INSTANCE = "http://169.254.169.254/computeMetadata/v1/instance/"
_PROJECT = "http://169.254.169.254/computeMetadata/v1/project/"

def _get(url):
    req = urllib2.Request(url)
    req.add_header("Metadata-Flavor", "Google")
    return urllib2.urlopen(req).readline()

<<<<<<< HEAD

=======
>>>>>>> pre-merge
def shutdown():

    auth = Auth()
    instance = _get(_INSTANCE + "hostname").split(".")[0]
    zone = _get(_INSTANCE + "zone").split("/")[-1]
    project = _get(_PROJECT + "project-id")

    sp = discovery.build("compute", "v1")
<<<<<<< HEAD

    req = sp.instances().delete(project=project, zone=zone, instance=instance)
    req.headers["Authorization"] = auth.header_str()
    req.execute()


if __name__ == "__main__":

    shutdown()
=======
    
    req = sp.instances().delete(project=project, zone=zone, instance=instance)
    req.headers["Authorization"] = auth.header_str()
    req.execute()
    
if __name__ == "__main__":
    
    shutdown()

>>>>>>> pre-merge
