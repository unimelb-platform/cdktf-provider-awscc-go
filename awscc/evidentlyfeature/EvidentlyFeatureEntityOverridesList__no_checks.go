//go:build no_runtime_type_checking

package evidentlyfeature

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvidentlyFeatureEntityOverridesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvidentlyFeatureEntityOverridesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureEntityOverridesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureEntityOverridesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureEntityOverridesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureEntityOverridesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvidentlyFeatureEntityOverridesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

