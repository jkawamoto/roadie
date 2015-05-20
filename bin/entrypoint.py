#!/usr/bin/env python
#
# entrypoint.py
#
# Copyright (c) 2015 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
import argparse
import glob
import itertools
<<<<<<< HEAD
=======
import mimetypes
>>>>>>> pre-merge
import os
import shutil
import subprocess
import sys
import zipfile
from common.docker.environment import MONGO
from os import path


TMP_DIR = "/tmp"
STDOUT = path.join(TMP_DIR, "stdout")
ZIPFILE = path.join(TMP_DIR, "archive.zip")


class Storage(object):
    """ Interface of Storage.
    """
    def copy_file(self, src):
        """ Copy a file into the storage.

        Args:
          src: A file name.
        """
        raise NotImplementedError()


class GCEStorage(Storage):
    """ Store data into Google Could Storage.
    """
<<<<<<< HEAD
    def __init__(self, bucket, prefix=None, mimetype="text/plain"):
=======
    def __init__(self, bucket, prefix=None):
>>>>>>> pre-merge
        """ Construct GCEStorage.

        Args:
          bucket: a bucket name.
          prefix: Prefix of file names (default: None)
<<<<<<< HEAD
          mimetype: MIME type (default: text/plain)
=======
>>>>>>> pre-merge
        """
        from common import gce
        self.__storage = gce.Storage(bucket)
        self.__prefix = prefix
<<<<<<< HEAD
        self.__mimetype = mimetype
=======
>>>>>>> pre-merge

    def copy_file(self, src):
        """ Copy a file into the storage.

        Args:
          src: A file name.
        """
        if self.__prefix:
            dst = self.__prefix + "/" + path.basename(src)
        else:
            dst = path.basename(src)
<<<<<<< HEAD
        self.__storage.upload_file(dst, src, self.__mimetype)
=======

        mtype, _ = mimetypes.guess_type(src)
        if not mtype:
            mtype = "application/octet-stream"
        self.__storage.upload_file(dst, src, mtype)
>>>>>>> pre-merge


class MongoStorage(Storage):
    """ Store data into MongoDB.
    """
    def __init__(self, db, collection, host=MONGO.hostname, port=MONGO.port):
        """
        db: a database name.
        collection: a collection name.
        host: a hostname of a MongoDB server.
        port: a port number of a MongoDB server.
        """
        from common import mongo
        self.__db = db
        self.__collection = collection
        self.__host = host
        self.__port = port

    def copy_file(self, src):
        """ Copy a file into the storage.

        Args:
          src: A file name.
        """
        dst = path.basename(src)
        with open(src, "r") as fp:
            mongo.push(self.__db, self.__collection, dst, fp, host=self.__host, port=self.__port, squash=True)


class LocalStorage(Storage):
    """ Store data to a local file system.
    """
    def __init__(self, dir):
        """ Construct LocalStorage.

        Args:
          dir: a directory path where outputs will be stored.
        """
        self.__cwd = dir

    def copy_file(self, src):
        """ Copy a file into the storage.

        Args:
          src: A file name.
        """
        dst = path.join(self.__cwd, path.basename(src))
        shutil.copyfile(src, dst)


def run(cmd, observe, storage, cwd="/data", output=sys.stdout, shutdown=False, zip=False, **kwargs):
    """ Run a command.

    Args:
      cmd: a command line to be run.
      observe: a UNIX like file pattern to be stored.
      storage: a storage type and must be an instance of Storage.
      cwd: specify a directory where the given command run.
      output: specify a file object to write outputs.
      shutdown: if true, vm will be shutdowned when the command end.
      zip: if true, observed files will be zipped.
    """
    try:

        # Constructing a storage object.
        s = storage(**kwargs)

        # Creating temporary files.
        with open(STDOUT, "w") as stdout:

            # Create a subprocess.
            p = subprocess.Popen(cmd, cwd=cwd, shell=True, bufsize=1, stdout=stdout, stderr=output)

            # Wait the process will end.
            p.wait()

        # Storing stdout
        if path.exists(STDOUT):
            try:
                s.copy_file(STDOUT)
                os.remove(STDOUT)

            except Exception as e:
                output.write("{0}\n".format(e))

        # Copy other files.
        if observe:

            try:

                if zip:
                    # Create a zip file and store observed files.
                    with zipfile.ZipFile(ZIPFILE, "w", compression=zipfile.ZIP_DEFLATED, allowZip64=True) as arc:

                        for pat in observe:
                            for src in itertools.ifilterfalse(path.isdir, glob.glob(pat)):
                                arc.write(src, arcname=path.basename(src))
                    s.copy_file(ZIPFILE)

                else:

                    for pat in observe:
                        for src in itertools.ifilterfalse(path.isdir, glob.glob(pat)):
                            s.copy_file(src)

            except Exception as e:
                output.write("{0}\n".format(e))

    finally:

        if shutdown:
            from common.gce import shutdown
            shutdown.shutdown()


def main():

    parser = argparse.ArgumentParser()
    parser.add_argument("--observe", nargs="*", help="File pattern to be stored.")
    parser.add_argument("cmd", nargs="?", default="main.py", help="Command to be run.")
    parser.add_argument("--cwd", default="/data", help="Change working directory (default: /data).")
    parser.add_argument("--output", default=sys.stdout, help="Store output into a file.")
    parser.add_argument("--shutdown", default=False, action="store_true",
                        help="Shutdown after the program ends (working only in Google Compute Engine)")
    parser.add_argument("--zip", default=False, action="store_true", help="Zip observed file.")


    subparsers = parser.add_subparsers()

    gcs_cmd = subparsers.add_parser("gcs", help="Storing outputs into GCS.")
    gcs_cmd.add_argument("--mimetype", default="text/plain", help="MIME type of outputs.")
    gcs_cmd.add_argument("bucket", help="Bucket name.")
    gcs_cmd.add_argument("--prefix", help="Prefix of stored files.")
    gcs_cmd.set_defaults(storage=GCEStorage)

    mongo_cmd = subparsers.add_parser("mongo", help="Storing outputs into MongoDB.")
    mongo_cmd.add_argument("--host", default=MONGO.hostname, help="Host name of MongoDB server.")
    mongo_cmd.add_argument("--port", default=MONGO.port, type=int, help="Port number of MongoDB server.")
    mongo_cmd.add_argument("db", help="Database name.")
    mongo_cmd.add_argument("collection", help="Collection name.")
    mongo_cmd.set_defaults(storage=MongoStorage)

    local_cmd = subparsers.add_parser("local", help="Storing outputs into local filesystem.")
    local_cmd.add_argument("dir", help="Directory for storing output.")
    local_cmd.set_defaults(storage=LocalStorage)

    run(**vars(parser.parse_args()))


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        sys.exit(1)
