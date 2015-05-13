#! /bin/bash
git rm --cached bin/common
git rm .gitmodules
rm -rf bin/common/.git

git add bin/common/__init__.py
git add bin/common/decorator.py
git add bin/common/memoize.py
git add bin/common/mongo.py

git add bin/common/docker/__init__.py
git add bin/common/docker/environment.py

git add bin/common/gce/__init__.py
git add bin/common/gce/auth.py
git add bin/common/gce/shutdown.py
git add bin/common/gce/storage.py

git rm fabfile.py

git commit -m "Import files from submodules."
