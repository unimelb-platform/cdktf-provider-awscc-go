//go:build no_runtime_type_checking

package ecscluster

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsClusterTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsClusterTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsClusterTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsClusterTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsClusterTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsClusterTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsClusterTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

