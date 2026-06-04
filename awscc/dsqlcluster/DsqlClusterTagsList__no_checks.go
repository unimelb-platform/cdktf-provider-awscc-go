//go:build no_runtime_type_checking

package dsqlcluster

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DsqlClusterTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DsqlClusterTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DsqlClusterTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DsqlClusterTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DsqlClusterTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DsqlClusterTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DsqlClusterTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDsqlClusterTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

