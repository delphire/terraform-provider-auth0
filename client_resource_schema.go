package main

import "github.com/hashicorp/terraform/helper/schema"

func clientResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":                       schemaString(true),
		"description":                schemaString(false),
		"logo_uri":                   schemaString(false),
		"callbacks":                  schemaList(false, schemaString(nil)),
		"allowed_origins":            schemaList(false, schemaString(nil)),
		"client_aliases":             schemaList(false, schemaString(nil)),
		"allowed_clients":            schemaList(false, schemaString(nil)),
		"allowed_logout_urls":        schemaList(false, schemaString(nil)),
		"grant_types":                schemaList(false, schemaString(nil)),
		"token_endpoint_auth_method": schemaString(false),
		"app_type":                   schemaString(false),
		"oidc_conformant":            schemaString(false),
		"jwt_configuration": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     clientJwtConfigurationResource(),
		},
		"encryption_key": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     clientEncryptionKeyResource(),
		},
		"sso":                       schemaString(false),
		"sso_disabled":              schemaString(false),
		"cross_origin_auth":         schemaString(false),
		"cross_origin_loc":          schemaString(false),
		"custom_login_page_on":      schemaString(false),
		"custom_login_page":         schemaString(false),
		"custom_login_page_preview": schemaString(false),
		"form_template":             schemaString(false),
		"is_heroku_app":             schemaString(false),
		"addons": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     clientAddonsResource(),
		},
		"client_metadata": schemaMap(false),
		// "android_app_package_name":         schemaString(true),
		// "android_sha256_cert_fingerprints": schemaList(true, schemaString(nil)),
		// "ios_team_id":                      schemaString(true),
		// "ios_app_bundle_identifier":        schemaString(true),
	}
}

func clientJwtConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"lifetime_in_seconds": schemaInt(true),
			"scopes":              schemaMap(false),
			"alg":                 schemaString(false),
		},
	}
}

func clientEncryptionKeyResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pub":  schemaString(true),
			"cert": schemaString(true),
		},
	}
}

func clientAddonsResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aws":                    schemaMap(false),
			"azure_blob":             schemaMap(false),
			"azure_sb":               schemaMap(false),
			"rms":                    schemaMap(false),
			"mscrm":                  schemaMap(false),
			"slack":                  schemaMap(false),
			"box":                    schemaMap(false),
			"cloudbees":              schemaMap(false),
			"concur":                 schemaMap(false),
			"dropbox":                schemaMap(false),
			"echosign":               schemaMap(false),
			"egnyte":                 schemaMap(false),
			"firebase":               schemaMap(false),
			"newrelic":               schemaMap(false),
			"office365":              schemaMap(false),
			"salesforce":             schemaMap(false),
			"salesforce_api":         schemaMap(false),
			"salesforce_sandbox_api": schemaMap(false),
			"samlp":                  schemaMap(false),
			"layer":                  schemaMap(false),
			"sap_api":                schemaMap(false),
			"sharepoint":             schemaMap(false),
			"springcm":               schemaMap(false),
			"wams":                   schemaMap(false),
			"wsfed":                  schemaMap(false),
			"zendesk":                schemaMap(false),
			"zoom":                   schemaMap(false),
		},
	}
}

func clientMobileResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"android": schemaResource(false, clientMobileAndroidResource()),
			"ios":     schemaResource(false, clientMobileIosResource()),
		},
	}
}

func clientMobileAndroidResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_package_name":         schemaString(true),
			"sha256_cert_fingerprints": schemaList(true, schemaString(nil)),
		},
	}
}

func clientMobileIosResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"team_id":               schemaString(true),
			"app_bundle_identifier": schemaString(true),
		},
	}
}
