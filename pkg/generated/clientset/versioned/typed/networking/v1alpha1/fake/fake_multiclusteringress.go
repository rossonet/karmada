/*
Copyright The Karmada Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMultiClusterIngresses implements MultiClusterIngressInterface
type FakeMultiClusterIngresses struct {
	Fake *FakeNetworkingV1alpha1
	ns   string
}

var multiclusteringressesResource = v1alpha1.SchemeGroupVersion.WithResource("multiclusteringresses")

var multiclusteringressesKind = v1alpha1.SchemeGroupVersion.WithKind("MultiClusterIngress")

// Get takes name of the multiClusterIngress, and returns the corresponding multiClusterIngress object, and an error if there is any.
func (c *FakeMultiClusterIngresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MultiClusterIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(multiclusteringressesResource, c.ns, name), &v1alpha1.MultiClusterIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterIngress), err
}

// List takes label and field selectors, and returns the list of MultiClusterIngresses that match those selectors.
func (c *FakeMultiClusterIngresses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MultiClusterIngressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(multiclusteringressesResource, multiclusteringressesKind, c.ns, opts), &v1alpha1.MultiClusterIngressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MultiClusterIngressList{ListMeta: obj.(*v1alpha1.MultiClusterIngressList).ListMeta}
	for _, item := range obj.(*v1alpha1.MultiClusterIngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested multiClusterIngresses.
func (c *FakeMultiClusterIngresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(multiclusteringressesResource, c.ns, opts))

}

// Create takes the representation of a multiClusterIngress and creates it.  Returns the server's representation of the multiClusterIngress, and an error, if there is any.
func (c *FakeMultiClusterIngresses) Create(ctx context.Context, multiClusterIngress *v1alpha1.MultiClusterIngress, opts v1.CreateOptions) (result *v1alpha1.MultiClusterIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(multiclusteringressesResource, c.ns, multiClusterIngress), &v1alpha1.MultiClusterIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterIngress), err
}

// Update takes the representation of a multiClusterIngress and updates it. Returns the server's representation of the multiClusterIngress, and an error, if there is any.
func (c *FakeMultiClusterIngresses) Update(ctx context.Context, multiClusterIngress *v1alpha1.MultiClusterIngress, opts v1.UpdateOptions) (result *v1alpha1.MultiClusterIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(multiclusteringressesResource, c.ns, multiClusterIngress), &v1alpha1.MultiClusterIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterIngress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMultiClusterIngresses) UpdateStatus(ctx context.Context, multiClusterIngress *v1alpha1.MultiClusterIngress, opts v1.UpdateOptions) (*v1alpha1.MultiClusterIngress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(multiclusteringressesResource, "status", c.ns, multiClusterIngress), &v1alpha1.MultiClusterIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterIngress), err
}

// Delete takes name of the multiClusterIngress and deletes it. Returns an error if one occurs.
func (c *FakeMultiClusterIngresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(multiclusteringressesResource, c.ns, name, opts), &v1alpha1.MultiClusterIngress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMultiClusterIngresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(multiclusteringressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MultiClusterIngressList{})
	return err
}

// Patch applies the patch and returns the patched multiClusterIngress.
func (c *FakeMultiClusterIngresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiClusterIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(multiclusteringressesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MultiClusterIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterIngress), err
}
