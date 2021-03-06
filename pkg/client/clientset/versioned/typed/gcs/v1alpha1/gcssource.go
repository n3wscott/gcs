/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/vaikas-google/gcs/pkg/apis/gcs/v1alpha1"
	scheme "github.com/vaikas-google/gcs/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GCSSourcesGetter has a method to return a GCSSourceInterface.
// A group's client should implement this interface.
type GCSSourcesGetter interface {
	GCSSources(namespace string) GCSSourceInterface
}

// GCSSourceInterface has methods to work with GCSSource resources.
type GCSSourceInterface interface {
	Create(*v1alpha1.GCSSource) (*v1alpha1.GCSSource, error)
	Update(*v1alpha1.GCSSource) (*v1alpha1.GCSSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GCSSource, error)
	List(opts v1.ListOptions) (*v1alpha1.GCSSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GCSSource, err error)
	GCSSourceExpansion
}

// gCSSources implements GCSSourceInterface
type gCSSources struct {
	client rest.Interface
	ns     string
}

// newGCSSources returns a GCSSources
func newGCSSources(c *SourcesV1alpha1Client, namespace string) *gCSSources {
	return &gCSSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gCSSource, and returns the corresponding gCSSource object, and an error if there is any.
func (c *gCSSources) Get(name string, options v1.GetOptions) (result *v1alpha1.GCSSource, err error) {
	result = &v1alpha1.GCSSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gcssources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GCSSources that match those selectors.
func (c *gCSSources) List(opts v1.ListOptions) (result *v1alpha1.GCSSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GCSSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gcssources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gCSSources.
func (c *gCSSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gcssources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a gCSSource and creates it.  Returns the server's representation of the gCSSource, and an error, if there is any.
func (c *gCSSources) Create(gCSSource *v1alpha1.GCSSource) (result *v1alpha1.GCSSource, err error) {
	result = &v1alpha1.GCSSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gcssources").
		Body(gCSSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gCSSource and updates it. Returns the server's representation of the gCSSource, and an error, if there is any.
func (c *gCSSources) Update(gCSSource *v1alpha1.GCSSource) (result *v1alpha1.GCSSource, err error) {
	result = &v1alpha1.GCSSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gcssources").
		Name(gCSSource.Name).
		Body(gCSSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the gCSSource and deletes it. Returns an error if one occurs.
func (c *gCSSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gcssources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gCSSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gcssources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gCSSource.
func (c *gCSSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GCSSource, err error) {
	result = &v1alpha1.GCSSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gcssources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
