package main

import (
	"github.com/bocodigitalmedia/go-auth0/auth0mgmt"
	"github.com/hashicorp/terraform/helper/schema"
)

type Config struct {
	Domain      string
	AccessToken string
}

func Configure(d *schema.ResourceData) (interface{}, error) {

	api := auth0mgmt.NewApi(&auth0mgmt.NewApiParams{
		Domain:      d.Get("domain").(string),
		AccessToken: d.Get("access_token").(string),
	})

	return NewMeta(api), nil
}
