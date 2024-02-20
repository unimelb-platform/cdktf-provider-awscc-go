//go:build no_runtime_type_checking

package kinesisstream

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisStreamTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KinesisStreamTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KinesisStreamTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KinesisStreamTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KinesisStreamTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KinesisStreamTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKinesisStreamTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

