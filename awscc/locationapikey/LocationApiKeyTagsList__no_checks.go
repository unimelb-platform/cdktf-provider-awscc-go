//go:build no_runtime_type_checking

package locationapikey

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LocationApiKeyTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LocationApiKeyTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LocationApiKeyTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LocationApiKeyTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LocationApiKeyTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LocationApiKeyTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLocationApiKeyTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

