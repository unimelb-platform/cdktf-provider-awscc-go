//go:build no_runtime_type_checking

package ivschannel

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IvsChannelTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IvsChannelTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IvsChannelTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IvsChannelTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IvsChannelTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IvsChannelTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIvsChannelTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

