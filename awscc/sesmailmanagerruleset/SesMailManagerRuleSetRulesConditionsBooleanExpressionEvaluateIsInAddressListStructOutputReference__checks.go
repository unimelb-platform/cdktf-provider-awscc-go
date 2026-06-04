//go:build !no_runtime_type_checking

package sesmailmanagerruleset

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (s *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateSetAddressListsParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateSetAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStruct:
		val := val.(*SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStruct)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStruct:
		val_ := val.(SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStruct)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStruct; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewSesMailManagerRuleSetRulesConditionsBooleanExpressionEvaluateIsInAddressListStructOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

