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
import mimetypes
import logging
import logging.config
import os
import shutil
import subprocess
import sys
import zipfile
from common.docker.environment import MONGO
from os import path


TMP_DIR = "/tmp"
STDOUT = path.join(TMP_DIR, "stdout.txt")
ZIPFILE = path.join(TMP_DIR, "archive.zip")

logger = logging.getLogger(__name__)

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
    def __init__(self, bucket, prefix=None):
        """ Construct GCEStorage.

        Args:
          bucket: a bucket name.
          prefix: Prefix of file names (default: None)
        """
        from common import gce
        self.__storage = gce.Storage(bucket)
        self.__prefix = prefix

    def copy_file(self, src):
        """ Copy a file into the storage.

        Args:
          src: A file name.
        """
        if self.__prefix:
            dst = self.__prefix + "/" + path.basename(src)
        else:
            dst = path.basename(src)

        mtype, _ = mimetypes.guess_type(src)
        if not mtype:
            mtype = "application/octet-stream"
        self.__storage.upload_file(dst, src, mtype)


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


def run(cmd, observe, storage, log=None, stderr=sys.stdout, cwd="/data", shutdown=False, zip=False, **kwargs):
    """ Run a command.

    Args:
      cmd: a command line to be run.
      observe: a UNIX like file pattern to be stored.
      storage: a storage type and must be an instance of Storage.
      cwd: specify a directory where the given command run.
      shutdown: if true, vm will be shutdowned when the command end.
      zip: if true, observed files will be zipped.
    """
    if log:
        logging.config.fileConfig(log)

    try:

        # Constructing a storage object.
        s = storage(**kwargs)

        # Creating temporary files.
        with open(STDOUT, "w") as stdout:

            logger.info("Start a sub process.")
            # Create a subprocess.
            p = subprocess.Popen(cmd, cwd=cwd, shell=True, bufsize=1, stdout=stdout, stderr=stderr)

            # Wait the process will end.
            p.wait()

            logger.info("The sub process has ended.")

        # Storing stdout
        if path.exists(STDOUT):
            logger.info("Upload %s.", STDOUT)

            try:
                s.copy_file(STDOUT)
                os.remove(STDOUT)

            except Exception:
                logger.exception("An error occurred when storing stdout.")

        # Copy other files.
        if observe:

            try:

                if zip:
                    # Create a zip file and store observed files.
                    with zipfile.ZipFile(ZIPFILE, "w", compression=zipfile.ZIP_DEFLATED, allowZip64=True) as arc:

                        for pat in observe:
                            for src in itertools.ifilterfalse(path.isdir, glob.glob(pat)):
                                arc.write(src, arcname=path.basename(src))

                    logger.info("Upload %s.", ZIPFILE)
                    s.copy_file(ZIPFILE)

                else:

                    for pat in observe:
                        for src in itertools.ifilterfalse(path.isdir, glob.glob(pat)):
                            logger.info("Upload %s.", src)
                            s.copy_file(src)

            except Exception:
                logger.exception("An error occurred when storing observed files.")

    finally:

        if shutdown:
            from common.gce import shutdown
            logger.info("Now shuddown.")
            shutdown.shutdown()


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--observe", nargs="*", help="glob patterns of files to be stored.")
    parser.add_argument("--cwd", default="/data", help="Change working directory (default: /data).")
    parser.add_argument("--shutdown", default=False, action="store_true",
                        help="Shutdown after the program ends (working only in Google Compute Engine)")
    parser.add_argument("--zip", default=False, action="store_true", help="Files specified in overve option will be zipped.")
    parser.add_argument("--log", help="Config file of loggers.")
    parser.add_argument("--stderr", default=sys.stdout, help="Specify where stderr should be stored (default: stdout).")
    parser.add_argument("cmd", help="Command line to be run.")

    subparsers = parser.add_subparsers()

    gcs_cmd = subparsers.add_parser("gcs", help="Storing outputs into Google Cloud Storage.")
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
    logging.basicConfig(level=logging.INFO)
    try:
        main()
    except KeyboardInterrupt:
        sys.exit(1)
    finally:
        logging.shutdown()
