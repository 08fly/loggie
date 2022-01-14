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
	"context"
	"time"

	v1beta1 "github.com/loggie-io/loggie/pkg/discovery/kubernetes/apis/loggie/v1beta1"
	scheme "github.com/loggie-io/loggie/pkg/discovery/kubernetes/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SinksGetter has a method to return a SinkInterface.
// A group's client should implement this interface.
type SinksGetter interface {
	Sinks() SinkInterface
}

// SinkInterface has methods to work with Sink resources.
type SinkInterface interface {
	Create(ctx context.Context, sink *v1beta1.Sink, opts v1.CreateOptions) (*v1beta1.Sink, error)
	Update(ctx context.Context, sink *v1beta1.Sink, opts v1.UpdateOptions) (*v1beta1.Sink, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Sink, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.SinkList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Sink, err error)
	SinkExpansion
}

// sinks implements SinkInterface
type sinks struct {
	client rest.Interface
}

// newSinks returns a Sinks
func newSinks(c *LoggieV1beta1Client) *sinks {
	return &sinks{
		client: c.RESTClient(),
	}
}

// Get takes name of the sink, and returns the corresponding sink object, and an error if there is any.
func (c *sinks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Sink, err error) {
	result = &v1beta1.Sink{}
	err = c.client.Get().
		Resource("sinks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Sinks that match those selectors.
func (c *sinks) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SinkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.SinkList{}
	err = c.client.Get().
		Resource("sinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sinks.
func (c *sinks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("sinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sink and creates it.  Returns the server's representation of the sink, and an error, if there is any.
func (c *sinks) Create(ctx context.Context, sink *v1beta1.Sink, opts v1.CreateOptions) (result *v1beta1.Sink, err error) {
	result = &v1beta1.Sink{}
	err = c.client.Post().
		Resource("sinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sink).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sink and updates it. Returns the server's representation of the sink, and an error, if there is any.
func (c *sinks) Update(ctx context.Context, sink *v1beta1.Sink, opts v1.UpdateOptions) (result *v1beta1.Sink, err error) {
	result = &v1beta1.Sink{}
	err = c.client.Put().
		Resource("sinks").
		Name(sink.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sink).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sink and deletes it. Returns an error if one occurs.
func (c *sinks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("sinks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sinks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("sinks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sink.
func (c *sinks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Sink, err error) {
	result = &v1beta1.Sink{}
	err = c.client.Patch(pt).
		Resource("sinks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
