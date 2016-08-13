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

- A bug of uploaded filenames that all uploaded files were saved as a same name.


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
