//go:build no_runtime_type_checking

package s3storagelensgroup

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3StorageLensGroupTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3StorageLensGroupTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3StorageLensGroupTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3StorageLensGroupTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3StorageLensGroupTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3StorageLensGroupTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3StorageLensGroupTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

