package amplifydomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AmplifyDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_domain#app_id AmplifyDomain#app_id}.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_domain#domain_name AmplifyDomain#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_domain#sub_domain_settings AmplifyDomain#sub_domain_settings}.
	SubDomainSettings interface{} `field:"required" json:"subDomainSettings" yaml:"subDomainSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_domain#auto_sub_domain_creation_patterns AmplifyDomain#auto_sub_domain_creation_patterns}.
	AutoSubDomainCreationPatterns *[]*string `field:"optional" json:"autoSubDomainCreationPatterns" yaml:"autoSubDomainCreationPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_domain#auto_sub_domain_iam_role AmplifyDomain#auto_sub_domain_iam_role}.
	AutoSubDomainIamRole *string `field:"optional" json:"autoSubDomainIamRole" yaml:"autoSubDomainIamRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_domain#enable_auto_sub_domain AmplifyDomain#enable_auto_sub_domain}.
	EnableAutoSubDomain interface{} `field:"optional" json:"enableAutoSubDomain" yaml:"enableAutoSubDomain"`
}

