/*
Copyright 2022 The Knative Authors

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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/eventing-blockchain/pkg/apis/sources/v1alpha1"
	scheme "knative.dev/eventing-blockchain/pkg/client/clientset/versioned/scheme"
)

// GitHubSourcesGetter has a method to return a GitHubSourceInterface.
// A group's client should implement this interface.
type GitHubSourcesGetter interface {
	GitHubSources(namespace string) GitHubSourceInterface
}

// GitHubSourceInterface has methods to work with GitHubSource resources.
type GitHubSourceInterface interface {
	Create(ctx context.Context, gitHubSource *v1alpha1.GitHubSource, opts v1.CreateOptions) (*v1alpha1.GitHubSource, error)
	Update(ctx context.Context, gitHubSource *v1alpha1.GitHubSource, opts v1.UpdateOptions) (*v1alpha1.GitHubSource, error)
	UpdateStatus(ctx context.Context, gitHubSource *v1alpha1.GitHubSource, opts v1.UpdateOptions) (*v1alpha1.GitHubSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.GitHubSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.GitHubSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GitHubSource, err error)
	GitHubSourceExpansion
}

// gitHubSources implements GitHubSourceInterface
type gitHubSources struct {
	client rest.Interface
	ns     string
}

// newGitHubSources returns a GitHubSources
func newGitHubSources(c *SourcesV1alpha1Client, namespace string) *gitHubSources {
	return &gitHubSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gitHubSource, and returns the corresponding gitHubSource object, and an error if there is any.
func (c *gitHubSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GitHubSource, err error) {
	result = &v1alpha1.GitHubSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("githubsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GitHubSources that match those selectors.
func (c *gitHubSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GitHubSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GitHubSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("githubsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gitHubSources.
func (c *gitHubSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("githubsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gitHubSource and creates it.  Returns the server's representation of the gitHubSource, and an error, if there is any.
func (c *gitHubSources) Create(ctx context.Context, gitHubSource *v1alpha1.GitHubSource, opts v1.CreateOptions) (result *v1alpha1.GitHubSource, err error) {
	result = &v1alpha1.GitHubSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("githubsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gitHubSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gitHubSource and updates it. Returns the server's representation of the gitHubSource, and an error, if there is any.
func (c *gitHubSources) Update(ctx context.Context, gitHubSource *v1alpha1.GitHubSource, opts v1.UpdateOptions) (result *v1alpha1.GitHubSource, err error) {
	result = &v1alpha1.GitHubSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("githubsources").
		Name(gitHubSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gitHubSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *gitHubSources) UpdateStatus(ctx context.Context, gitHubSource *v1alpha1.GitHubSource, opts v1.UpdateOptions) (result *v1alpha1.GitHubSource, err error) {
	result = &v1alpha1.GitHubSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("githubsources").
		Name(gitHubSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gitHubSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gitHubSource and deletes it. Returns an error if one occurs.
func (c *gitHubSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("githubsources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gitHubSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("githubsources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gitHubSource.
func (c *gitHubSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GitHubSource, err error) {
	result = &v1alpha1.GitHubSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("githubsources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
