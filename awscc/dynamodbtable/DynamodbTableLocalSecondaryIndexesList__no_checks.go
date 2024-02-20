//go:build no_runtime_type_checking

package dynamodbtable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbTableLocalSecondaryIndexesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

