// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/shipwright-io/build/pkg/apis/build/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuildStrategies implements BuildStrategyInterface
type FakeBuildStrategies struct {
	Fake *FakeShipwrightV1beta1
	ns   string
}

var buildstrategiesResource = v1beta1.SchemeGroupVersion.WithResource("buildstrategies")

var buildstrategiesKind = v1beta1.SchemeGroupVersion.WithKind("BuildStrategy")

// Get takes name of the buildStrategy, and returns the corresponding buildStrategy object, and an error if there is any.
func (c *FakeBuildStrategies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BuildStrategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(buildstrategiesResource, c.ns, name), &v1beta1.BuildStrategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BuildStrategy), err
}

// List takes label and field selectors, and returns the list of BuildStrategies that match those selectors.
func (c *FakeBuildStrategies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BuildStrategyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(buildstrategiesResource, buildstrategiesKind, c.ns, opts), &v1beta1.BuildStrategyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BuildStrategyList{ListMeta: obj.(*v1beta1.BuildStrategyList).ListMeta}
	for _, item := range obj.(*v1beta1.BuildStrategyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested buildStrategies.
func (c *FakeBuildStrategies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(buildstrategiesResource, c.ns, opts))

}

// Create takes the representation of a buildStrategy and creates it.  Returns the server's representation of the buildStrategy, and an error, if there is any.
func (c *FakeBuildStrategies) Create(ctx context.Context, buildStrategy *v1beta1.BuildStrategy, opts v1.CreateOptions) (result *v1beta1.BuildStrategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(buildstrategiesResource, c.ns, buildStrategy), &v1beta1.BuildStrategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BuildStrategy), err
}

// Update takes the representation of a buildStrategy and updates it. Returns the server's representation of the buildStrategy, and an error, if there is any.
func (c *FakeBuildStrategies) Update(ctx context.Context, buildStrategy *v1beta1.BuildStrategy, opts v1.UpdateOptions) (result *v1beta1.BuildStrategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(buildstrategiesResource, c.ns, buildStrategy), &v1beta1.BuildStrategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BuildStrategy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBuildStrategies) UpdateStatus(ctx context.Context, buildStrategy *v1beta1.BuildStrategy, opts v1.UpdateOptions) (*v1beta1.BuildStrategy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(buildstrategiesResource, "status", c.ns, buildStrategy), &v1beta1.BuildStrategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BuildStrategy), err
}

// Delete takes name of the buildStrategy and deletes it. Returns an error if one occurs.
func (c *FakeBuildStrategies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(buildstrategiesResource, c.ns, name, opts), &v1beta1.BuildStrategy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuildStrategies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(buildstrategiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.BuildStrategyList{})
	return err
}

// Patch applies the patch and returns the patched buildStrategy.
func (c *FakeBuildStrategies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BuildStrategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildstrategiesResource, c.ns, name, pt, data, subresources...), &v1beta1.BuildStrategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BuildStrategy), err
}
