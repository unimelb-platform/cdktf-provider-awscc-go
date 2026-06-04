//go:build no_runtime_type_checking

package pcscluster

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PcsClusterEndpointsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PcsClusterEndpointsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PcsClusterEndpointsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PcsClusterEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PcsClusterEndpointsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PcsClusterEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPcsClusterEndpointsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

