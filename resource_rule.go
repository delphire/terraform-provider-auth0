package main

import (
	"github.com/bocodigitalmedia/go-auth0/auth0mgmt"
	"github.com/bocodigitalmedia/go-auth0/auth0rule"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceRule() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRuleExists,
		Create: resourceRuleCreate,
		Read:   resourceRuleRead,
		Update: resourceRuleUpdate,
		Delete: resourceRuleDelete,
		Schema: resourceSchemaMap(),
	}
}

func resourceSchemaMap() map[string]*schema.Schema {
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

	if _, ok := err.(*auth0mgmt.ApiError); ok {
		return false, nil
	}

	return err == nil, err
}

func resourceRuleCreate(d *schema.ResourceData, meta interface{}) error {
	svc := metaRuleService(meta)
	params := getRuleProperties(d)
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

	setRuleData(r, d)

	return nil
}

func resourceRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	svc := metaRuleService(meta)
	params := getRuleProperties(d)
	r, _, err := svc.Update(d.Id(), params)

	if err != nil {
		return err
	}

	setRuleData(r, d)
	return nil
}

func resourceRuleDelete(d *schema.ResourceData, meta interface{}) error {
	svc := metaRuleService(meta)
	_, err := svc.Delete(d.Id())
	return err
}

func getRuleProperties(d *schema.ResourceData) *auth0rule.Properties {
	return &auth0rule.Properties{
		Name:    d.Get("name").(string),
		Script:  d.Get("script").(string),
		Enabled: d.Get("enabled").(bool),
		Order:   d.Get("order"),
	}
}

func setRuleData(p *auth0rule.Rule, d *schema.ResourceData) {
	d.SetId(p.Id.(string))
	d.Set("name", p.Name.(string))
	d.Set("script", p.Script.(string))
	d.Set("enabled", p.Enabled.(bool))
	d.Set("order", p.Order)
}
