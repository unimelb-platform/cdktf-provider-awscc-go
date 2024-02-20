//go:build no_runtime_type_checking

package wafv2ipset

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_Wafv2IpSetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_Wafv2IpSetTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Wafv2IpSetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Wafv2IpSetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Wafv2IpSetTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Wafv2IpSetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWafv2IpSetTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

