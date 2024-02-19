//go:build no_runtime_type_checking

package s3accessgrant

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3AccessGrantTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessGrantTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3AccessGrantTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3AccessGrantTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3AccessGrantTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3AccessGrantTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3AccessGrantTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3AccessGrantTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

