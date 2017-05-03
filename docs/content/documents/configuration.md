---
title: Configuration
---
To update configuration of roadie, use `roadie config` command.
It provides the following sub commands:

- **project** shows and updates project ID of Google Cloud Platform,
- **region** shows and updates region used to run scripts,
- **machine** shows and updates machine type used to run scripts,

Note that every configurations are stored to `roadie.yml`
in the current working directory by default.
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

### region
In Google Cloud Platform, the platform is divided into several zones
based on actual locations where virtual machine will run.
You can find current available zones by running

```sh
$ roadie config region list
```

By default, `us-central1-b` is chosen.
To check current zone, run

```sh
$ roadie config region
```

and to set another `ZONE`, run

```sh
$ roadie config region set ZONE
```

### machine
There are some options about machine type on which your program runs.
Each machine type has different number of virtual CPUs and RAM.
You can find available machine types by running

```sh
$ roadie config machine list
```

and [here](https://cloud.google.com/compute/pricing) is more information about
machine types and their pricing.
Available machine types might be depended to which zone you choose.
You should set zone before checking machine types.

By default, `n1-standard-1`, which has 1 vCPU and 3.75 GB RAM, is selected.
To check current machine type, run

```sh
$ roadie config machine
```

and to set another `TYPE`, run

```sh
$ roadie config machine set TYPE
```
