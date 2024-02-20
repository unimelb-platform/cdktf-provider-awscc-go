package lookoutmetricsanomalydetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lookoutmetricsanomalydetector/internal"
)

type LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference interface {
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
	DatabaseHost() *string
	SetDatabaseHost(val *string)
	DatabaseHostInput() *string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DatabasePort() *float64
	SetDatabasePort(val *float64)
	DatabasePortInput() *float64
	DbInstanceIdentifier() *string
	SetDbInstanceIdentifier(val *string)
	DbInstanceIdentifierInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SecretManagerArn() *string
	SetSecretManagerArn(val *string)
	SecretManagerArnInput() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcConfiguration() LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigVpcConfigurationOutputReference
	VpcConfigurationInput() *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigVpcConfiguration
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
	PutVpcConfiguration(value *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigVpcConfiguration)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference
type jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) DatabaseHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) DatabaseHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) DatabasePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"databasePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) DatabasePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"databasePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) DbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) DbInstanceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) SecretManagerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) SecretManagerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) VpcConfiguration() LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigVpcConfigurationOutputReference {
	var returns LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigVpcConfigurationOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) VpcConfigurationInput() *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigVpcConfiguration {
	var returns *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigVpcConfiguration
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


func NewLookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference{}

	_jsii_.Create(
		"awscc.lookoutmetricsAnomalyDetector.LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference_Override(l LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lookoutmetricsAnomalyDetector.LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference)SetDatabaseHost(val *string) {
	if err := j.validateSetDatabaseHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseHost",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference)SetDatabasePort(val *float64) {
	if err := j.validateSetDatabasePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databasePort",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference)SetDbInstanceIdentifier(val *string) {
	if err := j.validateSetDbInstanceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference)SetSecretManagerArn(val *string) {
	if err := j.validateSetSecretManagerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretManagerArn",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) PutVpcConfiguration(value *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigVpcConfiguration) {
	if err := l.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

