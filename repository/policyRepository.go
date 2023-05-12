package repository

import (
	"github.com/nepackage/nepackage/database"
	"github.com/nepackage/nepackage/models"
)

func GetPolicies() ([]models.Policy, error) {
	var policies []models.Policy
	if err := database.DB.Find(&policies).Error; err != nil {
		return nil, err
	}
	return policies, nil
}

func GetPolicyById(policyId uint) (*models.Policy, error) {
	var policy *models.Policy
	if err := database.DB.Where("id = ?", policyId).First(&policy).Error; err != nil {
		return nil, err
	}

	return policy, nil
}

func CreatePolicy(policy *models.Policy) (bool, error) {
	if err := database.DB.Create(&policy).Error; err != nil {
		return false, err
	}

	return true, nil
}

func UpdatePolicy(policyId uint, policyInput models.PolicyUpdate) (*models.PolicyUpdate, error) {
	var policy models.Policy

	if err := database.DB.Where("id = ?", policyId).Model(&policy).Updates(policyInput).Error; err != nil {
		return nil, err
	}

	return &policyInput, nil
}

func DeletePolicy(policyId uint) (bool, error) {
	var policy models.Policy
	if err := database.DB.Where("id = ?", policyId).First(&policy).Error; err != nil {
		return false, err
	}
	if err := database.DB.Delete(&policy).Error; err != nil {
		return false, err
	}

	return true, nil
}

func GetPoliciesByIDs(policiesIDs []int) ([]models.Policy, error) {
	var policies []models.Policy
	if err := database.DB.Where("id IN ?", policiesIDs).Find(&policies).Error; err != nil {
		return nil, err
	}
	return policies, nil
}

func GetPoliciesWithMatchedPath(policiesID []int, path string) ([]models.Policy, error) {
	var policies []models.Policy
	if err := database.DB.Where("id IN ? AND path = ?", policies, path).Find(&policies).Error; err != nil {
		return nil, err
	}

	return policies, nil
}

func GetRolePolicies(policiesIDs []int) ([]models.Policy, error) {
	var rolePolicies []models.Policy
	rolePolicies, err := GetPoliciesByIDs(policiesIDs)
	if err != nil {
		return nil, err
	}
	return rolePolicies, nil
}
