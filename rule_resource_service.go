package main

import (
	"github.com/bocodigitalmedia/go-auth0/auth0rule"
	"github.com/hashicorp/terraform/helper/schema"
)

type RuleResourceService struct {
	*schema.ResourceData
	meta interface{}
}

func (r *RuleResourceService) Meta() *Meta {
	return r.meta.(*Meta)
}

func (r *RuleResourceService) Service() *auth0rule.Service {
	return r.Meta().RuleService
}

func (r *RuleResourceService) Properties() *auth0rule.Properties {
	return &auth0rule.Properties{
		Name:    r.Get("name").(string),
		Script:  r.Get("script").(string),
		Order:   r.Get("order").(int),
		Enabled: r.Get("enabled").(bool),
	}
}

func (r *RuleResourceService) FromRule(c *auth0rule.Rule) {
	r.SetId(c.Id.(string))
	r.Set("name", c.Name)
	r.Set("script", c.Script)
	r.Set("order", c.Order)
	r.Set("enabled", c.Enabled)
}

func (r *RuleResourceService) Create() error {
	params := r.Properties()

	if rule, _, err := r.Service().Create(params); err != nil {
		return err
	} else {
		r.FromRule(rule)
		return nil
	}
}

func (r *RuleResourceService) Read() error {
	if rule, _, err := r.Service().Read(r.Id(), nil); err != nil {
		return err
	} else {
		r.FromRule(rule)
		return nil
	}
}

func (r *RuleResourceService) Update() error {
	if rule, _, err := r.Service().Update(r.Id(), r.Properties()); err != nil {
		return err
	} else {
		r.FromRule(rule)
		return nil
	}
}

func (r *RuleResourceService) Delete() error {
	if _, err := r.Service().Delete(r.Id()); err != nil && !isApi404Err(err) {
		return err
	} else {
		return nil
	}
}

func (r *RuleResourceService) Exists() (bool, error) {
	if _, _, err := r.Service().Read(r.Id(), nil); err != nil && isApi404Err(err) {
		return false, nil
	} else {
		return err == nil, err
	}
}

func NewRuleResourceService(d *schema.ResourceData, meta interface{}) *RuleResourceService {
	return &RuleResourceService{d, meta}
}
