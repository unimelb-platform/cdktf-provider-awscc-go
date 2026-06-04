package gluecrawler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/gluecrawler/internal"
)

type GlueCrawlerTargetsDeltaTargetsOutputReference interface {
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
	ConnectionName() *string
	SetConnectionName(val *string)
	ConnectionNameInput() *string
	CreateNativeDeltaTable() interface{}
	SetCreateNativeDeltaTable(val interface{})
	CreateNativeDeltaTableInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeltaTables() *[]*string
	SetDeltaTables(val *[]*string)
	DeltaTablesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WriteManifest() interface{}
	SetWriteManifest(val interface{})
	WriteManifestInput() interface{}
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
	ResetConnectionName()
	ResetCreateNativeDeltaTable()
	ResetDeltaTables()
	ResetWriteManifest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueCrawlerTargetsDeltaTargetsOutputReference
type jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) ConnectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) CreateNativeDeltaTable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createNativeDeltaTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) CreateNativeDeltaTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createNativeDeltaTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) DeltaTables() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deltaTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) DeltaTablesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deltaTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) WriteManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeManifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) WriteManifestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeManifestInput",
		&returns,
	)
	return returns
}


func NewGlueCrawlerTargetsDeltaTargetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GlueCrawlerTargetsDeltaTargetsOutputReference {
	_init_.Initialize()

	if err := validateNewGlueCrawlerTargetsDeltaTargetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference{}

	_jsii_.Create(
		"awscc.glueCrawler.GlueCrawlerTargetsDeltaTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGlueCrawlerTargetsDeltaTargetsOutputReference_Override(g GlueCrawlerTargetsDeltaTargetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.glueCrawler.GlueCrawlerTargetsDeltaTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference)SetConnectionName(val *string) {
	if err := j.validateSetConnectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionName",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference)SetCreateNativeDeltaTable(val interface{}) {
	if err := j.validateSetCreateNativeDeltaTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createNativeDeltaTable",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference)SetDeltaTables(val *[]*string) {
	if err := j.validateSetDeltaTablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaTables",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference)SetWriteManifest(val interface{}) {
	if err := j.validateSetWriteManifestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeManifest",
		val,
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) ResetConnectionName() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) ResetCreateNativeDeltaTable() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateNativeDeltaTable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) ResetDeltaTables() {
	_jsii_.InvokeVoid(
		g,
		"resetDeltaTables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) ResetWriteManifest() {
	_jsii_.InvokeVoid(
		g,
		"resetWriteManifest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsDeltaTargetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

