//go:build no_runtime_type_checking

package eksaccessentry

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksAccessEntryAccessPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksAccessEntryAccessPoliciesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksAccessEntryAccessPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksAccessEntryAccessPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksAccessEntryAccessPoliciesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksAccessEntryAccessPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksAccessEntryAccessPoliciesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

