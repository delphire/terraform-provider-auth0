package main

import "github.com/hashicorp/terraform/helper/schema"

func ruleResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":    schemaString(true),
		"script":  schemaString(true),
		"order":   schemaInt(false),
		"enabled": schemaString(false),
	}
}
