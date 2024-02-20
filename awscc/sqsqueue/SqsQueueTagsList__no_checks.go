//go:build no_runtime_type_checking

package sqsqueue

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsQueueTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqsQueueTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqsQueueTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SqsQueueTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqsQueueTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqsQueueTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqsQueueTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

