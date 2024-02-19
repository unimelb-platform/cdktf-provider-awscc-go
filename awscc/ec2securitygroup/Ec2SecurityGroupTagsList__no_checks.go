//go:build no_runtime_type_checking

package ec2securitygroup

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2SecurityGroupTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_Ec2SecurityGroupTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2SecurityGroupTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2SecurityGroupTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2SecurityGroupTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2SecurityGroupTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2SecurityGroupTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2SecurityGroupTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

