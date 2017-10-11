---
title: Queue
description: >-
  This document explains about queue management system in Roadie.
date: 2016-08-14
lastmod: 2017-10-10
slug: queue
---
Roadie supports queue to keep a bunch of script.
Enqueued scripts are run sequentially by default,
but you can add instances working with the queue.

### Add new script to a queue
Enqueue a script to a queue, use `roadie run` command with `--queue` flag.
The `--queue` flag takes a queue name and the given script will be enqueued
the queue.
If there are such queues, it will be created.
If there are no instances working with the queue, one instance will be created.
The other flags in `roadie run` command are as same as the case of starting
script without queue.

#### example
```shell
$ roadie run --local . --name instance-1 --queue queue-1 script.yml
```

The above command enqueues `script.yml` to a queue `queue-1`.
The script is named `instance-1` so that you can refer results and logs with the
name.

If there are no instances working with the queue, one instance will be created.


### Queue management
To find existing queues, use `roadie queue list` command.
On the other hand, to find enqueued scripts in a queue,
use

```shell
$ roadie queue show <queue name>
```

If you want to stop executing a queue,

```shell
$ roadie queue stop <queue name>
```

do it.
But scripts already running won't be stopped.
To restart stopped queue, use

```shell
$ roadie queue restart <queue name>
```

It restarts the queue and creates one instance to handle scripts.


### Instance management
By default, scripts in a queue are handled by one instance.
You can add other instances to handle any queue.

```shell
$ roadie queue instance add --instances N <queue name>
```

command adds `N`
instances to the given named queue. If you omit `--instances` flag,
just one instance will be added.
This command also support `--disk-size` flag as same as `roadie run` command.
You can modify disk size of instances according to your script.

```shell
$ roadie queue instance list <queue name>
```

shows instances working with the given named queue.


### Retrieve logging information
To retrieve logging information from tasks, use `roadie queue log` command
instead of `roadie log` command.

The following example shows all logging information from tasks in a queue:

```shell
$ roadie queue log <queue name>
```

If you want to see logging information from a specific task, give the task name
, too:

```shell
$ roadie queue log <queue name> <task name>
```
