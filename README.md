# roadie
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
[![Build Status](https://travis-ci.org/jkawamoto/roadie.svg?branch=master)](https://travis-ci.org/jkawamoto/roadie)
[![Code Climate](https://codeclimate.com/github/jkawamoto/roadie/badges/gpa.svg)](https://codeclimate.com/github/jkawamoto/roadie)
[![Release](https://img.shields.io/badge/release-0.2.7-brightgreen.svg)](https://github.com/jkawamoto/roadie/releases/tag/v0.2.7)
[![Japanese](https://img.shields.io/badge/qiita-%E6%97%A5%E6%9C%AC%E8%AA%9E-brightgreen.svg)](http://qiita.com/jkawamoto/items/751558536a597a33ae2a)

A easy way to run your programs on
[Google Cloud Platform](https://cloud.google.com/).
See [documents](https://github.com/jkawamoto/roadie/wiki) for more information.

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

## Install
Compiled binary files for some platforms are uploaded in [release page](https://github.com/jkawamoto/roadie/releases).
To install in other platforms following the go manner, use `go get`:

```sh
$ go get github.com/jkawamoto/roadie
```

For mac user, `roadie` is available in [Homebrew](http://brew.sh/).

```sh
$ brew tap jkawamoto/roadie
$ brew install roadie
```

## License
This software except files in `docker` folder is released under The GNU General Public License Version 3, see [COPYING](COPYING) for more detail.
