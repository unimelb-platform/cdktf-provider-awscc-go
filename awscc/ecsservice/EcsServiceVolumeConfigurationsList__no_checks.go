//go:build no_runtime_type_checking

package ecsservice

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsServiceVolumeConfigurationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsServiceVolumeConfigurationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

