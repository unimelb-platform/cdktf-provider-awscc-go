//go:build no_runtime_type_checking

package evsenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvsEnvironmentHostsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EvsEnvironmentHostsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvsEnvironmentHostsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentHostsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentHostsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentHostsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentHostsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvsEnvironmentHostsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

