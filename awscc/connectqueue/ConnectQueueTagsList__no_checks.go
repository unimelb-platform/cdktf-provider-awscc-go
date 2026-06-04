//go:build no_runtime_type_checking

package connectqueue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectQueueTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConnectQueueTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConnectQueueTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConnectQueueTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConnectQueueTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConnectQueueTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConnectQueueTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConnectQueueTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

