package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/matthewmueller/terraform-provider-url/url"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: url.Provider,
	})
}
