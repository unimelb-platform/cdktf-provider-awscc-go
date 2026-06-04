//go:build no_runtime_type_checking

package locationmap

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LocationMapTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LocationMapTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LocationMapTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LocationMapTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LocationMapTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LocationMapTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LocationMapTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLocationMapTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

