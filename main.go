package rancheros

import (
	"github.com/dot1q/terraform-provider-rancheros/rancheros"
        "github.com/hashicorp/terraform/plugin"
        "github.com/hashicorp/terraform/terraform"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() terraform.ResourceProvider {
                        return Provider()
                },
        })
}
