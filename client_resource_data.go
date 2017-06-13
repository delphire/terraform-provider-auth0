package main

import (
	"github.com/bocodigitalmedia/go-auth0/auth0client"
	"github.com/hashicorp/terraform/helper/schema"
)

func clientResourceDataToProperties(d *schema.ResourceData) *auth0client.Properties {
	p := &auth0client.Properties{}

	if raw, ok := d.GetOk("name"); ok {
		p.Name = raw.(string)
	}

	if raw, ok := d.GetOk("description"); ok {
		p.Description = raw.(string)
	}

	if raw, ok := d.GetOk("logo_uri"); ok {
		p.LogoUri = raw.(string)
	}

	if raw, ok := d.GetOk("callbacks"); ok {
		val := raw.([]interface{})
		p.Callbacks = &val
	}

	if raw, ok := d.GetOk("allowed_origins"); ok {
		val := raw.([]interface{})
		p.AllowedOrigins = &val
	}

	if raw, ok := d.GetOk("client_aliases"); ok {
		val := raw.([]interface{})
		p.ClientAliases = &val
	}

	if raw, ok := d.GetOk("allowed_clients"); ok {
		val := raw.([]interface{})
		p.AllowedClients = &val
	}

	if raw, ok := d.GetOk("allowed_logout_urls"); ok {
		val := raw.([]interface{})
		p.AllowedLogoutUrls = &val
	}

	if raw, ok := d.GetOk("grant_types"); ok {
		val := raw.([]interface{})
		p.GrantTypes = &val
	}

	if raw, ok := d.GetOk("token_endpoint_auth_method"); ok {
		p.TokenEndpointAuthMethod = raw.(string)
	}

	if raw, ok := d.GetOk("app_type"); ok {
		p.AppType = raw.(string)
	}

	if raw, ok := d.GetOk("oidc_conformant"); ok {
		p.OidcConformant = raw.(string) == "1"
	}

	if raw, ok := d.GetOk("jwt_configuration"); ok {
		val := raw.([]interface{})
		if len(val) == 1 {
			p.JwtConfiguration = clientJwtConfigurationFromData(val)
		}
	}

	if raw, ok := d.GetOk("encryption_key"); ok {
		val := raw.([]interface{})
		if len(val) == 1 {
			p.EncryptionKey = clientEncryptionKeyFromData(val)
		}
	}

	if raw, ok := d.GetOk("sso"); ok {
		p.Sso = raw.(string) == "1"
	}

	if raw, ok := d.GetOk("sso_disabled"); ok {
		p.SsoDisabled = raw.(string) == "1"
	}

	if raw, ok := d.GetOk("cross_origin_auth"); ok {
		p.CrossOriginAuth = raw.(string) == "1"
	}

	if raw, ok := d.GetOk("cross_origin_loc"); ok {
		p.CrossOriginLoc = raw.(string)
	}

	if raw, ok := d.GetOk("custom_login_page_on"); ok {
		p.CustomLoginPageOn = raw.(string) == "1"
	}

	if raw, ok := d.GetOk("custom_login_page"); ok {
		p.CustomLoginPage = raw.(string)
	}

	if raw, ok := d.GetOk("addons"); ok {
		val := raw.([]interface{})
		if len(val) == 1 {
			p.Addons = clientAddonsFromData(val)
		}
	}

	if raw, ok := d.GetOk("custom_login_page_preview"); ok {
		p.CustomLoginPagePreview = raw.(string) == "1"
	}

	if raw, ok := d.GetOk("form_template"); ok {
		p.FormTemplate = raw.(string)
	}

	if raw, ok := d.GetOk("is_heroku_app"); ok {
		p.IsHerokuApp = raw.(string) == "1"
	}

	if raw, ok := d.GetOk("client_metadata"); ok {
		val := raw.(map[string]interface{})
		p.ClientMetadata = &val
	}

	return p
}

