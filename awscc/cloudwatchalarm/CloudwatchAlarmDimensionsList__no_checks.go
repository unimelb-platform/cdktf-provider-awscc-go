//go:build no_runtime_type_checking

package cloudwatchalarm

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudwatchAlarmDimensionsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudwatchAlarmDimensionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlarmDimensionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlarmDimensionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlarmDimensionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlarmDimensionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudwatchAlarmDimensionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

