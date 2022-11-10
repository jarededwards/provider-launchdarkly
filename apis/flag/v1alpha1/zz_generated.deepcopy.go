//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientSideAvailabilityObservation) DeepCopyInto(out *ClientSideAvailabilityObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientSideAvailabilityObservation.
func (in *ClientSideAvailabilityObservation) DeepCopy() *ClientSideAvailabilityObservation {
	if in == nil {
		return nil
	}
	out := new(ClientSideAvailabilityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientSideAvailabilityParameters) DeepCopyInto(out *ClientSideAvailabilityParameters) {
	*out = *in
	if in.UsingEnvironmentID != nil {
		in, out := &in.UsingEnvironmentID, &out.UsingEnvironmentID
		*out = new(bool)
		**out = **in
	}
	if in.UsingMobileKey != nil {
		in, out := &in.UsingMobileKey, &out.UsingMobileKey
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientSideAvailabilityParameters.
func (in *ClientSideAvailabilityParameters) DeepCopy() *ClientSideAvailabilityParameters {
	if in == nil {
		return nil
	}
	out := new(ClientSideAvailabilityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomPropertiesObservation) DeepCopyInto(out *CustomPropertiesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomPropertiesObservation.
func (in *CustomPropertiesObservation) DeepCopy() *CustomPropertiesObservation {
	if in == nil {
		return nil
	}
	out := new(CustomPropertiesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomPropertiesParameters) DeepCopyInto(out *CustomPropertiesParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomPropertiesParameters.
func (in *CustomPropertiesParameters) DeepCopy() *CustomPropertiesParameters {
	if in == nil {
		return nil
	}
	out := new(CustomPropertiesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultsObservation) DeepCopyInto(out *DefaultsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultsObservation.
func (in *DefaultsObservation) DeepCopy() *DefaultsObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultsParameters) DeepCopyInto(out *DefaultsParameters) {
	*out = *in
	if in.OffVariation != nil {
		in, out := &in.OffVariation, &out.OffVariation
		*out = new(float64)
		**out = **in
	}
	if in.OnVariation != nil {
		in, out := &in.OnVariation, &out.OnVariation
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultsParameters.
func (in *DefaultsParameters) DeepCopy() *DefaultsParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flag) DeepCopyInto(out *Flag) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flag.
func (in *Flag) DeepCopy() *Flag {
	if in == nil {
		return nil
	}
	out := new(Flag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Flag) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagList) DeepCopyInto(out *FlagList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Flag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagList.
func (in *FlagList) DeepCopy() *FlagList {
	if in == nil {
		return nil
	}
	out := new(FlagList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlagList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagObservation) DeepCopyInto(out *FlagObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagObservation.
func (in *FlagObservation) DeepCopy() *FlagObservation {
	if in == nil {
		return nil
	}
	out := new(FlagObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagParameters) DeepCopyInto(out *FlagParameters) {
	*out = *in
	if in.Archived != nil {
		in, out := &in.Archived, &out.Archived
		*out = new(bool)
		**out = **in
	}
	if in.ClientSideAvailability != nil {
		in, out := &in.ClientSideAvailability, &out.ClientSideAvailability
		*out = make([]ClientSideAvailabilityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomProperties != nil {
		in, out := &in.CustomProperties, &out.CustomProperties
		*out = make([]CustomPropertiesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Defaults != nil {
		in, out := &in.Defaults, &out.Defaults
		*out = make([]DefaultsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IncludeInSnippet != nil {
		in, out := &in.IncludeInSnippet, &out.IncludeInSnippet
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.MaintainerID != nil {
		in, out := &in.MaintainerID, &out.MaintainerID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProjectKey != nil {
		in, out := &in.ProjectKey, &out.ProjectKey
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Temporary != nil {
		in, out := &in.Temporary, &out.Temporary
		*out = new(bool)
		**out = **in
	}
	if in.VariationType != nil {
		in, out := &in.VariationType, &out.VariationType
		*out = new(string)
		**out = **in
	}
	if in.Variations != nil {
		in, out := &in.Variations, &out.Variations
		*out = make([]VariationsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagParameters.
func (in *FlagParameters) DeepCopy() *FlagParameters {
	if in == nil {
		return nil
	}
	out := new(FlagParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagSpec) DeepCopyInto(out *FlagSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagSpec.
func (in *FlagSpec) DeepCopy() *FlagSpec {
	if in == nil {
		return nil
	}
	out := new(FlagSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagStatus) DeepCopyInto(out *FlagStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagStatus.
func (in *FlagStatus) DeepCopy() *FlagStatus {
	if in == nil {
		return nil
	}
	out := new(FlagStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VariationsObservation) DeepCopyInto(out *VariationsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VariationsObservation.
func (in *VariationsObservation) DeepCopy() *VariationsObservation {
	if in == nil {
		return nil
	}
	out := new(VariationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VariationsParameters) DeepCopyInto(out *VariationsParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VariationsParameters.
func (in *VariationsParameters) DeepCopy() *VariationsParameters {
	if in == nil {
		return nil
	}
	out := new(VariationsParameters)
	in.DeepCopyInto(out)
	return out
}