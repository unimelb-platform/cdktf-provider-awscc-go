//go:build no_runtime_type_checking

package s3bucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

