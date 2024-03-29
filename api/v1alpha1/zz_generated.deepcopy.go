// +build !ignore_autogenerated

/*


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
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerDeployment) DeepCopyInto(out *CertManagerDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerDeployment.
func (in *CertManagerDeployment) DeepCopy() *CertManagerDeployment {
	if in == nil {
		return nil
	}
	out := new(CertManagerDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertManagerDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerDeploymentCondition) DeepCopyInto(out *CertManagerDeploymentCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerDeploymentCondition.
func (in *CertManagerDeploymentCondition) DeepCopy() *CertManagerDeploymentCondition {
	if in == nil {
		return nil
	}
	out := new(CertManagerDeploymentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerDeploymentList) DeepCopyInto(out *CertManagerDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertManagerDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerDeploymentList.
func (in *CertManagerDeploymentList) DeepCopy() *CertManagerDeploymentList {
	if in == nil {
		return nil
	}
	out := new(CertManagerDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertManagerDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerDeploymentSpec) DeepCopyInto(out *CertManagerDeploymentSpec) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	in.DangerZone.DeepCopyInto(&out.DangerZone)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerDeploymentSpec.
func (in *CertManagerDeploymentSpec) DeepCopy() *CertManagerDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(CertManagerDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerDeploymentStatus) DeepCopyInto(out *CertManagerDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CertManagerDeploymentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeploymentConditions != nil {
		in, out := &in.DeploymentConditions, &out.DeploymentConditions
		*out = make([]ManagedDeploymentWithConditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CRDConditions != nil {
		in, out := &in.CRDConditions, &out.CRDConditions
		*out = make([]ManagedCRDWithConditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerDeploymentStatus.
func (in *CertManagerDeploymentStatus) DeepCopy() *CertManagerDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(CertManagerDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerArgOverrides) DeepCopyInto(out *ContainerArgOverrides) {
	*out = *in
	in.Controller.DeepCopyInto(&out.Controller)
	in.Webhook.DeepCopyInto(&out.Webhook)
	in.CAInjector.DeepCopyInto(&out.CAInjector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerArgOverrides.
func (in *ContainerArgOverrides) DeepCopy() *ContainerArgOverrides {
	if in == nil {
		return nil
	}
	out := new(ContainerArgOverrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DangerZone) DeepCopyInto(out *DangerZone) {
	*out = *in
	if in.ImageOverrides != nil {
		in, out := &in.ImageOverrides, &out.ImageOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.ContainerArgOverrides.DeepCopyInto(&out.ContainerArgOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DangerZone.
func (in *DangerZone) DeepCopy() *DangerZone {
	if in == nil {
		return nil
	}
	out := new(DangerZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedCRDWithConditions) DeepCopyInto(out *ManagedCRDWithConditions) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.CustomResourceDefinitionCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedCRDWithConditions.
func (in *ManagedCRDWithConditions) DeepCopy() *ManagedCRDWithConditions {
	if in == nil {
		return nil
	}
	out := new(ManagedCRDWithConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedDeploymentWithConditions) DeepCopyInto(out *ManagedDeploymentWithConditions) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]appsv1.DeploymentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedDeploymentWithConditions.
func (in *ManagedDeploymentWithConditions) DeepCopy() *ManagedDeploymentWithConditions {
	if in == nil {
		return nil
	}
	out := new(ManagedDeploymentWithConditions)
	in.DeepCopyInto(out)
	return out
}
