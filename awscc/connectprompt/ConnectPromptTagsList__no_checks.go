//go:build no_runtime_type_checking

package connectprompt

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectPromptTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConnectPromptTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConnectPromptTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConnectPromptTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConnectPromptTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConnectPromptTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConnectPromptTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

