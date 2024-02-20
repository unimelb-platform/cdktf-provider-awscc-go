package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRedshiftSourceConfigVpcConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#security_group_id_list LookoutmetricsAnomalyDetector#security_group_id_list}.
	SecurityGroupIdList *[]*string `field:"required" json:"securityGroupIdList" yaml:"securityGroupIdList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#subnet_id_list LookoutmetricsAnomalyDetector#subnet_id_list}.
	SubnetIdList *[]*string `field:"required" json:"subnetIdList" yaml:"subnetIdList"`
}

