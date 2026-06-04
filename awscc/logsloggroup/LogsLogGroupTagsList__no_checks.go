//go:build no_runtime_type_checking

package logsloggroup

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogsLogGroupTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LogsLogGroupTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogsLogGroupTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogsLogGroupTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogsLogGroupTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogsLogGroupTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogsLogGroupTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogsLogGroupTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

