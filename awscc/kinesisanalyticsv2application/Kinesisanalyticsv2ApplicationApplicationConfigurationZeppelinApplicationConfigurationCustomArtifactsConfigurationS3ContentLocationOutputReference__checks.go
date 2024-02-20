//go:build !no_runtime_type_checking

package kinesisanalyticsv2application

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateSetBucketArnParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateSetFileKeyParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocation:
		val := val.(*Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocation)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocation:
		val_ := val.(Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocation)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocation; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateSetObjectVersionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewKinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocationOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

