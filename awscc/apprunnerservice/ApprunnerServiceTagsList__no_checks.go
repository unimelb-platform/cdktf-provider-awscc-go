//go:build no_runtime_type_checking

package apprunnerservice

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApprunnerServiceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApprunnerServiceTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApprunnerServiceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApprunnerServiceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApprunnerServiceTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApprunnerServiceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApprunnerServiceTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

