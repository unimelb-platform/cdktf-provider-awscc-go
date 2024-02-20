package ec2instance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2instance/internal"
)

type Ec2InstanceBlockDeviceMappingsEbsOutputReference interface {
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
	DeleteOnTermination() interface{}
	SetDeleteOnTermination(val interface{})
	DeleteOnTerminationInput() interface{}
	Encrypted() interface{}
	SetEncrypted(val interface{})
	EncryptedInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	SnapshotId() *string
	SetSnapshotId(val *string)
	SnapshotIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeSize() *float64
	SetVolumeSize(val *float64)
	VolumeSizeInput() *float64
	VolumeType() *string
	SetVolumeType(val *string)
	VolumeTypeInput() *string
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
	ResetDeleteOnTermination()
	ResetEncrypted()
	ResetIops()
	ResetKmsKeyId()
	ResetSnapshotId()
	ResetVolumeSize()
	ResetVolumeType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2InstanceBlockDeviceMappingsEbsOutputReference
type jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) DeleteOnTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) DeleteOnTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) Encrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) EncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) VolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) VolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}


func NewEc2InstanceBlockDeviceMappingsEbsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2InstanceBlockDeviceMappingsEbsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2InstanceBlockDeviceMappingsEbsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference{}

	_jsii_.Create(
		"awscc.ec2Instance.Ec2InstanceBlockDeviceMappingsEbsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2InstanceBlockDeviceMappingsEbsOutputReference_Override(e Ec2InstanceBlockDeviceMappingsEbsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2Instance.Ec2InstanceBlockDeviceMappingsEbsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference)SetDeleteOnTermination(val interface{}) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference)SetEncrypted(val interface{}) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference)SetVolumeSize(val *float64) {
	if err := j.validateSetVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSize",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference)SetVolumeType(val *string) {
	if err := j.validateSetVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		e,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) ResetEncrypted() {
	_jsii_.InvokeVoid(
		e,
		"resetEncrypted",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) ResetIops() {
	_jsii_.InvokeVoid(
		e,
		"resetIops",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		e,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) ResetVolumeSize() {
	_jsii_.InvokeVoid(
		e,
		"resetVolumeSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) ResetVolumeType() {
	_jsii_.InvokeVoid(
		e,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsEbsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

