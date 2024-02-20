//go:build no_runtime_type_checking

package sagemakerdomain

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SagemakerDomainTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SagemakerDomainTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SagemakerDomainTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SagemakerDomainTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SagemakerDomainTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SagemakerDomainTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSagemakerDomainTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

