//go:build no_runtime_type_checking

package amplifyapp

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAmplifyAppEnvironmentVariablesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

