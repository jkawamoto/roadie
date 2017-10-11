---
title: Installation
description: >-
  This document explains installation and initialization of Roadie.
  Currently, Roadie supports Google Cloud Platform and Microsoft Azure.
  Each cloud service requires different account information.
  This document also shows where you can get such information, too.
date: 2016-12-20
lastmod: 2017-10-10
slug: installation
---
### Install
Roadie's compiled binary files for some platforms are found in GitHub's
[release page](https://github.com/jkawamoto/roadie/releases).
Download one of them according to your environment and put the binary into
a directory in your `$PATH`, or put it as same directory as your current project.

If you're a [Homebrew](http://brew.sh/) or [Linuxbrew](http://linuxbrew.sh/)
user, you can install Roadie by the following commands:

```sh
$ brew tap jkawamoto/roadie
$ brew install roadie
```

### Initialization
The following command initializes Roadie:

```shell
$ roadie init
```

Since Roadie currently supports the following cloud services,

- <i class="fa fa-google" aria-hidden="true"></i> [Google Cloud Platform](https://cloud.google.com/),
- <i class="fa fa-windows" aria-hidden="true"></i> [Microsoft Azure](https://azure.microsoft.com/),

the initialization command asks which cloud service you will use first.

Then you need to tell Roadie some information depended on cloud services.

#### Google Cloud Platform
Google Cloud Platform manages resources based on project,
and Roadie needs to know which project it will work with.
After choosing Google Cloud Platform in the initialization, Roadie asks a project ID.

You can find available project IDs in [the cloud console](https://console.cloud.google.com/project).

<img src="img/gcp-projects-en.png" style="width: 80%;" title="Google Cloud Platform Projects"/>

Note that, Roadie requires a *project ID*, not *project name*.

After entering a project ID, Roadie starts authorization process and shows a URL.
You need to open the URL by a web browser on the same computer Roadie running
and grant permission.

Roadie then creates configuration file `roadie.yml` in the current directory.
Roadie has other configurations, see [configuration page](documents/configuration) for detailed information.


#### Microsoft Azure
In Microsoft Azure, each user account belongs to an Active Directory,
which is identified by a tenant ID,
and each user has several subscriptions.
Roadie hence needs to know both tenant ID and subscription ID.

To find your tenant ID, open Azure Active Directory tab
in the [portal site](https://portal.azure.com/).

<img src="img/active-directory-en.png" style="width: 80%;" title="Azure Active Directory"/>

Then, open Properties tab.
The *directory ID* in the tab is the your *tenant ID*.

<img src="img/tenant-id-en.png" style="width: 80%;" title="Tenant ID"/>

To find your subscription IDs, open Cost Management + Billing in the portal site.

<img src="img/subscription-en.png" style="width: 80%;" title="Cost Management + Billing"/>

In Overview tab, there is a list of your subscriptions.

<img src="img/subscription2-en.png" style="width: 80%;" title="Subscription ID"/>

Roadie also uses *project ID* to manage resources.
This ID must be unique in the world and can have lower alphabets and numbers,
but you can use any ID.

After entering a tenant ID, subscription ID, and project ID,
Roadie starts authorization process.
It shows a URL and a pass code.
You need to open the URL by a web browser and enter the code.

Roadie then creates configuration file `roadie.yml` in the current directory.
Roadie has other configurations, see [configuration page](documents/configuration) for detailed information.
