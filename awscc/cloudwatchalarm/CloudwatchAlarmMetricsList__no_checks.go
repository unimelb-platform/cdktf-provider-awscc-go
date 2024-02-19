//go:build no_runtime_type_checking

package cloudwatchalarm

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudwatchAlarmMetricsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudwatchAlarmMetricsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlarmMetricsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlarmMetricsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlarmMetricsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlarmMetricsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudwatchAlarmMetricsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

