//go:build !no_runtime_type_checking

package applicationautoscalingscalingpolicy

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecification:
		val := val.(*ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecification)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecification:
		val_ := val.(ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecification)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecification; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateSetPredefinedMetricTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateSetResourceLabelParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecificationOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

