//go:build no_runtime_type_checking

package connectuser

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectUserTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConnectUserTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConnectUserTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConnectUserTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConnectUserTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConnectUserTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConnectUserTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

