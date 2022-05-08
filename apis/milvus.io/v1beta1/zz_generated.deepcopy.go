// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1beta1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	in.ComponentSpec.DeepCopyInto(&out.ComponentSpec)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSpec) DeepCopyInto(out *ComponentSpec) {
	*out = *in
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
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
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpec.
func (in *ComponentSpec) DeepCopy() *ComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InClusterConfig) DeepCopyInto(out *InClusterConfig) {
	*out = *in
	in.Values.DeepCopyInto(&out.Values)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InClusterConfig.
func (in *InClusterConfig) DeepCopy() *InClusterConfig {
	if in == nil {
		return nil
	}
	out := new(InClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Milvus) DeepCopyInto(out *Milvus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Milvus.
func (in *Milvus) DeepCopy() *Milvus {
	if in == nil {
		return nil
	}
	out := new(Milvus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Milvus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusComponents) DeepCopyInto(out *MilvusComponents) {
	*out = *in
	in.ComponentSpec.DeepCopyInto(&out.ComponentSpec)
	in.Proxy.DeepCopyInto(&out.Proxy)
	in.RootCoord.DeepCopyInto(&out.RootCoord)
	in.IndexCoord.DeepCopyInto(&out.IndexCoord)
	in.DataCoord.DeepCopyInto(&out.DataCoord)
	in.QueryCoord.DeepCopyInto(&out.QueryCoord)
	in.IndexNode.DeepCopyInto(&out.IndexNode)
	in.DataNode.DeepCopyInto(&out.DataNode)
	in.QueryNode.DeepCopyInto(&out.QueryNode)
	in.Standalone.DeepCopyInto(&out.Standalone)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusComponents.
func (in *MilvusComponents) DeepCopy() *MilvusComponents {
	if in == nil {
		return nil
	}
	out := new(MilvusComponents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusCondition) DeepCopyInto(out *MilvusCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusCondition.
func (in *MilvusCondition) DeepCopy() *MilvusCondition {
	if in == nil {
		return nil
	}
	out := new(MilvusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusDataCoord) DeepCopyInto(out *MilvusDataCoord) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusDataCoord.
func (in *MilvusDataCoord) DeepCopy() *MilvusDataCoord {
	if in == nil {
		return nil
	}
	out := new(MilvusDataCoord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusDataNode) DeepCopyInto(out *MilvusDataNode) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusDataNode.
func (in *MilvusDataNode) DeepCopy() *MilvusDataNode {
	if in == nil {
		return nil
	}
	out := new(MilvusDataNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusDependencies) DeepCopyInto(out *MilvusDependencies) {
	*out = *in
	in.Etcd.DeepCopyInto(&out.Etcd)
	in.Pulsar.DeepCopyInto(&out.Pulsar)
	in.Kafka.DeepCopyInto(&out.Kafka)
	in.RocksMQ.DeepCopyInto(&out.RocksMQ)
	in.Storage.DeepCopyInto(&out.Storage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusDependencies.
func (in *MilvusDependencies) DeepCopy() *MilvusDependencies {
	if in == nil {
		return nil
	}
	out := new(MilvusDependencies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusEtcd) DeepCopyInto(out *MilvusEtcd) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InCluster != nil {
		in, out := &in.InCluster, &out.InCluster
		*out = new(InClusterConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusEtcd.
func (in *MilvusEtcd) DeepCopy() *MilvusEtcd {
	if in == nil {
		return nil
	}
	out := new(MilvusEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusIndexCoord) DeepCopyInto(out *MilvusIndexCoord) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusIndexCoord.
func (in *MilvusIndexCoord) DeepCopy() *MilvusIndexCoord {
	if in == nil {
		return nil
	}
	out := new(MilvusIndexCoord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusIndexNode) DeepCopyInto(out *MilvusIndexNode) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusIndexNode.
func (in *MilvusIndexNode) DeepCopy() *MilvusIndexNode {
	if in == nil {
		return nil
	}
	out := new(MilvusIndexNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusIngress) DeepCopyInto(out *MilvusIngress) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IngressClassName != nil {
		in, out := &in.IngressClassName, &out.IngressClassName
		*out = new(string)
		**out = **in
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLSSecretRefs != nil {
		in, out := &in.TLSSecretRefs, &out.TLSSecretRefs
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusIngress.
func (in *MilvusIngress) DeepCopy() *MilvusIngress {
	if in == nil {
		return nil
	}
	out := new(MilvusIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusKafka) DeepCopyInto(out *MilvusKafka) {
	*out = *in
	if in.InCluster != nil {
		in, out := &in.InCluster, &out.InCluster
		*out = new(InClusterConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.BrokerList != nil {
		in, out := &in.BrokerList, &out.BrokerList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusKafka.
func (in *MilvusKafka) DeepCopy() *MilvusKafka {
	if in == nil {
		return nil
	}
	out := new(MilvusKafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusList) DeepCopyInto(out *MilvusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Milvus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusList.
func (in *MilvusList) DeepCopy() *MilvusList {
	if in == nil {
		return nil
	}
	out := new(MilvusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MilvusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusProxy) DeepCopyInto(out *MilvusProxy) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
	if in.ServiceLabels != nil {
		in, out := &in.ServiceLabels, &out.ServiceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceAnnotations != nil {
		in, out := &in.ServiceAnnotations, &out.ServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(MilvusIngress)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusProxy.
func (in *MilvusProxy) DeepCopy() *MilvusProxy {
	if in == nil {
		return nil
	}
	out := new(MilvusProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusPulsar) DeepCopyInto(out *MilvusPulsar) {
	*out = *in
	if in.InCluster != nil {
		in, out := &in.InCluster, &out.InCluster
		*out = new(InClusterConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusPulsar.
func (in *MilvusPulsar) DeepCopy() *MilvusPulsar {
	if in == nil {
		return nil
	}
	out := new(MilvusPulsar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusQueryCoord) DeepCopyInto(out *MilvusQueryCoord) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusQueryCoord.
func (in *MilvusQueryCoord) DeepCopy() *MilvusQueryCoord {
	if in == nil {
		return nil
	}
	out := new(MilvusQueryCoord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusQueryNode) DeepCopyInto(out *MilvusQueryNode) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusQueryNode.
func (in *MilvusQueryNode) DeepCopy() *MilvusQueryNode {
	if in == nil {
		return nil
	}
	out := new(MilvusQueryNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusReplicas) DeepCopyInto(out *MilvusReplicas) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusReplicas.
func (in *MilvusReplicas) DeepCopy() *MilvusReplicas {
	if in == nil {
		return nil
	}
	out := new(MilvusReplicas)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusRocksMQ) DeepCopyInto(out *MilvusRocksMQ) {
	*out = *in
	in.Persistence.DeepCopyInto(&out.Persistence)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusRocksMQ.
func (in *MilvusRocksMQ) DeepCopy() *MilvusRocksMQ {
	if in == nil {
		return nil
	}
	out := new(MilvusRocksMQ)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusRootCoord) DeepCopyInto(out *MilvusRootCoord) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusRootCoord.
func (in *MilvusRootCoord) DeepCopy() *MilvusRootCoord {
	if in == nil {
		return nil
	}
	out := new(MilvusRootCoord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusSpec) DeepCopyInto(out *MilvusSpec) {
	*out = *in
	in.Com.DeepCopyInto(&out.Com)
	in.Dep.DeepCopyInto(&out.Dep)
	in.Conf.DeepCopyInto(&out.Conf)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusSpec.
func (in *MilvusSpec) DeepCopy() *MilvusSpec {
	if in == nil {
		return nil
	}
	out := new(MilvusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusStandalone) DeepCopyInto(out *MilvusStandalone) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusStandalone.
func (in *MilvusStandalone) DeepCopy() *MilvusStandalone {
	if in == nil {
		return nil
	}
	out := new(MilvusStandalone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusStatus) DeepCopyInto(out *MilvusStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MilvusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.IngressStatus.DeepCopyInto(&out.IngressStatus)
	out.Replicas = in.Replicas
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusStatus.
func (in *MilvusStatus) DeepCopy() *MilvusStatus {
	if in == nil {
		return nil
	}
	out := new(MilvusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusStorage) DeepCopyInto(out *MilvusStorage) {
	*out = *in
	if in.InCluster != nil {
		in, out := &in.InCluster, &out.InCluster
		*out = new(InClusterConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusStorage.
func (in *MilvusStorage) DeepCopy() *MilvusStorage {
	if in == nil {
		return nil
	}
	out := new(MilvusStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence) DeepCopyInto(out *Persistence) {
	*out = *in
	in.PersistentVolumeClaim.DeepCopyInto(&out.PersistentVolumeClaim)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence.
func (in *Persistence) DeepCopy() *Persistence {
	if in == nil {
		return nil
	}
	out := new(Persistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaim) DeepCopyInto(out *PersistentVolumeClaim) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaim.
func (in *PersistentVolumeClaim) DeepCopy() *PersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Values.
func (in *Values) DeepCopy() *Values {
	if in == nil {
		return nil
	}
	out := new(Values)
	in.DeepCopyInto(out)
	return out
}
