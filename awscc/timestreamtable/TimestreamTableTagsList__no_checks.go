//go:build no_runtime_type_checking

package timestreamtable

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TimestreamTableTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TimestreamTableTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TimestreamTableTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TimestreamTableTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TimestreamTableTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TimestreamTableTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTimestreamTableTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

