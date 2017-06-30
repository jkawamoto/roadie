# roadie
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
[![Build Status](https://travis-ci.org/jkawamoto/roadie.svg?branch=master)](https://travis-ci.org/jkawamoto/roadie)
[![wercker status](https://app.wercker.com/status/6c499024136e7067b86bef4bd07d7f62/s/master "wercker status")](https://app.wercker.com/project/byKey/6c499024136e7067b86bef4bd07d7f62)
[![Go Report](https://goreportcard.com/badge/github.com/jkawamoto/roadie)](https://goreportcard.com/report/github.com/jkawamoto/roadie)
[![Release](https://img.shields.io/badge/release-0.3.11-brightgreen.svg)](https://github.com/jkawamoto/roadie/releases/tag/v0.3.11)
[![Japanese](https://img.shields.io/badge/qiita-%E6%97%A5%E6%9C%AC%E8%AA%9E-brightgreen.svg)](http://qiita.com/jkawamoto/items/751558536a597a33ae2a)

[![Logo](https://jkawamoto.github.io/roadie/img/banner.png)](https://jkawamoto.github.io/roadie/)

A easy way to run your programs on
[Google Cloud Platform](https://cloud.google.com/).
See [documents](https://jkawamoto.github.io/roadie/) for more information.

## Description
`roadie` helps you to upload your source codes to the cloud, create and delete
instances, and manage outputs.

For example,

```sh
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
to run those two commands; unrar the downloaded file, analyze the obtained
data files.

`roadie` uploads results of such commands to a cloud storage after they finish.
You can access those results by

```sh
$ roadie result get analyze-wowah "*" -o ./res
```

Then, `roadie` downloads all result files into `./res` directory.

## Installation
Compiled binary files for some platforms are uploaded in
[release page](https://github.com/jkawamoto/roadie/releases).

For mac user, `roadie` is available in [Homebrew](http://brew.sh/).

```sh
$ brew tap jkawamoto/roadie
$ brew install roadie
```

## Initialization
After installing `roadie`, the following initialization is required in order to
authorize `roadie` to access cloud services.

```sh
$ roadie init
```

## License
This software except files in `docker` folder is released under The GNU General Public License Version 3, see [COPYING](COPYING) for more detail.
