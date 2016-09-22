---
title: Contents
---
`roadie` is a easy way to run your programs on [Google Cloud Platform](https://cloud.google.com/).
It helps you to upload your source codes to the cloud, create and delete
instances, and manage outputs.

- [Installation](documents/installation)
- [Configuration](documents/configuration)
- [Run your program](documents/execution)
- [Logging](documents/logging)
- [Data handling](documents/data)
- [Taks queue](documents/queue)

### Simple example
Suppose your are in a directory which has your source codes and `script.yml`,
then run

```sh
$ roadie run --local . --name analyze-wowah script.yml
```

uploads your source codes in the current directory,
and run them in such a manner that `script.yml` specifies.

The `script.yml` is a simple YAML file like

```yaml
apt:
- unrar
data:
- http://mmnet.iis.sinica.edu.tw/dl/wowah/wowah.rar
run:
- unrar x -r wowah.rar
- analyze WoWAH
upload:
- *.png
```

The above `script.yml` asks roadie to install apt package `unrar` and
download a data file from such URL as the preparation. Then, it directs
to run those two commands; unrar the downloaded file, analyze the obtained
data files.

You can check your program is still running or ends by

```sh
$ roadie status
```

After the program finishes,
`roadie` uploads results of such commands to a cloud storage.
You can access those results by

```sh
$ roadie result get analyze-wowah "*" -o ./res
```

`roadie` will download all result files into `./res` directory.
