//go:build no_runtime_type_checking

package iamuser

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IamUserPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IamUserPoliciesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IamUserPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IamUserPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IamUserPoliciesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IamUserPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIamUserPoliciesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

