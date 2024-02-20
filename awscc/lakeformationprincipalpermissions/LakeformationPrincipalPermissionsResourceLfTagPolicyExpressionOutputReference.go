package lakeformationprincipalpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lakeformationprincipalpermissions/internal"
)

type LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *LakeformationPrincipalPermissionsResourceLfTagPolicyExpression
	SetInternalValue(val *LakeformationPrincipalPermissionsResourceLfTagPolicyExpression)
	TagKey() *string
	SetTagKey(val *string)
	TagKeyInput() *string
	TagValues() *[]*string
	SetTagValues(val *[]*string)
	TagValuesInput() *[]*string
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
	ResetTagKey()
	ResetTagValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference
type jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) InternalValue() *LakeformationPrincipalPermissionsResourceLfTagPolicyExpression {
	var returns *LakeformationPrincipalPermissionsResourceLfTagPolicyExpression
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) TagKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) TagKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) TagValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) TagValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference {
	_init_.Initialize()

	if err := validateNewLakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference{}

	_jsii_.Create(
		"awscc.lakeformationPrincipalPermissions.LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference_Override(l LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lakeformationPrincipalPermissions.LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference)SetInternalValue(val *LakeformationPrincipalPermissionsResourceLfTagPolicyExpression) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference)SetTagKey(val *string) {
	if err := j.validateSetTagKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagKey",
		val,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference)SetTagValues(val *[]*string) {
	if err := j.validateSetTagValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagValues",
		val,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) ResetTagKey() {
	_jsii_.InvokeVoid(
		l,
		"resetTagKey",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) ResetTagValues() {
	_jsii_.InvokeVoid(
		l,
		"resetTagValues",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceLfTagPolicyExpressionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

