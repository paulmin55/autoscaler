/*
Copyright 2018 The Kubernetes Authors.

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

package nodegroups

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/context"
	"k8s.io/kubernetes/pkg/scheduler/schedulercache"
)

// NodeGroupListProcessor processes lists of NodeGroups considered in scale-up.
type NodeGroupListProcessor interface {
	Process(context *context.AutoscalingContext, nodeGroups []cloudprovider.NodeGroup,
		nodeInfos map[string]*schedulercache.NodeInfo,
		unschedulablePods []*apiv1.Pod) ([]cloudprovider.NodeGroup, map[string]*schedulercache.NodeInfo, error)
}

// NoOpNodeGroupListProcessor is returning pod lists without processing them.
type NoOpNodeGroupListProcessor struct {
}

// NewDefaultNodeGroupListProcessor creates an instance of NodeGroupListProcessor.
func NewDefaultNodeGroupListProcessor() NodeGroupListProcessor {
	// TODO(maciekpytel): Use a better default
	return &AutoprovisioningNodeGroupListProcessor{}
}

// Process processes lists of unschedulable and sheduled pods before scaling of the cluster.
func (p *NoOpNodeGroupListProcessor) Process(context *context.AutoscalingContext, nodeGroups []cloudprovider.NodeGroup, nodeInfos map[string]*schedulercache.NodeInfo,
	unschedulablePods []*apiv1.Pod) ([]cloudprovider.NodeGroup, map[string]*schedulercache.NodeInfo, error) {
	return nodeGroups, nodeInfos, nil
}
