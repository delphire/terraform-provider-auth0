package main

import "github.com/hashicorp/terraform/helper/schema"

func schemaOfType(t schema.ValueType, required interface{}) *schema.Schema {
	s := &schema.Schema{Type: t}
	val, isBool := required.(bool)

	if isBool && val {
		s.Required = true
	}

	if isBool && !val {
		s.Optional = true
	}

	return s
}

func schemaString(required interface{}) *schema.Schema {
	return schemaOfType(schema.TypeString, required)
}

func schemaBool(required interface{}) *schema.Schema {
	return schemaOfType(schema.TypeBool, required)
}

func schemaInt(required interface{}) *schema.Schema {
	return schemaOfType(schema.TypeInt, required)
}

func schemaFloat(required interface{}) *schema.Schema {
	return schemaOfType(schema.TypeFloat, required)
}

func schemaMap(required bool) *schema.Schema {
	return schemaOfType(schema.TypeMap, required)
}

func schemaList(required bool, elem *schema.Schema) *schema.Schema {
	s := schemaOfType(schema.TypeList, required)
	s.Elem = elem
	return s
}

func schemaSet(required bool, elem *schema.Schema) *schema.Schema {
	s := schemaOfType(schema.TypeSet, required)
	s.Elem = elem
	return s
}

func schemaResource(required bool, elem *schema.Resource) *schema.Schema {
	s := schemaOfType(schema.TypeSet, required)
	s.MaxItems = 1
	s.Elem = elem

	return s
}

func boolString(v bool) string {
	if v {
		return "1"
	} else {
		return "0"
	}
}