func clientResourceDataFromClient(d *schema.ResourceData, c *auth0client.Client) error {
	if raw := c.ClientId; raw != nil {
		d.SetId(raw.(string))
	}
	if raw := c.Name; raw != nil {
		d.Set("name", raw.(string))
	}
	if raw := c.Description; raw != nil {
		d.Set("description", raw.(string))
	}
	if raw := c.LogoUri; raw != nil {
		d.Set("logo_uri", raw.(string))
	}
	if raw := c.Callbacks; raw != nil {
		d.Set("callbacks", *raw)
	}
	if raw := c.AllowedOrigins; raw != nil {
		d.Set("allowed_origins", *raw)
	}
	if raw := c.ClientAliases; raw != nil {
		d.Set("client_aliases", *raw)
	}
	if raw := c.AllowedClients; raw != nil {
		d.Set("allowed_clients", *raw)
	}
	if raw := c.AllowedLogoutUrls; raw != nil {
		d.Set("allowed_logout_urls", *raw)
	}
	if raw := c.GrantTypes; raw != nil {
		d.Set("grant_types", *raw)
	}
	if raw := c.TokenEndpointAuthMethod; raw != nil {
		d.Set("token_endpoint_auth_method", raw.(string))
	}
	if raw := c.AppType; raw != nil {
		d.Set("app_type", raw.(string))
	}
	if raw := c.OidcConformant; raw != nil {
		d.Set("oidc_conformant", boolString(raw.(bool)))
	}
	if raw := c.JwtConfiguration; raw != nil {
		d.Set("jwt_configuration", dataFromClientJwtConfiguration(raw))
	}
	if raw := c.EncryptionKey; raw != nil {
		d.Set("encryption_key", dataFromClientEncryptionKey(raw))
	}
	if raw := c.Sso; raw != nil {
		d.Set("sso", boolString(raw.(bool)))
	}
	if raw := c.SsoDisabled; raw != nil {
		d.Set("sso_disabled", boolString(raw.(bool)))
	}
	if raw := c.CrossOriginAuth; raw != nil {
		d.Set("cross_origin_auth", boolString(raw.(bool)))
	}
	if raw := c.CrossOriginLoc; raw != nil {
		d.Set("cross_origin_loc", raw.(string))
	}
	if raw := c.CustomLoginPageOn; raw != nil {
		d.Set("custom_login_page_on", boolString(raw.(bool)))
	}
	if raw := c.CustomLoginPage; raw != nil {
		d.Set("custom_login_page", raw.(string))
	}
	if raw := c.Addons; raw != nil {
		d.Set("addons", dataFromClientAddons(raw))
	}
	if raw := c.CustomLoginPagePreview; raw != nil {
		d.Set("custom_login_page_preview", raw.(string))
	}
	if raw := c.FormTemplate; raw != nil {
		d.Set("form_template", raw.(string))
	}
	if raw := c.IsHerokuApp; raw != nil {
		d.Set("is_heroku_app", boolString(raw.(bool)))
	}
	if raw := c.ClientMetadata; raw != nil {
		d.Set("client_metadata", *raw)
	}

	return nil
}

func clientAddonsFromData(src []interface{}) *auth0client.Addons {
	out := new(auth0client.Addons)
	data := src[0].(map[string]interface{})
	addon := func(key string) *map[string]interface{} {
		if raw, ok := data[key]; ok {
			val := raw.(map[string]interface{})
			return &val
		} else {
			return nil
		}
	}

	out.Aws = addon("aws")
	out.Firebase = addon("firebase")
	out.AzureBlob = addon("azure_blob")
	out.AzureSb = addon("azure_sb")
	out.Rms = addon("rms")
	out.MsCrm = addon("mscrm")
	out.Slack = addon("slack")
	out.Box = addon("box")
	out.CloudBees = addon("cloudbees")
	out.Concur = addon("concur")
	out.Dropbox = addon("dropbox")
	out.EchoSign = addon("echosign")
	out.Egnyte = addon("egnyte")
	out.NewRelic = addon("newrelic")
	out.Office365 = addon("office365")
	out.Salesforce = addon("salesforce")
	out.SalesforceApi = addon("salesforce_api")
	out.SalesforceSandboxApi = addon("salesforce_sandbox_api")
	out.SamlP = addon("samlp")
	out.Layer = addon("layer")
	out.SapApi = addon("sap_api")
	out.SharePoint = addon("sharepoint")
	out.SpringCm = addon("springcm")
	out.Wams = addon("wams")
	out.WsFed = addon("wsfed")
	out.Zendesk = addon("zendesk")
	out.Zoom = addon("zoom")

	return out
}

