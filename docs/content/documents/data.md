---
title: Data
---
roadie manages three kinds of data in a cloud storage;
source codes, input data, outputted results.
To access those data, roadie has commands `source`, `data`,
and `result`, respectively.

### source command
`source` command provides a way to access your source codes
uploaded by `roadie run` command with `--local` flag.

To check source files uploaded, run

```sh
$ roadie source list
```

and to delete some file named `FILENAME`, run

```sh
$ roadie source delete FILENAME
```

### data command
`data` command manages input data files which may be used in
`data` section in your script file.

To upload a data file `FILENAME`, run

```sh
$ roadie data put FILENAME
```

To check uploaded files and their URL, run

```sh
$ roadie data list --url
```

The URLs shown by the above command, which start with `gs://`,
can be used in `data` section in your script file.

To delete some file `FILENAME`, run

```sh
$ roadie data delete FILENAME
```

### result command
`result` command manages outputs from instances and downloads them to your PC.

To check instance names which have results, run

```sh
$ roadie result list
```

and to check file names which instance `INSTANCE` has as its result, run

```sh
$ roadie result INSTANCE
```

To download those files into directory `./res`, run

```sh
$ roadie result get INSTANCE "*" -o ./res
```
