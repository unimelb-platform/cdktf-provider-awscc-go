//go:build no_runtime_type_checking

package configstoredquery

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigStoredQueryTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConfigStoredQueryTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfigStoredQueryTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfigStoredQueryTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigStoredQueryTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigStoredQueryTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfigStoredQueryTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfigStoredQueryTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

