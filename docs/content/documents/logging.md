---
title: Logging
description: >-
  This document explains how to obtain statuses of your programs and receive
  log data.
date: 2016-08-14
lastmod: 2017-10-10
slug: logging
---
After running your program, it is helpful to check status of your program,
i.e. instance, and check logs.
Roadie provides two commands:

- **status** shows status of all instances and kill an instance.
- **log** shows logs from an instance.

### status
Status command `roadie status` shows your instances are running or have ended.
If your instance has already ended and you also deleted all result data
from the instance, such instance name will be omitted to print.
To print all instance status including such deleted instances, use `--all` flag.
However, status of old instances will be deleted after a certain period of time.

The another property of status command is to kill some instance.
To kill instance `INSTANCE`, run

```shell
$ roadie status kill INSTANCE
```

If you kill an instance, outputs from the instance might not be stored.


### log
Log command prints log messages from an instance.
To see log messages from instance `INSTANCE`, run

```shell
$ roadie log INSTANCE
```

The log messages consist of logs about preprocess and post process,
and outputs your programs write in standard err (`stderr`).

Roadie's execution model treats outputs written in standard output (`stdout`)
and `stderr` in different way.
Outputs written in `stdout` are uploaded to a bucket as parts of results
from your program.
You can see and download such outputs by `roadie result` command but
you cannot check them while your instance is running.

On the other hand, outputs written in `stderr` are not stored in any persistent
disks but treated as prompt logs, which means you can check such logs while
your instance is still running.
Because outputs written in `stderr` cause of network traffic, it isn't
recommended to write huge messages there.

#### options
Log command has two option flags; `--no-timestamp` and `--follow`.
If `--no-timestamp` flag is set, roadie omits to print time stamps.
If `--follow` flag is set, roadie will not end and keep waiting new logs coming. To stop it, use `control + c`.