func clientEncryptionKeyFromData(src []interface{}) *auth0client.EncryptionKey {
	data := src[0].(map[string]interface{})
	out := new(auth0client.EncryptionKey)

	if raw, ok := data["pub"]; ok {
		out.Pub = raw.(string)
	}

	if raw, ok := data["cert"]; ok {
		out.Cert = raw.(string)
	}

	return out
}

func clientJwtConfigurationFromData(src []interface{}) *auth0client.JwtConfiguration {
	data := src[0].(map[string]interface{})
	out := new(auth0client.JwtConfiguration)

	if raw, ok := data["lifetime_in_seconds"]; ok {
		out.LifetimeInSeconds = float64(raw.(int))
	}

	if raw, ok := data["scopes"]; ok {
		val := raw.(map[string]interface{})
		out.Scopes = &val
	}

	if raw, ok := data["alg"]; ok {
		out.Alg = raw.(string)
	}

	return out
}

func dataFromClientAddons(src *auth0client.Addons) []interface{} {
	out := make([]interface{}, 1)
	data := make(map[string]interface{}, 0)
	addon := func(key string, raw *map[string]interface{}) {
		if raw != nil {
			data[key] = *raw
		}
	}

	addon("aws", src.Aws)
	addon("firebase", src.Firebase)
	addon("azure_blob", src.AzureBlob)
	addon("azure_sb", src.AzureSb)
	addon("rms", src.Rms)
	addon("mscrm", src.MsCrm)
	addon("slack", src.Slack)
	addon("box", src.Box)
	addon("cloudbees", src.CloudBees)
	addon("concur", src.Concur)
	addon("dropbox", src.Dropbox)
	addon("echosign", src.EchoSign)
	addon("egnyte", src.Egnyte)
	addon("newrelic", src.NewRelic)
	addon("office365", src.Office365)
	addon("salesforce", src.Salesforce)
	addon("salesforce_api", src.SalesforceApi)
	addon("salesforce_sandbox_api", src.SalesforceSandboxApi)
	addon("samlp", src.SamlP)
	addon("layer", src.Layer)
	addon("sap_api", src.SapApi)
	addon("sharepoint", src.SharePoint)
	addon("springcm", src.SpringCm)
	addon("wams", src.Wams)
	addon("wsfed", src.WsFed)
	addon("zendesk", src.Zendesk)
	addon("zoom", src.Zoom)

	out[0] = data
	return out
}

func dataFromClientEncryptionKey(src *auth0client.EncryptionKey) []interface{} {
	out := make([]interface{}, 1)
	data := make(map[string]interface{}, 0)

	if val := src.Pub; val != nil {
		data["pub"] = val.(string)
	}
	if val := src.Cert; val != nil {
		data["cert"] = val.(string)
	}

	out[0] = data
	return out
}

func dataFromClientJwtConfiguration(src *auth0client.JwtConfiguration) []interface{} {
	out := make([]interface{}, 1)
	data := make(map[string]interface{}, 0)

	if val := src.Alg; val != nil {
		data["alg"] = val.(string)
	}

	if val := src.LifetimeInSeconds; val != nil {
		data["lifetime_in_seconds"] = int(val.(float64))
	}

	if val := src.Scopes; val != nil {
		data["scopes"] = *val
	}

	out[0] = data
	return out
}
