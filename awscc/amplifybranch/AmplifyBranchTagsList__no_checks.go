//go:build no_runtime_type_checking

package amplifybranch

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmplifyBranchTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AmplifyBranchTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AmplifyBranchTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AmplifyBranchTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AmplifyBranchTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AmplifyBranchTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAmplifyBranchTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

