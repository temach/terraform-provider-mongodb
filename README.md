# Terraform Provider Mongodb

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/fabiovpcaumo/terraform-provider-mongodb?logo=go&style=flat-square)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/fabiovpcaumo/terraform-provider-mongodb?logo=git&style=flat-square)
![GitHub](https://img.shields.io/github/license/fabiovpcaumo/terraform-provider-mongodb?color=yellow&style=flat-square)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/fabiovpcaumo/terraform-provider-mongodb/golangci?logo=github&style=flat-square)
![GitHub issues](https://img.shields.io/github/issues/fabiovpcaumo/terraform-provider-mongodb?logo=github&style=flat-square)

This repository is a [Terraform](https://www.terraform.io) MongoDB/DocumentDB provider forked from Kaginari/terraform-provider-mongodb.

# Using this provider from GitHub directly

```terraform
module "mongodb" {
    source = "github.com/fabiovpcaumo/terraform-provider-mongodb.git?ref=v<desired_version>"
}

```

### Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13
- [Go](https://golang.org/doc/install) >= 1.18

### Installation

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the `make install` command:

```bash
git clone https://github.com/fabiovpcaumo/terraform-provider-mongodb
cd terraform-provider-mongodb
make install
```

### To test the provider locally

**1.1: start the docker-compose**

```bash
cd docker
docker-compose up -d
```

**1.2 : create admin user in mongo**

```bash
$ docker exec -it mongo bash
> mongo
> use admin
> db.createUser({ user: "root" , pwd: "root", roles: ["userAdminAnyDatabase", "dbAdminAnyDatabase", "readWriteAnyDatabase"]})
```

**1.3 : accessing the local MongoDB via Mongo Express**

By default, the docker compose exposes a Mongo Express container at localhost:8081.

**2: Build the provider**

follow the [Installation](#Installation)

**3: Use the provider**

You are now ready to use the local provider as you like.

For an example code you can use:

```bash
cd mongodb
make apply
```
