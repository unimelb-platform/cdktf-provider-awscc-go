//go:build no_runtime_type_checking

package syntheticscanary

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsCanaryTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SyntheticsCanaryTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsCanaryTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsCanaryTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SyntheticsCanaryTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsCanaryTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsCanaryTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsCanaryTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

