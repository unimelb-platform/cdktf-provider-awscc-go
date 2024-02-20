//go:build no_runtime_type_checking

package xraygroup

// Building without runtime type checking enabled, so all the below just return nil

func (x *jsiiProxy_XrayGroupTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (x *jsiiProxy_XrayGroupTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_XrayGroupTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_XrayGroupTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_XrayGroupTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_XrayGroupTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewXrayGroupTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

