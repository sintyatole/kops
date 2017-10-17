/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	kops "k8s.io/kops/pkg/apis/kops"
)

// FakeSSHCredentials implements SSHCredentialInterface
type FakeSSHCredentials struct {
	Fake *FakeKops
	ns   string
}

var sshcredentialsResource = schema.GroupVersionResource{Group: "kops", Version: "", Resource: "sshcredentials"}

var sshcredentialsKind = schema.GroupVersionKind{Group: "kops", Version: "", Kind: "SSHCredential"}

// Get takes name of the sSHCredential, and returns the corresponding sSHCredential object, and an error if there is any.
func (c *FakeSSHCredentials) Get(name string, options v1.GetOptions) (result *kops.SSHCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sshcredentialsResource, c.ns, name), &kops.SSHCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kops.SSHCredential), err
}

// List takes label and field selectors, and returns the list of SSHCredentials that match those selectors.
func (c *FakeSSHCredentials) List(opts v1.ListOptions) (result *kops.SSHCredentialList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sshcredentialsResource, sshcredentialsKind, c.ns, opts), &kops.SSHCredentialList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kops.SSHCredentialList{}
	for _, item := range obj.(*kops.SSHCredentialList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sSHCredentials.
func (c *FakeSSHCredentials) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sshcredentialsResource, c.ns, opts))

}

// Create takes the representation of a sSHCredential and creates it.  Returns the server's representation of the sSHCredential, and an error, if there is any.
func (c *FakeSSHCredentials) Create(sSHCredential *kops.SSHCredential) (result *kops.SSHCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sshcredentialsResource, c.ns, sSHCredential), &kops.SSHCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kops.SSHCredential), err
}

// Update takes the representation of a sSHCredential and updates it. Returns the server's representation of the sSHCredential, and an error, if there is any.
func (c *FakeSSHCredentials) Update(sSHCredential *kops.SSHCredential) (result *kops.SSHCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sshcredentialsResource, c.ns, sSHCredential), &kops.SSHCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kops.SSHCredential), err
}

// Delete takes name of the sSHCredential and deletes it. Returns an error if one occurs.
func (c *FakeSSHCredentials) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sshcredentialsResource, c.ns, name), &kops.SSHCredential{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSSHCredentials) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sshcredentialsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &kops.SSHCredentialList{})
	return err
}

// Patch applies the patch and returns the patched sSHCredential.
func (c *FakeSSHCredentials) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *kops.SSHCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sshcredentialsResource, c.ns, name, data, subresources...), &kops.SSHCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kops.SSHCredential), err
}