//go:build no_runtime_type_checking

package ssmpatchbaseline

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmPatchBaselineTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsmPatchBaselineTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsmPatchBaselineTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

