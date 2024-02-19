//go:build no_runtime_type_checking

package amplifyapp

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmplifyAppTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AmplifyAppTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AmplifyAppTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAmplifyAppTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

