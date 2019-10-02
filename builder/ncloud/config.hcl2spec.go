// Code generated by "hcl2-schema -type Config"; DO NOT EDIT.\n

package ncloud

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"AccessKey":                         &hcldec.AttrSpec{Name: "access_key", Type: cty.String, Required: false},
		"SecretKey":                         &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"ServerImageProductCode":            &hcldec.AttrSpec{Name: "server_image_product_code", Type: cty.String, Required: false},
		"ServerProductCode":                 &hcldec.AttrSpec{Name: "server_product_code", Type: cty.String, Required: false},
		"MemberServerImageNo":               &hcldec.AttrSpec{Name: "member_server_image_no", Type: cty.String, Required: false},
		"ServerImageName":                   &hcldec.AttrSpec{Name: "server_image_name", Type: cty.String, Required: false},
		"ServerImageDescription":            &hcldec.AttrSpec{Name: "server_image_description", Type: cty.String, Required: false},
		"UserData":                          &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"UserDataFile":                      &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"BlockStorageSize":                  &hcldec.AttrSpec{Name: "block_storage_size", Type: cty.Number, Required: false},
		"Region":                            &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"AccessControlGroupConfigurationNo": &hcldec.AttrSpec{Name: "access_control_group_configuration_no", Type: cty.String, Required: false},
	}
	return s
}
