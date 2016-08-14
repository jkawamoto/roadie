---
title: Configuration
---
To update configuration of roadie, use `roadie config` command.
It provides four sub commands:

- **project** shows and updates project ID of Google Cloud Platform,
- **zone** shows and updates zone used to run scripts,
- **type** shows and updates machine type used to run scripts,
- **bucket** shows and updates bucket name used to store related files.

Note that every configurations are stored to '.roadie'
in the current working directory.
You can also update configurations without this command by editing that file.

### project
Project ID is an ID registered in Google Cloud Platform.
You can find your project ID [here](https://console.cloud.google.com/project).

To check the project ID currently set to roadie, run

```sh
$ roadie config project
```

and to set another project ID `PROJECT`, run

```sh
$ roadie config project set PROJECT
```

Valid project ID is required to access Google Cloud Platform.
If you set correct project ID but any commands such as
`roadie config type list` fail,
make sure you have authenticated your computer by checking `gcloud auth list`
shows your account is credentialed.

### zone
In Google Cloud Platform, the platform is divided into several zones
based on actual locations where virtual machine will run.
You can find current available zones by running

```sh
$ roadie config zone list
```

By default, `us-central1-b` is chosen.
To check current zone, run

```sh
$ roadie config zone
```

and to set another `ZONE`, run

```sh
$ roadie config zone set ZONE
```

### type
There are some options about machine type on which your program runs.
Each machine type has different number of virtual CPUs and RAM.
You can find available machine types by running

```sh
$ roadie config type list
```

and [here](https://cloud.google.com/compute/pricing) is more information about
machine types and their pricing.
Available machine types might be depended to which zone you choose.
You should set zone before checking machine types.

By default, `n1-standard-1`, which has 1 vCPU and 3.75 GB RAM, is selected.
To check current machine type, run

```sh
$ roadie config type
```

and to set another `TYPE`, run

```sh
$ roadie config type set TYPE
```

### bucket
Bucket is a place to store any data including source codes and outputs from
your program.
Each bucket is identified by bucket name and it must be unique in the world.
You can choose any name but it is recommended to use same name as your project ID.
Actually, by default, roadie sets the project ID to the bucket name.

To check current bucket name, run

```sh
$ roadie config bucket
```

and to set another bucket name `NAME`, run

```sh
$ roadie config bucket set NAME
```
