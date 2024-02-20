package lambdacodesigningconfig


type LambdaCodeSigningConfigCodeSigningPolicies struct {
	// Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_code_signing_config#untrusted_artifact_on_deployment LambdaCodeSigningConfig#untrusted_artifact_on_deployment}
	UntrustedArtifactOnDeployment *string `field:"optional" json:"untrustedArtifactOnDeployment" yaml:"untrustedArtifactOnDeployment"`
}

