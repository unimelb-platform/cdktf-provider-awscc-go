//go:build no_runtime_type_checking

package m2environment

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_M2EnvironmentStorageConfigurationsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_M2EnvironmentStorageConfigurationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_M2EnvironmentStorageConfigurationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_M2EnvironmentStorageConfigurationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_M2EnvironmentStorageConfigurationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_M2EnvironmentStorageConfigurationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewM2EnvironmentStorageConfigurationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

