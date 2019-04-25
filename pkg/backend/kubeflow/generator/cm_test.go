// Copyright 2018 Caicloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import (
	"testing"

	common "github.com/kubeflow/tf-operator/pkg/apis/common/v1beta2"
	pytorchv1beta2 "github.com/kubeflow/pytorch-operator/pkg/apis/pytorch/v1beta2"
	tfv1beta2 "github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1beta2"

	"github.com/caicloud/ciao/pkg/types"
)

const (
	defaultNamespace = "default"
)

func TestCMNewTFJob(t *testing.T) {
	cm := NewCM(defaultNamespace)

	expectedPSCount := 1
	expectedWorkerCount := 1
	expectedCM := "image"
	expectedCleanPolicy := types.CleanPodPolicyAll

	param := &types.Parameter{
		PSCount:     expectedPSCount,
		WorkerCount: expectedWorkerCount,
		Image:       expectedCM,
		CleanPolicy: types.CleanPodPolicyAll,
	}

	tfJob := cm.GenerateTFJob(param)
	actualPSCount := *tfJob.Spec.TFReplicaSpecs[tfv1beta2.TFReplicaTypePS].Replicas
	actualWorkerCount := *tfJob.Spec.TFReplicaSpecs[tfv1beta2.TFReplicaTypeWorker].Replicas
	actualCM := tfJob.Spec.TFReplicaSpecs[tfv1beta2.TFReplicaTypePS].Template.Spec.Containers[0].VolumeMounts[0].Name
	actualCleanPolicy := *tfJob.Spec.CleanPodPolicy
	if actualPSCount != int32(expectedPSCount) {
		t.Errorf("Expected %d ps, got %d", expectedPSCount, actualPSCount)
	}
	if actualWorkerCount != int32(expectedWorkerCount) {
		t.Errorf("Expected %d workers, got %d", expectedWorkerCount, actualWorkerCount)
	}
	if actualCM != expectedCM {
		t.Errorf("Expected configmap name %s, got %s", expectedCM, actualCM)
	}
	if actualCleanPolicy != common.CleanPodPolicy(expectedCleanPolicy) {
		t.Errorf("Expected clean policy %s, got %s", expectedCleanPolicy, actualCleanPolicy)
	}
}

func TestCMNewPyTorchJob(t *testing.T) {
	cm := NewCM(defaultNamespace)

	expectedMasterCount := 1
	expectedWorkerCount := 1
	expectedCM := "image"
	expectedCleanPolicy := types.CleanPodPolicyAll

	param := &types.Parameter{
		MasterCount: expectedMasterCount,
		WorkerCount: expectedWorkerCount,
		Image:       expectedCM,
		CleanPolicy: types.CleanPodPolicyAll,
	}

	pytorchJob := cm.GeneratePyTorchJob(param)
	actualMasterCount := *pytorchJob.Spec.PyTorchReplicaSpecs[pytorchv1beta2.PyTorchReplicaTypeMaster].Replicas
	actualWorkerCount := *pytorchJob.Spec.PyTorchReplicaSpecs[pytorchv1beta2.PyTorchReplicaTypeWorker].Replicas
	actualCM := pytorchJob.Spec.PyTorchReplicaSpecs[pytorchv1beta2.PyTorchReplicaTypeMaster].Template.Spec.Containers[0].VolumeMounts[0].Name
	actualCleanPolicy := *pytorchJob.Spec.CleanPodPolicy
	if actualMasterCount != int32(expectedMasterCount) {
		t.Errorf("Expected %d masters, got %d", expectedMasterCount, actualMasterCount)
	}
	if actualWorkerCount != int32(expectedWorkerCount) {
		t.Errorf("Expected %d workers, got %d", expectedWorkerCount, actualWorkerCount)
	}
	if actualCM != expectedCM {
		t.Errorf("Expected configmap name %s, got %s", expectedCM, actualCM)
	}
	if actualCleanPolicy != common.CleanPodPolicy(expectedCleanPolicy) {
		t.Errorf("Expected clean policy %s, got %s", expectedCleanPolicy, actualCleanPolicy)
	}
}
