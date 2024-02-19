//go:build no_runtime_type_checking

package kendradatasource

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KendraDataSourceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KendraDataSourceTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KendraDataSourceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KendraDataSourceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KendraDataSourceTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KendraDataSourceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKendraDataSourceTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

