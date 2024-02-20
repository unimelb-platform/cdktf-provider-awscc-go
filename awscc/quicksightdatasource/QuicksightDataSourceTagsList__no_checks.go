//go:build no_runtime_type_checking

package quicksightdatasource

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QuicksightDataSourceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QuicksightDataSourceTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QuicksightDataSourceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QuicksightDataSourceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QuicksightDataSourceTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QuicksightDataSourceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQuicksightDataSourceTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

