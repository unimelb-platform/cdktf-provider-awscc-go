package quicksightdataset


type QuicksightDataSetColumnGroups struct {
	// <p>Geospatial column group that denotes a hierarchy.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#geo_spatial_column_group QuicksightDataSet#geo_spatial_column_group}
	GeoSpatialColumnGroup *QuicksightDataSetColumnGroupsGeoSpatialColumnGroup `field:"optional" json:"geoSpatialColumnGroup" yaml:"geoSpatialColumnGroup"`
}

