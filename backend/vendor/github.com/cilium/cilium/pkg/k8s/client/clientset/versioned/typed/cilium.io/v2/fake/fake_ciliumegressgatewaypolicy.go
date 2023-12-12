// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCiliumEgressGatewayPolicies implements CiliumEgressGatewayPolicyInterface
type FakeCiliumEgressGatewayPolicies struct {
	Fake *FakeCiliumV2
}

var ciliumegressgatewaypoliciesResource = v2.SchemeGroupVersion.WithResource("ciliumegressgatewaypolicies")

var ciliumegressgatewaypoliciesKind = v2.SchemeGroupVersion.WithKind("CiliumEgressGatewayPolicy")

// Get takes name of the ciliumEgressGatewayPolicy, and returns the corresponding ciliumEgressGatewayPolicy object, and an error if there is any.
func (c *FakeCiliumEgressGatewayPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.CiliumEgressGatewayPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ciliumegressgatewaypoliciesResource, name), &v2.CiliumEgressGatewayPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEgressGatewayPolicy), err
}

// List takes label and field selectors, and returns the list of CiliumEgressGatewayPolicies that match those selectors.
func (c *FakeCiliumEgressGatewayPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v2.CiliumEgressGatewayPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ciliumegressgatewaypoliciesResource, ciliumegressgatewaypoliciesKind, opts), &v2.CiliumEgressGatewayPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.CiliumEgressGatewayPolicyList{ListMeta: obj.(*v2.CiliumEgressGatewayPolicyList).ListMeta}
	for _, item := range obj.(*v2.CiliumEgressGatewayPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumEgressGatewayPolicies.
func (c *FakeCiliumEgressGatewayPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ciliumegressgatewaypoliciesResource, opts))
}

// Create takes the representation of a ciliumEgressGatewayPolicy and creates it.  Returns the server's representation of the ciliumEgressGatewayPolicy, and an error, if there is any.
func (c *FakeCiliumEgressGatewayPolicies) Create(ctx context.Context, ciliumEgressGatewayPolicy *v2.CiliumEgressGatewayPolicy, opts v1.CreateOptions) (result *v2.CiliumEgressGatewayPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ciliumegressgatewaypoliciesResource, ciliumEgressGatewayPolicy), &v2.CiliumEgressGatewayPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEgressGatewayPolicy), err
}

// Update takes the representation of a ciliumEgressGatewayPolicy and updates it. Returns the server's representation of the ciliumEgressGatewayPolicy, and an error, if there is any.
func (c *FakeCiliumEgressGatewayPolicies) Update(ctx context.Context, ciliumEgressGatewayPolicy *v2.CiliumEgressGatewayPolicy, opts v1.UpdateOptions) (result *v2.CiliumEgressGatewayPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ciliumegressgatewaypoliciesResource, ciliumEgressGatewayPolicy), &v2.CiliumEgressGatewayPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEgressGatewayPolicy), err
}

// Delete takes name of the ciliumEgressGatewayPolicy and deletes it. Returns an error if one occurs.
func (c *FakeCiliumEgressGatewayPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ciliumegressgatewaypoliciesResource, name, opts), &v2.CiliumEgressGatewayPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumEgressGatewayPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ciliumegressgatewaypoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2.CiliumEgressGatewayPolicyList{})
	return err
}

// Patch applies the patch and returns the patched ciliumEgressGatewayPolicy.
func (c *FakeCiliumEgressGatewayPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumEgressGatewayPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ciliumegressgatewaypoliciesResource, name, pt, data, subresources...), &v2.CiliumEgressGatewayPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEgressGatewayPolicy), err
}
