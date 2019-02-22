package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jayakumarc/terraform-provider-simplehttp/simplehttp"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: simplehttp.HttpProvider,
	})
}
