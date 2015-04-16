#
# decorator.py
#
# Copyright (c) 2015 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
from functools import wraps

def print_args(func):
    @wraps(func)
    def _(*args, **kwargs):
        print args, kwargs
        return func(*args, **kwargs)
    return _


def print_return(func):
    @wraps(func)
    def _(*args, **kwargs):
        res = func(*args, **kwargs)
        print res
        return res
    return _


def constant(func):
    @wraps(func)
    def _(*args, **kwargs):
        if not _.res:
            _.res = func(*args, **kwargs)
        return _.res
    _.res = None
    return _
