//go:build no_runtime_type_checking

package iotdimension

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotDimensionTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotDimensionTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotDimensionTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotDimensionTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotDimensionTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotDimensionTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotDimensionTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

