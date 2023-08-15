// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func NewNamespacePolicyAttachment() resource.Resource {
	return NewResource(
		"namespace_policy_attachment",
		namespacePolicyAttachmentSchema(),
		&NamespacePolicyAttachment{},
	)
}

type NamespacePolicyAttachment struct{}

func (r *NamespacePolicyAttachment) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *structs.NamespacePolicyAttachment
	resp.Diagnostics.Append(models.DecodeNamespacePolicyAttachment(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	namespace, _, err := client.Namespaces().Read(config.Namespace, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read ACL token", err.Error())
		return
	}
	if namespace == nil {
		resp.Diagnostics.AddError("namespace not found", fmt.Sprintf("namespace %q not found", config.Namespace))
		return
	}

	if namespace.ACLs == nil {
		namespace.ACLs = &api.NamespaceACLConfig{}
	}

	link := api.ACLLink{
		ID:   config.Policy.ID,
		Name: config.Policy.Name,
	}
	if namespaceAttachmentFound(namespace.ACLs.PolicyDefaults, &link) != nil {
		ident := link.ID + link.Name
		resp.Diagnostics.AddError("policy already attached to the namespace", fmt.Sprintf("policy %q is already attached to namespace %s", ident, config.Namespace))
		return
	}

	namespace.ACLs.PolicyDefaults = append(namespace.ACLs.PolicyDefaults, link)

	namespace, _, err = client.Namespaces().Update(namespace, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to attach policy", err.Error())
		return
	}

	config.ID = fmt.Sprintf("%s-%s", config.Namespace, config.Policy.ID+config.Policy.Name)
	newLink := namespaceAttachmentFound(namespace.ACLs.PolicyDefaults, &link)
	if newLink == nil {
		resp.Diagnostics.AddError("failed to attach policy", "failed to find policy attached to token")
		return
	}
	config.Policy.ID = newLink.ID
	config.Policy.Name = newLink.Name
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *NamespacePolicyAttachment) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *structs.NamespacePolicyAttachment
	resp.Diagnostics.Append(models.DecodeNamespacePolicyAttachment(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	namespace, _, err := client.Namespaces().Read(state.Namespace, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read namespace", err.Error())
		return
	}

	link := &api.ACLLink{
		ID:   state.Policy.ID,
		Name: state.Policy.Name,
	}
	if namespaceAttachmentFound(namespace.ACLs.PolicyDefaults, link) == nil {
		resp.State.RemoveResource(ctx)
		return
	}
}

func (r *NamespacePolicyAttachment) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "namespace_policy_attachment cannot be updated")
}

func (r *NamespacePolicyAttachment) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *structs.NamespacePolicyAttachment
	resp.Diagnostics.Append(models.DecodeNamespacePolicyAttachment(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	namespace, _, err := client.Namespaces().Read(state.Namespace, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to detach ACL policy", err.Error())
		return
	}
	for i, policy := range namespace.ACLs.PolicyDefaults {
		if state.Policy.ID == policy.ID {
			namespace.ACLs.PolicyDefaults = append(namespace.ACLs.PolicyDefaults[:i], namespace.ACLs.PolicyDefaults[i+1:]...)
			break
		}
	}

	_, _, err = client.Namespaces().Update(namespace, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to detach ACL policy", err.Error())
		return
	}
}
