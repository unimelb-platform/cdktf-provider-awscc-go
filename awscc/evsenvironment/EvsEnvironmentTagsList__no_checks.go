//go:build no_runtime_type_checking

package evsenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvsEnvironmentTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EvsEnvironmentTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvsEnvironmentTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvsEnvironmentTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvsEnvironmentTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

