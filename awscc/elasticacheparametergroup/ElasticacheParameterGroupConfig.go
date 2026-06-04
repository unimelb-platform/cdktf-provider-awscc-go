package elasticacheparametergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticacheParameterGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the cache parameter group family that this cache parameter group is compatible with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticache_parameter_group#cache_parameter_group_family ElasticacheParameterGroup#cache_parameter_group_family}
	CacheParameterGroupFamily *string `field:"required" json:"cacheParameterGroupFamily" yaml:"cacheParameterGroupFamily"`
	// The description for this cache parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticache_parameter_group#description ElasticacheParameterGroup#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// A comma-delimited list of parameter name/value pairs. For more information see ModifyCacheParameterGroup in the Amazon ElastiCache API Reference Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticache_parameter_group#properties ElasticacheParameterGroup#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Tags are composed of a Key/Value pair.
	//
	// You can use tags to categorize and track each parameter group. The tag value null is permitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticache_parameter_group#tags ElasticacheParameterGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

