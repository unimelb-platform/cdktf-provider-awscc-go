//go:build no_runtime_type_checking

package dynamodbtable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbTableAttributeDefinitionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableAttributeDefinitionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableAttributeDefinitionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableAttributeDefinitionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableAttributeDefinitionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableAttributeDefinitionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableAttributeDefinitionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbTableAttributeDefinitionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

