//go:build no_runtime_type_checking

package rdseventsubscription

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsEventSubscriptionTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsEventSubscriptionTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsEventSubscriptionTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RdsEventSubscriptionTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsEventSubscriptionTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsEventSubscriptionTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsEventSubscriptionTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

