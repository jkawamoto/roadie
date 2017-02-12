## 0.3.4 (2017-02-11)
### Updated
- Switch to use Google Cloud for Go instead of Google APIs Client Library for
  Go to access storages,
- Use https URLs for git repositories even if ssh URLs.

### Fixed
- Broken tables to print download status,
- Complete a wild card if any glob pattern is not given in result get command,
- Complete missing .tar.gz in source flag of run command.


## 0.3.3 (2016-12-29)
### Fixed
- Cloud datastore access according to the new API.
- Stackdriver Logging client to use the new [logadmin](https://godoc.org/cloud.google.com/go/logging/logadmin) package.

### Added
- `assets.go` to compile roadie without go-bindata command.


## 0.3.2 (2016-10-11)
### Fixed
- `roadie result show` command won't output escape sequences.


## 0.3.1 (2016-10-11)
### Added
- Creating instance function waits until operation done message appears in log.
- Logs from worker instances have the instance name instead of the queue name.


## 0.3.0 (2016-09-22)
### Added
- Support queue based task management.


## 0.2.7 (2016-09-22)
### Update
- Startup script waits until fluentd is ready.


## 0.2.6 (2016-09-21)
### Fixed
- Not specifying default zone and machine type.


## 0.2.5 (2016-09-17)
### Added
- Support parallel uploading data files.


## 0.2.4 (2016-09-14)
### Added
- Support a new URL scheme `roadie://` which is a shortcut to
  `gs://<your bucket>/.roadie/<section>` in script files.

For example, a URL `roadie://sample.dat` in `data` section of your script file,
will be treated as `gs://<your bucket>/.roadie/data/sample.dat`.
It will reduce your types to make script files.


## 0.2.3 (2016-09-13)
### Fixed
- Output log entries of old instances when reusing same instance name.


## 0.2.2 (2016-09-08)
### Added
- Support to restart program after rebooting containers.

### Fixed
- Retry number in run command.


## 0.2.1 (2016-08-13)
### Fixed
- Issue #15: Delete old container before retrying.


## 0.2.0 (2016-08-12)
### Added
- Support `--retry` option by default it is set to 10.

By the option, roadie will retry executing a given program such times
when GCP error happens.


## 0.1.4 (2016-07-17)
### Fixed
- Look for configuration files.
- Print warning if Project ID is not set in configure file.
- Project ID and Bucket name do not allow empty strings.
- Project name has been renamed to Project ID.
- Update to use current zone name in order to search available machine types.
- Run command creates a bucket if necessary.


## 0.1.3 (2016-07-15)
### Feature

- In addition to `--name` flag, support `-n` flag as a short version.

### Fixed

- A bug of uploaded file names that all uploaded files were saved as a same name.


## 0.1.2 (2016-07-13)
### Fixed

- Disable font color and style in Windows.

In Windows, command prompt does not support escape scenes.
From this version, in windows, such escape sequences are not used.


## 0.1.1 (2016-07-10)
### Fixed

- Fixed a bug when gcloud does not return any project IDs.


## 0.1.0 (2016-07-08)
Initial release
