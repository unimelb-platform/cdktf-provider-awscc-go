//go:build no_runtime_type_checking

package snstopic

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsTopicSubscriptionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SnsTopicSubscriptionList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SnsTopicSubscriptionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SnsTopicSubscriptionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnsTopicSubscriptionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SnsTopicSubscriptionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SnsTopicSubscriptionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSnsTopicSubscriptionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

