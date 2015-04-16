# -*- coding: utf-8 -*-
#
# memoize.py
#
# Copyright (c) 2015 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
from functools import wraps

def memoized(func):
  cache = {}
  @wraps(func)
  def memoized_function(*args):
    try:
      return cache[args]
    except KeyError:
      value = func(*args)
      cache[args] = value
      return value
  return memoized_function
