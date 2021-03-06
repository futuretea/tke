/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	platform "tkestack.io/tke/api/platform"
)

// LBCFsGetter has a method to return a LBCFInterface.
// A group's client should implement this interface.
type LBCFsGetter interface {
	LBCFs() LBCFInterface
}

// LBCFInterface has methods to work with LBCF resources.
type LBCFInterface interface {
	Create(ctx context.Context, lBCF *platform.LBCF, opts v1.CreateOptions) (*platform.LBCF, error)
	Update(ctx context.Context, lBCF *platform.LBCF, opts v1.UpdateOptions) (*platform.LBCF, error)
	UpdateStatus(ctx context.Context, lBCF *platform.LBCF, opts v1.UpdateOptions) (*platform.LBCF, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*platform.LBCF, error)
	List(ctx context.Context, opts v1.ListOptions) (*platform.LBCFList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *platform.LBCF, err error)
	LBCFExpansion
}

// lBCFs implements LBCFInterface
type lBCFs struct {
	client rest.Interface
}

// newLBCFs returns a LBCFs
func newLBCFs(c *PlatformClient) *lBCFs {
	return &lBCFs{
		client: c.RESTClient(),
	}
}

// Get takes name of the lBCF, and returns the corresponding lBCF object, and an error if there is any.
func (c *lBCFs) Get(ctx context.Context, name string, options v1.GetOptions) (result *platform.LBCF, err error) {
	result = &platform.LBCF{}
	err = c.client.Get().
		Resource("lbcfs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LBCFs that match those selectors.
func (c *lBCFs) List(ctx context.Context, opts v1.ListOptions) (result *platform.LBCFList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &platform.LBCFList{}
	err = c.client.Get().
		Resource("lbcfs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lBCFs.
func (c *lBCFs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("lbcfs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a lBCF and creates it.  Returns the server's representation of the lBCF, and an error, if there is any.
func (c *lBCFs) Create(ctx context.Context, lBCF *platform.LBCF, opts v1.CreateOptions) (result *platform.LBCF, err error) {
	result = &platform.LBCF{}
	err = c.client.Post().
		Resource("lbcfs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lBCF).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a lBCF and updates it. Returns the server's representation of the lBCF, and an error, if there is any.
func (c *lBCFs) Update(ctx context.Context, lBCF *platform.LBCF, opts v1.UpdateOptions) (result *platform.LBCF, err error) {
	result = &platform.LBCF{}
	err = c.client.Put().
		Resource("lbcfs").
		Name(lBCF.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lBCF).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *lBCFs) UpdateStatus(ctx context.Context, lBCF *platform.LBCF, opts v1.UpdateOptions) (result *platform.LBCF, err error) {
	result = &platform.LBCF{}
	err = c.client.Put().
		Resource("lbcfs").
		Name(lBCF.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lBCF).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the lBCF and deletes it. Returns an error if one occurs.
func (c *lBCFs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("lbcfs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched lBCF.
func (c *lBCFs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *platform.LBCF, err error) {
	result = &platform.LBCF{}
	err = c.client.Patch(pt).
		Resource("lbcfs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
