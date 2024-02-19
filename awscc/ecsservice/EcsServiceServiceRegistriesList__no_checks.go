//go:build no_runtime_type_checking

package ecsservice

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsServiceServiceRegistriesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EcsServiceServiceRegistriesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsServiceServiceRegistriesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsServiceServiceRegistriesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsServiceServiceRegistriesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsServiceServiceRegistriesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsServiceServiceRegistriesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsServiceServiceRegistriesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

