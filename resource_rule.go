package main

import (
	"github.com/bocodigitalmedia/go-auth0/auth0mgmt"
	"github.com/bocodigitalmedia/go-auth0/auth0rule"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceRule() *schema.Resource {
	return &schema.Resource{
		Exists:   resourceRuleExists,
		Create:   resourceRuleCreate,
		Read:     resourceRuleRead,
		Update:   resourceRuleUpdate,
		Delete:   resourceRuleDelete,
		Importer: &schema.ResourceImporter{},
		Schema:   resourceRuleSchema(),
	}
}

func resourceRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":    schemaString(),
		"enabled": schemaBool(),
		"script":  schemaString(),
		"order":   schemaIntOptional(),
	}
}

func resourceRuleExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	svc := metaRuleService(meta)
	_, _, err := svc.Read(d.Id(), nil)

	if resourceRuleIsApi404Err(err) {
		return false, nil
	}

	return err == nil, err
}

func resourceRuleCreate(d *schema.ResourceData, meta interface{}) error {
	svc := metaRuleService(meta)
	params := resourceRuleGetProperties(d)
	r, _, err := svc.Create(params)

	if err != nil {
		return err
	}

	d.SetId(r.Id.(string))

	return nil
}

func resourceRuleRead(d *schema.ResourceData, meta interface{}) error {
	svc := metaRuleService(meta)
	r, _, err := svc.Read(d.Id(), nil)

	if err != nil {
		return err
	}

	resourceRuleSetData(r, d)

	return nil
}

func resourceRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	svc := metaRuleService(meta)
	params := resourceRuleGetProperties(d)
	r, _, err := svc.Update(d.Id(), params)

	if err != nil {
		return err
	}

	resourceRuleSetData(r, d)
	return nil
}

func resourceRuleDelete(d *schema.ResourceData, meta interface{}) error {
	svc := metaRuleService(meta)
	_, err := svc.Delete(d.Id())

	if resourceRuleIsApi404Err(err) {
		return nil
	} else {
		return err
	}
}

func resourceRuleGetProperties(d *schema.ResourceData) *auth0rule.Properties {
	return &auth0rule.Properties{
		Name:    d.Get("name").(string),
		Script:  d.Get("script").(string),
		Enabled: d.Get("enabled").(bool),
		Order:   d.Get("order"),
	}
}

func resourceRuleSetData(p *auth0rule.Rule, d *schema.ResourceData) {
	// d.SetId(p.Id.(string))
	d.Set("name", p.Name.(string))
	d.Set("script", p.Script.(string))
	d.Set("enabled", p.Enabled.(bool))
	d.Set("order", p.Order)
}

func resourceRuleIsApi404Err(err error) bool {
	apiErr, ok := err.(*auth0mgmt.ApiError)
	return ok && apiErr.StatusCode == 404
}
