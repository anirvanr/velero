/*
Copyright the Velero contributors.

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

package v1

import (
	v1 "github.com/heptio/velero/pkg/apis/ark/v1"
	scheme "github.com/heptio/velero/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumeSnapshotLocationsGetter has a method to return a VolumeSnapshotLocationInterface.
// A group's client should implement this interface.
type VolumeSnapshotLocationsGetter interface {
	VolumeSnapshotLocations(namespace string) VolumeSnapshotLocationInterface
}

// VolumeSnapshotLocationInterface has methods to work with VolumeSnapshotLocation resources.
type VolumeSnapshotLocationInterface interface {
	Create(*v1.VolumeSnapshotLocation) (*v1.VolumeSnapshotLocation, error)
	Update(*v1.VolumeSnapshotLocation) (*v1.VolumeSnapshotLocation, error)
	UpdateStatus(*v1.VolumeSnapshotLocation) (*v1.VolumeSnapshotLocation, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.VolumeSnapshotLocation, error)
	List(opts metav1.ListOptions) (*v1.VolumeSnapshotLocationList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.VolumeSnapshotLocation, err error)
	VolumeSnapshotLocationExpansion
}

// volumeSnapshotLocations implements VolumeSnapshotLocationInterface
type volumeSnapshotLocations struct {
	client rest.Interface
	ns     string
}

// newVolumeSnapshotLocations returns a VolumeSnapshotLocations
func newVolumeSnapshotLocations(c *ArkV1Client, namespace string) *volumeSnapshotLocations {
	return &volumeSnapshotLocations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the volumeSnapshotLocation, and returns the corresponding volumeSnapshotLocation object, and an error if there is any.
func (c *volumeSnapshotLocations) Get(name string, options metav1.GetOptions) (result *v1.VolumeSnapshotLocation, err error) {
	result = &v1.VolumeSnapshotLocation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumesnapshotlocations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeSnapshotLocations that match those selectors.
func (c *volumeSnapshotLocations) List(opts metav1.ListOptions) (result *v1.VolumeSnapshotLocationList, err error) {
	result = &v1.VolumeSnapshotLocationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumesnapshotlocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeSnapshotLocations.
func (c *volumeSnapshotLocations) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("volumesnapshotlocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a volumeSnapshotLocation and creates it.  Returns the server's representation of the volumeSnapshotLocation, and an error, if there is any.
func (c *volumeSnapshotLocations) Create(volumeSnapshotLocation *v1.VolumeSnapshotLocation) (result *v1.VolumeSnapshotLocation, err error) {
	result = &v1.VolumeSnapshotLocation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("volumesnapshotlocations").
		Body(volumeSnapshotLocation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a volumeSnapshotLocation and updates it. Returns the server's representation of the volumeSnapshotLocation, and an error, if there is any.
func (c *volumeSnapshotLocations) Update(volumeSnapshotLocation *v1.VolumeSnapshotLocation) (result *v1.VolumeSnapshotLocation, err error) {
	result = &v1.VolumeSnapshotLocation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumesnapshotlocations").
		Name(volumeSnapshotLocation.Name).
		Body(volumeSnapshotLocation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *volumeSnapshotLocations) UpdateStatus(volumeSnapshotLocation *v1.VolumeSnapshotLocation) (result *v1.VolumeSnapshotLocation, err error) {
	result = &v1.VolumeSnapshotLocation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumesnapshotlocations").
		Name(volumeSnapshotLocation.Name).
		SubResource("status").
		Body(volumeSnapshotLocation).
		Do().
		Into(result)
	return
}

// Delete takes name of the volumeSnapshotLocation and deletes it. Returns an error if one occurs.
func (c *volumeSnapshotLocations) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumesnapshotlocations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeSnapshotLocations) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumesnapshotlocations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched volumeSnapshotLocation.
func (c *volumeSnapshotLocations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.VolumeSnapshotLocation, err error) {
	result = &v1.VolumeSnapshotLocation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("volumesnapshotlocations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
