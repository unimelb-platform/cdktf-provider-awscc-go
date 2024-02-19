package iotdomainconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iotdomainconfiguration/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_domain_configuration awscc_iot_domain_configuration}.
type IotDomainConfiguration interface {
	cdktf.TerraformResource
	Arn() *string
	AuthorizerConfig() IotDomainConfigurationAuthorizerConfigOutputReference
	AuthorizerConfigInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DomainConfigurationName() *string
	SetDomainConfigurationName(val *string)
	DomainConfigurationNameInput() *string
	DomainConfigurationStatus() *string
	SetDomainConfigurationStatus(val *string)
	DomainConfigurationStatusInput() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	DomainType() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ServerCertificateArns() *[]*string
	SetServerCertificateArns(val *[]*string)
	ServerCertificateArnsInput() *[]*string
	ServerCertificates() IotDomainConfigurationServerCertificatesList
	ServiceType() *string
	SetServiceType(val *string)
	ServiceTypeInput() *string
	Tags() IotDomainConfigurationTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TlsConfig() IotDomainConfigurationTlsConfigOutputReference
	TlsConfigInput() interface{}
	ValidationCertificateArn() *string
	SetValidationCertificateArn(val *string)
	ValidationCertificateArnInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAuthorizerConfig(value *IotDomainConfigurationAuthorizerConfig)
	PutTags(value interface{})
	PutTlsConfig(value *IotDomainConfigurationTlsConfig)
	ResetAuthorizerConfig()
	ResetDomainConfigurationName()
	ResetDomainConfigurationStatus()
	ResetDomainName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServerCertificateArns()
	ResetServiceType()
	ResetTags()
	ResetTlsConfig()
	ResetValidationCertificateArn()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for IotDomainConfiguration
type jsiiProxy_IotDomainConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotDomainConfiguration) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) AuthorizerConfig() IotDomainConfigurationAuthorizerConfigOutputReference {
	var returns IotDomainConfigurationAuthorizerConfigOutputReference
	_jsii_.Get(
		j,
		"authorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) AuthorizerConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainConfigurationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainConfigurationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainConfigurationStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainConfigurationStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ServerCertificateArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverCertificateArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ServerCertificateArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverCertificateArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ServerCertificates() IotDomainConfigurationServerCertificatesList {
	var returns IotDomainConfigurationServerCertificatesList
	_jsii_.Get(
		j,
		"serverCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ServiceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ServiceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Tags() IotDomainConfigurationTagsList {
	var returns IotDomainConfigurationTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) TlsConfig() IotDomainConfigurationTlsConfigOutputReference {
	var returns IotDomainConfigurationTlsConfigOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) TlsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ValidationCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ValidationCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationCertificateArnInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_domain_configuration awscc_iot_domain_configuration} Resource.
func NewIotDomainConfiguration(scope constructs.Construct, id *string, config *IotDomainConfigurationConfig) IotDomainConfiguration {
	_init_.Initialize()

	if err := validateNewIotDomainConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotDomainConfiguration{}

	_jsii_.Create(
		"awscc.iotDomainConfiguration.IotDomainConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_domain_configuration awscc_iot_domain_configuration} Resource.
func NewIotDomainConfiguration_Override(i IotDomainConfiguration, scope constructs.Construct, id *string, config *IotDomainConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotDomainConfiguration.IotDomainConfiguration",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetDomainConfigurationName(val *string) {
	if err := j.validateSetDomainConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainConfigurationName",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetDomainConfigurationStatus(val *string) {
	if err := j.validateSetDomainConfigurationStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainConfigurationStatus",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetServerCertificateArns(val *[]*string) {
	if err := j.validateSetServerCertificateArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCertificateArns",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetServiceType(val *string) {
	if err := j.validateSetServiceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceType",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetValidationCertificateArn(val *string) {
	if err := j.validateSetValidationCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationCertificateArn",
		val,
	)
}

// Generates CDKTF code for importing a IotDomainConfiguration resource upon running "cdktf plan <stack-name>".
func IotDomainConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIotDomainConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.iotDomainConfiguration.IotDomainConfiguration",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func IotDomainConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotDomainConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotDomainConfiguration.IotDomainConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotDomainConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotDomainConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotDomainConfiguration.IotDomainConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotDomainConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotDomainConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotDomainConfiguration.IotDomainConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotDomainConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.iotDomainConfiguration.IotDomainConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotDomainConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotDomainConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotDomainConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotDomainConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotDomainConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotDomainConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotDomainConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotDomainConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotDomainConfiguration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) PutAuthorizerConfig(value *IotDomainConfigurationAuthorizerConfig) {
	if err := i.validatePutAuthorizerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAuthorizerConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) PutTlsConfig(value *IotDomainConfigurationTlsConfig) {
	if err := i.validatePutTlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTlsConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetAuthorizerConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthorizerConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetDomainConfigurationName() {
	_jsii_.InvokeVoid(
		i,
		"resetDomainConfigurationName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetDomainConfigurationStatus() {
	_jsii_.InvokeVoid(
		i,
		"resetDomainConfigurationStatus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetDomainName() {
	_jsii_.InvokeVoid(
		i,
		"resetDomainName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetServerCertificateArns() {
	_jsii_.InvokeVoid(
		i,
		"resetServerCertificateArns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetServiceType() {
	_jsii_.InvokeVoid(
		i,
		"resetServiceType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetValidationCertificateArn() {
	_jsii_.InvokeVoid(
		i,
		"resetValidationCertificateArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

