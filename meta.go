package main

import (
	"github.com/bocodigitalmedia/go-auth0/auth0client"
	"github.com/bocodigitalmedia/go-auth0/auth0connection"
	"github.com/bocodigitalmedia/go-auth0/auth0mgmt"
	"github.com/bocodigitalmedia/go-auth0/auth0rule"
)

type Meta struct {
	RuleService       *auth0rule.Service
	ClientService     *auth0client.Service
	ConnectionService *auth0connection.Service
}

func NewMeta(api *auth0mgmt.Api) *Meta {
	return &Meta{
		RuleService:       &auth0rule.Service{api},
		ClientService:     &auth0client.Service{api},
		ConnectionService: &auth0connection.Service{api},
	}
}

func metaClientService(meta interface{}) *auth0client.Service {
	return meta.(*Meta).ClientService
}

func metaRuleService(meta interface{}) *auth0rule.Service {
	return meta.(*Meta).RuleService
}

func metaConnectionService(meta interface{}) *auth0connection.Service {
	return meta.(*Meta).ConnectionService
}
