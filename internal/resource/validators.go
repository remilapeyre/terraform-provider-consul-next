// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type MutuallyExclusiveValidator struct {
	fields []string
}

func NewMutuallyExclusiveValidator(fields ...string) *MutuallyExclusiveValidator {
	return &MutuallyExclusiveValidator{fields: fields}
}

func (m *MutuallyExclusiveValidator) Description(context.Context) string {
	return ""
}

func (m *MutuallyExclusiveValidator) MarkdownDescription(context.Context) string {
	return ""
}

func (m *MutuallyExclusiveValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() {
		return
	}

	for _, field := range m.fields {
		var s basetypes.StringValue
		path := req.Path.ParentPath().AtName(field)
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.ParentPath().AtName(field), &s)...)
		if resp.Diagnostics.HasError() {
			return
		}

		if !s.IsNull() {
			resp.Diagnostics.AddAttributeError(req.Path, "mutually exclusive attributes", fmt.Sprintf("%s and %s cannot be both set", req.Path.String(), path))
			return
		}
	}
}
