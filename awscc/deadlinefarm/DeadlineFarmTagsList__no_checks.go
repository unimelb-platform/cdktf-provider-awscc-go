//go:build no_runtime_type_checking

package deadlinefarm

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeadlineFarmTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DeadlineFarmTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeadlineFarmTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeadlineFarmTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeadlineFarmTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeadlineFarmTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeadlineFarmTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeadlineFarmTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

