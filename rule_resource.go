package main

import "github.com/hashicorp/terraform/helper/schema"

func NewRuleResource() *schema.Resource {
	return &schema.Resource{
		Exists:   ruleResourceExists,
		Create:   ruleResourceCreate,
		Read:     ruleResourceRead,
		Update:   ruleResourceUpdate,
		Delete:   ruleResourceDelete,
		Importer: &schema.ResourceImporter{},
		Schema:   ruleResourceSchema(),
	}
}

func ruleResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":    schemaString(true),
		"script":  schemaString(false),
		"order":   schemaInt(false),
		"enabled": schemaBool(false),
	}
}

func ruleResourceCreate(d *schema.ResourceData, meta interface{}) error {
	return NewRuleResourceService(d, meta).Create()
}

func ruleResourceRead(d *schema.ResourceData, meta interface{}) error {
	return NewRuleResourceService(d, meta).Read()
}

func ruleResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	return NewRuleResourceService(d, meta).Update()
}

func ruleResourceDelete(d *schema.ResourceData, meta interface{}) error {
	return NewRuleResourceService(d, meta).Delete()
}

func ruleResourceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return NewRuleResourceService(d, meta).Exists()
}
