//go:build no_runtime_type_checking

package ec2ipampool

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2IpamPoolTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_Ec2IpamPoolTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2IpamPoolTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2IpamPoolTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2IpamPoolTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2IpamPoolTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2IpamPoolTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2IpamPoolTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

