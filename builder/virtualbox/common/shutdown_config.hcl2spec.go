// Code generated by "hcl2-schema -type ShutdownConfig"; DO NOT EDIT.\n

package common

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*ShutdownConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"ShutdownCommand":      &hcldec.AttrSpec{Name: "shutdown_command", Type: cty.String, Required: false},
		"RawShutdownTimeout":   &hcldec.AttrSpec{Name: "shutdown_timeout", Type: cty.String, Required: false},
		"RawPostShutdownDelay": &hcldec.AttrSpec{Name: "post_shutdown_delay", Type: cty.String, Required: false},
		"ShutdownTimeout":      &hcldec.AttrSpec{Name: "shutdowntimeout", Type: cty.String, Required: false},
		"PostShutdownDelay":    &hcldec.AttrSpec{Name: "postshutdowndelay", Type: cty.String, Required: false},
	}
	return s
}
