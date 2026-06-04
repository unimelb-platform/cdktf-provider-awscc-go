//go:build no_runtime_type_checking

package cloudwatchalarm

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudwatchAlarmTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudwatchAlarmTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudwatchAlarmTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlarmTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlarmTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlarmTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlarmTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudwatchAlarmTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

