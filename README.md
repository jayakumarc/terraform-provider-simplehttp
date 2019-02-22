A simple data source provider
=============================

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.11.x
- [Go](https://golang.org/doc/install) 1.11.x

Building the provider
---------------------

```bash
git clone https://github.com/jayakumarc/terraform-provider-simplehttp.git
cd terraform-provider-simplehttp
go build -o ~/.terraform.d/plugins/terraform-provider-simplehttp
```

Example usage
-------------

```bash
cd example
terraform init
terraform apply
```