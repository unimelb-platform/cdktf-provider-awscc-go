//go:build !no_runtime_type_checking

package lexbot

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validatePutButtonsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardButtons:
		value := value.(*[]*LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardButtons)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardButtons:
		value_ := value.([]*LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardButtons)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardButtons; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateSetImageUrlParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCard:
		val := val.(*LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCard)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCard:
		val_ := val.(LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCard)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCard; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateSetSubtitleParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReference) validateSetTitleParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCardOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

