package sagemakercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakercluster/internal"
)

type SagemakerClusterInstanceGroupsOutputReference interface {
	cdktf.ComplexObject
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
	CurrentCount() *float64
	SetCurrentCount(val *float64)
	CurrentCountInput() *float64
	ExecutionRole() *string
	SetExecutionRole(val *string)
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceGroupName() *string
	SetInstanceGroupName(val *string)
	InstanceGroupNameInput() *string
	InstanceStorageConfigs() SagemakerClusterInstanceGroupsInstanceStorageConfigsList
	InstanceStorageConfigsInput() interface{}
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LifeCycleConfig() SagemakerClusterInstanceGroupsLifeCycleConfigOutputReference
	LifeCycleConfigInput() interface{}
	OnStartDeepHealthChecks() *[]*string
	SetOnStartDeepHealthChecks(val *[]*string)
	OnStartDeepHealthChecksInput() *[]*string
	OverrideVpcConfig() SagemakerClusterInstanceGroupsOverrideVpcConfigOutputReference
	OverrideVpcConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThreadsPerCore() *float64
	SetThreadsPerCore(val *float64)
	ThreadsPerCoreInput() *float64
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
	PutInstanceStorageConfigs(value interface{})
	PutLifeCycleConfig(value *SagemakerClusterInstanceGroupsLifeCycleConfig)
	PutOverrideVpcConfig(value *SagemakerClusterInstanceGroupsOverrideVpcConfig)
	ResetCurrentCount()
	ResetInstanceStorageConfigs()
	ResetOnStartDeepHealthChecks()
	ResetOverrideVpcConfig()
	ResetThreadsPerCore()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerClusterInstanceGroupsOutputReference
type jsiiProxy_SagemakerClusterInstanceGroupsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) CurrentCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) CurrentCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceStorageConfigs() SagemakerClusterInstanceGroupsInstanceStorageConfigsList {
	var returns SagemakerClusterInstanceGroupsInstanceStorageConfigsList
	_jsii_.Get(
		j,
		"instanceStorageConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceStorageConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceStorageConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) LifeCycleConfig() SagemakerClusterInstanceGroupsLifeCycleConfigOutputReference {
	var returns SagemakerClusterInstanceGroupsLifeCycleConfigOutputReference
	_jsii_.Get(
		j,
		"lifeCycleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) LifeCycleConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifeCycleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) OnStartDeepHealthChecks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onStartDeepHealthChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) OnStartDeepHealthChecksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onStartDeepHealthChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) OverrideVpcConfig() SagemakerClusterInstanceGroupsOverrideVpcConfigOutputReference {
	var returns SagemakerClusterInstanceGroupsOverrideVpcConfigOutputReference
	_jsii_.Get(
		j,
		"overrideVpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) OverrideVpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideVpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ThreadsPerCore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ThreadsPerCoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCoreInput",
		&returns,
	)
	return returns
}


func NewSagemakerClusterInstanceGroupsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerClusterInstanceGroupsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerClusterInstanceGroupsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerClusterInstanceGroupsOutputReference{}

	_jsii_.Create(
		"awscc.sagemakerCluster.SagemakerClusterInstanceGroupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerClusterInstanceGroupsOutputReference_Override(s SagemakerClusterInstanceGroupsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerCluster.SagemakerClusterInstanceGroupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetCurrentCount(val *float64) {
	if err := j.validateSetCurrentCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currentCount",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetInstanceGroupName(val *string) {
	if err := j.validateSetInstanceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceGroupName",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetOnStartDeepHealthChecks(val *[]*string) {
	if err := j.validateSetOnStartDeepHealthChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onStartDeepHealthChecks",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetThreadsPerCore(val *float64) {
	if err := j.validateSetThreadsPerCoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadsPerCore",
		val,
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutInstanceStorageConfigs(value interface{}) {
	if err := s.validatePutInstanceStorageConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInstanceStorageConfigs",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutLifeCycleConfig(value *SagemakerClusterInstanceGroupsLifeCycleConfig) {
	if err := s.validatePutLifeCycleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLifeCycleConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutOverrideVpcConfig(value *SagemakerClusterInstanceGroupsOverrideVpcConfig) {
	if err := s.validatePutOverrideVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOverrideVpcConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetCurrentCount() {
	_jsii_.InvokeVoid(
		s,
		"resetCurrentCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetInstanceStorageConfigs() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceStorageConfigs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetOnStartDeepHealthChecks() {
	_jsii_.InvokeVoid(
		s,
		"resetOnStartDeepHealthChecks",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetOverrideVpcConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideVpcConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetThreadsPerCore() {
	_jsii_.InvokeVoid(
		s,
		"resetThreadsPerCore",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

