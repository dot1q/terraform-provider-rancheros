package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
        return &schema.Provider{
                ResourcesMap: map[string]*schema.Resource{
			"rancheros_apikey": resourceApiKey(),
			"rancheros_localauthconfig": resourceLocalAuthConfig(),
		},
        }
}
