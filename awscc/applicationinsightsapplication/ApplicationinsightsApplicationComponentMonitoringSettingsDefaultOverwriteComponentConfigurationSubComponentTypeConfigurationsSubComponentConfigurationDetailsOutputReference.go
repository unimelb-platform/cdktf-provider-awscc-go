package applicationinsightsapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/applicationinsightsapplication/internal"
)

type ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference interface {
	cdktf.ComplexObject
	AlarmMetrics() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsAlarmMetricsList
	AlarmMetricsInput() interface{}
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
	InternalValue() *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetails
	SetInternalValue(val *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetails)
	Logs() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsLogsList
	LogsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsEvents() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsWindowsEventsList
	WindowsEventsInput() interface{}
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
	PutAlarmMetrics(value interface{})
	PutLogs(value interface{})
	PutWindowsEvents(value interface{})
	ResetAlarmMetrics()
	ResetLogs()
	ResetWindowsEvents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference
type jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) AlarmMetrics() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsAlarmMetricsList {
	var returns ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsAlarmMetricsList
	_jsii_.Get(
		j,
		"alarmMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) AlarmMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) InternalValue() *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetails {
	var returns *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) Logs() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsLogsList {
	var returns ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsLogsList
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) LogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) WindowsEvents() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsWindowsEventsList {
	var returns ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsWindowsEventsList
	_jsii_.Get(
		j,
		"windowsEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) WindowsEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"windowsEventsInput",
		&returns,
	)
	return returns
}


func NewApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference{}

	_jsii_.Create(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference_Override(a ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference)SetInternalValue(val *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) PutAlarmMetrics(value interface{}) {
	if err := a.validatePutAlarmMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlarmMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) PutLogs(value interface{}) {
	if err := a.validatePutLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) PutWindowsEvents(value interface{}) {
	if err := a.validatePutWindowsEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWindowsEvents",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) ResetAlarmMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAlarmMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) ResetLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) ResetWindowsEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetWindowsEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

