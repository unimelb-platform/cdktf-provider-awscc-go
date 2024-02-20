package kinesisanalyticsv2application

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/kinesisanalyticsv2application/internal"
)

type Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference interface {
	cdktf.ComplexObject
	CatalogConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCatalogConfigurationOutputReference
	CatalogConfigurationInput() interface{}
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
	CustomArtifactsConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationList
	CustomArtifactsConfigurationInput() interface{}
	DeployAsApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationDeployAsApplicationConfigurationOutputReference
	DeployAsApplicationConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MonitoringConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationMonitoringConfigurationOutputReference
	MonitoringConfigurationInput() interface{}
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
	PutCatalogConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCatalogConfiguration)
	PutCustomArtifactsConfiguration(value interface{})
	PutDeployAsApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationDeployAsApplicationConfiguration)
	PutMonitoringConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationMonitoringConfiguration)
	ResetCatalogConfiguration()
	ResetCustomArtifactsConfiguration()
	ResetDeployAsApplicationConfiguration()
	ResetMonitoringConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference
type jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) CatalogConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCatalogConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCatalogConfigurationOutputReference
	_jsii_.Get(
		j,
		"catalogConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) CatalogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) CustomArtifactsConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationList {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationList
	_jsii_.Get(
		j,
		"customArtifactsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) CustomArtifactsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customArtifactsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) DeployAsApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationDeployAsApplicationConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationDeployAsApplicationConfigurationOutputReference
	_jsii_.Get(
		j,
		"deployAsApplicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) DeployAsApplicationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deployAsApplicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) MonitoringConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationMonitoringConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationMonitoringConfigurationOutputReference
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) MonitoringConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.kinesisanalyticsv2Application.Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference_Override(k Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.kinesisanalyticsv2Application.Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) PutCatalogConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCatalogConfiguration) {
	if err := k.validatePutCatalogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCatalogConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) PutCustomArtifactsConfiguration(value interface{}) {
	if err := k.validatePutCustomArtifactsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCustomArtifactsConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) PutDeployAsApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationDeployAsApplicationConfiguration) {
	if err := k.validatePutDeployAsApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDeployAsApplicationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) PutMonitoringConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationMonitoringConfiguration) {
	if err := k.validatePutMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) ResetCatalogConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetCatalogConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) ResetCustomArtifactsConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetCustomArtifactsConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) ResetDeployAsApplicationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetDeployAsApplicationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) ResetMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetMonitoringConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

