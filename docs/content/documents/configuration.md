---
title: Configuration
description: >-
  This document explains configurations to use Google Cloud Platform and
  Microsoft Azure.
date: 2016-08-14
lastmod: 2017-10-09
slug: configuration
---
Roadie's configuration is stored as a YAML document; by default `roadie.yml`,
and it is editable by any text editors.
In addition, Roadie provides `roadie config` command to edit configurations
interactively.

This document first explain `roadie config` command, and then the YAML based
configuration file.


### config command
`roadie config` command shows current configurations and available options, and
update them.
It provides the following sub commands:

- **project** shows and updates project ID,
- **region** shows and updates region where programs will be run,
- **machine** shows and updates machine type used to run scripts,

Every configurations are stored to `roadie.yml`
in the current working directory by default.
If you want to use another file, use `-c` flag with a file path.

For example, the following command sets "another" as the project ID and
stores it to `another.yml`:

```shell
$ roadie -c another.yml config project set another
```

#### project
To check the project ID currently set, run

```shell
$ roadie config project
```

and to set another project ID `PROJECT`, run

```shell
$ roadie config project set PROJECT
```

#### region
In most of cloud services, they have lots of servers in the world and
those servers are grouped based on actual locations.
Roadie needs to know which group of servers it should use to run programs.

You can find available regions in the cloud platform you're using by running
the following command:

```shell
$ roadie config region list
```

To find current selected region, run

```shell
$ roadie config region
```

and to set another `REGION`, run

```shell
$ roadie config region set REGION
```

### machine
There are some options about machine type on which your program runs.
Each machine type has different number of virtual CPUs and RAM.
You can find available machine types by running

```shell
$ roadie config machine list
```

Available machine types might be depended to which zone you choose.
You should set zone before checking machine types.

To check current machine type, run

```shell
$ roadie config machine
```

and to set another `TYPE`, run

```shell
$ roadie config machine set TYPE
```


### Configuration file
Roadie's configuration file is a YAML document.
The elements in the document are different based on which cloud provider,
Google Cloud Platform or Microsoft Azure.

#### Google Cloud Platform
If you choose [Google Cloud Platform](https://cloud.google.com/), the root
element of the configuration file is `gcp`, and it has child elements as shown
below:

```yaml
gcp:
  project: <project ID>
  bucket: <bucket ID>
  zone: us-central1-b
  machine_type: n1-standard-1
  disk_size: 10
```

- `project`: project ID registered in [Google Cloud Platform](https://console.cloud.google.com/project).
- `bucket`: bucket ID of [Cloud Storage](https://cloud.google.com/storage/) to
  store resources. This ID also has to be unique in the world. By default, the
  same ID as the project ID is set.
- `zone`: name of region where programs will be run. By default, one of the
  cheapest region `us-central1-b` is selected.
- `machine_type`: machine type. By default, one of the cheapest type
  `n1-standard-1` is selected.
- `disk_size`: disk size in GB of virtual machines hosting programs. By default,
  10 GB is selected.

#### Microsoft Azure
If you choose [Microsoft Azure](https://azure.microsoft.com/), the root element
of the configuration file is `azure`, and it has child elements as shown below:

```yaml
azure:
  tenant_id: <directory ID (tenant ID)>
  subscription_id: <subscription ID>
  project_id: <project ID>
  location: westus2
  machine_type: Standard_A2
  os:
    publisher_name: Canonical
    offer: UbuntuServer
    skus: "17.04"
    version: latest
```

- `tenant_id`: your tenant ID.
- `subscription_id`: your subscription ID.
- `project_id`: project ID.
- `location`: location where programs will be run (default: westus2).
- `machine_type`: machine type (default: Standard_A2).
- `os`: OS type to be used. Currently, only default values are supported.
