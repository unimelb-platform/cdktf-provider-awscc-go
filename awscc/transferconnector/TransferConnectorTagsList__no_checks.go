//go:build no_runtime_type_checking

package transferconnector

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransferConnectorTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TransferConnectorTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TransferConnectorTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TransferConnectorTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TransferConnectorTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TransferConnectorTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTransferConnectorTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

