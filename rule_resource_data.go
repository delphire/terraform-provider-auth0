package main

import (
	"github.com/bocodigitalmedia/go-auth0/auth0rule"
	"github.com/hashicorp/terraform/helper/schema"
)

func ruleResourceDataToProperties(d *schema.ResourceData) *auth0rule.Properties {
	p := new(auth0rule.Properties)

	if raw, ok := d.GetOk("name"); ok {
		p.Name = raw.(string)
	}

	if raw, ok := d.GetOk("script"); ok {
		p.Script = raw.(string)
	}

	if raw, ok := d.GetOk("order"); ok {
		p.Order = float64(raw.(int))
	}

	if raw, ok := d.GetOk("enabled"); ok {
		if val, ok := raw.(string); ok && val != "" {
			p.Enabled = val == "1"
		}
	}

	return p
}

func ruleResourceDataFromRule(d *schema.ResourceData, r *auth0rule.Rule) error {
	if raw := r.Id; raw != nil {
		d.SetId(raw.(string))
	}

	if raw := r.Name; raw != nil {
		d.Set("name", raw.(string))
	}

	if raw := r.Script; raw != nil {
		d.Set("script", raw.(string))
	}

	if raw := r.Order; raw != nil {
		d.Set("order", int(raw.(float64)))
	}

	if raw := r.Enabled; raw != nil {
		d.Set("enabled", boolString(raw.(bool)))
	}

	return nil
}
