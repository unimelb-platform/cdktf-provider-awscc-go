//go:build no_runtime_type_checking

package ssmdocument

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmDocumentAttachmentsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsmDocumentAttachmentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentAttachmentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentAttachmentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentAttachmentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentAttachmentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsmDocumentAttachmentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

