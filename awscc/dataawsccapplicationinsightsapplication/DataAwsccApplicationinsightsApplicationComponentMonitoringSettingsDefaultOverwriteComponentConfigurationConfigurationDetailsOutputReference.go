package dataawsccapplicationinsightsapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccapplicationinsightsapplication/internal"
)

type DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference interface {
	cdktf.ComplexObject
	AlarmMetrics() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarmMetricsList
	Alarms() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarmsList
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HaClusterPrometheusExporter() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHaClusterPrometheusExporterOutputReference
	HanaPrometheusExporter() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHanaPrometheusExporterOutputReference
	InternalValue() *DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetails
	SetInternalValue(val *DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetails)
	JmxPrometheusExporter() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference
	Logs() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsLogsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsEvents() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsWindowsEventsList
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference
type jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) AlarmMetrics() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarmMetricsList {
	var returns DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarmMetricsList
	_jsii_.Get(
		j,
		"alarmMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) Alarms() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarmsList {
	var returns DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarmsList
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) HaClusterPrometheusExporter() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHaClusterPrometheusExporterOutputReference {
	var returns DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHaClusterPrometheusExporterOutputReference
	_jsii_.Get(
		j,
		"haClusterPrometheusExporter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) HanaPrometheusExporter() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHanaPrometheusExporterOutputReference {
	var returns DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHanaPrometheusExporterOutputReference
	_jsii_.Get(
		j,
		"hanaPrometheusExporter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) InternalValue() *DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetails {
	var returns *DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) JmxPrometheusExporter() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference {
	var returns DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference
	_jsii_.Get(
		j,
		"jmxPrometheusExporter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) Logs() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsLogsList {
	var returns DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsLogsList
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) WindowsEvents() DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsWindowsEventsList {
	var returns DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsWindowsEventsList
	_jsii_.Get(
		j,
		"windowsEvents",
		&returns,
	)
	return returns
}


func NewDataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccApplicationinsightsApplication.DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference_Override(d DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccApplicationinsightsApplication.DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference)SetInternalValue(val *DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

