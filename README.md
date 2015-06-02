Roadie
========
A helper container to execute a program and store results.

How to use
-----------

### 1. Build a container including your code.

You need to build a container based on Roadie.
In your container, your codes should be put in `/data`.
You can choose [Google Cloud Storage]([https://cloud.google.com/storage/) or [MongoDB](http://www.mongodb.org/) as a place to store outputs.

Here is a sample `Dockerfile` assuming your code is `main.py` and you choose `sample-bucket` in Google Cloud Storage as the storage.

```dockerfile:Dockerfile
FROM jkawamoto/roadie

# Import your code.
ADD main.py /data

# Set configurations.
CMD ["main.py", "gcs", "sample-bucket"]
```

Then, you can build the container named `yourname/main`.

```sh
$ docker build -t yourname/main .
```

### 2. Run the container.
Assuming the name of your container is `yourname/main`, you can run the container by the following command.

```sh
$ docker run -d yourname/main
```

Then, the executed result will be stored in a place specified in the Dockerfile.
In the above example, a file `stdout` consisting of outputs written in stdout will be stored in sample-bucket.
On the other hand, stderr will be connected to stdout of the container so that you can watch it via such as `docker logs`.
If your program produces other output files, see detailed information below.


Detailed information
---------------------
Containers delivered from roadie will execute a helper script at a run time.
Then, this script invoke your program. Actually, the script will invoke the first argument.
For example, you can run simple `echo` command by the following command.

```sh
$ docker run -d junpei/roadie "echo 'Hello world!'" gce "sample-bucket"
```

The format of the script is as follows.

```
usage: entrypoint.py [-h] [--observe [OBSERVE [OBSERVE ...]]] [--cwd CWD]
                     [--shutdown] [--zip] [--log LOG] [--stderr STDERR]
                     cmd {gcs,mongo,local} ...

positional arguments:
  cmd                   Command line to be run.
  {gcs,mongo,local}
    gcs                 Storing outputs into Google Cloud Storage.
    mongo               Storing outputs into MongoDB.
    local               Storing outputs into local filesystem.

optional arguments:
  -h, --help            show this help message and exit
  --observe [OBSERVE [OBSERVE ...]]
                        glob patterns of files to be stored.
  --cwd CWD             Change working directory (default: /data).
  --shutdown            Shutdown after the program ends (working only in
                        Google Compute Engine)
  --zip                 Files specified in overve option will be zipped.
  --log LOG             Config file of loggers.
  --stderr STDERR       Specify where stderr should be stored (default:
                        stdout).
```

If your program will produce some files, specify file-name patterns at `--observe` argument to be stored.
UNIX based patterns such as `result-*.out`, etc. are supported.
`cmd` should be quoted if it has spaces.

### Store to Google Cloud Storage
To store outputs to GCS, use `gcs` option.

```
usage: cmd gcs [-h] [--prefix PREFIX] bucket

positional arguments:
  bucket           Bucket name.

optional arguments:
  -h, --help       show this help message and exit
  --prefix PREFIX  Prefix of stored files.
```

### Store to MongoDB
To store outputs to MongoDB, use `mongo` option.

```
usage: cmd mongo [-h] [--host HOST] [--port PORT] db collection

positional arguments:
  db           Database name.
  collection   Collection name.

optional arguments:
  -h, --help   show this help message and exit
  --host HOST  Host name of MongoDB server.
  --port PORT  Port number of MongoDB server.
```

### Store to local filesystem
To store outputs to local, use `local` option.

```
usage: cmd local [-h] dir

positional arguments:
  dir         Directory for storing output.

optional arguments:
  -h, --help  show this help message and exit
```

License
--------
This software is released under the MIT License, see LICENSE.
