//go:build no_runtime_type_checking

package fmspolicy

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FmsPolicyResourceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FmsPolicyResourceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FmsPolicyResourceTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FmsPolicyResourceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FmsPolicyResourceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FmsPolicyResourceTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FmsPolicyResourceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFmsPolicyResourceTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

