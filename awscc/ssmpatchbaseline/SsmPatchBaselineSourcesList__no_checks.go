//go:build no_runtime_type_checking

package ssmpatchbaseline

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmPatchBaselineSourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsmPatchBaselineSourcesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineSourcesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineSourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineSourcesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineSourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsmPatchBaselineSourcesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

