package servicecatalogcloudformationprovisionedproduct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/servicecatalogcloudformationprovisionedproduct/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/servicecatalog_cloudformation_provisioned_product awscc_servicecatalog_cloudformation_provisioned_product}.
type ServicecatalogCloudformationProvisionedProduct interface {
	cdktf.TerraformResource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudformationStackArn() *string
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
	NotificationArns() *[]*string
	SetNotificationArns(val *[]*string)
	NotificationArnsInput() *[]*string
	Outputs() cdktf.StringMap
	PathId() *string
	SetPathId(val *string)
	PathIdInput() *string
	PathName() *string
	SetPathName(val *string)
	PathNameInput() *string
	ProductId() *string
	SetProductId(val *string)
	ProductIdInput() *string
	ProductName() *string
	SetProductName(val *string)
	ProductNameInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProvisionedProductId() *string
	ProvisionedProductName() *string
	SetProvisionedProductName(val *string)
	ProvisionedProductNameInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProvisioningArtifactId() *string
	SetProvisioningArtifactId(val *string)
	ProvisioningArtifactIdInput() *string
	ProvisioningArtifactName() *string
	SetProvisioningArtifactName(val *string)
	ProvisioningArtifactNameInput() *string
	ProvisioningParameters() ServicecatalogCloudformationProvisionedProductProvisioningParametersList
	ProvisioningParametersInput() interface{}
	ProvisioningPreferences() ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference
	ProvisioningPreferencesInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RecordId() *string
	Tags() ServicecatalogCloudformationProvisionedProductTagsList
	TagsInput() interface{}
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutProvisioningParameters(value interface{})
	PutProvisioningPreferences(value *ServicecatalogCloudformationProvisionedProductProvisioningPreferences)
	PutTags(value interface{})
	ResetAcceptLanguage()
	ResetNotificationArns()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPathId()
	ResetPathName()
	ResetProductId()
	ResetProductName()
	ResetProvisionedProductName()
	ResetProvisioningArtifactId()
	ResetProvisioningArtifactName()
	ResetProvisioningParameters()
	ResetProvisioningPreferences()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicecatalogCloudformationProvisionedProduct
type jsiiProxy_ServicecatalogCloudformationProvisionedProduct struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) CloudformationStackArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudformationStackArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) NotificationArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) Outputs() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) PathId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) PathIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) PathName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) PathNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProductName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProductNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProvisionedProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisionedProductId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProvisionedProductName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisionedProductName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProvisionedProductNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisionedProductNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProvisioningArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProvisioningArtifactIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProvisioningArtifactName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProvisioningArtifactNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProvisioningParameters() ServicecatalogCloudformationProvisionedProductProvisioningParametersList {
	var returns ServicecatalogCloudformationProvisionedProductProvisioningParametersList
	_jsii_.Get(
		j,
		"provisioningParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProvisioningParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisioningParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProvisioningPreferences() ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference {
	var returns ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference
	_jsii_.Get(
		j,
		"provisioningPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ProvisioningPreferencesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisioningPreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) RecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) Tags() ServicecatalogCloudformationProvisionedProductTagsList {
	var returns ServicecatalogCloudformationProvisionedProductTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/servicecatalog_cloudformation_provisioned_product awscc_servicecatalog_cloudformation_provisioned_product} Resource.
func NewServicecatalogCloudformationProvisionedProduct(scope constructs.Construct, id *string, config *ServicecatalogCloudformationProvisionedProductConfig) ServicecatalogCloudformationProvisionedProduct {
	_init_.Initialize()

	if err := validateNewServicecatalogCloudformationProvisionedProductParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicecatalogCloudformationProvisionedProduct{}

	_jsii_.Create(
		"awscc.servicecatalogCloudformationProvisionedProduct.ServicecatalogCloudformationProvisionedProduct",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/servicecatalog_cloudformation_provisioned_product awscc_servicecatalog_cloudformation_provisioned_product} Resource.
func NewServicecatalogCloudformationProvisionedProduct_Override(s ServicecatalogCloudformationProvisionedProduct, scope constructs.Construct, id *string, config *ServicecatalogCloudformationProvisionedProductConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.servicecatalogCloudformationProvisionedProduct.ServicecatalogCloudformationProvisionedProduct",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetAcceptLanguage(val *string) {
	if err := j.validateSetAcceptLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetNotificationArns(val *[]*string) {
	if err := j.validateSetNotificationArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationArns",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetPathId(val *string) {
	if err := j.validateSetPathIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetPathName(val *string) {
	if err := j.validateSetPathNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathName",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetProductId(val *string) {
	if err := j.validateSetProductIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetProductName(val *string) {
	if err := j.validateSetProductNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productName",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetProvisionedProductName(val *string) {
	if err := j.validateSetProvisionedProductNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedProductName",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetProvisioningArtifactId(val *string) {
	if err := j.validateSetProvisioningArtifactIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningArtifactId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProduct)SetProvisioningArtifactName(val *string) {
	if err := j.validateSetProvisioningArtifactNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningArtifactName",
		val,
	)
}

// Generates CDKTF code for importing a ServicecatalogCloudformationProvisionedProduct resource upon running "cdktf plan <stack-name>".
func ServicecatalogCloudformationProvisionedProduct_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServicecatalogCloudformationProvisionedProduct_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.servicecatalogCloudformationProvisionedProduct.ServicecatalogCloudformationProvisionedProduct",
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
func ServicecatalogCloudformationProvisionedProduct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicecatalogCloudformationProvisionedProduct_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.servicecatalogCloudformationProvisionedProduct.ServicecatalogCloudformationProvisionedProduct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicecatalogCloudformationProvisionedProduct_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicecatalogCloudformationProvisionedProduct_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.servicecatalogCloudformationProvisionedProduct.ServicecatalogCloudformationProvisionedProduct",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicecatalogCloudformationProvisionedProduct_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicecatalogCloudformationProvisionedProduct_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.servicecatalogCloudformationProvisionedProduct.ServicecatalogCloudformationProvisionedProduct",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogCloudformationProvisionedProduct_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.servicecatalogCloudformationProvisionedProduct.ServicecatalogCloudformationProvisionedProduct",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) PutProvisioningParameters(value interface{}) {
	if err := s.validatePutProvisioningParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProvisioningParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) PutProvisioningPreferences(value *ServicecatalogCloudformationProvisionedProductProvisioningPreferences) {
	if err := s.validatePutProvisioningPreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProvisioningPreferences",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) PutTags(value interface{}) {
	if err := s.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetNotificationArns() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationArns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetPathId() {
	_jsii_.InvokeVoid(
		s,
		"resetPathId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetPathName() {
	_jsii_.InvokeVoid(
		s,
		"resetPathName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetProductId() {
	_jsii_.InvokeVoid(
		s,
		"resetProductId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetProductName() {
	_jsii_.InvokeVoid(
		s,
		"resetProductName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetProvisionedProductName() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisionedProductName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetProvisioningArtifactId() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisioningArtifactId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetProvisioningArtifactName() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisioningArtifactName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetProvisioningParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisioningParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetProvisioningPreferences() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisioningPreferences",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProduct) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

