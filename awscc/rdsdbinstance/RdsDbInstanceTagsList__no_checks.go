//go:build no_runtime_type_checking

package rdsdbinstance

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsDbInstanceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsDbInstanceTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsDbInstanceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RdsDbInstanceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsDbInstanceTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsDbInstanceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsDbInstanceTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

