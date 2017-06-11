---
title: Execution
---
### Execution model
roadie runs your program on a container.
The image of the container is
[jkawamoto/roadie-gcp](https://github.com/jkawamoto/roadie-gcp)
which is based on [Ubuntu](http://www.ubuntu.com/).
That means most of Ubuntu packages are available in roadie, too.
You can install such packages via `apt` section in script file.

On the running container, your program will be set in `/data`,
and other data specified `data` section in script file will also be set in
`/data` by default.
If you provide your source codes as an archived file and
roadie reports an error that it couldn't find your program,
make sure that your archive file does not create additional directory.

For example, you provide your program `main` in `archive.zip` and
you expect your program will be set in `/data/main`.
But, depending on a way to create such archive file, your program may be set
in `/data/archive/main`.
For debugging, add `ls -la` on top of the run section in your script.

Your program has three ways to outputs; writing to `stdout`, `stderr`, and
creating other files.

All outputs written in `stdout` are treated as official results.
They are stored to persistent storage,
i.e. [Google Cloud Storage](https://cloud.google.com/storage/),
so that you can access them any time.
Each command in the `run` section of your script makes one file to store
outputs written in `stdout`.
More precisely, *i*-th command creates `stdout{i}.txt`,
where *i* is a zero-origin integer.

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

### Script file
Script file is a YAML file which consists of five sections;
`apt`, `source`, `data`, `run`, `upload`.
Here is a simple example of the script.

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

Roughly speaking, the above script means

1. Install apt package `unrar`,
2. Clone source codes from a git repository,
3. Download a data file from a certain URL.
4. Execute two commands; expand the downloaded file and run an analyze program.
5. Upload created file matching a glob pattern `*.png`
  in addition to outputs written in `stdout`

#### apt
`apt` section takes a list of apt package names.

```yaml
apt:
  - python-numpy
  - python-scipy
  - python-matplotlib
```

The above example installs three major python packages.
Note that, if you want to install apt packages after running some commands,
you can write normal `apt-get install` command in your `run` section.

#### source
`source` section takes one URL which points the location your program stored.
This URL can take several forms.

- If the URL ends with `.git`, roadie treats it a Git repository URL.
  roadie use `git clone` to obtain your source codes.
- If the URL starts with `dropbox://` instead of `http://` and `https://`,
  roadie thinks your source codes are provided in public directory in
  [Dropbox](https://www.dropbox.com/).
  You can make such URL easily by replacing `https://` to `dropbox://`
  in your shared link made by Dropbox.
- If the URL starts with `gs://`, which means your source codes are stored in
  Google Cloud Storage, roadie downloads them.
- If the URL likes `roadie://<name>`, which means your source codes are
  maintained in `roadie` and specifies to use as same source codes as ones
  used in instance `<name>`,
- Otherwise, roadie downloads the URL via http or https.

In any cases, if the URL ends with `.zip`, `.tar`, or `.tar.gz`,
downloaded files are expanded by `unzip` or `tar` command.

For example,
```yaml
source: https://github.com/jkawamoto/roadie.git
```
clones the Git repository.

However, if your program is still work in progress,
it could be trouble to make an archive file and upload it,
or push your codes to some repository.
roadie has a function to make such archive file containing files in a directory
and upload it instead of you. See **specify source codes on the fly** section.

The downloaded source codes are stored in `/data` directory.

If your source codes are written in python
and they have `requirements.txt` or `requirements.in`,
those required packages will be installed automatically.

#### data
`data` section takes a list of URLs.
Those URL support, `http`, `https`, `gs`, and `dropbox`.
As same as `source` section, if the URL ends with `.zip`, `.tar`, or `.tar.gz`,
those files are expanded as expected.

By default, downloaded files are stored in `/data` directory,
which is same directory as source codes.
You can customize destinations by adding `:` plus destination path to each URL.
For example,

```yaml
data:
  - https://www.sample.com/program.zip:/data/input
```

downloads `program.zip` and stored files in the zip into `/data/input`.
Here is another example,

```yaml
data:
  - gs://your-project/dataset/some_data_v2.json:some_data.json
```

downloads `somr_data_v2.json` into `/data` from a bucket in
Google Cloud Storage, and rename it to `some_data.json`.

In this section, URL scheme `roadie://` is also supported.
`roadie://somefile.dat` means `gs://<your bucket>/.roadie/data/somefile.dat`.
The place `gs://<your bucket>/.roadie/data/` is the default place,
`roadie` uploads your fils via `roadie data put` command.

#### run
`run` section takes a list of commands.
You can write any commands such as running your program,
installing any packages,
downloading any files (you should use `data` section, though), etc.

Note that, you may need to start your command with `./`
if the running commands are in your source codes and set in `/data`.
roadie doesn't add `/data` to `$PATH`.

For example, if your program is written in [node.js](https://nodejs.org/en/),
the first command may `npm install`.
Of course, you need.js to install node in `apt` section.

Each command listed in the `run` section has a zero-origin number,
i,e, the first command has 0.
This number is used to store outputs written in `stdout` and
the outputs written in `stdout` from *i*-th command are stored
in `stdout{i}.txt` file.
Those files will be accessed via `roadie result`.

#### upload
`upload` section takes a list of
[glob](https://en.wikipedia.org/wiki/Glob_(programming)) patterns.
Files matching one of those patterns are treated of results
and uploaded to a cloud storage.
To access those uploaded files, use `roadie result` command.

### Specify source codes on the fly
Sometimes, it is difficult to provide your source codes from the web,
such as Git repository, Dropbox, and some web site.
roadie helps to upload your source code from a local PC to a cloud storage,
which is a private place.
If you use this function, you can omit `source` section in your script file.

#### Upload source codes from a local directory
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
