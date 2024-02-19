//go:build !no_runtime_type_checking

package mediapackagev2originendpoint

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validatePutFilterConfigurationParameters(value *Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfiguration) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validatePutScteHlsParameters(value *Mediapackagev2OriginEndpointLowLatencyHlsManifestsScteHls) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateSetChildManifestNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *Mediapackagev2OriginEndpointLowLatencyHlsManifests:
		val := val.(*Mediapackagev2OriginEndpointLowLatencyHlsManifests)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case Mediapackagev2OriginEndpointLowLatencyHlsManifests:
		val_ := val.(Mediapackagev2OriginEndpointLowLatencyHlsManifests)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *Mediapackagev2OriginEndpointLowLatencyHlsManifests; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateSetManifestNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateSetManifestWindowSecondsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateSetProgramDateTimeIntervalSecondsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) validateSetUrlParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewMediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

