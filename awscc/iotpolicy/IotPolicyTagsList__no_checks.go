//go:build no_runtime_type_checking

package iotpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotPolicyTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IotPolicyTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotPolicyTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotPolicyTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotPolicyTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotPolicyTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotPolicyTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotPolicyTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

