#!/usr/bin/env python
#
# mongo.py
#
# Copyright (c) 2015 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
import argparse
import itertools
import json
import sys
import urlparse
from datetime import datetime
from docker.environment import MONGO
from pymongo import MongoClient


def _dispatch(func, **kwargs):
    func(**kwargs)


def _formatter(parse):
    if parse:
        def _(text):
            return json.loads(text.replace("'", "\""))
    else:
        def _(text):
            return text.rstrip()
    return _


def translater(js):
    if js:
        def _(s):
            return json.loads(s)
    else:
        def _(s):
            return s.rstrip()
    return _


def _make_object(l):
    res = {}
    for k, v in [kv.split("=") for kv in l]:
        res[k] = v
    return res


def _make_item(name, key, value, properties):

    item = {
        "time": datetime.now(),
        "name": name,
        key: value
    }

    if properties:
        item["properties"] = properties

    return item


def push(db, collection, name, input, host=MONGO.hostname, port=MONGO.port, key="data", squash=False, properties=[], json=False, quiet=False):

    with MongoClient(host, port) as client:
        db = client[db]
        col = db[collection]

        props = None
        if properties:
            props = _make_object(properties)

        if squash:
            data = map(translater(json), input)
            if data:
                col.insert(_make_item(name, key, data, props))

                if not quiet:
                    for s in data:
                        print s

        else:
            for s in itertools.imap(translater(json), input):

                col.insert(_make_item(name, key, s, props))

                if not quiet:
                    print s


def pull(db, collection, name, output=sys.stdout, host=MONGO.hostname, port=MONGO.port, key="data", query=""):

    with MongoClient(host, port) as client:
        db = client[db]
        col = db[collection]

        if query:
            q = json.loads(query.replace("'", "\""))
        else:
            q = {"name": name}

        for res in ifilter(lambda res: key in res, col.find(q)):
            if isinstance(res[key], list):
                for i in res[key]:
                    output.write(i)
                    output.write("\n")
            else:
                output.write(res[key])
                output.write("\n")

    output.flush()


def main():

    parser = argparse.ArgumentParser()
    parser.add_argument("--host", default=MONGO.hostname, help="Host name of a MongoDB Server.")
    parser.add_argument("--port", default=MONGO.port, type=int, help="Port number of a MongoDB Server.")
    parser.add_argument("--quiet", default=False, action="store_true")
    parser.add_argument("--key", default="data")
    parser.add_argument("db", help="Database name.")
    parser.add_argument("collection", help="Collection name.")
    parser.add_argument("name", help="Document name.")

    subparsers = parser.add_subparsers()

    push_cmd = subparsers.add_parser("push")
    push_cmd.add_argument("--input", default=sys.stdin, type=argparse.FileType("r"), help="Source of pushing data.")
    push_cmd.add_argument("-p", "--properties", nargs="*")
    push_cmd.add_argument("--squash", default=False, action="store_true")
    push_cmd.add_argument("--json", default=False, action="store_true")
    push_cmd.set_defaults(func=push)

    pull_cmd = subparsers.add_parser("pull")
    pull_cmd.add_argument("--query", default="")
    pull_cmd.add_argument("--output", default=sys.stdout, type=argparse.FileType("w"))
    pull_cmd.set_defaults(func=pull)

    _dispatch(**vars(parser.parse_args()))


if __name__ == "__main__":
    main()
