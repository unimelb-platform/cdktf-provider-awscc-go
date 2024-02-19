package athenaworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/athenaworkgroup/internal"
)

type AthenaWorkGroupWorkGroupConfigurationOutputReference interface {
	cdktf.ComplexObject
	AdditionalConfiguration() *string
	SetAdditionalConfiguration(val *string)
	AdditionalConfigurationInput() *string
	BytesScannedCutoffPerQuery() *float64
	SetBytesScannedCutoffPerQuery(val *float64)
	BytesScannedCutoffPerQueryInput() *float64
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
	CustomerContentEncryptionConfiguration() AthenaWorkGroupWorkGroupConfigurationCustomerContentEncryptionConfigurationOutputReference
	CustomerContentEncryptionConfigurationInput() interface{}
	EnforceWorkGroupConfiguration() interface{}
	SetEnforceWorkGroupConfiguration(val interface{})
	EnforceWorkGroupConfigurationInput() interface{}
	EngineVersion() AthenaWorkGroupWorkGroupConfigurationEngineVersionOutputReference
	EngineVersionInput() interface{}
	ExecutionRole() *string
	SetExecutionRole(val *string)
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PublishCloudwatchMetricsEnabled() interface{}
	SetPublishCloudwatchMetricsEnabled(val interface{})
	PublishCloudwatchMetricsEnabledInput() interface{}
	RequesterPaysEnabled() interface{}
	SetRequesterPaysEnabled(val interface{})
	RequesterPaysEnabledInput() interface{}
	ResultConfiguration() AthenaWorkGroupWorkGroupConfigurationResultConfigurationOutputReference
	ResultConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutCustomerContentEncryptionConfiguration(value *AthenaWorkGroupWorkGroupConfigurationCustomerContentEncryptionConfiguration)
	PutEngineVersion(value *AthenaWorkGroupWorkGroupConfigurationEngineVersion)
	PutResultConfiguration(value *AthenaWorkGroupWorkGroupConfigurationResultConfiguration)
	ResetAdditionalConfiguration()
	ResetBytesScannedCutoffPerQuery()
	ResetCustomerContentEncryptionConfiguration()
	ResetEnforceWorkGroupConfiguration()
	ResetEngineVersion()
	ResetExecutionRole()
	ResetPublishCloudwatchMetricsEnabled()
	ResetRequesterPaysEnabled()
	ResetResultConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AthenaWorkGroupWorkGroupConfigurationOutputReference
type jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) AdditionalConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) AdditionalConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) BytesScannedCutoffPerQuery() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) BytesScannedCutoffPerQueryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) CustomerContentEncryptionConfiguration() AthenaWorkGroupWorkGroupConfigurationCustomerContentEncryptionConfigurationOutputReference {
	var returns AthenaWorkGroupWorkGroupConfigurationCustomerContentEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"customerContentEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) CustomerContentEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerContentEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) EnforceWorkGroupConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceWorkGroupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) EnforceWorkGroupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceWorkGroupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) EngineVersion() AthenaWorkGroupWorkGroupConfigurationEngineVersionOutputReference {
	var returns AthenaWorkGroupWorkGroupConfigurationEngineVersionOutputReference
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) EngineVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) PublishCloudwatchMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) PublishCloudwatchMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) RequesterPaysEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) RequesterPaysEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ResultConfiguration() AthenaWorkGroupWorkGroupConfigurationResultConfigurationOutputReference {
	var returns AthenaWorkGroupWorkGroupConfigurationResultConfigurationOutputReference
	_jsii_.Get(
		j,
		"resultConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ResultConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resultConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAthenaWorkGroupWorkGroupConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AthenaWorkGroupWorkGroupConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewAthenaWorkGroupWorkGroupConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.athenaWorkGroup.AthenaWorkGroupWorkGroupConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAthenaWorkGroupWorkGroupConfigurationOutputReference_Override(a AthenaWorkGroupWorkGroupConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.athenaWorkGroup.AthenaWorkGroupWorkGroupConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference)SetAdditionalConfiguration(val *string) {
	if err := j.validateSetAdditionalConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalConfiguration",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference)SetBytesScannedCutoffPerQuery(val *float64) {
	if err := j.validateSetBytesScannedCutoffPerQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bytesScannedCutoffPerQuery",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference)SetEnforceWorkGroupConfiguration(val interface{}) {
	if err := j.validateSetEnforceWorkGroupConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceWorkGroupConfiguration",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference)SetPublishCloudwatchMetricsEnabled(val interface{}) {
	if err := j.validateSetPublishCloudwatchMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishCloudwatchMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference)SetRequesterPaysEnabled(val interface{}) {
	if err := j.validateSetRequesterPaysEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterPaysEnabled",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) PutCustomerContentEncryptionConfiguration(value *AthenaWorkGroupWorkGroupConfigurationCustomerContentEncryptionConfiguration) {
	if err := a.validatePutCustomerContentEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomerContentEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) PutEngineVersion(value *AthenaWorkGroupWorkGroupConfigurationEngineVersion) {
	if err := a.validatePutEngineVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEngineVersion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) PutResultConfiguration(value *AthenaWorkGroupWorkGroupConfigurationResultConfiguration) {
	if err := a.validatePutResultConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResultConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ResetAdditionalConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ResetBytesScannedCutoffPerQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetBytesScannedCutoffPerQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ResetCustomerContentEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerContentEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ResetEnforceWorkGroupConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforceWorkGroupConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ResetExecutionRole() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ResetPublishCloudwatchMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPublishCloudwatchMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ResetRequesterPaysEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRequesterPaysEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ResetResultConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetResultConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

