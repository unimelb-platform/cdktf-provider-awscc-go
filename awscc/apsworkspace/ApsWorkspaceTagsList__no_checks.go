//go:build no_runtime_type_checking

package apsworkspace

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApsWorkspaceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApsWorkspaceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApsWorkspaceTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApsWorkspaceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApsWorkspaceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApsWorkspaceTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApsWorkspaceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApsWorkspaceTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

