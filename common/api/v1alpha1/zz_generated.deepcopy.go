// +build !ignore_autogenerated

/*
Copyright 2021 Google LLC

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialSpec) DeepCopyInto(out *CredentialSpec) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.SecretReference)
		**out = **in
	}
	if in.GsmSecretRef != nil {
		in, out := &in.GsmSecretRef, &out.GsmSecretRef
		*out = new(GsmSecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialSpec.
func (in *CredentialSpec) DeepCopy() *CredentialSpec {
	if in == nil {
		return nil
	}
	out := new(CredentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSpec) DeepCopyInto(out *DatabaseSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSpec.
func (in *DatabaseSpec) DeepCopy() *DatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseStatus) DeepCopyInto(out *DatabaseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseStatus.
func (in *DatabaseStatus) DeepCopy() *DatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSpec) DeepCopyInto(out *DiskSpec) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(DiskType)
		**out = **in
	}
	out.Size = in.Size.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSpec.
func (in *DiskSpec) DeepCopy() *DiskSpec {
	if in == nil {
		return nil
	}
	out := new(DiskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GsmSecretReference) DeepCopyInto(out *GsmSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GsmSecretReference.
func (in *GsmSecretReference) DeepCopy() *GsmSecretReference {
	if in == nil {
		return nil
	}
	out := new(GsmSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]DiskSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SourceCidrRanges != nil {
		in, out := &in.SourceCidrRanges, &out.SourceCidrRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Patching != nil {
		in, out := &in.Patching, &out.Patching
		*out = new(PatchingSpec)
		**out = **in
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make(map[Service]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindowSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindowSpec) DeepCopyInto(out *MaintenanceWindowSpec) {
	*out = *in
	if in.TimeRanges != nil {
		in, out := &in.TimeRanges, &out.TimeRanges
		*out = make([]TimeRange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindowSpec.
func (in *MaintenanceWindowSpec) DeepCopy() *MaintenanceWindowSpec {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeRange) DeepCopyInto(out *TimeRange) {
	*out = *in
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = (*in).DeepCopy()
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeRange.
func (in *TimeRange) DeepCopy() *TimeRange {
	if in == nil {
		return nil
	}
	out := new(TimeRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpec) DeepCopyInto(out *UserSpec) {
	*out = *in
	in.CredentialSpec.DeepCopyInto(&out.CredentialSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpec.
func (in *UserSpec) DeepCopy() *UserSpec {
	if in == nil {
		return nil
	}
	out := new(UserSpec)
	in.DeepCopyInto(out)
	return out
}
