//go:build no_runtime_type_checking

package iotcommand

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotCommandTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IotCommandTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotCommandTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotCommandTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotCommandTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotCommandTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotCommandTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotCommandTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

