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

func ruleResourceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	_, _, err := metaRuleService(meta).Read(d.Id(), nil)

	if err != nil && isApi404Err(err) {
		return false, nil
	} else {
		return err == nil, err
	}
}

func ruleResourceCreate(d *schema.ResourceData, meta interface{}) error {
	props := ruleResourceDataToProperties(d)

	if rule, _, err := metaRuleService(meta).Create(props); err != nil {
		return err
	} else {
		d.SetId(rule.Id.(string))
		return nil
	}
}

func ruleResourceRead(d *schema.ResourceData, meta interface{}) error {
	if rule, _, err := metaRuleService(meta).Read(d.Id(), nil); err != nil {
		return err
	} else {
		return ruleResourceDataFromRule(d, rule)
	}
}

func ruleResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	props := ruleResourceDataToProperties(d)

	if _, _, err := metaRuleService(meta).Update(d.Id(), props); err != nil {
		return err
	} else {
		return nil
	}
}

func ruleResourceDelete(d *schema.ResourceData, meta interface{}) error {
	if _, err := metaRuleService(meta).Delete(d.Id()); err != nil {
		return err
	} else {
		return nil
	}
}
