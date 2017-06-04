package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func ProviderFunc() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:        providerSchema(),
		ResourcesMap:  providerResourcesMap(),
		ConfigureFunc: ProviderConfigureFunc,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"domain":       schemaEnvRequiredString("AUTH0_DOMAIN"),
		"access_token": schemaEnvRequiredString("AUTH0_ACCESS_TOKEN"),
	}
}

func providerResourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"auth0_rule": resourceRule(),
	}
}
