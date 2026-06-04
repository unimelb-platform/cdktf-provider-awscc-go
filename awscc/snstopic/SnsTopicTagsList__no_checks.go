//go:build no_runtime_type_checking

package snstopic

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsTopicTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SnsTopicTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SnsTopicTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SnsTopicTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnsTopicTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SnsTopicTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SnsTopicTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSnsTopicTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

