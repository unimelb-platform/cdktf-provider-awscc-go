//go:build no_runtime_type_checking

package evsenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvsEnvironmentCredentialsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EvsEnvironmentCredentialsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvsEnvironmentCredentialsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentCredentialsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentCredentialsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentCredentialsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvsEnvironmentCredentialsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

