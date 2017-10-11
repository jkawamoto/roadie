---
title: Data
description: >-
  This document explains how to manage user programs and data files those
  programs will use. It also explains how to obtain computation results from
  cloud services.
date: 2016-08-14
lastmod: 2017-10-10
slug: data
---
Roadie manages three kinds of data in a cloud storage:

- source code,
- data files,
- outputted results.

To access those data, Roadie has commands `source`, `data`,
and `result`, respectively.

### source code
Source code uploaded by `roadie run` command with `--local` flag are stored in
`roadie://source/`. `source` command provides methods to manage those source
code.

#### list
To find source code archives stored in `roadie://source/`, use `list` sub command.
The following example prints all stored archives:

```shell
$ roadie source list
```

#### delete
To delete an archive file `FILENAME`, use `delete` sub command:

```shell
$ roadie source delete FILENAME
```

#### get
`get` sub command downloads a stored archive. The following example downloads
`FILENAME` in the current directory:

```shell
$ roadie source get FILENAME
```

If you want to download to another directory, such as `~/path`, give a path
with `-o` flag:

```shell
$ roadie source get -o ~/path FILENAME
```

#### put
`put` sub command uploads your source code.
The following example archives files in `~/source` into `source.tar.gz`, and
uploads it:

```shell
$ roadie source put ~/source source.tar.gz
```

### data files
Data files are stored in `roadie://data/` and referred in `data` of Roadie's
script files.

#### put
`put` sub command uploads a file to a cloud storage.
The following example uploads `FILENAME`:

```shell
$ roadie data put FILENAME
```

and after uploading is succeeded, it shows the URL of the upload file.

#### list
`list` sub command shows uploaded files and those URLs.

```shell
$ roadie data list --url
```

If `--url` flag is not give, only file names are shown.

#### delete
`delete` sub command deletes an uploaded file.
The following example deletes `FILENAME`:

```shell
$ roadie data delete FILENAME
```

#### get
`get` sub command downloads an uploaded file.
The following example download `FILENAME` into the current directory:

```shell
$ roadie data get FILENAME
```

If you want to download to another directory, such as `~/path`, give a path
with `-o` flag:

```shell
$ roadie data get -o ~/path FILENAME
```

### Result files
Messages written in the standard output `stdout` and files specified in `result`
of Roadie's script file will be stored in `roadie://resutl/<instance name>/`.

#### list
`list` sub command without any options shows instance names.
The following example prints a list of instance names:

```shell
$ roadie result list
```

`list` sub command with an instance name shows result files uploaded from the
specified instance.
The following example shows result file names uploaded from `INSTANCE`:

```shell
$ roadie result list INSTANCE
```

#### get
`get` sub command takes a glob pattern and downloads result files matching the
given pattern.
The following example downloads all result file into the current directory
by using a wild card pattern `*`:

```shell
$ roadie result get INSTANCE "*"
```

If you want to download them to another directory, use `-o` flag.
For example, the following example downloads files start with `stdout` into
`~/path`:

```shell
$ roadie result get INSTANCE "stdout*" -o `~/path`
```

#### delete
`delete` sub command deletes result files matching a given glob pattern.
For example, the following example deletes files end with `.png`:

```shell
$ roadie result delete INSTANCE "*.png"
```

If the glob pattern is omitted, all result files *including log files* will be
deleted.

#### show
`show` sub command shows messages written in the standard output.
The following example shows all messages in `INSTANCE`:

```shell
$ roadie result show INSTANCE
```

If you want to see outputted massages from *i*-th command in `run` of your
script file, give the number `i` like

```shell
$ roadie result show INSTANCE i
```
