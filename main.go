package main

import (
	"github.com/Wikia/terraform-provider-dyn/dyn"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dyn.Provider})
}
