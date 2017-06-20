package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func NewProvider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:        providerSchema(),
		ResourcesMap:  providerResourcesMap(),
		ConfigureFunc: Configure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"domain": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"access_token": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func providerResourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"auth0_rule":       NewRuleResource(),
		"auth0_client":     NewClientResource(),
		"auth0_connection": NewConnectionResource(),
	}
}
