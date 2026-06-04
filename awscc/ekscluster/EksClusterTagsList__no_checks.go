//go:build no_runtime_type_checking

package ekscluster

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksClusterTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EksClusterTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksClusterTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksClusterTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksClusterTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksClusterTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksClusterTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksClusterTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

