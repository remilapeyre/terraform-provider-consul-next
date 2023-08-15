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

func NewNamespaceRoleAttachment() resource.Resource {
	return NewResource(
		"namespace_role_attachment",
		namespaceRoleAttachmentSchema(),
		&NamespaceRoleAttachment{},
	)
}

type NamespaceRoleAttachment struct{}

func (r *NamespaceRoleAttachment) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *structs.NamespaceRoleAttachment
	resp.Diagnostics.Append(models.DecodeNamespaceRoleAttachment(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	namespace, _, err := client.Namespaces().Read(config.Namespace, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read namespace", err.Error())
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
		ID:   config.Role.ID,
		Name: config.Role.Name,
	}
	if namespaceAttachmentFound(namespace.ACLs.RoleDefaults, &link) != nil {
		ident := link.ID + link.Name
		resp.Diagnostics.AddError("role already attached to the namespace", fmt.Sprintf("role %q is already attached to namespace %s", ident, config.Namespace))
		return
	}

	namespace.ACLs.RoleDefaults = append(namespace.ACLs.RoleDefaults, link)

	namespace, _, err = client.Namespaces().Update(namespace, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to attach role", err.Error())
		return
	}

	config.ID = fmt.Sprintf("%s-%s", config.Namespace, config.Role.ID+config.Role.Name)
	newLink := namespaceAttachmentFound(namespace.ACLs.RoleDefaults, &link)
	if newLink == nil {
		resp.Diagnostics.AddError("failed to attach role", "failed to find role attached to namespace")
		return
	}
	config.Role.ID = newLink.ID
	config.Role.Name = newLink.Name
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *NamespaceRoleAttachment) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *structs.NamespaceRoleAttachment
	resp.Diagnostics.Append(models.DecodeNamespaceRoleAttachment(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	namespace, _, err := client.Namespaces().Read(state.Namespace, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read namespace", err.Error())
		return
	}

	link := &api.ACLLink{
		ID:   state.Role.ID,
		Name: state.Role.Name,
	}
	if namespaceAttachmentFound(namespace.ACLs.RoleDefaults, link) == nil {
		resp.State.RemoveResource(ctx)
		return
	}
}

func (r *NamespaceRoleAttachment) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "namespace_role_attachment cannot be updated")
}

func (r *NamespaceRoleAttachment) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *structs.NamespaceRoleAttachment
	resp.Diagnostics.Append(models.DecodeNamespaceRoleAttachment(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	namespace, _, err := client.Namespaces().Read(state.Namespace, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to detach ACL role", err.Error())
		return
	}
	for i, role := range namespace.ACLs.RoleDefaults {
		if state.Role.ID == role.ID {
			namespace.ACLs.RoleDefaults = append(namespace.ACLs.RoleDefaults[:i], namespace.ACLs.RoleDefaults[i+1:]...)
			break
		}
	}

	_, _, err = client.Namespaces().Update(namespace, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to detach ACL role", err.Error())
		return
	}
}

func namespaceAttachmentFound(links []api.ACLLink, link *api.ACLLink) *api.ACLLink {
	for _, l := range links {
		if l.ID == link.ID || l.Name == link.Name {
			return &l
		}
	}
	return nil
}
