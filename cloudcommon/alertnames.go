// Copyright 2022 MobiledgeX, Inc
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

package cloudcommon

// Alert names
const (
	AlertAutoScaleUp                         = "AutoScaleUp"
	AlertAutoScaleDown                       = "AutoScaleDown"
	AlertClusterAutoScale                    = "ClusterAutoScale"
	AlertAppInstDown                         = "AppInstDown"
	AlertClusterSvcAppInstFailure            = "ClusterSvcAppInstFailure"
	AlertAutoUndeploy                        = "AutoProvUndeploy"
	AlertCloudletDown                        = "CloudletDown"
	AlertCloudletDownDescription             = "Cloudlet resource manager is offline"
	AlertClusterSvcAppInstFailureDescription = "Cluster-svc create AppInst failed"
	AlertCloudletResourceUsage               = "CloudletResourceUsage"
	AlertTypeUserDefined                     = "UserDefined"
)

// Alert types
const (
	AlertAnnotationTitle       = "title"
	AlertAnnotationDescription = "description"
	AlertAnnotationErrorString = "error"
	AlertScopeTypeTag          = "scope"
	AlertSeverityLabel         = "severity"
	AlertScopeApp              = "Application"
	AlertScopeCloudlet         = "Cloudlet"
	AlertTypeLabel             = "type"
	AlertScopePlatform         = "Platform"
)

// Alert annotation keys
// for autoscale:
var AlertKeyNodeCount = "nodecount"
var AlertKeyLowCpuNodeCount = "lowcpunodecount"
var AlertKeyMinNodes = "minnodes"

// for healthCheck:
var AlertHealthCheckStatus = "status"

const (
	AlertSeverityError = "error"
	AlertSeverityWarn  = "warning"
	AlertSeverityInfo  = "info"
	// List in the order of increasing severity
	ValidAlertSeverityString = `"info", "warning", "error"`
)

var AlertSeverityTypes = map[string]struct{}{
	AlertSeverityError: struct{}{},
	AlertSeverityWarn:  struct{}{},
	AlertSeverityInfo:  struct{}{},
}

// Map represents severities for the specific alerts that the platfrom generates
var AlertSeverityValues = map[string]string{
	AlertAppInstDown:              AlertSeverityError,
	AlertCloudletDown:             AlertSeverityError,
	AlertCloudletResourceUsage:    AlertSeverityWarn,
	AlertClusterSvcAppInstFailure: AlertSeverityError,
}

func GetSeverityForAlert(alertname string) string {
	if severity, found := AlertSeverityValues[alertname]; found {
		return severity
	}
	// default to "info"
	return AlertSeverityInfo
}

func IsMonitoredAlert(labels map[string]string) bool {
	alertName, found := labels["alertname"]
	// Alertnames with empty alertnames, or no alertnames are not monitored
	if !found || alertName == "" {
		return false
	}
	alertScope, _ := labels[AlertScopeTypeTag]
	// All App/Cloudlet alerts are monitored
	if alertScope == AlertScopeApp ||
		alertScope == AlertScopeCloudlet ||
		alertScope == AlertScopePlatform {
		return true
	}
	alertType, _ := labels[AlertTypeLabel]
	// user defined alerts are always monitored
	if alertType == AlertTypeUserDefined {
		return true
	}
	if alertName == AlertClusterAutoScale ||
		alertName == AlertAutoScaleUp ||
		alertName == AlertAutoScaleDown ||
		alertName == AlertAppInstDown ||
		alertName == AlertCloudletDown ||
		alertName == AlertAutoUndeploy ||
		alertName == AlertCloudletResourceUsage ||
		alertName == AlertClusterSvcAppInstFailure {
		return true
	}
	return false
}

// Prune labels we don't want to show on the alerts sent to the external alert integrations
func IsLabelInternal(label string) bool {
	if label == "instance" ||
		label == "job" ||
		label == AlertTypeLabel {
		return true
	}
	return false
}

func IsInternalAlert(labels map[string]string) bool {
	alertName, _ := labels["alertname"]
	if alertName == AlertAppInstDown ||
		alertName == AlertCloudletDown ||
		alertName == AlertCloudletResourceUsage ||
		alertName == AlertClusterSvcAppInstFailure {
		return false
	}
	alertType, _ := labels[AlertTypeLabel]
	// user defined alerts are external
	if alertType == AlertTypeUserDefined {
		return false
	}
	return true
}

func IsAlertSeverityValid(severity string) bool {
	if _, found := AlertSeverityTypes[severity]; found {
		return true
	}
	return false
}

// Helper function - returns the string representations of all valid severities
func GetValidAlertSeverityString() string {
	return ValidAlertSeverityString
}
