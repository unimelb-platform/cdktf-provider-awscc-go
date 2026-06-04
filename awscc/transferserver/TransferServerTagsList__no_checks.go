//go:build no_runtime_type_checking

package transferserver

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransferServerTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TransferServerTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TransferServerTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TransferServerTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TransferServerTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TransferServerTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TransferServerTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTransferServerTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

