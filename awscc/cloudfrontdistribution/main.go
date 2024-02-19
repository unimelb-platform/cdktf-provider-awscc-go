package cloudfrontdistribution

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistribution",
		reflect.TypeOf((*CloudfrontDistribution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "distributionConfig", GoGetter: "DistributionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "distributionConfigInput", GoGetter: "DistributionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putDistributionConfig", GoMethod: "PutDistributionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putTags", GoMethod: "PutTags"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistribution{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionConfig",
		reflect.TypeOf((*CloudfrontDistributionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfig",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviors",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviors)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValues",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValues)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookies",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookies)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "forward", GoGetter: "Forward"},
			_jsii_.MemberProperty{JsiiProperty: "forwardInput", GoGetter: "ForwardInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhitelistedNames", GoMethod: "ResetWhitelistedNames"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistedNames", GoGetter: "WhitelistedNames"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistedNamesInput", GoGetter: "WhitelistedNamesInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cookies", GoGetter: "Cookies"},
			_jsii_.MemberProperty{JsiiProperty: "cookiesInput", GoGetter: "CookiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putCookies", GoMethod: "PutCookies"},
			_jsii_.MemberProperty{JsiiProperty: "queryString", GoGetter: "QueryString"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCacheKeys", GoGetter: "QueryStringCacheKeys"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCacheKeysInput", GoGetter: "QueryStringCacheKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringInput", GoGetter: "QueryStringInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCookies", GoMethod: "ResetCookies"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryStringCacheKeys", GoMethod: "ResetQueryStringCacheKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsFunctionAssociations",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviorsFunctionAssociations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsFunctionAssociationsList",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviorsFunctionAssociationsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsFunctionAssociationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsFunctionAssociationsOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviorsFunctionAssociationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eventType", GoGetter: "EventType"},
			_jsii_.MemberProperty{JsiiProperty: "eventTypeInput", GoGetter: "EventTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionArnInput", GoGetter: "FunctionArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventType", GoMethod: "ResetEventType"},
			_jsii_.MemberMethod{JsiiMethod: "resetFunctionArn", GoMethod: "ResetFunctionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsFunctionAssociationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociations",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsList",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eventType", GoGetter: "EventType"},
			_jsii_.MemberProperty{JsiiProperty: "eventTypeInput", GoGetter: "EventTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "includeBody", GoGetter: "IncludeBody"},
			_jsii_.MemberProperty{JsiiProperty: "includeBodyInput", GoGetter: "IncludeBodyInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionArn", GoGetter: "LambdaFunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionArnInput", GoGetter: "LambdaFunctionArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventType", GoMethod: "ResetEventType"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeBody", GoMethod: "ResetIncludeBody"},
			_jsii_.MemberMethod{JsiiMethod: "resetLambdaFunctionArn", GoMethod: "ResetLambdaFunctionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsList",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviorsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCacheBehaviorsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedMethods", GoGetter: "AllowedMethods"},
			_jsii_.MemberProperty{JsiiProperty: "allowedMethodsInput", GoGetter: "AllowedMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cachedMethods", GoGetter: "CachedMethods"},
			_jsii_.MemberProperty{JsiiProperty: "cachedMethodsInput", GoGetter: "CachedMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyId", GoGetter: "CachePolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyIdInput", GoGetter: "CachePolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compress", GoGetter: "Compress"},
			_jsii_.MemberProperty{JsiiProperty: "compressInput", GoGetter: "CompressInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtl", GoGetter: "DefaultTtl"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtlInput", GoGetter: "DefaultTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLevelEncryptionId", GoGetter: "FieldLevelEncryptionId"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLevelEncryptionIdInput", GoGetter: "FieldLevelEncryptionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedValues", GoGetter: "ForwardedValues"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedValuesInput", GoGetter: "ForwardedValuesInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "functionAssociations", GoGetter: "FunctionAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "functionAssociationsInput", GoGetter: "FunctionAssociationsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionAssociations", GoGetter: "LambdaFunctionAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionAssociationsInput", GoGetter: "LambdaFunctionAssociationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtl", GoGetter: "MaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtlInput", GoGetter: "MaxTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "minTtl", GoGetter: "MinTtl"},
			_jsii_.MemberProperty{JsiiProperty: "minTtlInput", GoGetter: "MinTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyId", GoGetter: "OriginRequestPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyIdInput", GoGetter: "OriginRequestPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "pathPattern", GoGetter: "PathPattern"},
			_jsii_.MemberProperty{JsiiProperty: "pathPatternInput", GoGetter: "PathPatternInput"},
			_jsii_.MemberMethod{JsiiMethod: "putForwardedValues", GoMethod: "PutForwardedValues"},
			_jsii_.MemberMethod{JsiiMethod: "putFunctionAssociations", GoMethod: "PutFunctionAssociations"},
			_jsii_.MemberMethod{JsiiMethod: "putLambdaFunctionAssociations", GoMethod: "PutLambdaFunctionAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigArn", GoGetter: "RealtimeLogConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigArnInput", GoGetter: "RealtimeLogConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedMethods", GoMethod: "ResetAllowedMethods"},
			_jsii_.MemberMethod{JsiiMethod: "resetCachedMethods", GoMethod: "ResetCachedMethods"},
			_jsii_.MemberMethod{JsiiMethod: "resetCachePolicyId", GoMethod: "ResetCachePolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompress", GoMethod: "ResetCompress"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTtl", GoMethod: "ResetDefaultTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetFieldLevelEncryptionId", GoMethod: "ResetFieldLevelEncryptionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetForwardedValues", GoMethod: "ResetForwardedValues"},
			_jsii_.MemberMethod{JsiiMethod: "resetFunctionAssociations", GoMethod: "ResetFunctionAssociations"},
			_jsii_.MemberMethod{JsiiMethod: "resetLambdaFunctionAssociations", GoMethod: "ResetLambdaFunctionAssociations"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTtl", GoMethod: "ResetMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinTtl", GoMethod: "ResetMinTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginRequestPolicyId", GoMethod: "ResetOriginRequestPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRealtimeLogConfigArn", GoMethod: "ResetRealtimeLogConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersPolicyId", GoMethod: "ResetResponseHeadersPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSmoothStreaming", GoMethod: "ResetSmoothStreaming"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedKeyGroups", GoMethod: "ResetTrustedKeyGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedSigners", GoMethod: "ResetTrustedSigners"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyId", GoGetter: "ResponseHeadersPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyIdInput", GoGetter: "ResponseHeadersPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "smoothStreaming", GoGetter: "SmoothStreaming"},
			_jsii_.MemberProperty{JsiiProperty: "smoothStreamingInput", GoGetter: "SmoothStreamingInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetOriginId", GoGetter: "TargetOriginId"},
			_jsii_.MemberProperty{JsiiProperty: "targetOriginIdInput", GoGetter: "TargetOriginIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustedKeyGroups", GoGetter: "TrustedKeyGroups"},
			_jsii_.MemberProperty{JsiiProperty: "trustedKeyGroupsInput", GoGetter: "TrustedKeyGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSigners", GoGetter: "TrustedSigners"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSignersInput", GoGetter: "TrustedSignersInput"},
			_jsii_.MemberProperty{JsiiProperty: "viewerProtocolPolicy", GoGetter: "ViewerProtocolPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "viewerProtocolPolicyInput", GoGetter: "ViewerProtocolPolicyInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCustomErrorResponses",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCustomErrorResponses)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCustomErrorResponsesList",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCustomErrorResponsesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigCustomErrorResponsesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCustomErrorResponsesOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCustomErrorResponsesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "errorCachingMinTtl", GoGetter: "ErrorCachingMinTtl"},
			_jsii_.MemberProperty{JsiiProperty: "errorCachingMinTtlInput", GoGetter: "ErrorCachingMinTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorCode", GoGetter: "ErrorCode"},
			_jsii_.MemberProperty{JsiiProperty: "errorCodeInput", GoGetter: "ErrorCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorCachingMinTtl", GoMethod: "ResetErrorCachingMinTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseCode", GoMethod: "ResetResponseCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponsePagePath", GoMethod: "ResetResponsePagePath"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseCode", GoGetter: "ResponseCode"},
			_jsii_.MemberProperty{JsiiProperty: "responseCodeInput", GoGetter: "ResponseCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "responsePagePath", GoGetter: "ResponsePagePath"},
			_jsii_.MemberProperty{JsiiProperty: "responsePagePathInput", GoGetter: "ResponsePagePathInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigCustomErrorResponsesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCustomOrigin",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCustomOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCustomOriginOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigCustomOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dnsName", GoGetter: "DnsName"},
			_jsii_.MemberProperty{JsiiProperty: "dnsNameInput", GoGetter: "DnsNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpPort", GoGetter: "HttpPort"},
			_jsii_.MemberProperty{JsiiProperty: "httpPortInput", GoGetter: "HttpPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpsPort", GoGetter: "HttpsPort"},
			_jsii_.MemberProperty{JsiiProperty: "httpsPortInput", GoGetter: "HttpsPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "originProtocolPolicy", GoGetter: "OriginProtocolPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "originProtocolPolicyInput", GoGetter: "OriginProtocolPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "originSslProtocols", GoGetter: "OriginSslProtocols"},
			_jsii_.MemberProperty{JsiiProperty: "originSslProtocolsInput", GoGetter: "OriginSslProtocolsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpPort", GoMethod: "ResetHttpPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpsPort", GoMethod: "ResetHttpsPort"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehavior",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigDefaultCacheBehavior)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValues",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValues)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesCookies",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesCookies)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesCookiesOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesCookiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "forward", GoGetter: "Forward"},
			_jsii_.MemberProperty{JsiiProperty: "forwardInput", GoGetter: "ForwardInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhitelistedNames", GoMethod: "ResetWhitelistedNames"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistedNames", GoGetter: "WhitelistedNames"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistedNamesInput", GoGetter: "WhitelistedNamesInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesCookiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cookies", GoGetter: "Cookies"},
			_jsii_.MemberProperty{JsiiProperty: "cookiesInput", GoGetter: "CookiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putCookies", GoMethod: "PutCookies"},
			_jsii_.MemberProperty{JsiiProperty: "queryString", GoGetter: "QueryString"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCacheKeys", GoGetter: "QueryStringCacheKeys"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCacheKeysInput", GoGetter: "QueryStringCacheKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringInput", GoGetter: "QueryStringInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCookies", GoMethod: "ResetCookies"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryStringCacheKeys", GoMethod: "ResetQueryStringCacheKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociations",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociationsList",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociationsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociationsOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eventType", GoGetter: "EventType"},
			_jsii_.MemberProperty{JsiiProperty: "eventTypeInput", GoGetter: "EventTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionArnInput", GoGetter: "FunctionArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventType", GoMethod: "ResetEventType"},
			_jsii_.MemberMethod{JsiiMethod: "resetFunctionArn", GoMethod: "ResetFunctionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociations",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsList",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eventType", GoGetter: "EventType"},
			_jsii_.MemberProperty{JsiiProperty: "eventTypeInput", GoGetter: "EventTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "includeBody", GoGetter: "IncludeBody"},
			_jsii_.MemberProperty{JsiiProperty: "includeBodyInput", GoGetter: "IncludeBodyInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionArn", GoGetter: "LambdaFunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionArnInput", GoGetter: "LambdaFunctionArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventType", GoMethod: "ResetEventType"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeBody", GoMethod: "ResetIncludeBody"},
			_jsii_.MemberMethod{JsiiMethod: "resetLambdaFunctionArn", GoMethod: "ResetLambdaFunctionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedMethods", GoGetter: "AllowedMethods"},
			_jsii_.MemberProperty{JsiiProperty: "allowedMethodsInput", GoGetter: "AllowedMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cachedMethods", GoGetter: "CachedMethods"},
			_jsii_.MemberProperty{JsiiProperty: "cachedMethodsInput", GoGetter: "CachedMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyId", GoGetter: "CachePolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyIdInput", GoGetter: "CachePolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compress", GoGetter: "Compress"},
			_jsii_.MemberProperty{JsiiProperty: "compressInput", GoGetter: "CompressInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtl", GoGetter: "DefaultTtl"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtlInput", GoGetter: "DefaultTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLevelEncryptionId", GoGetter: "FieldLevelEncryptionId"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLevelEncryptionIdInput", GoGetter: "FieldLevelEncryptionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedValues", GoGetter: "ForwardedValues"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedValuesInput", GoGetter: "ForwardedValuesInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "functionAssociations", GoGetter: "FunctionAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "functionAssociationsInput", GoGetter: "FunctionAssociationsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionAssociations", GoGetter: "LambdaFunctionAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionAssociationsInput", GoGetter: "LambdaFunctionAssociationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtl", GoGetter: "MaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtlInput", GoGetter: "MaxTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "minTtl", GoGetter: "MinTtl"},
			_jsii_.MemberProperty{JsiiProperty: "minTtlInput", GoGetter: "MinTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyId", GoGetter: "OriginRequestPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyIdInput", GoGetter: "OriginRequestPolicyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "putForwardedValues", GoMethod: "PutForwardedValues"},
			_jsii_.MemberMethod{JsiiMethod: "putFunctionAssociations", GoMethod: "PutFunctionAssociations"},
			_jsii_.MemberMethod{JsiiMethod: "putLambdaFunctionAssociations", GoMethod: "PutLambdaFunctionAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigArn", GoGetter: "RealtimeLogConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigArnInput", GoGetter: "RealtimeLogConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedMethods", GoMethod: "ResetAllowedMethods"},
			_jsii_.MemberMethod{JsiiMethod: "resetCachedMethods", GoMethod: "ResetCachedMethods"},
			_jsii_.MemberMethod{JsiiMethod: "resetCachePolicyId", GoMethod: "ResetCachePolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompress", GoMethod: "ResetCompress"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTtl", GoMethod: "ResetDefaultTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetFieldLevelEncryptionId", GoMethod: "ResetFieldLevelEncryptionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetForwardedValues", GoMethod: "ResetForwardedValues"},
			_jsii_.MemberMethod{JsiiMethod: "resetFunctionAssociations", GoMethod: "ResetFunctionAssociations"},
			_jsii_.MemberMethod{JsiiMethod: "resetLambdaFunctionAssociations", GoMethod: "ResetLambdaFunctionAssociations"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTtl", GoMethod: "ResetMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinTtl", GoMethod: "ResetMinTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginRequestPolicyId", GoMethod: "ResetOriginRequestPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRealtimeLogConfigArn", GoMethod: "ResetRealtimeLogConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersPolicyId", GoMethod: "ResetResponseHeadersPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSmoothStreaming", GoMethod: "ResetSmoothStreaming"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedKeyGroups", GoMethod: "ResetTrustedKeyGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedSigners", GoMethod: "ResetTrustedSigners"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyId", GoGetter: "ResponseHeadersPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyIdInput", GoGetter: "ResponseHeadersPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "smoothStreaming", GoGetter: "SmoothStreaming"},
			_jsii_.MemberProperty{JsiiProperty: "smoothStreamingInput", GoGetter: "SmoothStreamingInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetOriginId", GoGetter: "TargetOriginId"},
			_jsii_.MemberProperty{JsiiProperty: "targetOriginIdInput", GoGetter: "TargetOriginIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustedKeyGroups", GoGetter: "TrustedKeyGroups"},
			_jsii_.MemberProperty{JsiiProperty: "trustedKeyGroupsInput", GoGetter: "TrustedKeyGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSigners", GoGetter: "TrustedSigners"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSignersInput", GoGetter: "TrustedSignersInput"},
			_jsii_.MemberProperty{JsiiProperty: "viewerProtocolPolicy", GoGetter: "ViewerProtocolPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "viewerProtocolPolicyInput", GoGetter: "ViewerProtocolPolicyInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigLogging",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigLogging)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigLoggingOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigLoggingOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "bucketInput", GoGetter: "BucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "includeCookies", GoGetter: "IncludeCookies"},
			_jsii_.MemberProperty{JsiiProperty: "includeCookiesInput", GoGetter: "IncludeCookiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeCookies", GoMethod: "ResetIncludeCookies"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigLoggingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroups",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroups)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsItems",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsItemsFailoverCriteria",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsItemsFailoverCriteria)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsItemsFailoverCriteriaOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsItemsFailoverCriteriaOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putStatusCodes", GoMethod: "PutStatusCodes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "statusCodes", GoGetter: "StatusCodes"},
			_jsii_.MemberProperty{JsiiProperty: "statusCodesInput", GoGetter: "StatusCodesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginGroupsItemsFailoverCriteriaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsItemsFailoverCriteriaStatusCodes",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsItemsFailoverCriteriaStatusCodes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsItemsFailoverCriteriaStatusCodesOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsItemsFailoverCriteriaStatusCodesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberProperty{JsiiProperty: "quantity", GoGetter: "Quantity"},
			_jsii_.MemberProperty{JsiiProperty: "quantityInput", GoGetter: "QuantityInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginGroupsItemsFailoverCriteriaStatusCodesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsItemsList",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsItemsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginGroupsItemsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsItemsMembers",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsItemsMembers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsItemsMembersItems",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsItemsMembersItems)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsItemsMembersItemsList",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsItemsMembersItemsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginGroupsItemsMembersItemsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsItemsMembersItemsOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsItemsMembersItemsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "originId", GoGetter: "OriginId"},
			_jsii_.MemberProperty{JsiiProperty: "originIdInput", GoGetter: "OriginIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginGroupsItemsMembersItemsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsItemsMembersOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsItemsMembersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putItems", GoMethod: "PutItems"},
			_jsii_.MemberProperty{JsiiProperty: "quantity", GoGetter: "Quantity"},
			_jsii_.MemberProperty{JsiiProperty: "quantityInput", GoGetter: "QuantityInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginGroupsItemsMembersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsItemsOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsItemsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "failoverCriteria", GoGetter: "FailoverCriteria"},
			_jsii_.MemberProperty{JsiiProperty: "failoverCriteriaInput", GoGetter: "FailoverCriteriaInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "members", GoGetter: "Members"},
			_jsii_.MemberProperty{JsiiProperty: "membersInput", GoGetter: "MembersInput"},
			_jsii_.MemberMethod{JsiiMethod: "putFailoverCriteria", GoMethod: "PutFailoverCriteria"},
			_jsii_.MemberMethod{JsiiMethod: "putMembers", GoMethod: "PutMembers"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginGroupsItemsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginGroupsOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginGroupsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putItems", GoMethod: "PutItems"},
			_jsii_.MemberProperty{JsiiProperty: "quantity", GoGetter: "Quantity"},
			_jsii_.MemberProperty{JsiiProperty: "quantityInput", GoGetter: "QuantityInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetItems", GoMethod: "ResetItems"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginGroupsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOrigins",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOrigins)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsCustomOriginConfig",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginsCustomOriginConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsCustomOriginConfigOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginsCustomOriginConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpPort", GoGetter: "HttpPort"},
			_jsii_.MemberProperty{JsiiProperty: "httpPortInput", GoGetter: "HttpPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpsPort", GoGetter: "HttpsPort"},
			_jsii_.MemberProperty{JsiiProperty: "httpsPortInput", GoGetter: "HttpsPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "originKeepaliveTimeout", GoGetter: "OriginKeepaliveTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "originKeepaliveTimeoutInput", GoGetter: "OriginKeepaliveTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "originProtocolPolicy", GoGetter: "OriginProtocolPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "originProtocolPolicyInput", GoGetter: "OriginProtocolPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "originReadTimeout", GoGetter: "OriginReadTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "originReadTimeoutInput", GoGetter: "OriginReadTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "originSslProtocols", GoGetter: "OriginSslProtocols"},
			_jsii_.MemberProperty{JsiiProperty: "originSslProtocolsInput", GoGetter: "OriginSslProtocolsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpPort", GoMethod: "ResetHttpPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpsPort", GoMethod: "ResetHttpsPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginKeepaliveTimeout", GoMethod: "ResetOriginKeepaliveTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginReadTimeout", GoMethod: "ResetOriginReadTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginSslProtocols", GoMethod: "ResetOriginSslProtocols"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginsCustomOriginConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsList",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsOriginCustomHeaders",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginsOriginCustomHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsOriginCustomHeadersList",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginsOriginCustomHeadersList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOriginCustomHeadersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsOriginCustomHeadersOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginsOriginCustomHeadersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOriginCustomHeadersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsOriginShield",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginsOriginShield)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsOriginShieldOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginsOriginShieldOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "originShieldRegion", GoGetter: "OriginShieldRegion"},
			_jsii_.MemberProperty{JsiiProperty: "originShieldRegionInput", GoGetter: "OriginShieldRegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginShieldRegion", GoMethod: "ResetOriginShieldRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOriginShieldOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectionAttempts", GoGetter: "ConnectionAttempts"},
			_jsii_.MemberProperty{JsiiProperty: "connectionAttemptsInput", GoGetter: "ConnectionAttemptsInput"},
			_jsii_.MemberProperty{JsiiProperty: "connectionTimeout", GoGetter: "ConnectionTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "connectionTimeoutInput", GoGetter: "ConnectionTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customOriginConfig", GoGetter: "CustomOriginConfig"},
			_jsii_.MemberProperty{JsiiProperty: "customOriginConfigInput", GoGetter: "CustomOriginConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameInput", GoGetter: "DomainNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessControlId", GoGetter: "OriginAccessControlId"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessControlIdInput", GoGetter: "OriginAccessControlIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "originCustomHeaders", GoGetter: "OriginCustomHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "originCustomHeadersInput", GoGetter: "OriginCustomHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "originPath", GoGetter: "OriginPath"},
			_jsii_.MemberProperty{JsiiProperty: "originPathInput", GoGetter: "OriginPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "originShield", GoGetter: "OriginShield"},
			_jsii_.MemberProperty{JsiiProperty: "originShieldInput", GoGetter: "OriginShieldInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomOriginConfig", GoMethod: "PutCustomOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putOriginCustomHeaders", GoMethod: "PutOriginCustomHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "putOriginShield", GoMethod: "PutOriginShield"},
			_jsii_.MemberMethod{JsiiMethod: "putS3OriginConfig", GoMethod: "PutS3OriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionAttempts", GoMethod: "ResetConnectionAttempts"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionTimeout", GoMethod: "ResetConnectionTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomOriginConfig", GoMethod: "ResetCustomOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginAccessControlId", GoMethod: "ResetOriginAccessControlId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginCustomHeaders", GoMethod: "ResetOriginCustomHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginPath", GoMethod: "ResetOriginPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginShield", GoMethod: "ResetOriginShield"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3OriginConfig", GoMethod: "ResetS3OriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3OriginConfig", GoGetter: "S3OriginConfig"},
			_jsii_.MemberProperty{JsiiProperty: "s3OriginConfigInput", GoGetter: "S3OriginConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsS3OriginConfig",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginsS3OriginConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsS3OriginConfigOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOriginsS3OriginConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessIdentity", GoGetter: "OriginAccessIdentity"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessIdentityInput", GoGetter: "OriginAccessIdentityInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginAccessIdentity", GoMethod: "ResetOriginAccessIdentity"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginsS3OriginConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aliases", GoGetter: "Aliases"},
			_jsii_.MemberProperty{JsiiProperty: "aliasesInput", GoGetter: "AliasesInput"},
			_jsii_.MemberProperty{JsiiProperty: "cacheBehaviors", GoGetter: "CacheBehaviors"},
			_jsii_.MemberProperty{JsiiProperty: "cacheBehaviorsInput", GoGetter: "CacheBehaviorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cnames", GoGetter: "Cnames"},
			_jsii_.MemberProperty{JsiiProperty: "cnamesInput", GoGetter: "CnamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "continuousDeploymentPolicyId", GoGetter: "ContinuousDeploymentPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "continuousDeploymentPolicyIdInput", GoGetter: "ContinuousDeploymentPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customErrorResponses", GoGetter: "CustomErrorResponses"},
			_jsii_.MemberProperty{JsiiProperty: "customErrorResponsesInput", GoGetter: "CustomErrorResponsesInput"},
			_jsii_.MemberProperty{JsiiProperty: "customOrigin", GoGetter: "CustomOrigin"},
			_jsii_.MemberProperty{JsiiProperty: "customOriginInput", GoGetter: "CustomOriginInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCacheBehavior", GoGetter: "DefaultCacheBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCacheBehaviorInput", GoGetter: "DefaultCacheBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRootObject", GoGetter: "DefaultRootObject"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRootObjectInput", GoGetter: "DefaultRootObjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpVersion", GoGetter: "HttpVersion"},
			_jsii_.MemberProperty{JsiiProperty: "httpVersionInput", GoGetter: "HttpVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Enabled", GoGetter: "Ipv6Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6EnabledInput", GoGetter: "Ipv6EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "loggingInput", GoGetter: "LoggingInput"},
			_jsii_.MemberProperty{JsiiProperty: "originGroups", GoGetter: "OriginGroups"},
			_jsii_.MemberProperty{JsiiProperty: "originGroupsInput", GoGetter: "OriginGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "origins", GoGetter: "Origins"},
			_jsii_.MemberProperty{JsiiProperty: "originsInput", GoGetter: "OriginsInput"},
			_jsii_.MemberProperty{JsiiProperty: "priceClass", GoGetter: "PriceClass"},
			_jsii_.MemberProperty{JsiiProperty: "priceClassInput", GoGetter: "PriceClassInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCacheBehaviors", GoMethod: "PutCacheBehaviors"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomErrorResponses", GoMethod: "PutCustomErrorResponses"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomOrigin", GoMethod: "PutCustomOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultCacheBehavior", GoMethod: "PutDefaultCacheBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "putLogging", GoMethod: "PutLogging"},
			_jsii_.MemberMethod{JsiiMethod: "putOriginGroups", GoMethod: "PutOriginGroups"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigins", GoMethod: "PutOrigins"},
			_jsii_.MemberMethod{JsiiMethod: "putRestrictions", GoMethod: "PutRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "putS3Origin", GoMethod: "PutS3Origin"},
			_jsii_.MemberMethod{JsiiMethod: "putViewerCertificate", GoMethod: "PutViewerCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetAliases", GoMethod: "ResetAliases"},
			_jsii_.MemberMethod{JsiiMethod: "resetCacheBehaviors", GoMethod: "ResetCacheBehaviors"},
			_jsii_.MemberMethod{JsiiMethod: "resetCnames", GoMethod: "ResetCnames"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetContinuousDeploymentPolicyId", GoMethod: "ResetContinuousDeploymentPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomErrorResponses", GoMethod: "ResetCustomErrorResponses"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomOrigin", GoMethod: "ResetCustomOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultRootObject", GoMethod: "ResetDefaultRootObject"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpVersion", GoMethod: "ResetHttpVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6Enabled", GoMethod: "ResetIpv6Enabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogging", GoMethod: "ResetLogging"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginGroups", GoMethod: "ResetOriginGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigins", GoMethod: "ResetOrigins"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriceClass", GoMethod: "ResetPriceClass"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestrictions", GoMethod: "ResetRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3Origin", GoMethod: "ResetS3Origin"},
			_jsii_.MemberMethod{JsiiMethod: "resetStaging", GoMethod: "ResetStaging"},
			_jsii_.MemberMethod{JsiiMethod: "resetViewerCertificate", GoMethod: "ResetViewerCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebAclId", GoMethod: "ResetWebAclId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictions", GoGetter: "Restrictions"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionsInput", GoGetter: "RestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3Origin", GoGetter: "S3Origin"},
			_jsii_.MemberProperty{JsiiProperty: "s3OriginInput", GoGetter: "S3OriginInput"},
			_jsii_.MemberProperty{JsiiProperty: "staging", GoGetter: "Staging"},
			_jsii_.MemberProperty{JsiiProperty: "stagingInput", GoGetter: "StagingInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "viewerCertificate", GoGetter: "ViewerCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "viewerCertificateInput", GoGetter: "ViewerCertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "webAclId", GoGetter: "WebAclId"},
			_jsii_.MemberProperty{JsiiProperty: "webAclIdInput", GoGetter: "WebAclIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigRestrictions",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigRestrictions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigRestrictionsGeoRestriction",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigRestrictionsGeoRestriction)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigRestrictionsGeoRestrictionOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigRestrictionsGeoRestrictionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "locations", GoGetter: "Locations"},
			_jsii_.MemberProperty{JsiiProperty: "locationsInput", GoGetter: "LocationsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocations", GoMethod: "ResetLocations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionType", GoGetter: "RestrictionType"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionTypeInput", GoGetter: "RestrictionTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigRestrictionsGeoRestrictionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigRestrictionsOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigRestrictionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "geoRestriction", GoGetter: "GeoRestriction"},
			_jsii_.MemberProperty{JsiiProperty: "geoRestrictionInput", GoGetter: "GeoRestrictionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putGeoRestriction", GoMethod: "PutGeoRestriction"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigRestrictionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigS3Origin",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigS3Origin)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigS3OriginOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigS3OriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dnsName", GoGetter: "DnsName"},
			_jsii_.MemberProperty{JsiiProperty: "dnsNameInput", GoGetter: "DnsNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessIdentity", GoGetter: "OriginAccessIdentity"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessIdentityInput", GoGetter: "OriginAccessIdentityInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginAccessIdentity", GoMethod: "ResetOriginAccessIdentity"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigS3OriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigViewerCertificate",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigViewerCertificate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigViewerCertificateOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDistributionConfigViewerCertificateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acmCertificateArn", GoGetter: "AcmCertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "acmCertificateArnInput", GoGetter: "AcmCertificateArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudfrontDefaultCertificate", GoGetter: "CloudfrontDefaultCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "cloudfrontDefaultCertificateInput", GoGetter: "CloudfrontDefaultCertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iamCertificateId", GoGetter: "IamCertificateId"},
			_jsii_.MemberProperty{JsiiProperty: "iamCertificateIdInput", GoGetter: "IamCertificateIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "minimumProtocolVersion", GoGetter: "MinimumProtocolVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minimumProtocolVersionInput", GoGetter: "MinimumProtocolVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcmCertificateArn", GoMethod: "ResetAcmCertificateArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudfrontDefaultCertificate", GoMethod: "ResetCloudfrontDefaultCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetIamCertificateId", GoMethod: "ResetIamCertificateId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumProtocolVersion", GoMethod: "ResetMinimumProtocolVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetSslSupportMethod", GoMethod: "ResetSslSupportMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sslSupportMethod", GoGetter: "SslSupportMethod"},
			_jsii_.MemberProperty{JsiiProperty: "sslSupportMethodInput", GoGetter: "SslSupportMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDistributionConfigViewerCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudfrontDistribution.CloudfrontDistributionTags",
		reflect.TypeOf((*CloudfrontDistributionTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionTagsList",
		reflect.TypeOf((*CloudfrontDistributionTagsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.cloudfrontDistribution.CloudfrontDistributionTagsOutputReference",
		reflect.TypeOf((*CloudfrontDistributionTagsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
