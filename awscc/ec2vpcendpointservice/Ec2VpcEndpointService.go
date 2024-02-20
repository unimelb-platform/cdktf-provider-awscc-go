package ec2vpcendpointservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2vpcendpointservice/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_endpoint_service awscc_ec2_vpc_endpoint_service}.
type Ec2VpcEndpointService interface {
	cdktf.TerraformResource
	AcceptanceRequired() interface{}
	SetAcceptanceRequired(val interface{})
	AcceptanceRequiredInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContributorInsightsEnabled() interface{}
	SetContributorInsightsEnabled(val interface{})
	ContributorInsightsEnabledInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayLoadBalancerArns() *[]*string
	SetGatewayLoadBalancerArns(val *[]*string)
	GatewayLoadBalancerArnsInput() *[]*string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NetworkLoadBalancerArns() *[]*string
	SetNetworkLoadBalancerArns(val *[]*string)
	NetworkLoadBalancerArnsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	PayerResponsibility() *string
	SetPayerResponsibility(val *string)
	PayerResponsibilityInput() *string
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
	ServiceId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAcceptanceRequired()
	ResetContributorInsightsEnabled()
	ResetGatewayLoadBalancerArns()
	ResetNetworkLoadBalancerArns()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPayerResponsibility()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Ec2VpcEndpointService
type jsiiProxy_Ec2VpcEndpointService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Ec2VpcEndpointService) AcceptanceRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptanceRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) AcceptanceRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptanceRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) ContributorInsightsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contributorInsightsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) ContributorInsightsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contributorInsightsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) GatewayLoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gatewayLoadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) GatewayLoadBalancerArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gatewayLoadBalancerArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) NetworkLoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkLoadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) NetworkLoadBalancerArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkLoadBalancerArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) PayerResponsibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payerResponsibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) PayerResponsibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payerResponsibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcEndpointService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_endpoint_service awscc_ec2_vpc_endpoint_service} Resource.
func NewEc2VpcEndpointService(scope constructs.Construct, id *string, config *Ec2VpcEndpointServiceConfig) Ec2VpcEndpointService {
	_init_.Initialize()

	if err := validateNewEc2VpcEndpointServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2VpcEndpointService{}

	_jsii_.Create(
		"awscc.ec2VpcEndpointService.Ec2VpcEndpointService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_endpoint_service awscc_ec2_vpc_endpoint_service} Resource.
func NewEc2VpcEndpointService_Override(e Ec2VpcEndpointService, scope constructs.Construct, id *string, config *Ec2VpcEndpointServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2VpcEndpointService.Ec2VpcEndpointService",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointService)SetAcceptanceRequired(val interface{}) {
	if err := j.validateSetAcceptanceRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptanceRequired",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointService)SetContributorInsightsEnabled(val interface{}) {
	if err := j.validateSetContributorInsightsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contributorInsightsEnabled",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointService)SetGatewayLoadBalancerArns(val *[]*string) {
	if err := j.validateSetGatewayLoadBalancerArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayLoadBalancerArns",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointService)SetNetworkLoadBalancerArns(val *[]*string) {
	if err := j.validateSetNetworkLoadBalancerArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkLoadBalancerArns",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointService)SetPayerResponsibility(val *string) {
	if err := j.validateSetPayerResponsibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payerResponsibility",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcEndpointService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a Ec2VpcEndpointService resource upon running "cdktf plan <stack-name>".
func Ec2VpcEndpointService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEc2VpcEndpointService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.ec2VpcEndpointService.Ec2VpcEndpointService",
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
func Ec2VpcEndpointService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2VpcEndpointService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2VpcEndpointService.Ec2VpcEndpointService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2VpcEndpointService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2VpcEndpointService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2VpcEndpointService.Ec2VpcEndpointService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2VpcEndpointService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2VpcEndpointService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2VpcEndpointService.Ec2VpcEndpointService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ec2VpcEndpointService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.ec2VpcEndpointService.Ec2VpcEndpointService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2VpcEndpointService) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Ec2VpcEndpointService) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Ec2VpcEndpointService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2VpcEndpointService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2VpcEndpointService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2VpcEndpointService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2VpcEndpointService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2VpcEndpointService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2VpcEndpointService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2VpcEndpointService) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2VpcEndpointService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2VpcEndpointService) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Ec2VpcEndpointService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpcEndpointService) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Ec2VpcEndpointService) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2VpcEndpointService) ResetAcceptanceRequired() {
	_jsii_.InvokeVoid(
		e,
		"resetAcceptanceRequired",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcEndpointService) ResetContributorInsightsEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetContributorInsightsEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcEndpointService) ResetGatewayLoadBalancerArns() {
	_jsii_.InvokeVoid(
		e,
		"resetGatewayLoadBalancerArns",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcEndpointService) ResetNetworkLoadBalancerArns() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkLoadBalancerArns",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcEndpointService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcEndpointService) ResetPayerResponsibility() {
	_jsii_.InvokeVoid(
		e,
		"resetPayerResponsibility",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcEndpointService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpcEndpointService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpcEndpointService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpcEndpointService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

