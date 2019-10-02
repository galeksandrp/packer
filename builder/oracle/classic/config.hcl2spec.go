// Code generated by "hcl2-schema -type Config"; DO NOT EDIT.\n

package classic

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Username":                 &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"Password":                 &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"IdentityDomain":           &hcldec.AttrSpec{Name: "identity_domain", Type: cty.String, Required: false},
		"APIEndpoint":              &hcldec.AttrSpec{Name: "api_endpoint", Type: cty.String, Required: false},
		"ImageName":                &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"Shape":                    &hcldec.AttrSpec{Name: "shape", Type: cty.String, Required: false},
		"SourceImageList":          &hcldec.AttrSpec{Name: "source_image_list", Type: cty.String, Required: false},
		"SourceImageListEntry":     &hcldec.AttrSpec{Name: "source_image_list_entry", Type: cty.Number, Required: false},
		"SnapshotTimeout":          &hcldec.AttrSpec{Name: "snapshot_timeout", Type: cty.String, Required: false},
		"DestImageList":            &hcldec.AttrSpec{Name: "dest_image_list", Type: cty.String, Required: false},
		"Attributes":               &hcldec.AttrSpec{Name: "attributes", Type: cty.String, Required: false},
		"AttributesFile":           &hcldec.AttrSpec{Name: "attributes_file", Type: cty.String, Required: false},
		"DestImageListDescription": &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"SSHSourceList":            &hcldec.AttrSpec{Name: "ssh_source_list", Type: cty.String, Required: false},
	}
	for k, v := range (&Config{}).PVConfig.HCL2Spec() {
		s[k] = v
	}
	return s
}
