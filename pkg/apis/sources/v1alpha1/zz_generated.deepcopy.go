//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockchainSource) DeepCopyInto(out *BlockchainSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockchainSource.
func (in *BlockchainSource) DeepCopy() *BlockchainSource {
	if in == nil {
		return nil
	}
	out := new(BlockchainSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlockchainSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockchainSourceList) DeepCopyInto(out *BlockchainSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BlockchainSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockchainSourceList.
func (in *BlockchainSourceList) DeepCopy() *BlockchainSourceList {
	if in == nil {
		return nil
	}
	out := new(BlockchainSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlockchainSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockchainSourceSpec) DeepCopyInto(out *BlockchainSourceSpec) {
	*out = *in
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.AccessToken.DeepCopyInto(&out.AccessToken)
	in.SecretToken.DeepCopyInto(&out.SecretToken)
	if in.Secure != nil {
		in, out := &in.Secure, &out.Secure
		*out = new(bool)
		**out = **in
	}
	in.SourceSpec.DeepCopyInto(&out.SourceSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockchainSourceSpec.
func (in *BlockchainSourceSpec) DeepCopy() *BlockchainSourceSpec {
	if in == nil {
		return nil
	}
	out := new(BlockchainSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockchainSourceStatus) DeepCopyInto(out *BlockchainSourceStatus) {
	*out = *in
	in.SourceStatus.DeepCopyInto(&out.SourceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockchainSourceStatus.
func (in *BlockchainSourceStatus) DeepCopy() *BlockchainSourceStatus {
	if in == nil {
		return nil
	}
	out := new(BlockchainSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretValueFromSource) DeepCopyInto(out *SecretValueFromSource) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretValueFromSource.
func (in *SecretValueFromSource) DeepCopy() *SecretValueFromSource {
	if in == nil {
		return nil
	}
	out := new(SecretValueFromSource)
	in.DeepCopyInto(out)
	return out
}