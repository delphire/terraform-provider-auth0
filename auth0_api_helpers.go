package main

import "github.com/bocodigitalmedia/go-auth0/auth0mgmt"

func isApi404Err(err error) bool {
	apiErr, ok := err.(*auth0mgmt.ApiError)
	return ok && apiErr.StatusCode == 404
}
