package utils

import (
	"errors"

	"github.com/nepackage/nepackage/models"
)

func GetPoliciesWithMatchedPath(rolePolicies []models.Policy, urlPathRequested string) (matched []models.Policy, err error) {
	var policiesMatched []models.Policy
	for _, policy := range rolePolicies {
		if MatchValidator(policy.Path, urlPathRequested) {
			policiesMatched = append(policiesMatched, policy)
		}
	}
	if policiesMatched == nil {
		err = errors.New("you're not able to access " + urlPathRequested + " resource")
		return nil, err
	}
	return policiesMatched, nil
}

func ValidateMethodRequestWithPolicyMethod(policies []models.Policy, methodRequested string) (isAuthorized string, err error) {
	var authorized string
	for _, policy := range policies {
		methodsAllowed := ConvertStringToStruct(policy.AuthorizedMethods)
		for _, method := range methodsAllowed {
			if MatchValidator(method, methodRequested) {
				authorized = "true"
			}
		}
	}
	if authorized == "" {
		err = errors.New("no methods matched with method requested")
		return "", err
	}

	return authorized, nil
}

// func GetAllowedMethodsInPolicy(policy models.Policy) (methodsAllowed []models.AuthorizedMethods) {
// 	methodsAllowed = ConvertStringToStruct(policy.AuthorizedMethods)
// 	return methodsAllowed
// }
