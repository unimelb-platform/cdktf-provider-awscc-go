//go:build no_runtime_type_checking

package connectinstance

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectInstanceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConnectInstanceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConnectInstanceTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConnectInstanceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConnectInstanceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConnectInstanceTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConnectInstanceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConnectInstanceTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

