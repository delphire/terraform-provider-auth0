package main

import (
	"github.com/bocodigitalmedia/go-auth0/auth0connection"
	"github.com/hashicorp/terraform/helper/schema"
)

func connectionResourceDataToProperties(d *schema.ResourceData) *auth0connection.Properties {
	p := new(auth0connection.Properties)

	if raw, ok := d.GetOk("name"); ok && d.HasChange("name") {
		p.Name = raw.(string)
	}

	if raw, ok := d.GetOk("strategy"); ok {
		strategy := raw.(string)

		if d.HasChange("strategy") {
			p.Strategy = strategy
		}

		switch strategy {
		case "google-apps":
			if raw, ok := d.GetOk("google_apps_strategy_options"); ok {
				val := raw.([]interface{})
				if len(val) == 1 {
					p.Options = connectionGoogleAppsStrategyOptionsFromData(val)
				}
			}
		}
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

	if raw, err := c.GetOptions(); err != nil {
		return err
	} else if raw != nil {
		switch c.Strategy {
		case "google-apps":
			val := raw.(*auth0connection.GoogleAppsStrategyOptions)
			d.Set("google_apps_strategy_options", dataFromConnectionGoogleAppsStrategyOptions(val))
		}
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

func dataFromConnectionGoogleAppsStrategyOptions(src *auth0connection.GoogleAppsStrategyOptions) []interface{} {
	out := make([]interface{}, 1)
	data := make(map[string]interface{}, 0)

	if val := src.ApiEnableUsers; val != nil {
		data["api_enable_users"] = boolString(val.(bool))
	}
	if val := src.ClientId; val != nil {
		data["client_id"] = val.(string)
	}
	if val := src.ClientSecret; val != nil {
		data["client_secret"] = val.(string)
	}
	if val := src.DomainAliases; val != nil {
		data["domain_aliases"] = *val
	}
	if val := src.Email; val != nil {
		data["email"] = boolString(val.(bool))
	}
	if val := src.ExtGroups; val != nil {
		data["ext_groups"] = boolString(val.(bool))
	}
	if val := src.ExtAgreedTerms; val != nil {
		data["ext_agreed_terms"] = boolString(val.(bool))
	}
	if val := src.ExtIsAdmin; val != nil {
		data["ext_is_admin"] = boolString(val.(bool))
	}
	if val := src.ExtIsSuspended; val != nil {
		data["ext_is_suspended"] = boolString(val.(bool))
	}
	if val := src.Global; val != nil {
		data["global"] = boolString(val.(bool))
	}
	if val := src.HandleLoginFromSocial; val != nil {
		data["handle_login_from_social"] = boolString(val.(bool))
	}
	if val := src.Profile; val != nil {
		data["profile"] = boolString(val.(bool))
	}
	if val := src.Scope; val != nil {
		data["scope"] = *val
	}
	if val := src.Status; val != nil {
		data["status"] = boolString(val.(bool))
	}
	if val := src.TenantDomain; val != nil {
		data["tenant_domain"] = val.(string)
	}

	out[0] = data
	return out
}

func connectionGoogleAppsStrategyOptionsFromData(src []interface{}) *auth0connection.GoogleAppsStrategyOptions {
	data := src[0].(map[string]interface{})
	out := new(auth0connection.GoogleAppsStrategyOptions)

	if raw, ok := data["api_enable_users"]; ok {
		out.ApiEnableUsers = raw.(string) == "1"
	}
	if raw, ok := data["client_id"]; ok {
		out.ClientId = raw.(string)
	}
	if raw, ok := data["client_secret"]; ok {
		out.ClientSecret = raw.(string)
	}
	if raw, ok := data["domain_aliases"]; ok {
		val := interfaceSliceToStringSlice(raw.([]interface{}))
		out.DomainAliases = &val
	}
	if raw, ok := data["email"]; ok {
		out.Email = raw.(string) == "1"
	}
	if raw, ok := data["ext_agreed_terms"]; ok {
		out.ExtAgreedTerms = raw.(string) == "1"
	}
	if raw, ok := data["ext_groups"]; ok {
		out.ExtGroups = raw.(string) == "1"
	}
	if raw, ok := data["ext_is_admin"]; ok {
		out.ExtIsAdmin = raw.(string) == "1"
	}
	if raw, ok := data["ext_is_suspended"]; ok {
		out.ExtIsSuspended = raw.(string) == "1"
	}
	if raw, ok := data["global"]; ok {
		out.Global = raw.(string) == "1"
	}
	if raw, ok := data["handle_login_from_social"]; ok {
		out.HandleLoginFromSocial = raw.(string) == "1"
	}
	if raw, ok := data["profile"]; ok {
		out.Profile = raw.(string) == "1"
	}
	if raw, ok := data["scope"]; ok {
		val := interfaceSliceToStringSlice(raw.([]interface{}))
		out.Scope = &val
	}
	if raw, ok := data["status"]; ok {
		out.Status = raw.(string) == "1"
	}
	if raw, ok := data["tenant_domain"]; ok {
		out.TenantDomain = raw.(string)
	}

	return out
}
