//go:build !no_runtime_type_checking

package acmpcacertificate

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validatePutQualifierParameters(value *AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersQualifier) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiers:
		val := val.(*AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiers)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiers:
		val_ := val.(AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiers)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiers; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateSetPolicyQualifierIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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
