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
	"github.com/apache/incubator-devlake/core/models/common"
)

type TapdTaskCustomFields struct {
	ConnectionId uint64 `gorm:"primaryKey;type:BIGINT  NOT NULL"`
	Id           uint64 `gorm:"primaryKey;type:BIGINT  NOT NULL" json:"id,string"`
	WorkspaceId  uint64 `json:"workspace_id,string"`
	EntryType    string `json:"entry_type" gorm:"type:varchar(255)"`
	CustomField  string `json:"custom_field" gorm:"type:varchar(255)"`
	Type         string `json:"type" gorm:"type:varchar(255)"`
	Name         string `json:"name" gorm:"type:varchar(255)"`
	Options      string `json:"options" gorm:"type:text"`
	Enabled      string `json:"enabled" gorm:"type:varchar(255)"`
	Sort         string `json:"sort" gorm:"type:varchar(255)"`
	common.NoPKModel
}

func (TapdTaskCustomFields) TableName() string {
	return "_tool_tapd_task_custom_fields"
}
