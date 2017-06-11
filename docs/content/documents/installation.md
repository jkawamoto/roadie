---
title: Installation
---
### Install
Roadie's compiled binary files for some platforms are found in Github's
[release page](https://github.com/jkawamoto/roadie/releases).
Download one of them according to your environment and put the binary into
a directory in your `$PATH`, or put it as same directory as your current project.

If you are a mac user, you can install `roadie` via [Homebrew](http://brew.sh/).

```shell
$ brew tap jkawamoto/roadie
$ brew install roadie
```

### Initialization
Each project needs to initialize roadie and sets your *project ID*.
Project ID is an ID registered in Google Cloud Platform.
You can find your project ID [here](https://console.cloud.google.com/project).
Note that *project name* is different from project ID.

```shell
$ roadie init
```

Although the configuration file is a text file and you can edit it,
roadie provides `config` command to edit it interactively.
See [configuration page](documents/configuration) for more information.
