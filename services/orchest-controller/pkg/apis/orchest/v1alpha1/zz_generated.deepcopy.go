//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The orchest Authors.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationConfig) DeepCopyInto(out *ApplicationConfig) {
	*out = *in
	if in.Helm != nil {
		in, out := &in.Helm, &out.Helm
		*out = new(ApplicationConfigHelm)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationConfig.
func (in *ApplicationConfig) DeepCopy() *ApplicationConfig {
	if in == nil {
		return nil
	}
	out := new(ApplicationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationConfigHelm) DeepCopyInto(out *ApplicationConfigHelm) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]HelmParameter, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationConfigHelm.
func (in *ApplicationConfigHelm) DeepCopy() *ApplicationConfigHelm {
	if in == nil {
		return nil
	}
	out := new(ApplicationConfigHelm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	if in.Needs != nil {
		in, out := &in.Needs, &out.Needs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmParameter) DeepCopyInto(out *HelmParameter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmParameter.
func (in *HelmParameter) DeepCopy() *HelmParameter {
	if in == nil {
		return nil
	}
	out := new(HelmParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestCluster) DeepCopyInto(out *OrchestCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(OrchestClusterStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestCluster.
func (in *OrchestCluster) DeepCopy() *OrchestCluster {
	if in == nil {
		return nil
	}
	out := new(OrchestCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrchestCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestClusterList) DeepCopyInto(out *OrchestClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*OrchestCluster, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(OrchestCluster)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestClusterList.
func (in *OrchestClusterList) DeepCopy() *OrchestClusterList {
	if in == nil {
		return nil
	}
	out := new(OrchestClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrchestClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestClusterSpec) DeepCopyInto(out *OrchestClusterSpec) {
	*out = *in
	if in.SingleNode != nil {
		in, out := &in.SingleNode, &out.SingleNode
		*out = new(bool)
		**out = **in
	}
	in.Orchest.DeepCopyInto(&out.Orchest)
	in.Postgres.DeepCopyInto(&out.Postgres)
	in.RabbitMq.DeepCopyInto(&out.RabbitMq)
	if in.Applications != nil {
		in, out := &in.Applications, &out.Applications
		*out = make([]ApplicationSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ControlNodeSelector != nil {
		in, out := &in.ControlNodeSelector, &out.ControlNodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.WorkerNodeSelector != nil {
		in, out := &in.WorkerNodeSelector, &out.WorkerNodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestClusterSpec.
func (in *OrchestClusterSpec) DeepCopy() *OrchestClusterSpec {
	if in == nil {
		return nil
	}
	out := new(OrchestClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestClusterStatus) DeepCopyInto(out *OrchestClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestClusterStatus.
func (in *OrchestClusterStatus) DeepCopy() *OrchestClusterStatus {
	if in == nil {
		return nil
	}
	out := new(OrchestClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestComponent) DeepCopyInto(out *OrchestComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(OrchestComponentStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestComponent.
func (in *OrchestComponent) DeepCopy() *OrchestComponent {
	if in == nil {
		return nil
	}
	out := new(OrchestComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrchestComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestComponentList) DeepCopyInto(out *OrchestComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*OrchestComponent, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(OrchestComponent)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestComponentList.
func (in *OrchestComponentList) DeepCopy() *OrchestComponentList {
	if in == nil {
		return nil
	}
	out := new(OrchestComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrchestComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestComponentSpec) DeepCopyInto(out *OrchestComponentSpec) {
	*out = *in
	if in.OrchestHost != nil {
		in, out := &in.OrchestHost, &out.OrchestHost
		*out = new(string)
		**out = **in
	}
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestComponentSpec.
func (in *OrchestComponentSpec) DeepCopy() *OrchestComponentSpec {
	if in == nil {
		return nil
	}
	out := new(OrchestComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestComponentStatus) DeepCopyInto(out *OrchestComponentStatus) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestComponentStatus.
func (in *OrchestComponentStatus) DeepCopy() *OrchestComponentStatus {
	if in == nil {
		return nil
	}
	out := new(OrchestComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestComponentTemplate) DeepCopyInto(out *OrchestComponentTemplate) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestComponentTemplate.
func (in *OrchestComponentTemplate) DeepCopy() *OrchestComponentTemplate {
	if in == nil {
		return nil
	}
	out := new(OrchestComponentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestResourcesSpec) DeepCopyInto(out *OrchestResourcesSpec) {
	*out = *in
	if in.UserDirVolume != nil {
		in, out := &in.UserDirVolume, &out.UserDirVolume
		*out = make([]Volume, len(*in))
		copy(*out, *in)
	}
	if in.OrchestStateVolume != nil {
		in, out := &in.OrchestStateVolume, &out.OrchestStateVolume
		*out = make([]Volume, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestResourcesSpec.
func (in *OrchestResourcesSpec) DeepCopy() *OrchestResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(OrchestResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestSpec) DeepCopyInto(out *OrchestSpec) {
	*out = *in
	if in.Pause != nil {
		in, out := &in.Pause, &out.Pause
		*out = new(bool)
		**out = **in
	}
	if in.OrchestHost != nil {
		in, out := &in.OrchestHost, &out.OrchestHost
		*out = new(string)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.OrchestApi.DeepCopyInto(&out.OrchestApi)
	in.OrchestWebServer.DeepCopyInto(&out.OrchestWebServer)
	in.CeleryWorker.DeepCopyInto(&out.CeleryWorker)
	in.NodeAgent.DeepCopyInto(&out.NodeAgent)
	in.BuildKitDaemon.DeepCopyInto(&out.BuildKitDaemon)
	in.AuthServer.DeepCopyInto(&out.AuthServer)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestSpec.
func (in *OrchestSpec) DeepCopy() *OrchestSpec {
	if in == nil {
		return nil
	}
	out := new(OrchestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}
