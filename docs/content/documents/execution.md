---
title: Execution
description: >-
  This document introduces how Roadie runs programs on cloud environment,
  and scheme of script files.
date: 2016-08-14
lastmod: 2017-10-10
slug: execution
---
### Script file
Roadie's script file is a YAML document which has five elements
`apt`, `source`, `data`, `run`, and `upload`.

Here is an example:

```yaml
apt:
  - unrar
source: https://github.com/abcdefg/some-program.git
data:
  - http://mmnet.iis.sinica.edu.tw/dl/wowah/wowah.rar
run:
  - unrar x -r wowah.rar
  - ./analyze WoWAH
upload:
  - *.png
```

The above example instructs Roadie to

1. install an apt package `unrar`,
2. clone a Git repository `abcdefg/some-program` from GitHub.com,
3. download an archived file in the URL,
4. run two commands (where `analyze` is a program supplied in the git repository),
5. upload messages written in the standard output and files matching `*.png` to
   a cloud storage.

Note that unnecessary elements (except for `run`) can be omitted in script files.

#### apt
`apt` takes a list of apt packages.

```yaml
apt:
  - python-numpy
  - python-scipy
  - python-matplotlib
```

In the above example, Roadie will install three Python packages.

If you need to update apt repositories, you need to do it and install packages
in `run`.

#### source
`source` takes a URL where your source code is provided.

Roadie retrieves your source code in the following manner.

- If the URL ends with `.git`, Roadie treats it is a git repository and uses
  `git clone` to obtain the source code.
- If the URL scheme is `dropbox://`, the source code will be downloaded from
  [Dropbox](https://www.dropbox.com/). This URL is a public link created by
  Dropbox but the scheme is replaced from `https://` to `dropbox://`.
- If the URL scheme is `gs://`, the source code will be downloaded from Google
  Cloud Storage (available if only you use Google Cloud Platform).
- If the URL scheme is `roadie://`, it means the file is managed by Roadie.
  See [Data](documents/data) for more information.
- Otherwise, URL schemes `http://` and `https://` are supported.

In any case, if the URL ends with `.zip`, `.tar`, or `.tar.gz`,
Roadie decompress such archived file.

Roadie also supports to
[upload your source code from your local computer directly](documents/execution#local-source-files).

If your source code is written in Python and it has `requirements.txt`,
required packages will be installed automatically.

#### data
`data` takes a list of URLs.
As same as `source`, URL schemes `http://`, `https://`, `dropbox://`,
`gs://` (only available with Google Cloud Platform), and `roadie://`.
If the URL ends with `.zip`, `.tar`, or `.tar.gz`,
the archived file will be decompressed as expected.

By default, downloaded files are stored in `/data`,
which is the same directory where source code is stored.
You can customize destinations by adding `:` plus destination path to each URL.

For example,

```yaml
data:
  - https://www.sample.com/program.zip:/data/input
```

instructs to download `program.zip` and store files in the archive into `/data/input`.

Here is another example,

```yaml
data:
  - roadie://data/some_data_v2.json:some_data.json
```

It instructs to download `some_data_v2.json`, which is managed by Roadie,
into `/data`, and rename it to `some_data.json`.

`roadie://data/` is the directory where files uploaded via `roadie data put`
are stored.

#### run
`run` takes a list of commands.
You can write any commands such as running your program,
installing any packages,
downloading any files (you should use `data`, though), etc.

Note that, you may need to start your command with `./`
if the running commands are in your source codes and set in `/data`.
roadie doesn't add `/data` to `$PATH`.

For example, if your program is written in [node.js](https://nodejs.org/en/),
the first command may be `npm install`.
Of course, you need to install node.js in `apt` section.

Each command listed in the `run` section has a zero-origin number,
i.e, the first command has 0.
This number is used to store outputs written in `stdout` and
the outputs written in `stdout` from *i*-th command are stored
in `stdout{i}.txt` file.
Those files will be accessed via `roadie result`.

#### upload
`upload` takes a list of
[glob](https://en.wikipedia.org/wiki/Glob_(programming)) patterns.
Files matching one of those patterns are treated of results
and uploaded to a cloud storage.
To access those uploaded files, use `roadie result` command.


### Execution model
Roadie runs your program in a Docker container.

[![Docker](img/small_h-trans.png)](https://www.docker.com/)

This container is based on [Ubuntu](http://www.ubuntu.com/) and you can use
most of packages supplied for Ubuntu in Roadie.
Roadie's script file has `apt` section which takes a list of apt packages.

Your program will be copied in `/data` in the running container.
Files listed up in `data` section of Roadie's script will also be copied in
`/data` by default.

Linux programs can output messages for the standard output `stdout` and
the standard error output `stderr`.
In Roadie, messages written in `stdout` will be treated as results of the program,
and stored in a cloud storage.
Each command in the `run` of your script makes one file to store
outputs written in `stdout`.
More precisely, *i*-th command creates `stdout{i}.txt`,
where *i* is a zero-origin integer.
Those files are stored in `/tmp` before all commands in `run` are done.

On the other hand, outputs written in `stderr` are not stored in any persistent
disks but treated as prompt logs, which means you can check such logs while
your instance is still running.
Because outputs written in `stderr` cause of network traffic, it isn't
recommended to write huge messages there.

By default, any other files created by your program will not be stored as
results.
To specify which files should be treated as results
and stored to persistent storage,
use `upload` section in the script file.


### run command
`roadie run` command creates an instance and runs your program on it.
This command requires one script file explained in the next section.
There are many option flags but one of the useful options is `--name`,
which sets a given name to the creating instance.
So, suppose you will create an instance named `instance1`
with script file `script.yml`, run

```shell
$ roadie run --name instance1 script.yml
```

If you don't set any names, roadie makes some name.
After creating the instance, roadie shows the name of the instance.
Such name is used to check instance status, see logs,
and download computation results.

If `-f` or `--follow` flag is set, `roadie run` command will print logs from the
created instance until it ends, as same as `roadie log` command with `-f` or
`--follow` flag.


### Specify source codes on the fly
Sometimes, it is difficult to provide your source codes from the web,
such as Git repository, Dropbox, and some web site.
roadie helps to upload your source code from a local PC to a cloud storage,
which is a private place.
If you use this function, you can omit `source` section in your script file.

#### Upload source codes from a local directory <a name="local-source-files"></a>
`--local` flag of `roadie run` command takes a path of your source codes.
For example,

```shell
$ roadie run --local . --name instance-1 script.yml
```

notifies roadie of the current path as the root path of your source codes.
roadie makes an archive file of the path and uploads it to a cloud storage.
Then, the created instance will use that file as the source codes.

If you give a path of one file with `--local` flag,
roadie uploads that file and the created instance will use it.

#### Previously uploaded files
`--source` flag of `roadie run` command takes an instance name
which run previously.
If the previous instance created with `--local` flag,
the created new instance will use same uploaded source file.
For example, you created an instance by

```shell
$ roadie run --local . --name instance-1 script.yml
```

and now you are creating another instance by

```shell
$ roadie run --source instance-1 --name instance2 script2.yml
```

the new instance named `instance-2` uses same source codes as `instance-1`.
