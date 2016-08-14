---
title: Installation
---
### Requirements
Since roadie currently supports only
[Google Cloud Platform](https://cloud.google.com/),
you need to register it at first.

roadie uses Google's [Cloud SDK](https://cloud.google.com/sdk/).
Please install and initialize it by yourself.
`gcloud auth list` shows an authorized account name, i.e. an email address,
after success of initialization.

### Install
roadie's compiled binary files for some platforms are found in Github's
[release page](https://github.com/jkawamoto/roadie/releases).
Download one of them according to your environment and put the binary into
a directory in your `$PATH`, or put it as same directory as your current project.

You can also install roadie by following the go manner.

```sh
$ go get github.com/jkawamoto/roadie
```

Then, the binary will be installed in `$GOPATH/bin`.
In this case, you need to install `go` before running `go get`.

If you are using mac, you can install `roadie` via [Homebrew](http://brew.sh/).

```sh
$ brew tap jkawamoto/roadie
$ brew install roadie
```

### Initialization
Each project needs to initialize roadie and notifies the *project ID* and
a bucket name.
Project ID is an ID registered in Google Cloud Platform.
You can find your project ID [here](https://console.cloud.google.com/project).
Note that *project name* is different from project ID.

```sh
$ roadie init
```

The initialization command asks you the current Google Cloud Platform's project
ID and other your preference, then makes configuration file `.roadie` in the
current directory.

Although the configuration file is a text file and you can edit it,
roadie provides `config` command to edit it interactively.
See [configuration page](documents/configuration) for more information.
