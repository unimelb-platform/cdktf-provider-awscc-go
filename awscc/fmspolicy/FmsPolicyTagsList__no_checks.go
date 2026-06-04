//go:build no_runtime_type_checking

package fmspolicy

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FmsPolicyTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FmsPolicyTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FmsPolicyTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FmsPolicyTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FmsPolicyTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FmsPolicyTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FmsPolicyTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFmsPolicyTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

