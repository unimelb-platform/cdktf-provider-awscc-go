//go:build no_runtime_type_checking

package appstreamapplication

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppstreamApplicationTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppstreamApplicationTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppstreamApplicationTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppstreamApplicationTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppstreamApplicationTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppstreamApplicationTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppstreamApplicationTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

