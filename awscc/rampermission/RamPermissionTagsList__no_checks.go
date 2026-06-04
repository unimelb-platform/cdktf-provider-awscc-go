//go:build no_runtime_type_checking

package rampermission

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RamPermissionTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RamPermissionTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RamPermissionTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RamPermissionTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RamPermissionTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RamPermissionTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RamPermissionTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRamPermissionTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

