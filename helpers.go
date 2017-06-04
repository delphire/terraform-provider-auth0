package main

import (
	"os"

	"github.com/hashicorp/terraform/helper/schema"
)

func schemaEnvRequiredString(key string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		DefaultFunc: envDefaultFunc(key),
	}
}

func envDefaultFunc(k string) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		if v := os.Getenv(k); v != "" {
			return v, nil
		}

		return nil, nil
	}
}

func envDefaultFuncAllowMissing(k string) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		v := os.Getenv(k)
		return v, nil
	}
}

func schemaString() *schema.Schema {
	return &schema.Schema{Type: schema.TypeString, Required: true}
}

func schemaBool() *schema.Schema {
	return &schema.Schema{Type: schema.TypeBool, Required: true}
}

func schemaIntOptional() *schema.Schema {
	return &schema.Schema{Type: schema.TypeInt, Optional: true}
}
