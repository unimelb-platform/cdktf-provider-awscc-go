package dataawsccbatchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccbatchjobdefinition/internal"
)

type DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference interface {
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
	Containers() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsPolicy() *string
	// Experimental.
	Fqn() *string
	HostNetwork() cdktf.IResolvable
	ImagePullSecrets() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesImagePullSecretsList
	InitContainers() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesInitContainersList
	InternalValue() *DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodProperties
	SetInternalValue(val *DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodProperties)
	Metadata() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesMetadataOutputReference
	ServiceAccountName() *string
	ShareProcessNamespace() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volumes() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesVolumesList
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

// The jsii proxy struct for DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference
type jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) Containers() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersList {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersList
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) HostNetwork() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) ImagePullSecrets() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesImagePullSecretsList {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesImagePullSecretsList
	_jsii_.Get(
		j,
		"imagePullSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) InitContainers() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesInitContainersList {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesInitContainersList
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) InternalValue() *DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodProperties {
	var returns *DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) Metadata() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesMetadataOutputReference {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesMetadataOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) ShareProcessNamespace() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"shareProcessNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) Volumes() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesVolumesList {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewDataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccBatchJobDefinition.DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference_Override(d DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccBatchJobDefinition.DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference)SetInternalValue(val *DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

