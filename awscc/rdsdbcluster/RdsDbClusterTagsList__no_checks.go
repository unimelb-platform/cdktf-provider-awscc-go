//go:build no_runtime_type_checking

package rdsdbcluster

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsDbClusterTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RdsDbClusterTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsDbClusterTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsDbClusterTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RdsDbClusterTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsDbClusterTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsDbClusterTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsDbClusterTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

