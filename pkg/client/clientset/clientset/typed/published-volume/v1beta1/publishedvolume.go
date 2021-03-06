/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	"time"

	v1beta1 "github.com/ofek/csi-gcs/pkg/apis/published-volume/v1beta1"
	scheme "github.com/ofek/csi-gcs/pkg/client/clientset/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PublishedVolumesGetter has a method to return a PublishedVolumeInterface.
// A group's client should implement this interface.
type PublishedVolumesGetter interface {
	PublishedVolumes() PublishedVolumeInterface
}

// PublishedVolumeInterface has methods to work with PublishedVolume resources.
type PublishedVolumeInterface interface {
	Create(*v1beta1.PublishedVolume) (*v1beta1.PublishedVolume, error)
	Update(*v1beta1.PublishedVolume) (*v1beta1.PublishedVolume, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.PublishedVolume, error)
	List(opts v1.ListOptions) (*v1beta1.PublishedVolumeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.PublishedVolume, err error)
	PublishedVolumeExpansion
}

// publishedVolumes implements PublishedVolumeInterface
type publishedVolumes struct {
	client rest.Interface
}

// newPublishedVolumes returns a PublishedVolumes
func newPublishedVolumes(c *GcsV1beta1Client) *publishedVolumes {
	return &publishedVolumes{
		client: c.RESTClient(),
	}
}

// Get takes name of the publishedVolume, and returns the corresponding publishedVolume object, and an error if there is any.
func (c *publishedVolumes) Get(name string, options v1.GetOptions) (result *v1beta1.PublishedVolume, err error) {
	result = &v1beta1.PublishedVolume{}
	err = c.client.Get().
		Resource("publishedvolumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PublishedVolumes that match those selectors.
func (c *publishedVolumes) List(opts v1.ListOptions) (result *v1beta1.PublishedVolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.PublishedVolumeList{}
	err = c.client.Get().
		Resource("publishedvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested publishedVolumes.
func (c *publishedVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("publishedvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a publishedVolume and creates it.  Returns the server's representation of the publishedVolume, and an error, if there is any.
func (c *publishedVolumes) Create(publishedVolume *v1beta1.PublishedVolume) (result *v1beta1.PublishedVolume, err error) {
	result = &v1beta1.PublishedVolume{}
	err = c.client.Post().
		Resource("publishedvolumes").
		Body(publishedVolume).
		Do().
		Into(result)
	return
}

// Update takes the representation of a publishedVolume and updates it. Returns the server's representation of the publishedVolume, and an error, if there is any.
func (c *publishedVolumes) Update(publishedVolume *v1beta1.PublishedVolume) (result *v1beta1.PublishedVolume, err error) {
	result = &v1beta1.PublishedVolume{}
	err = c.client.Put().
		Resource("publishedvolumes").
		Name(publishedVolume.Name).
		Body(publishedVolume).
		Do().
		Into(result)
	return
}

// Delete takes name of the publishedVolume and deletes it. Returns an error if one occurs.
func (c *publishedVolumes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("publishedvolumes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *publishedVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("publishedvolumes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched publishedVolume.
func (c *publishedVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.PublishedVolume, err error) {
	result = &v1beta1.PublishedVolume{}
	err = c.client.Patch(pt).
		Resource("publishedvolumes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
