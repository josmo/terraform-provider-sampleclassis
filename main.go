package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/josmo/terraform-provider-sampleclassis/classis"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: classis.Provider,
	})
}
