package panoramapackageversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PanoramaPackageVersionConfig struct {
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
	// A package ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/panorama_package_version#package_id PanoramaPackageVersion#package_id}
	PackageId *string `field:"required" json:"packageId" yaml:"packageId"`
	// A package version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/panorama_package_version#package_version PanoramaPackageVersion#package_version}
	PackageVersion *string `field:"required" json:"packageVersion" yaml:"packageVersion"`
	// A patch version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/panorama_package_version#patch_version PanoramaPackageVersion#patch_version}
	PatchVersion *string `field:"required" json:"patchVersion" yaml:"patchVersion"`
	// Whether to mark the new version as the latest version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/panorama_package_version#mark_latest PanoramaPackageVersion#mark_latest}
	MarkLatest interface{} `field:"optional" json:"markLatest" yaml:"markLatest"`
	// An owner account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/panorama_package_version#owner_account PanoramaPackageVersion#owner_account}
	OwnerAccount *string `field:"optional" json:"ownerAccount" yaml:"ownerAccount"`
	// If the version was marked latest, the new version to maker as latest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/panorama_package_version#updated_latest_patch_version PanoramaPackageVersion#updated_latest_patch_version}
	UpdatedLatestPatchVersion *string `field:"optional" json:"updatedLatestPatchVersion" yaml:"updatedLatestPatchVersion"`
}

