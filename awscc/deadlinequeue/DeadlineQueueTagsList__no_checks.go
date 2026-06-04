//go:build no_runtime_type_checking

package deadlinequeue

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeadlineQueueTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DeadlineQueueTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeadlineQueueTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeadlineQueueTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeadlineQueueTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeadlineQueueTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeadlineQueueTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeadlineQueueTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

