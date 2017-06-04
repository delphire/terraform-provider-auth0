package main

import (
	"github.com/bocodigitalmedia/go-auth0/auth0mgmt"
	"github.com/hashicorp/terraform/helper/schema"
)

type ProviderConfig struct {
	Domain      string
	AccessToken string
}

func (c *ProviderConfig) Meta() interface{} {
	api := auth0mgmt.NewApi(&auth0mgmt.NewApiParams{
		Domain:      c.Domain,
		AccessToken: c.AccessToken,
	})

	return NewMeta(api)
}

func ProviderConfigureFunc(d *schema.ResourceData) (interface{}, error) {
	config := &ProviderConfig{
		Domain:      d.Get("domain").(string),
		AccessToken: d.Get("access_token").(string),
	}

	return config.Meta(), nil
}
