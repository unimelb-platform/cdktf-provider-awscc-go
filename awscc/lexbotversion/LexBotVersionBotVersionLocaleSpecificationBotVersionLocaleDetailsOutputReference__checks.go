//go:build !no_runtime_type_checking

package lexbotversion

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (l *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetails:
		val := val.(*LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetails)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetails:
		val_ := val.(LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetails)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetails; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateSetSourceBotVersionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewLexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetailsOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

