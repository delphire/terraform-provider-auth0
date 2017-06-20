package main

import "github.com/hashicorp/terraform/helper/schema"

func connectionResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"strategy": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"options":         schemaMap(false),
		"enabled_clients": schemaList(false, schemaString(nil)),
		"realms": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     schemaString(nil),
		},
	}
}

func connectionOptionsResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"validation":               schemaMap(false),
			"password_policy":          schemaString(false),
			"password_history":         schemaMap(false),
			"password_no_personal_inf": schemaMap(false),
			"password_dictionary":      schemaMap(false),
		},
	}
}
