//go:build no_runtime_type_checking

package dynamodbtable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbTableTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbTableTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

