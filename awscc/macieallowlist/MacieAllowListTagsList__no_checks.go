//go:build no_runtime_type_checking

package macieallowlist

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MacieAllowListTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MacieAllowListTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MacieAllowListTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MacieAllowListTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MacieAllowListTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MacieAllowListTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMacieAllowListTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

