---
title: Description
description: >-
  Roadie helps you to upload your source codes to the cloud, create and delete
  instances, and manage outputs.
date: 2016-12-20
lastmod: 2017-09-22
slug: description
---
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
[![Build Status](https://travis-ci.org/jkawamoto/roadie.svg?branch=master)](https://travis-ci.org/jkawamoto/roadie)
[![wercker status](https://app.wercker.com/status/6c499024136e7067b86bef4bd07d7f62/s/master "wercker status")](https://app.wercker.com/project/byKey/6c499024136e7067b86bef4bd07d7f62)
[![Go Report](https://goreportcard.com/badge/github.com/jkawamoto/roadie)](https://goreportcard.com/report/github.com/jkawamoto/roadie)
[![Release](https://img.shields.io/badge/release-0.3.13-brightgreen.svg)](https://github.com/jkawamoto/roadie/releases/tag/v0.3.13)
[![Japanese](https://img.shields.io/badge/qiita-%E6%97%A5%E6%9C%AC%E8%AA%9E-brightgreen.svg)](http://qiita.com/jkawamoto/items/751558536a597a33ae2a)

Roadie helps you to upload your source codes to the cloud, create and delete
instances, and manage outputs.

For example,

```shell
$ roadie run --local . --name analyze-wowah script.yml
```

uploads your source codes in current directory, and run them in such a manner
that `script.yml` specifies. The `script.yml` is a simple YAML file like

```yaml
apt:
  - unrar
data:
  - http://mmnet.iis.sinica.edu.tw/dl/wowah/wowah.rar
run:
  - unrar x -r wowah.rar
  - analyze WoWAH
```

The above `script.yml` asks roadie to install apt package `unrar` and
download a data file from such URL as the preparation. Then, it directs
to run those two commands: unrar the downloaded file, analyze the obtained
data files.

Roadie uploads results of such commands to a cloud storage after they finish.
You can access those results by

```shell
$ roadie result get analyze-wowah "*" -o ./res
```

Then, Roadie downloads all result files into `./res` directory.

Currently, Roadie supports 

- <i class="fa fa-google" aria-hidden="true"></i> [Google Cloud Platform](https://cloud.google.com/),
- <i class="fa fa-windows" aria-hidden="true"></i> [Microsoft Azure](https://azure.microsoft.com/).


## Installation
Compiled binary files for some platforms are uploaded in
[release page](https://github.com/jkawamoto/roadie/releases).

If you're a [Homebrew](http://brew.sh/) or [Linuxbrew](http://linuxbrew.sh/)
user, you can install Roadie by the following commands:

```shell
$ brew tap jkawamoto/roadie
$ brew install roadie
```

## Initialization
After installing Roadie, the following initialization is required in order to
authorize `roadie` to access cloud services.

```shell
$ roadie init
```

## License
This software except files in `docker` folder is released under The GNU General Public License Version 3,
see [LICENSES](info/licenses/) for more detail.
