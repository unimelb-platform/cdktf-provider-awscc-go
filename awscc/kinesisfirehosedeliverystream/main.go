package kinesisfirehosedeliverystream

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStream",
		reflect.TypeOf((*KinesisfirehoseDeliveryStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "amazonOpenSearchServerlessDestinationConfiguration", GoGetter: "AmazonOpenSearchServerlessDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "amazonOpenSearchServerlessDestinationConfigurationInput", GoGetter: "AmazonOpenSearchServerlessDestinationConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "amazonopensearchserviceDestinationConfiguration", GoGetter: "AmazonopensearchserviceDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "amazonopensearchserviceDestinationConfigurationInput", GoGetter: "AmazonopensearchserviceDestinationConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamEncryptionConfigurationInput", GoGetter: "DeliveryStreamEncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamEncryptionConfigurationInputInput", GoGetter: "DeliveryStreamEncryptionConfigurationInputInput"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamName", GoGetter: "DeliveryStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamNameInput", GoGetter: "DeliveryStreamNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamType", GoGetter: "DeliveryStreamType"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamTypeInput", GoGetter: "DeliveryStreamTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchDestinationConfiguration", GoGetter: "ElasticsearchDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchDestinationConfigurationInput", GoGetter: "ElasticsearchDestinationConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "extendedS3DestinationConfiguration", GoGetter: "ExtendedS3DestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "extendedS3DestinationConfigurationInput", GoGetter: "ExtendedS3DestinationConfigurationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "httpEndpointDestinationConfiguration", GoGetter: "HttpEndpointDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "httpEndpointDestinationConfigurationInput", GoGetter: "HttpEndpointDestinationConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisStreamSourceConfiguration", GoGetter: "KinesisStreamSourceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisStreamSourceConfigurationInput", GoGetter: "KinesisStreamSourceConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "mskSourceConfiguration", GoGetter: "MskSourceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "mskSourceConfigurationInput", GoGetter: "MskSourceConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAmazonOpenSearchServerlessDestinationConfiguration", GoMethod: "PutAmazonOpenSearchServerlessDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putAmazonopensearchserviceDestinationConfiguration", GoMethod: "PutAmazonopensearchserviceDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putDeliveryStreamEncryptionConfigurationInput", GoMethod: "PutDeliveryStreamEncryptionConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putElasticsearchDestinationConfiguration", GoMethod: "PutElasticsearchDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putExtendedS3DestinationConfiguration", GoMethod: "PutExtendedS3DestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putHttpEndpointDestinationConfiguration", GoMethod: "PutHttpEndpointDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putKinesisStreamSourceConfiguration", GoMethod: "PutKinesisStreamSourceConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putMskSourceConfiguration", GoMethod: "PutMskSourceConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putRedshiftDestinationConfiguration", GoMethod: "PutRedshiftDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putS3DestinationConfiguration", GoMethod: "PutS3DestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putSnowflakeDestinationConfiguration", GoMethod: "PutSnowflakeDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putSplunkDestinationConfiguration", GoMethod: "PutSplunkDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putTags", GoMethod: "PutTags"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftDestinationConfiguration", GoGetter: "RedshiftDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftDestinationConfigurationInput", GoGetter: "RedshiftDestinationConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAmazonOpenSearchServerlessDestinationConfiguration", GoMethod: "ResetAmazonOpenSearchServerlessDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetAmazonopensearchserviceDestinationConfiguration", GoMethod: "ResetAmazonopensearchserviceDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeliveryStreamEncryptionConfigurationInput", GoMethod: "ResetDeliveryStreamEncryptionConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeliveryStreamName", GoMethod: "ResetDeliveryStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeliveryStreamType", GoMethod: "ResetDeliveryStreamType"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticsearchDestinationConfiguration", GoMethod: "ResetElasticsearchDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtendedS3DestinationConfiguration", GoMethod: "ResetExtendedS3DestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpEndpointDestinationConfiguration", GoMethod: "ResetHttpEndpointDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetKinesisStreamSourceConfiguration", GoMethod: "ResetKinesisStreamSourceConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetMskSourceConfiguration", GoMethod: "ResetMskSourceConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedshiftDestinationConfiguration", GoMethod: "ResetRedshiftDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3DestinationConfiguration", GoMethod: "ResetS3DestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnowflakeDestinationConfiguration", GoMethod: "ResetSnowflakeDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetSplunkDestinationConfiguration", GoMethod: "ResetSplunkDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberProperty{JsiiProperty: "s3DestinationConfiguration", GoGetter: "S3DestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "s3DestinationConfigurationInput", GoGetter: "S3DestinationConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeDestinationConfiguration", GoGetter: "SnowflakeDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeDestinationConfigurationInput", GoGetter: "SnowflakeDestinationConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "splunkDestinationConfiguration", GoGetter: "SplunkDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "splunkDestinationConfigurationInput", GoGetter: "SplunkDestinationConfigurationInput"},
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
			j := jsiiProxy_KinesisfirehoseDeliveryStream{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "collectionEndpoint", GoGetter: "CollectionEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "collectionEndpointInput", GoGetter: "CollectionEndpointInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "indexName", GoGetter: "IndexName"},
			_jsii_.MemberProperty{JsiiProperty: "indexNameInput", GoGetter: "IndexNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfiguration", GoGetter: "ProcessingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfigurationInput", GoGetter: "ProcessingConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessingConfiguration", GoMethod: "PutProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryOptions", GoMethod: "PutRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putS3Configuration", GoMethod: "PutS3Configuration"},
			_jsii_.MemberMethod{JsiiMethod: "putVpcConfiguration", GoMethod: "PutVpcConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCollectionEndpoint", GoMethod: "ResetCollectionEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessingConfiguration", GoMethod: "ResetProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryOptions", GoMethod: "ResetRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3BackupMode", GoMethod: "ResetS3BackupMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcConfiguration", GoMethod: "ResetVpcConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptions", GoGetter: "RetryOptions"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptionsInput", GoGetter: "RetryOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupMode", GoGetter: "S3BackupMode"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupModeInput", GoGetter: "S3BackupModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3Configuration", GoGetter: "S3Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "s3ConfigurationInput", GoGetter: "S3ConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfiguration", GoGetter: "VpcConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfigurationInput", GoGetter: "VpcConfigurationInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "processors", GoGetter: "Processors"},
			_jsii_.MemberProperty{JsiiProperty: "processorsInput", GoGetter: "ProcessorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessors", GoMethod: "PutProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessors", GoMethod: "ResetProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessors",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessors)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInput", GoGetter: "ParametersInput"},
			_jsii_.MemberMethod{JsiiMethod: "putParameters", GoMethod: "PutParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameters", GoMethod: "ResetParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsParameters",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsParametersList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsParametersList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsParametersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
			_jsii_.MemberProperty{JsiiProperty: "parameterNameInput", GoGetter: "ParameterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValue", GoGetter: "ParameterValue"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValueInput", GoGetter: "ParameterValueInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationRetryOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationRetryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationRetryOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationRetryOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSeconds", GoGetter: "DurationInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSecondsInput", GoGetter: "DurationInSecondsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDurationInSeconds", GoMethod: "ResetDurationInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationRetryOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3Configuration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3Configuration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationEncryptionConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationEncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArn", GoGetter: "AwskmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArnInput", GoGetter: "AwskmsKeyArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfig", GoGetter: "KmsEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfigInput", GoGetter: "KmsEncryptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfig", GoGetter: "NoEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfigInput", GoGetter: "NoEncryptionConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKmsEncryptionConfig", GoMethod: "PutKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsEncryptionConfig", GoMethod: "ResetKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoEncryptionConfig", GoMethod: "ResetNoEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArnInput", GoGetter: "BucketArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormatInput", GoGetter: "CompressionFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfigurationInput", GoGetter: "EncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefixInput", GoGetter: "ErrorOutputPrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionConfiguration", GoMethod: "PutEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompressionFormat", GoMethod: "ResetCompressionFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionConfiguration", GoMethod: "ResetEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorOutputPrefix", GoMethod: "ResetErrorOutputPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3ConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationVpcConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationVpcConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationVpcConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationVpcConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdsInput", GoGetter: "SubnetIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationVpcConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationDocumentIdOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationDocumentIdOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationDocumentIdOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationDocumentIdOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultDocumentIdFormat", GoGetter: "DefaultDocumentIdFormat"},
			_jsii_.MemberProperty{JsiiProperty: "defaultDocumentIdFormatInput", GoGetter: "DefaultDocumentIdFormatInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationDocumentIdOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpointInput", GoGetter: "ClusterEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "documentIdOptions", GoGetter: "DocumentIdOptions"},
			_jsii_.MemberProperty{JsiiProperty: "documentIdOptionsInput", GoGetter: "DocumentIdOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "domainArn", GoGetter: "DomainArn"},
			_jsii_.MemberProperty{JsiiProperty: "domainArnInput", GoGetter: "DomainArnInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "indexName", GoGetter: "IndexName"},
			_jsii_.MemberProperty{JsiiProperty: "indexNameInput", GoGetter: "IndexNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "indexRotationPeriod", GoGetter: "IndexRotationPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "indexRotationPeriodInput", GoGetter: "IndexRotationPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfiguration", GoGetter: "ProcessingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfigurationInput", GoGetter: "ProcessingConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putDocumentIdOptions", GoMethod: "PutDocumentIdOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessingConfiguration", GoMethod: "PutProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryOptions", GoMethod: "PutRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putS3Configuration", GoMethod: "PutS3Configuration"},
			_jsii_.MemberMethod{JsiiMethod: "putVpcConfiguration", GoMethod: "PutVpcConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetClusterEndpoint", GoMethod: "ResetClusterEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetDocumentIdOptions", GoMethod: "ResetDocumentIdOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomainArn", GoMethod: "ResetDomainArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetIndexRotationPeriod", GoMethod: "ResetIndexRotationPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessingConfiguration", GoMethod: "ResetProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryOptions", GoMethod: "ResetRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3BackupMode", GoMethod: "ResetS3BackupMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetTypeName", GoMethod: "ResetTypeName"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcConfiguration", GoMethod: "ResetVpcConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptions", GoGetter: "RetryOptions"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptionsInput", GoGetter: "RetryOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupMode", GoGetter: "S3BackupMode"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupModeInput", GoGetter: "S3BackupModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3Configuration", GoGetter: "S3Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "s3ConfigurationInput", GoGetter: "S3ConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "typeName", GoGetter: "TypeName"},
			_jsii_.MemberProperty{JsiiProperty: "typeNameInput", GoGetter: "TypeNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfiguration", GoGetter: "VpcConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfigurationInput", GoGetter: "VpcConfigurationInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "processors", GoGetter: "Processors"},
			_jsii_.MemberProperty{JsiiProperty: "processorsInput", GoGetter: "ProcessorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessors", GoMethod: "PutProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessors", GoMethod: "ResetProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessors",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessors)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInput", GoGetter: "ParametersInput"},
			_jsii_.MemberMethod{JsiiMethod: "putParameters", GoMethod: "PutParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameters", GoMethod: "ResetParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsParameters",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsParametersList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsParametersList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsParametersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
			_jsii_.MemberProperty{JsiiProperty: "parameterNameInput", GoGetter: "ParameterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValue", GoGetter: "ParameterValue"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValueInput", GoGetter: "ParameterValueInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationRetryOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationRetryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationRetryOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationRetryOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSeconds", GoGetter: "DurationInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSecondsInput", GoGetter: "DurationInSecondsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDurationInSeconds", GoMethod: "ResetDurationInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationRetryOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3Configuration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3Configuration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArn", GoGetter: "AwskmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArnInput", GoGetter: "AwskmsKeyArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfig", GoGetter: "KmsEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfigInput", GoGetter: "KmsEncryptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfig", GoGetter: "NoEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfigInput", GoGetter: "NoEncryptionConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKmsEncryptionConfig", GoMethod: "PutKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsEncryptionConfig", GoMethod: "ResetKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoEncryptionConfig", GoMethod: "ResetNoEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArnInput", GoGetter: "BucketArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormatInput", GoGetter: "CompressionFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfigurationInput", GoGetter: "EncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefixInput", GoGetter: "ErrorOutputPrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionConfiguration", GoMethod: "PutEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompressionFormat", GoMethod: "ResetCompressionFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionConfiguration", GoMethod: "ResetEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorOutputPrefix", GoMethod: "ResetErrorOutputPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationVpcConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationVpcConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationVpcConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationVpcConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdsInput", GoGetter: "SubnetIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationVpcConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamConfig",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInput",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInput)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInputOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInputOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "keyArn", GoGetter: "KeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "keyArnInput", GoGetter: "KeyArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyType", GoGetter: "KeyType"},
			_jsii_.MemberProperty{JsiiProperty: "keyTypeInput", GoGetter: "KeyTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyArn", GoMethod: "ResetKeyArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInputOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationDocumentIdOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationDocumentIdOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationDocumentIdOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationDocumentIdOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultDocumentIdFormat", GoGetter: "DefaultDocumentIdFormat"},
			_jsii_.MemberProperty{JsiiProperty: "defaultDocumentIdFormatInput", GoGetter: "DefaultDocumentIdFormatInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationDocumentIdOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpointInput", GoGetter: "ClusterEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "documentIdOptions", GoGetter: "DocumentIdOptions"},
			_jsii_.MemberProperty{JsiiProperty: "documentIdOptionsInput", GoGetter: "DocumentIdOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "domainArn", GoGetter: "DomainArn"},
			_jsii_.MemberProperty{JsiiProperty: "domainArnInput", GoGetter: "DomainArnInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "indexName", GoGetter: "IndexName"},
			_jsii_.MemberProperty{JsiiProperty: "indexNameInput", GoGetter: "IndexNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "indexRotationPeriod", GoGetter: "IndexRotationPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "indexRotationPeriodInput", GoGetter: "IndexRotationPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfiguration", GoGetter: "ProcessingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfigurationInput", GoGetter: "ProcessingConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putDocumentIdOptions", GoMethod: "PutDocumentIdOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessingConfiguration", GoMethod: "PutProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryOptions", GoMethod: "PutRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putS3Configuration", GoMethod: "PutS3Configuration"},
			_jsii_.MemberMethod{JsiiMethod: "putVpcConfiguration", GoMethod: "PutVpcConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetClusterEndpoint", GoMethod: "ResetClusterEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetDocumentIdOptions", GoMethod: "ResetDocumentIdOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomainArn", GoMethod: "ResetDomainArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetIndexRotationPeriod", GoMethod: "ResetIndexRotationPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessingConfiguration", GoMethod: "ResetProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryOptions", GoMethod: "ResetRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3BackupMode", GoMethod: "ResetS3BackupMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetTypeName", GoMethod: "ResetTypeName"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcConfiguration", GoMethod: "ResetVpcConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptions", GoGetter: "RetryOptions"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptionsInput", GoGetter: "RetryOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupMode", GoGetter: "S3BackupMode"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupModeInput", GoGetter: "S3BackupModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3Configuration", GoGetter: "S3Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "s3ConfigurationInput", GoGetter: "S3ConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "typeName", GoGetter: "TypeName"},
			_jsii_.MemberProperty{JsiiProperty: "typeNameInput", GoGetter: "TypeNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfiguration", GoGetter: "VpcConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfigurationInput", GoGetter: "VpcConfigurationInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "processors", GoGetter: "Processors"},
			_jsii_.MemberProperty{JsiiProperty: "processorsInput", GoGetter: "ProcessorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessors", GoMethod: "PutProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessors", GoMethod: "ResetProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessors",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessors)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInput", GoGetter: "ParametersInput"},
			_jsii_.MemberMethod{JsiiMethod: "putParameters", GoMethod: "PutParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameters", GoMethod: "ResetParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsParameters",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsParametersList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsParametersList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsParametersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
			_jsii_.MemberProperty{JsiiProperty: "parameterNameInput", GoGetter: "ParameterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValue", GoGetter: "ParameterValue"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValueInput", GoGetter: "ParameterValueInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationRetryOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationRetryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationRetryOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationRetryOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSeconds", GoGetter: "DurationInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSecondsInput", GoGetter: "DurationInSecondsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDurationInSeconds", GoMethod: "ResetDurationInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationRetryOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3Configuration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3Configuration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationEncryptionConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationEncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArn", GoGetter: "AwskmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArnInput", GoGetter: "AwskmsKeyArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfig", GoGetter: "KmsEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfigInput", GoGetter: "KmsEncryptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfig", GoGetter: "NoEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfigInput", GoGetter: "NoEncryptionConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKmsEncryptionConfig", GoMethod: "PutKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsEncryptionConfig", GoMethod: "ResetKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoEncryptionConfig", GoMethod: "ResetNoEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArnInput", GoGetter: "BucketArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormatInput", GoGetter: "CompressionFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfigurationInput", GoGetter: "EncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefixInput", GoGetter: "ErrorOutputPrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionConfiguration", GoMethod: "PutEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompressionFormat", GoMethod: "ResetCompressionFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionConfiguration", GoMethod: "ResetEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorOutputPrefix", GoMethod: "ResetErrorOutputPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3ConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationVpcConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationVpcConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationVpcConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationVpcConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdsInput", GoGetter: "SubnetIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationVpcConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJsonSerDe",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJsonSerDe)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJsonSerDeOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJsonSerDeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetTimestampFormats", GoMethod: "ResetTimestampFormats"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timestampFormats", GoGetter: "TimestampFormats"},
			_jsii_.MemberProperty{JsiiProperty: "timestampFormatsInput", GoGetter: "TimestampFormatsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJsonSerDeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDe",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDe)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDeOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "caseInsensitive", GoGetter: "CaseInsensitive"},
			_jsii_.MemberProperty{JsiiProperty: "caseInsensitiveInput", GoGetter: "CaseInsensitiveInput"},
			_jsii_.MemberProperty{JsiiProperty: "columnToJsonKeyMappings", GoGetter: "ColumnToJsonKeyMappings"},
			_jsii_.MemberProperty{JsiiProperty: "columnToJsonKeyMappingsInput", GoGetter: "ColumnToJsonKeyMappingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "convertDotsInJsonKeysToUnderscores", GoGetter: "ConvertDotsInJsonKeysToUnderscores"},
			_jsii_.MemberProperty{JsiiProperty: "convertDotsInJsonKeysToUnderscoresInput", GoGetter: "ConvertDotsInJsonKeysToUnderscoresInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCaseInsensitive", GoMethod: "ResetCaseInsensitive"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumnToJsonKeyMappings", GoMethod: "ResetColumnToJsonKeyMappings"},
			_jsii_.MemberMethod{JsiiMethod: "resetConvertDotsInJsonKeysToUnderscores", GoMethod: "ResetConvertDotsInJsonKeysToUnderscores"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "hiveJsonSerDe", GoGetter: "HiveJsonSerDe"},
			_jsii_.MemberProperty{JsiiProperty: "hiveJsonSerDeInput", GoGetter: "HiveJsonSerDeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "openXJsonSerDe", GoGetter: "OpenXJsonSerDe"},
			_jsii_.MemberProperty{JsiiProperty: "openXJsonSerDeInput", GoGetter: "OpenXJsonSerDeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putHiveJsonSerDe", GoMethod: "PutHiveJsonSerDe"},
			_jsii_.MemberMethod{JsiiMethod: "putOpenXJsonSerDe", GoMethod: "PutOpenXJsonSerDe"},
			_jsii_.MemberMethod{JsiiMethod: "resetHiveJsonSerDe", GoMethod: "ResetHiveJsonSerDe"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenXJsonSerDe", GoMethod: "ResetOpenXJsonSerDe"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deserializer", GoGetter: "Deserializer"},
			_jsii_.MemberProperty{JsiiProperty: "deserializerInput", GoGetter: "DeserializerInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putDeserializer", GoMethod: "PutDeserializer"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeserializer", GoMethod: "ResetDeserializer"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "putSerializer", GoMethod: "PutSerializer"},
			_jsii_.MemberMethod{JsiiMethod: "resetSerializer", GoMethod: "ResetSerializer"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serializer", GoGetter: "Serializer"},
			_jsii_.MemberProperty{JsiiProperty: "serializerInput", GoGetter: "SerializerInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDeOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "blockSizeBytes", GoGetter: "BlockSizeBytes"},
			_jsii_.MemberProperty{JsiiProperty: "blockSizeBytesInput", GoGetter: "BlockSizeBytesInput"},
			_jsii_.MemberProperty{JsiiProperty: "bloomFilterColumns", GoGetter: "BloomFilterColumns"},
			_jsii_.MemberProperty{JsiiProperty: "bloomFilterColumnsInput", GoGetter: "BloomFilterColumnsInput"},
			_jsii_.MemberProperty{JsiiProperty: "bloomFilterFalsePositiveProbability", GoGetter: "BloomFilterFalsePositiveProbability"},
			_jsii_.MemberProperty{JsiiProperty: "bloomFilterFalsePositiveProbabilityInput", GoGetter: "BloomFilterFalsePositiveProbabilityInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberProperty{JsiiProperty: "compressionInput", GoGetter: "CompressionInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dictionaryKeyThreshold", GoGetter: "DictionaryKeyThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "dictionaryKeyThresholdInput", GoGetter: "DictionaryKeyThresholdInput"},
			_jsii_.MemberProperty{JsiiProperty: "enablePadding", GoGetter: "EnablePadding"},
			_jsii_.MemberProperty{JsiiProperty: "enablePaddingInput", GoGetter: "EnablePaddingInput"},
			_jsii_.MemberProperty{JsiiProperty: "formatVersion", GoGetter: "FormatVersion"},
			_jsii_.MemberProperty{JsiiProperty: "formatVersionInput", GoGetter: "FormatVersionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "paddingTolerance", GoGetter: "PaddingTolerance"},
			_jsii_.MemberProperty{JsiiProperty: "paddingToleranceInput", GoGetter: "PaddingToleranceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlockSizeBytes", GoMethod: "ResetBlockSizeBytes"},
			_jsii_.MemberMethod{JsiiMethod: "resetBloomFilterColumns", GoMethod: "ResetBloomFilterColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetBloomFilterFalsePositiveProbability", GoMethod: "ResetBloomFilterFalsePositiveProbability"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompression", GoMethod: "ResetCompression"},
			_jsii_.MemberMethod{JsiiMethod: "resetDictionaryKeyThreshold", GoMethod: "ResetDictionaryKeyThreshold"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnablePadding", GoMethod: "ResetEnablePadding"},
			_jsii_.MemberMethod{JsiiMethod: "resetFormatVersion", GoMethod: "ResetFormatVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetPaddingTolerance", GoMethod: "ResetPaddingTolerance"},
			_jsii_.MemberMethod{JsiiMethod: "resetRowIndexStride", GoMethod: "ResetRowIndexStride"},
			_jsii_.MemberMethod{JsiiMethod: "resetStripeSizeBytes", GoMethod: "ResetStripeSizeBytes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rowIndexStride", GoGetter: "RowIndexStride"},
			_jsii_.MemberProperty{JsiiProperty: "rowIndexStrideInput", GoGetter: "RowIndexStrideInput"},
			_jsii_.MemberProperty{JsiiProperty: "stripeSizeBytes", GoGetter: "StripeSizeBytes"},
			_jsii_.MemberProperty{JsiiProperty: "stripeSizeBytesInput", GoGetter: "StripeSizeBytesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "orcSerDe", GoGetter: "OrcSerDe"},
			_jsii_.MemberProperty{JsiiProperty: "orcSerDeInput", GoGetter: "OrcSerDeInput"},
			_jsii_.MemberProperty{JsiiProperty: "parquetSerDe", GoGetter: "ParquetSerDe"},
			_jsii_.MemberProperty{JsiiProperty: "parquetSerDeInput", GoGetter: "ParquetSerDeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putOrcSerDe", GoMethod: "PutOrcSerDe"},
			_jsii_.MemberMethod{JsiiMethod: "putParquetSerDe", GoMethod: "PutParquetSerDe"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrcSerDe", GoMethod: "ResetOrcSerDe"},
			_jsii_.MemberMethod{JsiiMethod: "resetParquetSerDe", GoMethod: "ResetParquetSerDe"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDeOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "blockSizeBytes", GoGetter: "BlockSizeBytes"},
			_jsii_.MemberProperty{JsiiProperty: "blockSizeBytesInput", GoGetter: "BlockSizeBytesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberProperty{JsiiProperty: "compressionInput", GoGetter: "CompressionInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableDictionaryCompression", GoGetter: "EnableDictionaryCompression"},
			_jsii_.MemberProperty{JsiiProperty: "enableDictionaryCompressionInput", GoGetter: "EnableDictionaryCompressionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "maxPaddingBytes", GoGetter: "MaxPaddingBytes"},
			_jsii_.MemberProperty{JsiiProperty: "maxPaddingBytesInput", GoGetter: "MaxPaddingBytesInput"},
			_jsii_.MemberProperty{JsiiProperty: "pageSizeBytes", GoGetter: "PageSizeBytes"},
			_jsii_.MemberProperty{JsiiProperty: "pageSizeBytesInput", GoGetter: "PageSizeBytesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlockSizeBytes", GoMethod: "ResetBlockSizeBytes"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompression", GoMethod: "ResetCompression"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableDictionaryCompression", GoMethod: "ResetEnableDictionaryCompression"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxPaddingBytes", GoMethod: "ResetMaxPaddingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "resetPageSizeBytes", GoMethod: "ResetPageSizeBytes"},
			_jsii_.MemberMethod{JsiiMethod: "resetWriterVersion", GoMethod: "ResetWriterVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "writerVersion", GoGetter: "WriterVersion"},
			_jsii_.MemberProperty{JsiiProperty: "writerVersionInput", GoGetter: "WriterVersionInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "inputFormatConfiguration", GoGetter: "InputFormatConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "inputFormatConfigurationInput", GoGetter: "InputFormatConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "outputFormatConfiguration", GoGetter: "OutputFormatConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "outputFormatConfigurationInput", GoGetter: "OutputFormatConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putInputFormatConfiguration", GoMethod: "PutInputFormatConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putOutputFormatConfiguration", GoMethod: "PutOutputFormatConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putSchemaConfiguration", GoMethod: "PutSchemaConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetInputFormatConfiguration", GoMethod: "ResetInputFormatConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutputFormatConfiguration", GoMethod: "ResetOutputFormatConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchemaConfiguration", GoMethod: "ResetSchemaConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schemaConfiguration", GoGetter: "SchemaConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "schemaConfigurationInput", GoGetter: "SchemaConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationSchemaConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationSchemaConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationSchemaConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationSchemaConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "catalogId", GoGetter: "CatalogId"},
			_jsii_.MemberProperty{JsiiProperty: "catalogIdInput", GoGetter: "CatalogIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "databaseNameInput", GoGetter: "DatabaseNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCatalogId", GoMethod: "ResetCatalogId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatabaseName", GoMethod: "ResetDatabaseName"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleArn", GoMethod: "ResetRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTableName", GoMethod: "ResetTableName"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersionId", GoMethod: "ResetVersionId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableNameInput", GoGetter: "TableNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "versionId", GoGetter: "VersionId"},
			_jsii_.MemberProperty{JsiiProperty: "versionIdInput", GoGetter: "VersionIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationSchemaConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "putRetryOptions", GoMethod: "PutRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryOptions", GoMethod: "ResetRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptions", GoGetter: "RetryOptions"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptionsInput", GoGetter: "RetryOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSeconds", GoGetter: "DurationInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSecondsInput", GoGetter: "DurationInSecondsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDurationInSeconds", GoMethod: "ResetDurationInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfig",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArn", GoGetter: "AwskmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArnInput", GoGetter: "AwskmsKeyArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfig", GoGetter: "KmsEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfigInput", GoGetter: "KmsEncryptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfig", GoGetter: "NoEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfigInput", GoGetter: "NoEncryptionConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKmsEncryptionConfig", GoMethod: "PutKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsEncryptionConfig", GoMethod: "ResetKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoEncryptionConfig", GoMethod: "ResetNoEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArnInput", GoGetter: "BucketArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormatInput", GoGetter: "CompressionFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormatConversionConfiguration", GoGetter: "DataFormatConversionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormatConversionConfigurationInput", GoGetter: "DataFormatConversionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicPartitioningConfiguration", GoGetter: "DynamicPartitioningConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicPartitioningConfigurationInput", GoGetter: "DynamicPartitioningConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfigurationInput", GoGetter: "EncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefixInput", GoGetter: "ErrorOutputPrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfiguration", GoGetter: "ProcessingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfigurationInput", GoGetter: "ProcessingConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putDataFormatConversionConfiguration", GoMethod: "PutDataFormatConversionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putDynamicPartitioningConfiguration", GoMethod: "PutDynamicPartitioningConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionConfiguration", GoMethod: "PutEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessingConfiguration", GoMethod: "PutProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putS3BackupConfiguration", GoMethod: "PutS3BackupConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompressionFormat", GoMethod: "ResetCompressionFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataFormatConversionConfiguration", GoMethod: "ResetDataFormatConversionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetDynamicPartitioningConfiguration", GoMethod: "ResetDynamicPartitioningConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionConfiguration", GoMethod: "ResetEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorOutputPrefix", GoMethod: "ResetErrorOutputPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessingConfiguration", GoMethod: "ResetProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3BackupConfiguration", GoMethod: "ResetS3BackupConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3BackupMode", GoMethod: "ResetS3BackupMode"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupConfiguration", GoGetter: "S3BackupConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupConfigurationInput", GoGetter: "S3BackupConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupMode", GoGetter: "S3BackupMode"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupModeInput", GoGetter: "S3BackupModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "processors", GoGetter: "Processors"},
			_jsii_.MemberProperty{JsiiProperty: "processorsInput", GoGetter: "ProcessorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessors", GoMethod: "PutProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessors", GoMethod: "ResetProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessors",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessors)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInput", GoGetter: "ParametersInput"},
			_jsii_.MemberMethod{JsiiMethod: "putParameters", GoMethod: "PutParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameters", GoMethod: "ResetParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsParameters",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsParametersList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsParametersList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsParametersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
			_jsii_.MemberProperty{JsiiProperty: "parameterNameInput", GoGetter: "ParameterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValue", GoGetter: "ParameterValue"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValueInput", GoGetter: "ParameterValueInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationEncryptionConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationEncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationEncryptionConfigurationKmsEncryptionConfig",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationEncryptionConfigurationKmsEncryptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArn", GoGetter: "AwskmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArnInput", GoGetter: "AwskmsKeyArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationEncryptionConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationEncryptionConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfig", GoGetter: "KmsEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfigInput", GoGetter: "KmsEncryptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfig", GoGetter: "NoEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfigInput", GoGetter: "NoEncryptionConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKmsEncryptionConfig", GoMethod: "PutKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsEncryptionConfig", GoMethod: "ResetKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoEncryptionConfig", GoMethod: "ResetNoEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationEncryptionConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArnInput", GoGetter: "BucketArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormatInput", GoGetter: "CompressionFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfigurationInput", GoGetter: "EncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefixInput", GoGetter: "ErrorOutputPrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionConfiguration", GoMethod: "PutEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompressionFormat", GoMethod: "ResetCompressionFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionConfiguration", GoMethod: "ResetEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorOutputPrefix", GoMethod: "ResetErrorOutputPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationEndpointConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationEndpointConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationEndpointConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationEndpointConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessKey", GoGetter: "AccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "accessKeyInput", GoGetter: "AccessKeyInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessKey", GoMethod: "ResetAccessKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationEndpointConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfiguration", GoGetter: "EndpointConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfigurationInput", GoGetter: "EndpointConfigurationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "processingConfiguration", GoGetter: "ProcessingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfigurationInput", GoGetter: "ProcessingConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEndpointConfiguration", GoMethod: "PutEndpointConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessingConfiguration", GoMethod: "PutProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putRequestConfiguration", GoMethod: "PutRequestConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryOptions", GoMethod: "PutRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putS3Configuration", GoMethod: "PutS3Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "requestConfiguration", GoGetter: "RequestConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "requestConfigurationInput", GoGetter: "RequestConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessingConfiguration", GoMethod: "ResetProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestConfiguration", GoMethod: "ResetRequestConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryOptions", GoMethod: "ResetRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleArn", GoMethod: "ResetRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3BackupMode", GoMethod: "ResetS3BackupMode"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptions", GoGetter: "RetryOptions"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptionsInput", GoGetter: "RetryOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupMode", GoGetter: "S3BackupMode"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupModeInput", GoGetter: "S3BackupModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3Configuration", GoGetter: "S3Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "s3ConfigurationInput", GoGetter: "S3ConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "processors", GoGetter: "Processors"},
			_jsii_.MemberProperty{JsiiProperty: "processorsInput", GoGetter: "ProcessorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessors", GoMethod: "PutProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessors", GoMethod: "ResetProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessors",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessors)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInput", GoGetter: "ParametersInput"},
			_jsii_.MemberMethod{JsiiMethod: "putParameters", GoMethod: "PutParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameters", GoMethod: "ResetParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParameters",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
			_jsii_.MemberProperty{JsiiProperty: "parameterNameInput", GoGetter: "ParameterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValue", GoGetter: "ParameterValue"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValueInput", GoGetter: "ParameterValueInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationCommonAttributes",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationCommonAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationCommonAttributesList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationCommonAttributesList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationCommonAttributesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationCommonAttributesOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationCommonAttributesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attributeName", GoGetter: "AttributeName"},
			_jsii_.MemberProperty{JsiiProperty: "attributeNameInput", GoGetter: "AttributeNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "attributeValue", GoGetter: "AttributeValue"},
			_jsii_.MemberProperty{JsiiProperty: "attributeValueInput", GoGetter: "AttributeValueInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationCommonAttributesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "commonAttributes", GoGetter: "CommonAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "commonAttributesInput", GoGetter: "CommonAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "contentEncoding", GoGetter: "ContentEncoding"},
			_jsii_.MemberProperty{JsiiProperty: "contentEncodingInput", GoGetter: "ContentEncodingInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putCommonAttributes", GoMethod: "PutCommonAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommonAttributes", GoMethod: "ResetCommonAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentEncoding", GoMethod: "ResetContentEncoding"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRetryOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRetryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRetryOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRetryOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSeconds", GoGetter: "DurationInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSecondsInput", GoGetter: "DurationInSecondsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDurationInSeconds", GoMethod: "ResetDurationInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRetryOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3Configuration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3Configuration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationEncryptionConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationEncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArn", GoGetter: "AwskmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArnInput", GoGetter: "AwskmsKeyArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfig", GoGetter: "KmsEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfigInput", GoGetter: "KmsEncryptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfig", GoGetter: "NoEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfigInput", GoGetter: "NoEncryptionConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKmsEncryptionConfig", GoMethod: "PutKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsEncryptionConfig", GoMethod: "ResetKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoEncryptionConfig", GoMethod: "ResetNoEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArnInput", GoGetter: "BucketArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormatInput", GoGetter: "CompressionFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfigurationInput", GoGetter: "EncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefixInput", GoGetter: "ErrorOutputPrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionConfiguration", GoMethod: "PutEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompressionFormat", GoMethod: "ResetCompressionFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionConfiguration", GoMethod: "ResetEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorOutputPrefix", GoMethod: "ResetErrorOutputPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamKinesisStreamSourceConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamKinesisStreamSourceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamKinesisStreamSourceConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamKinesisStreamSourceConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kinesisStreamArn", GoGetter: "KinesisStreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisStreamArnInput", GoGetter: "KinesisStreamArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamKinesisStreamSourceConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamMskSourceConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamMskSourceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamMskSourceConfigurationAuthenticationConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamMskSourceConfigurationAuthenticationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamMskSourceConfigurationAuthenticationConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamMskSourceConfigurationAuthenticationConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectivity", GoGetter: "Connectivity"},
			_jsii_.MemberProperty{JsiiProperty: "connectivityInput", GoGetter: "ConnectivityInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamMskSourceConfigurationAuthenticationConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamMskSourceConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamMskSourceConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authenticationConfiguration", GoGetter: "AuthenticationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationConfigurationInput", GoGetter: "AuthenticationConfigurationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "mskClusterArn", GoGetter: "MskClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "mskClusterArnInput", GoGetter: "MskClusterArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthenticationConfiguration", GoMethod: "PutAuthenticationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "topicName", GoGetter: "TopicName"},
			_jsii_.MemberProperty{JsiiProperty: "topicNameInput", GoGetter: "TopicNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamMskSourceConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommand",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommand)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommandOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommandOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "copyOptions", GoGetter: "CopyOptions"},
			_jsii_.MemberProperty{JsiiProperty: "copyOptionsInput", GoGetter: "CopyOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataTableColumns", GoGetter: "DataTableColumns"},
			_jsii_.MemberProperty{JsiiProperty: "dataTableColumnsInput", GoGetter: "DataTableColumnsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataTableName", GoGetter: "DataTableName"},
			_jsii_.MemberProperty{JsiiProperty: "dataTableNameInput", GoGetter: "DataTableNameInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCopyOptions", GoMethod: "ResetCopyOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataTableColumns", GoMethod: "ResetDataTableColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommandOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "clusterJdbcurl", GoGetter: "ClusterJdbcurl"},
			_jsii_.MemberProperty{JsiiProperty: "clusterJdbcurlInput", GoGetter: "ClusterJdbcurlInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "copyCommand", GoGetter: "CopyCommand"},
			_jsii_.MemberProperty{JsiiProperty: "copyCommandInput", GoGetter: "CopyCommandInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfiguration", GoGetter: "ProcessingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfigurationInput", GoGetter: "ProcessingConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putCopyCommand", GoMethod: "PutCopyCommand"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessingConfiguration", GoMethod: "PutProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryOptions", GoMethod: "PutRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putS3BackupConfiguration", GoMethod: "PutS3BackupConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putS3Configuration", GoMethod: "PutS3Configuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessingConfiguration", GoMethod: "ResetProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryOptions", GoMethod: "ResetRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3BackupConfiguration", GoMethod: "ResetS3BackupConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3BackupMode", GoMethod: "ResetS3BackupMode"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptions", GoGetter: "RetryOptions"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptionsInput", GoGetter: "RetryOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupConfiguration", GoGetter: "S3BackupConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupConfigurationInput", GoGetter: "S3BackupConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupMode", GoGetter: "S3BackupMode"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupModeInput", GoGetter: "S3BackupModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3Configuration", GoGetter: "S3Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "s3ConfigurationInput", GoGetter: "S3ConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "usernameInput", GoGetter: "UsernameInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "processors", GoGetter: "Processors"},
			_jsii_.MemberProperty{JsiiProperty: "processorsInput", GoGetter: "ProcessorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessors", GoMethod: "PutProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessors", GoMethod: "ResetProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessors",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessors)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInput", GoGetter: "ParametersInput"},
			_jsii_.MemberMethod{JsiiMethod: "putParameters", GoMethod: "PutParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameters", GoMethod: "ResetParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsParameters",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsParametersList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsParametersList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsParametersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
			_jsii_.MemberProperty{JsiiProperty: "parameterNameInput", GoGetter: "ParameterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValue", GoGetter: "ParameterValue"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValueInput", GoGetter: "ParameterValueInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationRetryOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationRetryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationRetryOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationRetryOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSeconds", GoGetter: "DurationInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSecondsInput", GoGetter: "DurationInSecondsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDurationInSeconds", GoMethod: "ResetDurationInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationRetryOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationEncryptionConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationEncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationEncryptionConfigurationKmsEncryptionConfig",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationEncryptionConfigurationKmsEncryptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArn", GoGetter: "AwskmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArnInput", GoGetter: "AwskmsKeyArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationEncryptionConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationEncryptionConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfig", GoGetter: "KmsEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfigInput", GoGetter: "KmsEncryptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfig", GoGetter: "NoEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfigInput", GoGetter: "NoEncryptionConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKmsEncryptionConfig", GoMethod: "PutKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsEncryptionConfig", GoMethod: "ResetKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoEncryptionConfig", GoMethod: "ResetNoEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationEncryptionConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArnInput", GoGetter: "BucketArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormatInput", GoGetter: "CompressionFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfigurationInput", GoGetter: "EncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefixInput", GoGetter: "ErrorOutputPrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionConfiguration", GoMethod: "PutEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompressionFormat", GoMethod: "ResetCompressionFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionConfiguration", GoMethod: "ResetEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorOutputPrefix", GoMethod: "ResetErrorOutputPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3Configuration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3Configuration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationEncryptionConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationEncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArn", GoGetter: "AwskmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArnInput", GoGetter: "AwskmsKeyArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfig", GoGetter: "KmsEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfigInput", GoGetter: "KmsEncryptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfig", GoGetter: "NoEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfigInput", GoGetter: "NoEncryptionConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKmsEncryptionConfig", GoMethod: "PutKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsEncryptionConfig", GoMethod: "ResetKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoEncryptionConfig", GoMethod: "ResetNoEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArnInput", GoGetter: "BucketArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormatInput", GoGetter: "CompressionFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfigurationInput", GoGetter: "EncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefixInput", GoGetter: "ErrorOutputPrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionConfiguration", GoMethod: "PutEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompressionFormat", GoMethod: "ResetCompressionFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionConfiguration", GoMethod: "ResetEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorOutputPrefix", GoMethod: "ResetErrorOutputPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamS3DestinationConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamS3DestinationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamS3DestinationConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamS3DestinationConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamS3DestinationConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamS3DestinationConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamS3DestinationConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamS3DestinationConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamS3DestinationConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamS3DestinationConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfig",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArn", GoGetter: "AwskmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArnInput", GoGetter: "AwskmsKeyArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfig", GoGetter: "KmsEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfigInput", GoGetter: "KmsEncryptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfig", GoGetter: "NoEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfigInput", GoGetter: "NoEncryptionConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKmsEncryptionConfig", GoMethod: "PutKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsEncryptionConfig", GoMethod: "ResetKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoEncryptionConfig", GoMethod: "ResetNoEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamS3DestinationConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamS3DestinationConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArnInput", GoGetter: "BucketArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormatInput", GoGetter: "CompressionFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfigurationInput", GoGetter: "EncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefixInput", GoGetter: "ErrorOutputPrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionConfiguration", GoMethod: "PutEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompressionFormat", GoMethod: "ResetCompressionFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionConfiguration", GoMethod: "ResetEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorOutputPrefix", GoMethod: "ResetErrorOutputPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountUrl", GoGetter: "AccountUrl"},
			_jsii_.MemberProperty{JsiiProperty: "accountUrlInput", GoGetter: "AccountUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "contentColumnName", GoGetter: "ContentColumnName"},
			_jsii_.MemberProperty{JsiiProperty: "contentColumnNameInput", GoGetter: "ContentColumnNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInput", GoGetter: "DatabaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataLoadingOption", GoGetter: "DataLoadingOption"},
			_jsii_.MemberProperty{JsiiProperty: "dataLoadingOptionInput", GoGetter: "DataLoadingOptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "keyPassphrase", GoGetter: "KeyPassphrase"},
			_jsii_.MemberProperty{JsiiProperty: "keyPassphraseInput", GoGetter: "KeyPassphraseInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaDataColumnName", GoGetter: "MetaDataColumnName"},
			_jsii_.MemberProperty{JsiiProperty: "metaDataColumnNameInput", GoGetter: "MetaDataColumnNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyInput", GoGetter: "PrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfiguration", GoGetter: "ProcessingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfigurationInput", GoGetter: "ProcessingConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessingConfiguration", GoMethod: "PutProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryOptions", GoMethod: "PutRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putS3Configuration", GoMethod: "PutS3Configuration"},
			_jsii_.MemberMethod{JsiiMethod: "putSnowflakeRoleConfiguration", GoMethod: "PutSnowflakeRoleConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putSnowflakeVpcConfiguration", GoMethod: "PutSnowflakeVpcConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentColumnName", GoMethod: "ResetContentColumnName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataLoadingOption", GoMethod: "ResetDataLoadingOption"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyPassphrase", GoMethod: "ResetKeyPassphrase"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetaDataColumnName", GoMethod: "ResetMetaDataColumnName"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessingConfiguration", GoMethod: "ResetProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryOptions", GoMethod: "ResetRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3BackupMode", GoMethod: "ResetS3BackupMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnowflakeRoleConfiguration", GoMethod: "ResetSnowflakeRoleConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnowflakeVpcConfiguration", GoMethod: "ResetSnowflakeVpcConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptions", GoGetter: "RetryOptions"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptionsInput", GoGetter: "RetryOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupMode", GoGetter: "S3BackupMode"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupModeInput", GoGetter: "S3BackupModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3Configuration", GoGetter: "S3Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "s3ConfigurationInput", GoGetter: "S3ConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeRoleConfiguration", GoGetter: "SnowflakeRoleConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeRoleConfigurationInput", GoGetter: "SnowflakeRoleConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeVpcConfiguration", GoGetter: "SnowflakeVpcConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeVpcConfigurationInput", GoGetter: "SnowflakeVpcConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "userInput", GoGetter: "UserInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "processors", GoGetter: "Processors"},
			_jsii_.MemberProperty{JsiiProperty: "processorsInput", GoGetter: "ProcessorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessors", GoMethod: "PutProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessors", GoMethod: "ResetProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessors",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessors)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInput", GoGetter: "ParametersInput"},
			_jsii_.MemberMethod{JsiiMethod: "putParameters", GoMethod: "PutParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameters", GoMethod: "ResetParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsParameters",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsParametersList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsParametersList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsParametersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
			_jsii_.MemberProperty{JsiiProperty: "parameterNameInput", GoGetter: "ParameterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValue", GoGetter: "ParameterValue"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValueInput", GoGetter: "ParameterValueInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationRetryOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationRetryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationRetryOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationRetryOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSeconds", GoGetter: "DurationInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSecondsInput", GoGetter: "DurationInSecondsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDurationInSeconds", GoMethod: "ResetDurationInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationRetryOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3Configuration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3Configuration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationEncryptionConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationEncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArn", GoGetter: "AwskmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArnInput", GoGetter: "AwskmsKeyArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfig", GoGetter: "KmsEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfigInput", GoGetter: "KmsEncryptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfig", GoGetter: "NoEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfigInput", GoGetter: "NoEncryptionConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKmsEncryptionConfig", GoMethod: "PutKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsEncryptionConfig", GoMethod: "ResetKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoEncryptionConfig", GoMethod: "ResetNoEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArnInput", GoGetter: "BucketArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormatInput", GoGetter: "CompressionFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfigurationInput", GoGetter: "EncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefixInput", GoGetter: "ErrorOutputPrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionConfiguration", GoMethod: "PutEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompressionFormat", GoMethod: "ResetCompressionFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionConfiguration", GoMethod: "ResetEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorOutputPrefix", GoMethod: "ResetErrorOutputPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3ConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeRoleConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeRoleConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeRoleConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeRoleConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnowflakeRole", GoMethod: "ResetSnowflakeRole"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeRole", GoGetter: "SnowflakeRole"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeRoleInput", GoGetter: "SnowflakeRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeRoleConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeVpcConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeVpcConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeVpcConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeVpcConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "privateLinkVpceId", GoGetter: "PrivateLinkVpceId"},
			_jsii_.MemberProperty{JsiiProperty: "privateLinkVpceIdInput", GoGetter: "PrivateLinkVpceIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeVpcConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hecAcknowledgmentTimeoutInSeconds", GoGetter: "HecAcknowledgmentTimeoutInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "hecAcknowledgmentTimeoutInSecondsInput", GoGetter: "HecAcknowledgmentTimeoutInSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "hecEndpoint", GoGetter: "HecEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "hecEndpointInput", GoGetter: "HecEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "hecEndpointType", GoGetter: "HecEndpointType"},
			_jsii_.MemberProperty{JsiiProperty: "hecEndpointTypeInput", GoGetter: "HecEndpointTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "hecToken", GoGetter: "HecToken"},
			_jsii_.MemberProperty{JsiiProperty: "hecTokenInput", GoGetter: "HecTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfiguration", GoGetter: "ProcessingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "processingConfigurationInput", GoGetter: "ProcessingConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessingConfiguration", GoMethod: "PutProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryOptions", GoMethod: "PutRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putS3Configuration", GoMethod: "PutS3Configuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetHecAcknowledgmentTimeoutInSeconds", GoMethod: "ResetHecAcknowledgmentTimeoutInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessingConfiguration", GoMethod: "ResetProcessingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryOptions", GoMethod: "ResetRetryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3BackupMode", GoMethod: "ResetS3BackupMode"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptions", GoGetter: "RetryOptions"},
			_jsii_.MemberProperty{JsiiProperty: "retryOptionsInput", GoGetter: "RetryOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupMode", GoGetter: "S3BackupMode"},
			_jsii_.MemberProperty{JsiiProperty: "s3BackupModeInput", GoGetter: "S3BackupModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3Configuration", GoGetter: "S3Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "s3ConfigurationInput", GoGetter: "S3ConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "processors", GoGetter: "Processors"},
			_jsii_.MemberProperty{JsiiProperty: "processorsInput", GoGetter: "ProcessorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProcessors", GoMethod: "PutProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessors", GoMethod: "ResetProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessors",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessors)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInput", GoGetter: "ParametersInput"},
			_jsii_.MemberMethod{JsiiMethod: "putParameters", GoMethod: "PutParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameters", GoMethod: "ResetParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsParameters",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsParametersList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsParametersList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsParametersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
			_jsii_.MemberProperty{JsiiProperty: "parameterNameInput", GoGetter: "ParameterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValue", GoGetter: "ParameterValue"},
			_jsii_.MemberProperty{JsiiProperty: "parameterValueInput", GoGetter: "ParameterValueInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationRetryOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationRetryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationRetryOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationRetryOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSeconds", GoGetter: "DurationInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "durationInSecondsInput", GoGetter: "DurationInSecondsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDurationInSeconds", GoMethod: "ResetDurationInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationRetryOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3Configuration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3Configuration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationBufferingHints",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationBufferingHints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationBufferingHintsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationBufferingHintsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "intervalInSeconds", GoGetter: "IntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInSecondsInput", GoGetter: "IntervalInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntervalInSeconds", GoMethod: "ResetIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMBs", GoMethod: "ResetSizeInMBs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBs", GoGetter: "SizeInMBs"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMBsInput", GoGetter: "SizeInMBsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationBufferingHintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupNameInput", GoGetter: "LogGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroupName", GoMethod: "ResetLogGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationEncryptionConfiguration",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationEncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArn", GoGetter: "AwskmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "awskmsKeyArnInput", GoGetter: "AwskmsKeyArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfig", GoGetter: "KmsEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionConfigInput", GoGetter: "KmsEncryptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfig", GoGetter: "NoEncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "noEncryptionConfigInput", GoGetter: "NoEncryptionConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKmsEncryptionConfig", GoMethod: "PutKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsEncryptionConfig", GoMethod: "ResetKmsEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoEncryptionConfig", GoMethod: "ResetNoEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArnInput", GoGetter: "BucketArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHints", GoGetter: "BufferingHints"},
			_jsii_.MemberProperty{JsiiProperty: "bufferingHintsInput", GoGetter: "BufferingHintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptions", GoGetter: "CloudwatchLoggingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingOptionsInput", GoGetter: "CloudwatchLoggingOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormatInput", GoGetter: "CompressionFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfigurationInput", GoGetter: "EncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefixInput", GoGetter: "ErrorOutputPrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBufferingHints", GoMethod: "PutBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLoggingOptions", GoMethod: "PutCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionConfiguration", GoMethod: "PutEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBufferingHints", GoMethod: "ResetBufferingHints"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLoggingOptions", GoMethod: "ResetCloudwatchLoggingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompressionFormat", GoMethod: "ResetCompressionFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionConfiguration", GoMethod: "ResetEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorOutputPrefix", GoMethod: "ResetErrorOutputPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamTags",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamTagsList",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamTagsList)(nil)).Elem(),
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
			j := jsiiProxy_KinesisfirehoseDeliveryStreamTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamTagsOutputReference",
		reflect.TypeOf((*KinesisfirehoseDeliveryStreamTagsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisfirehoseDeliveryStreamTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
