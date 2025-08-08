# Terraform Provider Mongodb

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/fabiovpcaumo/terraform-provider-mongodb?logo=go&style=flat-square)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/fabiovpcaumo/terraform-provider-mongodb?logo=git&style=flat-square)
![GitHub](https://img.shields.io/github/license/fabiovpcaumo/terraform-provider-mongodb?color=yellow&style=flat-square)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/fabiovpcaumo/terraform-provider-mongodb/golangci?logo=github&style=flat-square)
![GitHub issues](https://img.shields.io/github/issues/fabiovpcaumo/terraform-provider-mongodb?logo=github&style=flat-square)

This repository is a [Terraform](https://www.terraform.io) MongoDB/DocumentDB provider forked from Kaginari/terraform-provider-mongodb and fabiovpcaumo/terraform-provider-mongodb.

This fork adds the ability to create users that use external authentication mechanisms, such as using AWS IAM user or role to authenticate to the database.

# Using the provider and example password-less user

```terraform
terraform {
  required_providers {
    mongodb = {
      source = "temach/mongodb"
      version = "x.y.z" # Specify your desired version here
    }
  }
}

provider "mongodb" {
  # Configuration options
}

# This example creates a database user with authentication via AWS IAM user or role. `auth_database` must be "$external"
# and `auth_mechanisms` must include ["MONGODB-AWS"] per documentation:
# https://docs.aws.amazon.com/documentdb/latest/developerguide/iam-identity-auth.html#iam-identity-auth-get-started
resource "mongodb_db_user" "passwordless_user" {
  name = "arn:aws:iam::123456789123:role/iamrole" # Or use an IAM user, example: "arn:aws:iam::123456789123:user/iamuser"
  auth_database = "$external"
  auth_mechanisms = ["MONGODB-AWS"]
  role {
    role = "read"
    db =   "readDB"
  }
  role {
    role = "readWrite"
    db =   "readWriteDB"
  }
}
```

### Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13
- [Go](https://golang.org/doc/install) >= 1.17

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
cd examples
make apply
```
