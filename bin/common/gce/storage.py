#!/usr/bin/env python
#
# storage.py
#
# Copyright (c) 2015 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
import argparse
import json
import sys
from apiclient import discovery
from apiclient.http import MediaIoBaseUpload
from apiclient.http import MediaFileUpload
from auth import Auth
from googleapiclient.errors import HttpError

class Storage(object):

    def __init__(self, bucket):

        self._bucket = bucket
        self._sp = discovery.build("storage", "v1")

        self._auth = None
        self._update_auth()

    def get(self, path, metadata=False):
        if metadata:
            req = self._sp.objects().get(bucket=self._bucket, object=path)
        else:
            req = self._sp.objects().get_media(bucket=self._bucket, object=path)

        req.headers["Authorization"] = self._auth.header_str()
        try:
            return req.execute()

        except HttpError as e:
            if e.resp.status == 401:
                self._update_auth()
                return self.get(path, metadata)

            else:
                sys.stderr.write("Error: " + path + "\n")
                raise e

    def upload(self, path, data, mimetype):
        media_body = MediaIoBaseUpload(data, mimetype, resumable=True)
        return self._do_upload(path, media_body)

    def upload_file(self, path, fname, mimetype):
        media_body = MediaFileUpload(fname, mimetype, resumable=True)
        return self._do_upload(path, media_body)

    def _update_auth(self):
        print "Auth"
        self._auth = Auth()

    def _do_upload(self, path, media_body):
        req = self._sp.objects().insert(bucket=self._bucket, body=dict(name=path), media_body=media_body)

        req.headers["Authorization"] = self._auth.header_str()
        try:
            res = None
            print "Start uploading " + path
            while res is None:
                status, res = req.next_chunk()
                if status:
                    print "Uploaded %d%%." % int(status.progress() * 100)
            print "Upload Complete!"

            return res

        except HttpError as e:
            if e.resp.status == 401:
                self._update_auth()
                return self._do_upload(path, media_body)

            else:
                sys.stderr.write("Error: " + path + "\n")
                raise e


def get(bucket, path, output, metadata, **kwargs):
    s = Storage(bucket)
    if metadata:
        output.write(s.get(path, metadata))
    else:
        for data in s.get(path, metadata):
            output.write(data)
    output.flush()


def upload(bucket, path, input, mimetype, **kwargs):
    s = Storage(bucket)
    s.upload(path, input, mimetype)


def main(**kwargs):
    kwargs["func"](**kwargs)
    return 0


if __name__ == "__main__":

    parser = argparse.ArgumentParser()
    subparsers = parser.add_subparsers()

    get_cmd = subparsers.add_parser("get")
    get_cmd.add_argument("bucket")
    get_cmd.add_argument("path")
    get_cmd.add_argument("--metadata", action="store_true")
    get_cmd.add_argument("--output", default=sys.stdout, type=argparse.FileType("w"))
    get_cmd.set_defaults(func=get)

    upload_cmd = subparsers.add_parser("upload")
    upload_cmd.add_argument("bucket")
    upload_cmd.add_argument("path")
    upload_cmd.add_argument("mimetype")
    upload_cmd.add_argument("--input", default=sys.stdin, type=argparse.FileType("r"))
    upload_cmd.set_defaults(func=upload)

    sys.exit(main(**vars(parser.parse_args())))
