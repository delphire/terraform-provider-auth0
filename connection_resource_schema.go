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
		"google_apps_strategy_options": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem:     connectionGoogleAppsOptionsResource(),
		},
		"enabled_clients": schemaList(false, schemaString(nil)),
		// "realms": {
		// 	Type:     schema.TypeList,
		// 	Optional: true,
		// 	Elem:     schemaString(nil),
		// },
	}
}

func connectionGoogleAppsOptionsResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_enable_users": schemaString(false), // bool
			"client_id":        schemaString(false),
			"client_secret":    schemaString(false),
			"domain_aliases": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     schemaString(nil),
			},
			"email":                    schemaString(false), // bool
			"ext_agreed_terms":         schemaString(false), // bool
			"ext_groups":               schemaString(false), // bool
			"ext_is_admin":             schemaString(false), // bool
			"ext_is_suspended":         schemaString(false), //bool
			"global":                   schemaString(false), //bool
			"handle_login_from_social": schemaString(false), // bool
			"profile":                  schemaString(false), // bool
			"scope": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     schemaString(nil),
			},
			"status":        schemaString(false), //bool
			"tenant_domain": schemaString(false),
		},
	}
}
