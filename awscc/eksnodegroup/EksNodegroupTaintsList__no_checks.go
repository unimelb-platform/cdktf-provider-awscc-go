//go:build no_runtime_type_checking

package eksnodegroup

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksNodegroupTaintsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksNodegroupTaintsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksNodegroupTaintsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksNodegroupTaintsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksNodegroupTaintsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksNodegroupTaintsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksNodegroupTaintsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

