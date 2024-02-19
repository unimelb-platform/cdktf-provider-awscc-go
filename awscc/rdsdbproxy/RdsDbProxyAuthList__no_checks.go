//go:build no_runtime_type_checking

package rdsdbproxy

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsDbProxyAuthList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RdsDbProxyAuthList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsDbProxyAuthList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsDbProxyAuthList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RdsDbProxyAuthList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsDbProxyAuthList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsDbProxyAuthList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsDbProxyAuthListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

