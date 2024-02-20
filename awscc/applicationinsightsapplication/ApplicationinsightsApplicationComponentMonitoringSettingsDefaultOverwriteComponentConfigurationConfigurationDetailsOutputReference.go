package applicationinsightsapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/applicationinsightsapplication/internal"
)

type ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference interface {
	cdktf.ComplexObject
	AlarmMetrics() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarmMetricsList
	AlarmMetricsInput() interface{}
	Alarms() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarmsList
	AlarmsInput() interface{}
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
	HaClusterPrometheusExporter() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHaClusterPrometheusExporterOutputReference
	HaClusterPrometheusExporterInput() interface{}
	HanaPrometheusExporter() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHanaPrometheusExporterOutputReference
	HanaPrometheusExporterInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JmxPrometheusExporter() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference
	JmxPrometheusExporterInput() interface{}
	Logs() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsLogsList
	LogsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsEvents() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsWindowsEventsList
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
	PutAlarms(value interface{})
	PutHaClusterPrometheusExporter(value *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHaClusterPrometheusExporter)
	PutHanaPrometheusExporter(value *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHanaPrometheusExporter)
	PutJmxPrometheusExporter(value *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsJmxPrometheusExporter)
	PutLogs(value interface{})
	PutWindowsEvents(value interface{})
	ResetAlarmMetrics()
	ResetAlarms()
	ResetHaClusterPrometheusExporter()
	ResetHanaPrometheusExporter()
	ResetJmxPrometheusExporter()
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

// The jsii proxy struct for ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference
type jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) AlarmMetrics() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarmMetricsList {
	var returns ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarmMetricsList
	_jsii_.Get(
		j,
		"alarmMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) AlarmMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) Alarms() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarmsList {
	var returns ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarmsList
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) AlarmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) HaClusterPrometheusExporter() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHaClusterPrometheusExporterOutputReference {
	var returns ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHaClusterPrometheusExporterOutputReference
	_jsii_.Get(
		j,
		"haClusterPrometheusExporter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) HaClusterPrometheusExporterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haClusterPrometheusExporterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) HanaPrometheusExporter() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHanaPrometheusExporterOutputReference {
	var returns ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHanaPrometheusExporterOutputReference
	_jsii_.Get(
		j,
		"hanaPrometheusExporter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) HanaPrometheusExporterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hanaPrometheusExporterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) JmxPrometheusExporter() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference {
	var returns ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference
	_jsii_.Get(
		j,
		"jmxPrometheusExporter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) JmxPrometheusExporterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jmxPrometheusExporterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) Logs() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsLogsList {
	var returns ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsLogsList
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) LogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) WindowsEvents() ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsWindowsEventsList {
	var returns ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsWindowsEventsList
	_jsii_.Get(
		j,
		"windowsEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) WindowsEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"windowsEventsInput",
		&returns,
	)
	return returns
}


func NewApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference{}

	_jsii_.Create(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference_Override(a ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) PutAlarmMetrics(value interface{}) {
	if err := a.validatePutAlarmMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlarmMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) PutAlarms(value interface{}) {
	if err := a.validatePutAlarmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlarms",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) PutHaClusterPrometheusExporter(value *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHaClusterPrometheusExporter) {
	if err := a.validatePutHaClusterPrometheusExporterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHaClusterPrometheusExporter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) PutHanaPrometheusExporter(value *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHanaPrometheusExporter) {
	if err := a.validatePutHanaPrometheusExporterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHanaPrometheusExporter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) PutJmxPrometheusExporter(value *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsJmxPrometheusExporter) {
	if err := a.validatePutJmxPrometheusExporterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJmxPrometheusExporter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) PutLogs(value interface{}) {
	if err := a.validatePutLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) PutWindowsEvents(value interface{}) {
	if err := a.validatePutWindowsEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWindowsEvents",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ResetAlarmMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAlarmMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ResetAlarms() {
	_jsii_.InvokeVoid(
		a,
		"resetAlarms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ResetHaClusterPrometheusExporter() {
	_jsii_.InvokeVoid(
		a,
		"resetHaClusterPrometheusExporter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ResetHanaPrometheusExporter() {
	_jsii_.InvokeVoid(
		a,
		"resetHanaPrometheusExporter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ResetJmxPrometheusExporter() {
	_jsii_.InvokeVoid(
		a,
		"resetJmxPrometheusExporter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ResetLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ResetWindowsEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetWindowsEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

