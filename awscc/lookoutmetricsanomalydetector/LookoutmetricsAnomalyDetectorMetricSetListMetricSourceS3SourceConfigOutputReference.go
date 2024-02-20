package lookoutmetricsanomalydetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lookoutmetricsanomalydetector/internal"
)

type LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference interface {
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
	FileFormatDescriptor() LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptorOutputReference
	FileFormatDescriptorInput() *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptor
	// Experimental.
	Fqn() *string
	HistoricalDataPathList() *[]*string
	SetHistoricalDataPathList(val *[]*string)
	HistoricalDataPathListInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TemplatedPathList() *[]*string
	SetTemplatedPathList(val *[]*string)
	TemplatedPathListInput() *[]*string
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
	PutFileFormatDescriptor(value *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptor)
	ResetHistoricalDataPathList()
	ResetTemplatedPathList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference
type jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) FileFormatDescriptor() LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptorOutputReference {
	var returns LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptorOutputReference
	_jsii_.Get(
		j,
		"fileFormatDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) FileFormatDescriptorInput() *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptor {
	var returns *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptor
	_jsii_.Get(
		j,
		"fileFormatDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) HistoricalDataPathList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"historicalDataPathList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) HistoricalDataPathListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"historicalDataPathListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) TemplatedPathList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"templatedPathList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) TemplatedPathListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"templatedPathListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference{}

	_jsii_.Create(
		"awscc.lookoutmetricsAnomalyDetector.LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference_Override(l LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lookoutmetricsAnomalyDetector.LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference)SetHistoricalDataPathList(val *[]*string) {
	if err := j.validateSetHistoricalDataPathListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"historicalDataPathList",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference)SetTemplatedPathList(val *[]*string) {
	if err := j.validateSetTemplatedPathListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templatedPathList",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) PutFileFormatDescriptor(value *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptor) {
	if err := l.validatePutFileFormatDescriptorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFileFormatDescriptor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) ResetHistoricalDataPathList() {
	_jsii_.InvokeVoid(
		l,
		"resetHistoricalDataPathList",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) ResetTemplatedPathList() {
	_jsii_.InvokeVoid(
		l,
		"resetTemplatedPathList",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

