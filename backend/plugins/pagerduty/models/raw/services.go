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

package raw

import "time"

type Services struct {
	// AcknowledgementTimeout corresponds to the JSON schema field
	// "acknowledgement_timeout".
	AcknowledgementTimeout *int `json:"acknowledgement_timeout,omitempty"`

	// AlertCreation corresponds to the JSON schema field "alert_creation".
	AlertCreation *string `json:"alert_creation,omitempty"`

	// AlertGrouping corresponds to the JSON schema field "alert_grouping".
	AlertGrouping *string `json:"alert_grouping,omitempty"`

	// AlertGroupingTimeout corresponds to the JSON schema field
	// "alert_grouping_timeout".
	AlertGroupingTimeout *int `json:"alert_grouping_timeout,omitempty"`

	// AutoResolveTimeout corresponds to the JSON schema field "auto_resolve_timeout".
	AutoResolveTimeout *int `json:"auto_resolve_timeout,omitempty"`

	// CreatedAt corresponds to the JSON schema field "created_at".
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Description corresponds to the JSON schema field "description".
	Description *string `json:"description,omitempty"`

	// EscalationPolicy corresponds to the JSON schema field "escalation_policy".
	EscalationPolicy *ServicesEscalationPolicy `json:"escalation_policy,omitempty"`

	// HtmlUrl corresponds to the JSON schema field "html_url".
	HtmlUrl *string `json:"html_url,omitempty"`

	// Id corresponds to the JSON schema field "id".
	Id *string `json:"id,omitempty"`

	// IncidentUrgencyRule corresponds to the JSON schema field
	// "incident_urgency_rule".
	IncidentUrgencyRule *ServicesIncidentUrgencyRule `json:"incident_urgency_rule,omitempty"`

	// Integrations corresponds to the JSON schema field "integrations".
	Integrations []ServicesIntegrationsElem `json:"integrations,omitempty"`

	// LastIncidentTimestamp corresponds to the JSON schema field
	// "last_incident_timestamp".
	LastIncidentTimestamp *time.Time `json:"last_incident_timestamp,omitempty"`

	// Name corresponds to the JSON schema field "name".
	Name *string `json:"name,omitempty"`

	// ScheduledActions corresponds to the JSON schema field "scheduled_actions".
	ScheduledActions []ServicesScheduledActionsElem `json:"scheduled_actions,omitempty"`

	// Self corresponds to the JSON schema field "self".
	Self *string `json:"self,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status *string `json:"status,omitempty"`

	// Summary corresponds to the JSON schema field "summary".
	Summary *string `json:"summary,omitempty"`

	// SupportHours corresponds to the JSON schema field "support_hours".
	SupportHours *ServicesSupportHours `json:"support_hours,omitempty"`

	// Teams corresponds to the JSON schema field "teams".
	Teams []ServicesTeamsElem `json:"teams,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type *string `json:"type,omitempty"`
}

type ServicesEscalationPolicy struct {
	// HtmlUrl corresponds to the JSON schema field "html_url".
	HtmlUrl *string `json:"html_url,omitempty"`

	// Id corresponds to the JSON schema field "id".
	Id *string `json:"id,omitempty"`

	// Self corresponds to the JSON schema field "self".
	Self *string `json:"self,omitempty"`

	// Summary corresponds to the JSON schema field "summary".
	Summary *string `json:"summary,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type *string `json:"type,omitempty"`
}

type ServicesIncidentUrgencyRule struct {
	// DuringSupportHours corresponds to the JSON schema field "during_support_hours".
	DuringSupportHours *ServicesIncidentUrgencyRuleDuringSupportHours `json:"during_support_hours,omitempty"`

	// OutsideSupportHours corresponds to the JSON schema field
	// "outside_support_hours".
	OutsideSupportHours *ServicesIncidentUrgencyRuleOutsideSupportHours `json:"outside_support_hours,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type *string `json:"type,omitempty"`
}

type ServicesIncidentUrgencyRuleDuringSupportHours struct {
	// Type corresponds to the JSON schema field "type".
	Type *string `json:"type,omitempty"`

	// Urgency corresponds to the JSON schema field "urgency".
	Urgency *string `json:"urgency,omitempty"`
}

type ServicesIncidentUrgencyRuleOutsideSupportHours struct {
	// Type corresponds to the JSON schema field "type".
	Type *string `json:"type,omitempty"`

	// Urgency corresponds to the JSON schema field "urgency".
	Urgency *string `json:"urgency,omitempty"`
}

type ServicesIntegrationsElem struct {
	// HtmlUrl corresponds to the JSON schema field "html_url".
	HtmlUrl *string `json:"html_url,omitempty"`

	// Id corresponds to the JSON schema field "id".
	Id *string `json:"id,omitempty"`

	// Self corresponds to the JSON schema field "self".
	Self *string `json:"self,omitempty"`

	// Summary corresponds to the JSON schema field "summary".
	Summary *string `json:"summary,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type *string `json:"type,omitempty"`
}

type ServicesScheduledActionsElem struct {
	// At corresponds to the JSON schema field "at".
	At *ServicesScheduledActionsElemAt `json:"at,omitempty"`

	// ToUrgency corresponds to the JSON schema field "to_urgency".
	ToUrgency *string `json:"to_urgency,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type *string `json:"type,omitempty"`
}

type ServicesScheduledActionsElemAt struct {
	// Name corresponds to the JSON schema field "name".
	Name *string `json:"name,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type *string `json:"type,omitempty"`
}

type ServicesSupportHours struct {
	// DaysOfWeek corresponds to the JSON schema field "days_of_week".
	DaysOfWeek []int `json:"days_of_week,omitempty"`

	// EndTime corresponds to the JSON schema field "end_time".
	EndTime *string `json:"end_time,omitempty"`

	// StartTime corresponds to the JSON schema field "start_time".
	StartTime *string `json:"start_time,omitempty"`

	// TimeZone corresponds to the JSON schema field "time_zone".
	TimeZone *string `json:"time_zone,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type *string `json:"type,omitempty"`
}

type ServicesTeamsElem struct {
	// HtmlUrl corresponds to the JSON schema field "html_url".
	HtmlUrl *string `json:"html_url,omitempty"`

	// Id corresponds to the JSON schema field "id".
	Id *string `json:"id,omitempty"`

	// Self corresponds to the JSON schema field "self".
	Self *string `json:"self,omitempty"`

	// Summary corresponds to the JSON schema field "summary".
	Summary *string `json:"summary,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type *string `json:"type,omitempty"`
}
