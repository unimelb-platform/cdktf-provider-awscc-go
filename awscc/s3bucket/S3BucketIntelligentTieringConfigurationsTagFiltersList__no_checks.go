//go:build no_runtime_type_checking

package s3bucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketIntelligentTieringConfigurationsTagFiltersList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketIntelligentTieringConfigurationsTagFiltersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketIntelligentTieringConfigurationsTagFiltersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketIntelligentTieringConfigurationsTagFiltersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketIntelligentTieringConfigurationsTagFiltersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketIntelligentTieringConfigurationsTagFiltersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketIntelligentTieringConfigurationsTagFiltersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
