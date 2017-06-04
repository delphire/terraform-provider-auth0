package main

import (
	"github.com/bocodigitalmedia/go-auth0/auth0mgmt"
	"github.com/bocodigitalmedia/go-auth0/auth0rule"
)

type Meta struct {
	RuleService *auth0rule.Service
}

func NewMeta(api *auth0mgmt.Api) *Meta {
	return &Meta{
		RuleService: &auth0rule.Service{api},
	}
}

func metaRuleService(meta interface{}) *auth0rule.Service {
	return meta.(*Meta).RuleService
}
