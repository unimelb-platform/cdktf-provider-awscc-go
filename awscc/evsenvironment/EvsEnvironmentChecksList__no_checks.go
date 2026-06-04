//go:build no_runtime_type_checking

package evsenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvsEnvironmentChecksList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EvsEnvironmentChecksList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvsEnvironmentChecksList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentChecksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentChecksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentChecksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvsEnvironmentChecksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

