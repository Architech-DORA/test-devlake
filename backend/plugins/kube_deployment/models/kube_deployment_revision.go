/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

import (
	"time"

	"github.com/apache/incubator-devlake/core/models/common"
)

type KubeDeploymentRevision struct {
	common.NoPKModel  `json:"-"`
	ConnectionId      uint64    `gorm:"primaryKey"`
	Id                string    `json:"id" gorm:"primaryKey"`
	Namespace         string    `json:"namespace"`
	DeploymentName    string    `json:"deployment_name"`
	RevisionNumber    uint32    `json:"revision_number"`
	CreationTimestamp time.Time `json:"creation_timestamp"`
	CicdScopeId       string    `json:"cicd_scope_id"`
}

func (KubeDeploymentRevision) TableName() string {
	return "_tool_kube_deployment_revisions"
}
