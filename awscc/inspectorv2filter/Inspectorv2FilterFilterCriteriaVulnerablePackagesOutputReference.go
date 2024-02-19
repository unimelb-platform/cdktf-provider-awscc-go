package inspectorv2filter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/inspectorv2filter/internal"
)

type Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference interface {
	cdktf.ComplexObject
	Architecture() Inspectorv2FilterFilterCriteriaVulnerablePackagesArchitectureOutputReference
	ArchitectureInput() interface{}
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
	Epoch() Inspectorv2FilterFilterCriteriaVulnerablePackagesEpochOutputReference
	EpochInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() Inspectorv2FilterFilterCriteriaVulnerablePackagesNameOutputReference
	NameInput() interface{}
	Release() Inspectorv2FilterFilterCriteriaVulnerablePackagesReleaseOutputReference
	ReleaseInput() interface{}
	SourceLayerHash() Inspectorv2FilterFilterCriteriaVulnerablePackagesSourceLayerHashOutputReference
	SourceLayerHashInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() Inspectorv2FilterFilterCriteriaVulnerablePackagesVersionOutputReference
	VersionInput() interface{}
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
	PutArchitecture(value *Inspectorv2FilterFilterCriteriaVulnerablePackagesArchitecture)
	PutEpoch(value *Inspectorv2FilterFilterCriteriaVulnerablePackagesEpoch)
	PutName(value *Inspectorv2FilterFilterCriteriaVulnerablePackagesName)
	PutRelease(value *Inspectorv2FilterFilterCriteriaVulnerablePackagesRelease)
	PutSourceLayerHash(value *Inspectorv2FilterFilterCriteriaVulnerablePackagesSourceLayerHash)
	PutVersion(value *Inspectorv2FilterFilterCriteriaVulnerablePackagesVersion)
	ResetArchitecture()
	ResetEpoch()
	ResetName()
	ResetRelease()
	ResetSourceLayerHash()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference
type jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Architecture() Inspectorv2FilterFilterCriteriaVulnerablePackagesArchitectureOutputReference {
	var returns Inspectorv2FilterFilterCriteriaVulnerablePackagesArchitectureOutputReference
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ArchitectureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Epoch() Inspectorv2FilterFilterCriteriaVulnerablePackagesEpochOutputReference {
	var returns Inspectorv2FilterFilterCriteriaVulnerablePackagesEpochOutputReference
	_jsii_.Get(
		j,
		"epoch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) EpochInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"epochInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Name() Inspectorv2FilterFilterCriteriaVulnerablePackagesNameOutputReference {
	var returns Inspectorv2FilterFilterCriteriaVulnerablePackagesNameOutputReference
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) NameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Release() Inspectorv2FilterFilterCriteriaVulnerablePackagesReleaseOutputReference {
	var returns Inspectorv2FilterFilterCriteriaVulnerablePackagesReleaseOutputReference
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ReleaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"releaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) SourceLayerHash() Inspectorv2FilterFilterCriteriaVulnerablePackagesSourceLayerHashOutputReference {
	var returns Inspectorv2FilterFilterCriteriaVulnerablePackagesSourceLayerHashOutputReference
	_jsii_.Get(
		j,
		"sourceLayerHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) SourceLayerHashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceLayerHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Version() Inspectorv2FilterFilterCriteriaVulnerablePackagesVersionOutputReference {
	var returns Inspectorv2FilterFilterCriteriaVulnerablePackagesVersionOutputReference
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) VersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference {
	_init_.Initialize()

	if err := validateNewInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference{}

	_jsii_.Create(
		"awscc.inspectorv2Filter.Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference_Override(i Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.inspectorv2Filter.Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) PutArchitecture(value *Inspectorv2FilterFilterCriteriaVulnerablePackagesArchitecture) {
	if err := i.validatePutArchitectureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putArchitecture",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) PutEpoch(value *Inspectorv2FilterFilterCriteriaVulnerablePackagesEpoch) {
	if err := i.validatePutEpochParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEpoch",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) PutName(value *Inspectorv2FilterFilterCriteriaVulnerablePackagesName) {
	if err := i.validatePutNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putName",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) PutRelease(value *Inspectorv2FilterFilterCriteriaVulnerablePackagesRelease) {
	if err := i.validatePutReleaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRelease",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) PutSourceLayerHash(value *Inspectorv2FilterFilterCriteriaVulnerablePackagesSourceLayerHash) {
	if err := i.validatePutSourceLayerHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSourceLayerHash",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) PutVersion(value *Inspectorv2FilterFilterCriteriaVulnerablePackagesVersion) {
	if err := i.validatePutVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putVersion",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetArchitecture() {
	_jsii_.InvokeVoid(
		i,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetEpoch() {
	_jsii_.InvokeVoid(
		i,
		"resetEpoch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		i,
		"resetName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetRelease() {
	_jsii_.InvokeVoid(
		i,
		"resetRelease",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetSourceLayerHash() {
	_jsii_.InvokeVoid(
		i,
		"resetSourceLayerHash",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		i,
		"resetVersion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

