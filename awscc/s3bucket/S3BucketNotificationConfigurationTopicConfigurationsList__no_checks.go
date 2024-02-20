//go:build no_runtime_type_checking

package s3bucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketNotificationConfigurationTopicConfigurationsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketNotificationConfigurationTopicConfigurationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketNotificationConfigurationTopicConfigurationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketNotificationConfigurationTopicConfigurationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketNotificationConfigurationTopicConfigurationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketNotificationConfigurationTopicConfigurationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketNotificationConfigurationTopicConfigurationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

