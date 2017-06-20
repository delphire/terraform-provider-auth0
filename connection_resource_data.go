package main

import (
	"github.com/bocodigitalmedia/go-auth0/auth0connection"
	"github.com/hashicorp/terraform/helper/schema"
)

func connectionResourceDataToProperties(d *schema.ResourceData) *auth0connection.Properties {
	p := new(auth0connection.Properties)

	if raw, ok := d.GetOk("name"); ok {
		p.Name = raw.(string)
	}

	if raw, ok := d.GetOk("strategy"); ok {
		p.Strategy = raw.(string)
	}

	if raw, ok := d.GetOk("options"); ok {
		val := raw.(map[string]interface{})
		p.Options = &val
	}

	if raw, ok := d.GetOk("enabled_clients"); ok {
		val := raw.([]interface{})
		p.EnabledClients = &val
	}

	if raw, ok := d.GetOk("realms"); ok {
		val := raw.([]interface{})
		if len(val) > 0 {
			p.Realms = &val
		}
	}

	return p
}

func connectionResourceDataFromConnection(d *schema.ResourceData, c *auth0connection.Connection) error {

	if raw := c.Id; raw != nil {
		d.SetId(raw.(string))
	}

	if raw := c.Name; raw != nil {
		d.Set("name", raw.(string))
	}

	if raw := c.Strategy; raw != nil {
		d.Set("strategy", raw.(string))
	}

	if raw := c.Options; raw != nil {
		d.Set("options", *raw)
	}

	if raw := c.EnabledClients; raw != nil {
		d.Set("enabled_clients", *raw)
	}

	if raw := c.Realms; raw != nil {
		val := *raw
		if len(val) > 0 {
			d.Set("realms", val)
		}
	}

	return nil
}
