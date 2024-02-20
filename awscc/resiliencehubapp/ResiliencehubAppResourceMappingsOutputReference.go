package resiliencehubapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/resiliencehubapp/internal"
)

type ResiliencehubAppResourceMappingsOutputReference interface {
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
	EksSourceName() *string
	SetEksSourceName(val *string)
	EksSourceNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogicalStackName() *string
	SetLogicalStackName(val *string)
	LogicalStackNameInput() *string
	MappingType() *string
	SetMappingType(val *string)
	MappingTypeInput() *string
	PhysicalResourceId() ResiliencehubAppResourceMappingsPhysicalResourceIdOutputReference
	PhysicalResourceIdInput() interface{}
	ResourceName() *string
	SetResourceName(val *string)
	ResourceNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TerraformSourceName() *string
	SetTerraformSourceName(val *string)
	TerraformSourceNameInput() *string
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
	PutPhysicalResourceId(value *ResiliencehubAppResourceMappingsPhysicalResourceId)
	ResetEksSourceName()
	ResetLogicalStackName()
	ResetResourceName()
	ResetTerraformSourceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ResiliencehubAppResourceMappingsOutputReference
type jsiiProxy_ResiliencehubAppResourceMappingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) EksSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eksSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) EksSourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eksSourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) LogicalStackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalStackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) LogicalStackNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalStackNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) MappingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) MappingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) PhysicalResourceId() ResiliencehubAppResourceMappingsPhysicalResourceIdOutputReference {
	var returns ResiliencehubAppResourceMappingsPhysicalResourceIdOutputReference
	_jsii_.Get(
		j,
		"physicalResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) PhysicalResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"physicalResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) ResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) TerraformSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) TerraformSourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformSourceNameInput",
		&returns,
	)
	return returns
}


func NewResiliencehubAppResourceMappingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ResiliencehubAppResourceMappingsOutputReference {
	_init_.Initialize()

	if err := validateNewResiliencehubAppResourceMappingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResiliencehubAppResourceMappingsOutputReference{}

	_jsii_.Create(
		"awscc.resiliencehubApp.ResiliencehubAppResourceMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewResiliencehubAppResourceMappingsOutputReference_Override(r ResiliencehubAppResourceMappingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.resiliencehubApp.ResiliencehubAppResourceMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference)SetEksSourceName(val *string) {
	if err := j.validateSetEksSourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eksSourceName",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference)SetLogicalStackName(val *string) {
	if err := j.validateSetLogicalStackNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logicalStackName",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference)SetMappingType(val *string) {
	if err := j.validateSetMappingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mappingType",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference)SetResourceName(val *string) {
	if err := j.validateSetResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceName",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference)SetTerraformSourceName(val *string) {
	if err := j.validateSetTerraformSourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformSourceName",
		val,
	)
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) PutPhysicalResourceId(value *ResiliencehubAppResourceMappingsPhysicalResourceId) {
	if err := r.validatePutPhysicalResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPhysicalResourceId",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) ResetEksSourceName() {
	_jsii_.InvokeVoid(
		r,
		"resetEksSourceName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) ResetLogicalStackName() {
	_jsii_.InvokeVoid(
		r,
		"resetLogicalStackName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) ResetResourceName() {
	_jsii_.InvokeVoid(
		r,
		"resetResourceName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) ResetTerraformSourceName() {
	_jsii_.InvokeVoid(
		r,
		"resetTerraformSourceName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubAppResourceMappingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

