//go:build !no_runtime_type_checking

package lexbot

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validatePutCustomPayloadParameters(value *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsCustomPayload) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validatePutImageResponseCardParameters(value *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsImageResponseCard) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validatePutPlainTextMessageParameters(value *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsPlainTextMessage) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validatePutSsmlMessageParameters(value *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsSsmlMessage) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariations:
		val := val.(*LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariations)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariations:
		val_ := val.(LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariations)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariations; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewLexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListVariationsOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

