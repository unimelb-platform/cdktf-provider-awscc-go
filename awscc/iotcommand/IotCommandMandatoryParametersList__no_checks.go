//go:build no_runtime_type_checking

package iotcommand

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotCommandMandatoryParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IotCommandMandatoryParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotCommandMandatoryParametersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotCommandMandatoryParametersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotCommandMandatoryParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotCommandMandatoryParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotCommandMandatoryParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotCommandMandatoryParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

