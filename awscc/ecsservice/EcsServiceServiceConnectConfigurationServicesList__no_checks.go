//go:build no_runtime_type_checking

package ecsservice

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsServiceServiceConnectConfigurationServicesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

