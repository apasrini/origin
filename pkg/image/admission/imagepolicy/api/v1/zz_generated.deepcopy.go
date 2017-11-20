// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GroupResource).DeepCopyInto(out.(*GroupResource))
			return nil
		}, InType: reflect.TypeOf(new(GroupResource))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ImageCondition).DeepCopyInto(out.(*ImageCondition))
			return nil
		}, InType: reflect.TypeOf(new(ImageCondition))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ImageExecutionPolicyRule).DeepCopyInto(out.(*ImageExecutionPolicyRule))
			return nil
		}, InType: reflect.TypeOf(new(ImageExecutionPolicyRule))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ImagePolicyConfig).DeepCopyInto(out.(*ImagePolicyConfig))
			return nil
		}, InType: reflect.TypeOf(new(ImagePolicyConfig))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ImageResolutionPolicyRule).DeepCopyInto(out.(*ImageResolutionPolicyRule))
			return nil
		}, InType: reflect.TypeOf(new(ImageResolutionPolicyRule))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ImageResolutionType).DeepCopyInto(out.(*ImageResolutionType))
			return nil
		}, InType: reflect.TypeOf(new(ImageResolutionType))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ValueCondition).DeepCopyInto(out.(*ValueCondition))
			return nil
		}, InType: reflect.TypeOf(new(ValueCondition))},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupResource) DeepCopyInto(out *GroupResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupResource.
func (in *GroupResource) DeepCopy() *GroupResource {
	if in == nil {
		return nil
	}
	out := new(GroupResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageCondition) DeepCopyInto(out *ImageCondition) {
	*out = *in
	if in.OnResources != nil {
		in, out := &in.OnResources, &out.OnResources
		*out = make([]GroupResource, len(*in))
		copy(*out, *in)
	}
	if in.MatchRegistries != nil {
		in, out := &in.MatchRegistries, &out.MatchRegistries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchDockerImageLabels != nil {
		in, out := &in.MatchDockerImageLabels, &out.MatchDockerImageLabels
		*out = make([]ValueCondition, len(*in))
		copy(*out, *in)
	}
	if in.MatchImageLabels != nil {
		in, out := &in.MatchImageLabels, &out.MatchImageLabels
		*out = make([]meta_v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchImageAnnotations != nil {
		in, out := &in.MatchImageAnnotations, &out.MatchImageAnnotations
		*out = make([]ValueCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageCondition.
func (in *ImageCondition) DeepCopy() *ImageCondition {
	if in == nil {
		return nil
	}
	out := new(ImageCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageExecutionPolicyRule) DeepCopyInto(out *ImageExecutionPolicyRule) {
	*out = *in
	in.ImageCondition.DeepCopyInto(&out.ImageCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageExecutionPolicyRule.
func (in *ImageExecutionPolicyRule) DeepCopy() *ImageExecutionPolicyRule {
	if in == nil {
		return nil
	}
	out := new(ImageExecutionPolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicyConfig) DeepCopyInto(out *ImagePolicyConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ResolutionRules != nil {
		in, out := &in.ResolutionRules, &out.ResolutionRules
		*out = make([]ImageResolutionPolicyRule, len(*in))
		copy(*out, *in)
	}
	if in.ExecutionRules != nil {
		in, out := &in.ExecutionRules, &out.ExecutionRules
		*out = make([]ImageExecutionPolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicyConfig.
func (in *ImagePolicyConfig) DeepCopy() *ImagePolicyConfig {
	if in == nil {
		return nil
	}
	out := new(ImagePolicyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagePolicyConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageResolutionPolicyRule) DeepCopyInto(out *ImageResolutionPolicyRule) {
	*out = *in
	out.TargetResource = in.TargetResource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageResolutionPolicyRule.
func (in *ImageResolutionPolicyRule) DeepCopy() *ImageResolutionPolicyRule {
	if in == nil {
		return nil
	}
	out := new(ImageResolutionPolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageResolutionType) DeepCopyInto(out *ImageResolutionType) {
	{
		in := (*string)(unsafe.Pointer(in))
		out := (*string)(unsafe.Pointer(out))
		*out = *in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageResolutionType.
func (in *ImageResolutionType) DeepCopy() *ImageResolutionType {
	if in == nil {
		return nil
	}
	out := new(ImageResolutionType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueCondition) DeepCopyInto(out *ValueCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueCondition.
func (in *ValueCondition) DeepCopy() *ValueCondition {
	if in == nil {
		return nil
	}
	out := new(ValueCondition)
	in.DeepCopyInto(out)
	return out
}
