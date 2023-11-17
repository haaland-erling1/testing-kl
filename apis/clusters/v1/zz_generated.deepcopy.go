//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterConfig) DeepCopyInto(out *AWSClusterConfig) {
	*out = *in
	in.K3sMasters.DeepCopyInto(&out.K3sMasters)
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make(map[string]AwsEC2PoolConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.SpotNodePools != nil {
		in, out := &in.SpotNodePools, &out.SpotNodePools
		*out = make(map[string]AwsSpotPoolConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterConfig.
func (in *AWSClusterConfig) DeepCopy() *AWSClusterConfig {
	if in == nil {
		return nil
	}
	out := new(AWSClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSK3sMastersConfig) DeepCopyInto(out *AWSK3sMastersConfig) {
	*out = *in
	if in.IAMInstanceProfileRole != nil {
		in, out := &in.IAMInstanceProfileRole, &out.IAMInstanceProfileRole
		*out = new(string)
		**out = **in
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]MasterNodeProps, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSK3sMastersConfig.
func (in *AWSK3sMastersConfig) DeepCopy() *AWSK3sMastersConfig {
	if in == nil {
		return nil
	}
	out := new(AWSK3sMastersConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSNodePoolConfig) DeepCopyInto(out *AWSNodePoolConfig) {
	*out = *in
	if in.IAMInstanceProfileRole != nil {
		in, out := &in.IAMInstanceProfileRole, &out.IAMInstanceProfileRole
		*out = new(string)
		**out = **in
	}
	if in.EC2Pool != nil {
		in, out := &in.EC2Pool, &out.EC2Pool
		*out = new(AwsEC2PoolConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SpotPool != nil {
		in, out := &in.SpotPool, &out.SpotPool
		*out = new(AwsSpotPoolConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSNodePoolConfig.
func (in *AWSNodePoolConfig) DeepCopy() *AWSNodePoolConfig {
	if in == nil {
		return nil
	}
	out := new(AWSNodePoolConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountS3Bucket) DeepCopyInto(out *AccountS3Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountS3Bucket.
func (in *AccountS3Bucket) DeepCopy() *AccountS3Bucket {
	if in == nil {
		return nil
	}
	out := new(AccountS3Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountS3Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountS3BucketList) DeepCopyInto(out *AccountS3BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccountS3Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountS3BucketList.
func (in *AccountS3BucketList) DeepCopy() *AccountS3BucketList {
	if in == nil {
		return nil
	}
	out := new(AccountS3BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountS3BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountS3BucketSpec) DeepCopyInto(out *AccountS3BucketSpec) {
	*out = *in
	out.CredentialsRef = in.CredentialsRef
	if in.CredentialKeys != nil {
		in, out := &in.CredentialKeys, &out.CredentialKeys
		*out = new(CloudProviderCredentialKeys)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountS3BucketSpec.
func (in *AccountS3BucketSpec) DeepCopy() *AccountS3BucketSpec {
	if in == nil {
		return nil
	}
	out := new(AccountS3BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsEC2PoolConfig) DeepCopyInto(out *AwsEC2PoolConfig) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]NodeProps, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsEC2PoolConfig.
func (in *AwsEC2PoolConfig) DeepCopy() *AwsEC2PoolConfig {
	if in == nil {
		return nil
	}
	out := new(AwsEC2PoolConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsSpotCpuNode) DeepCopyInto(out *AwsSpotCpuNode) {
	*out = *in
	out.VCpu = in.VCpu
	out.MemoryPerVCpu = in.MemoryPerVCpu
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsSpotCpuNode.
func (in *AwsSpotCpuNode) DeepCopy() *AwsSpotCpuNode {
	if in == nil {
		return nil
	}
	out := new(AwsSpotCpuNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsSpotGpuNode) DeepCopyInto(out *AwsSpotGpuNode) {
	*out = *in
	if in.InstanceTypes != nil {
		in, out := &in.InstanceTypes, &out.InstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsSpotGpuNode.
func (in *AwsSpotGpuNode) DeepCopy() *AwsSpotGpuNode {
	if in == nil {
		return nil
	}
	out := new(AwsSpotGpuNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsSpotPoolConfig) DeepCopyInto(out *AwsSpotPoolConfig) {
	*out = *in
	if in.CpuNode != nil {
		in, out := &in.CpuNode, &out.CpuNode
		*out = new(AwsSpotCpuNode)
		**out = **in
	}
	if in.GpuNode != nil {
		in, out := &in.GpuNode, &out.GpuNode
		*out = new(AwsSpotGpuNode)
		(*in).DeepCopyInto(*out)
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]NodeProps, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsSpotPoolConfig.
func (in *AwsSpotPoolConfig) DeepCopy() *AwsSpotPoolConfig {
	if in == nil {
		return nil
	}
	out := new(AwsSpotPoolConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfig) DeepCopyInto(out *AzureConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfig.
func (in *AzureConfig) DeepCopy() *AzureConfig {
	if in == nil {
		return nil
	}
	out := new(AzureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOC) DeepCopyInto(out *BYOC) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOC.
func (in *BYOC) DeepCopy() *BYOC {
	if in == nil {
		return nil
	}
	out := new(BYOC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BYOC) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOCList) DeepCopyInto(out *BYOCList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BYOC, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOCList.
func (in *BYOCList) DeepCopy() *BYOCList {
	if in == nil {
		return nil
	}
	out := new(BYOCList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BYOCList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOCSpec) DeepCopyInto(out *BYOCSpec) {
	*out = *in
	if in.StorageClasses != nil {
		in, out := &in.StorageClasses, &out.StorageClasses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IngressClasses != nil {
		in, out := &in.IngressClasses, &out.IngressClasses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PublicIPs != nil {
		in, out := &in.PublicIPs, &out.PublicIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOCSpec.
func (in *BYOCSpec) DeepCopy() *BYOCSpec {
	if in == nil {
		return nil
	}
	out := new(BYOCSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProviderCredentialKeys) DeepCopyInto(out *CloudProviderCredentialKeys) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProviderCredentialKeys.
func (in *CloudProviderCredentialKeys) DeepCopy() *CloudProviderCredentialKeys {
	if in == nil {
		return nil
	}
	out := new(CloudProviderCredentialKeys)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOutput) DeepCopyInto(out *ClusterOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOutput.
func (in *ClusterOutput) DeepCopy() *ClusterOutput {
	if in == nil {
		return nil
	}
	out := new(ClusterOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	out.ClusterTokenRef = in.ClusterTokenRef
	out.CredentialsRef = in.CredentialsRef
	if in.CredentialKeys != nil {
		in, out := &in.CredentialKeys, &out.CredentialKeys
		*out = new(CloudProviderCredentialKeys)
		**out = **in
	}
	if in.ClusterInternalDnsHost != nil {
		in, out := &in.ClusterInternalDnsHost, &out.ClusterInternalDnsHost
		*out = new(string)
		**out = **in
	}
	if in.CloudflareEnabled != nil {
		in, out := &in.CloudflareEnabled, &out.CloudflareEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSClusterConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(ClusterOutput)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DigitalOceanConfig) DeepCopyInto(out *DigitalOceanConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DigitalOceanConfig.
func (in *DigitalOceanConfig) DeepCopy() *DigitalOceanConfig {
	if in == nil {
		return nil
	}
	out := new(DigitalOceanConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPConfig) DeepCopyInto(out *GCPConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPConfig.
func (in *GCPConfig) DeepCopy() *GCPConfig {
	if in == nil {
		return nil
	}
	out := new(GCPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastuctureAsCode) DeepCopyInto(out *InfrastuctureAsCode) {
	*out = *in
	out.CloudProviderAccessKey = in.CloudProviderAccessKey
	out.CloudProviderSecretKey = in.CloudProviderSecretKey
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastuctureAsCode.
func (in *InfrastuctureAsCode) DeepCopy() *InfrastuctureAsCode {
	if in == nil {
		return nil
	}
	out := new(InfrastuctureAsCode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterNodeProps) DeepCopyInto(out *MasterNodeProps) {
	*out = *in
	in.NodeProps.DeepCopyInto(&out.NodeProps)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterNodeProps.
func (in *MasterNodeProps) DeepCopy() *MasterNodeProps {
	if in == nil {
		return nil
	}
	out := new(MasterNodeProps)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Node) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeList) DeepCopyInto(out *NodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Node, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeList.
func (in *NodeList) DeepCopy() *NodeList {
	if in == nil {
		return nil
	}
	out := new(NodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePool) DeepCopyInto(out *NodePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePool.
func (in *NodePool) DeepCopy() *NodePool {
	if in == nil {
		return nil
	}
	out := new(NodePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolList) DeepCopyInto(out *NodePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolList.
func (in *NodePoolList) DeepCopy() *NodePoolList {
	if in == nil {
		return nil
	}
	out := new(NodePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpec) DeepCopyInto(out *NodePoolSpec) {
	*out = *in
	out.IAC = in.IAC
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSNodePoolConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpec.
func (in *NodePoolSpec) DeepCopy() *NodePoolSpec {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeProps) DeepCopyInto(out *NodeProps) {
	*out = *in
	if in.LastRecreatedAt != nil {
		in, out := &in.LastRecreatedAt, &out.LastRecreatedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeProps.
func (in *NodeProps) DeepCopy() *NodeProps {
	if in == nil {
		return nil
	}
	out := new(NodeProps)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSpec.
func (in *NodeSpec) DeepCopy() *NodeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeSpec)
	in.DeepCopyInto(out)
	return out
}
