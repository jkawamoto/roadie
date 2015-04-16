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
import os
import shutil
import subprocess
import sys
from common.docker.environment import MONGO
from contextlib import contextmanager
from os import path


@contextmanager
def closing_fd(fd):
    """ A context manager closing a file descriptor.

    Args:
      fd: a file descriptor opened by os.open.

    Yield:
      the file descriptor.
    """
    try:
        yield fd
    finally:
        os.close(fd)


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
    def __init__(self, bucket, prefix=None, mimetype="text/plain"):
        """ Construct GCEStorage.

        Args:
          bucket: a bucket name.
          prefix: Prefix of file names (default: None)
          mimetype: MIME type (default: text/plain)
        """
        from common import gce
        self.__storage = gce.Storage(bucket)
        self.__prefix = prefix
        self.__mimetype = mimetype

    def copy_file(self, src):
        """ Copy a file into the storage.

        Args:
          src: A file name.
        """
        if self.__prefix:
            dst = self.__prefix + "/" + path.basename(src)
        else:
            dst = path.basename(src)
        self.__storage.upload_file(dst, src, self.__mimetype)


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


def run(cmd, observe, storage, cwd, shutdown, **kwargs):

    # Temporary file names.
    STDOUT = path.join("/tmp", "stdout")
    STDERR = path.join("/tmp", "stderr")

    # Constructing a storage object.
    s = storage(**kwargs)

    try:

        # Creating temporary files.
        with closing_fd(os.open(STDOUT, os.O_CREAT | os.O_WRONLY | os.O_TRUNC, 0644)) as stdout:

            with closing_fd(os.open(STDERR, os.O_CREAT | os.O_WRONLY | os.O_TRUNC, 0644)) as stderr:

                # Create a subprocess.
                p = subprocess.Popen(cmd, cwd=cwd, shell=True, bufsize=1, stdout=stdout, stderr=stderr)

                # Wait the process will end.
                p.wait()

        # Storing stdout and stderr
        s.copy_file(STDOUT)
        s.copy_file(STDERR)

        os.remove(STDOUT)
        os.remove(STDERR)

        # Copy other files.
        for src in glob.glob(observe) if observe else []:
            s.copy_file(src)

    finally:

        if shutdown:
            from common.gce import shutdown
            shutdown.shutdown()


def main():

    parser = argparse.ArgumentParser()
    parser.add_argument("--observe", help="File pattern to be stored.")
    parser.add_argument("cmd", nargs="?", default="main.py", help="Command to be run.")
    parser.add_argument("--cwd", default="/data", help="Change working directory (default: /data).")
    parser.add_argument("--shutdown", default=False, action="store_true",
                        help="Shutdown after the program ends (working only in Google Compute Engine)")

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
    main()
