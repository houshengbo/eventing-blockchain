/*
Copyright 2022 The Knative Authors

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

package v1alpha1

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/webhook/resourcesemantics"
)

// Check that BlockchainSource can be validated and can be defaulted.
var _ runtime.Object = (*BlockchainSource)(nil)

var _ resourcesemantics.GenericCRD = (*BlockchainSource)(nil)

// Check that the type conforms to the duck Knative Resource shape.
var _ duckv1.KRShaped = (*BlockchainSource)(nil)

// BlockchainSourceSpec defines the desired state of BlockchainSource
// +kubebuilder:categories=all,knative,eventing,sources
type BlockchainSourceSpec struct {
	// ServiceAccountName holds the name of the Kubernetes service account
	// as which the underlying K8s resources should be run. If unspecified
	// this will default to the "default" service account for the namespace
	// in which the BlockchainSource exists.
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// OwnerAndRepository is the GitHub owner/org and repository to
	// receive events from. The repository may be left off to receive
	// events from an entire organization.
	// Examples:
	//  myuser/project
	//  myorganization
	// +kubebuilder:validation:MinLength=1
	OwnerAndRepository string `json:"ownerAndRepository"`

	// EventType is the type of event to receive from GitHub. These
	// correspond to the "Webhook event name" values listed at
	// https://developer.github.com/v3/activity/events/types/ - ie
	// "pull_request"
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:Enum=check_suite,commit_comment,create,delete,deployment,deployment_status,fork,gollum,installation,integration_installation,issue_comment,issues,label,member,membership,milestone,organization,org_block,page_build,ping,project_card,project_column,project,public,pull_request,pull_request_review,pull_request_review_comment,push,release,repository,status,team,team_add,watch
	EventTypes []string `json:"eventTypes"`

	// AccessToken is the Kubernetes secret containing the GitHub
	// access token
	AccessToken SecretValueFromSource `json:"accessToken"`

	// SecretToken is the Kubernetes secret containing the GitHub
	// secret token
	SecretToken SecretValueFromSource `json:"secretToken"`

	// API URL if using blockchain enterprise (default https://api.github.com)
	// +optional
	BlockchainAPIURL string `json:"blockchainAPIURL,omitempty"`

	// Secure can be set to true to configure the webhook to use https,
	// or false to use http.  Omitting it relies on the scheme of the
	// Knative Service created (e.g. if auto-TLS is enabled it should
	// do the right thing).
	// +optional
	Secure *bool `json:"secure,omitempty"`

	// inherits duck/v1 SourceSpec, which currently provides:
	// * Sink - a reference to an object that will resolve to a domain name or
	//   a URI directly to use as the sink.
	// * CloudEventOverrides - defines overrides to control the output format
	//   and modifications of the event sent to the sink.
	duckv1.SourceSpec `json:",inline"`
}

// SecretValueFromSource represents the source of a secret value
type SecretValueFromSource struct {
	// The Secret key to select from.
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

const (
	// GitHubEventTypePrefix is what all GitHub event types get
	// prefixed with when converting to CloudEvents.
	GitHubEventTypePrefix = "dev.knative.source.github"

	// GitHubEventSourcePrefix is what all GitHub event sources get
	// prefixed with when converting to CloudEvents.
	GitHubEventSourcePrefix = "https://github.com"
)

// GitHubEventType returns an event type emitted by a BlockchainSource suitable for
// the value of a CloudEvent's "type" context attribute.
func GitHubEventType(ghEventType string) string {
	return fmt.Sprintf("%s.%s", GitHubEventTypePrefix, ghEventType)
}

// GitHubEventSource returns a unique representation of a BlockchainSource suitable
// for the value of a CloudEvent's "source" context attribute.
func GitHubEventSource(ownerAndRepo string) string {
	return fmt.Sprintf("%s/%s", GitHubEventSourcePrefix, ownerAndRepo)
}

const (
	// BlockchainSourceConditionReady has status True when the
	// BlockchainSource is ready to send events.
	BlockchainSourceConditionReady = apis.ConditionReady

	// BlockchainSourceConditionSecretsProvided has status True when the
	// BlockchainSource has valid secret references
	BlockchainSourceConditionSecretsProvided apis.ConditionType = "SecretsProvided"

	// BlockchainSourceConditionSinkProvided has status True when the
	// BlockchainSource has been configured with a sink target.
	BlockchainSourceConditionSinkProvided apis.ConditionType = "SinkProvided"

	// BlockchainSourceConditionWebhookConfigured has a status True when the
	// BlockchainSource has been configured with a webhook.
	BlockchainSourceConditionWebhookConfigured apis.ConditionType = "WebhookConfigured"

	// GitHubServiceconditiondeployed has status True when then
	// BlockchainSource Service has been deployed
	//	GitHubServiceConditionDeployed apis.ConditionType = "Deployed"

	// BlockchainSourceReconciled has status True when the
	// BlockchainSource has been properly reconciled
	GitHub
)

var BlockchainSourceCondSet = apis.NewLivingConditionSet(
	BlockchainSourceConditionSecretsProvided,
	BlockchainSourceConditionSinkProvided,
	BlockchainSourceConditionWebhookConfigured)

//	GitHubServiceConditionDeployed)

// BlockchainSourceStatus defines the observed state of BlockchainSource
type BlockchainSourceStatus struct {
	// inherits duck/v1 SourceStatus, which currently provides:
	// * ObservedGeneration - the 'Generation' of the Service that was last
	//   processed by the controller.
	// * Conditions - the latest available observations of a resource's current
	//   state.
	// * SinkURI - the current active sink URI that has been configured for the
	//   Source.
	duckv1.SourceStatus `json:",inline"`

	// WebhookIDKey is the ID of the webhook registered with GitHub
	WebhookIDKey string `json:"webhookIDKey,omitempty"`
}

func (*BlockchainSource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("BlockchainSource")
}

// GetConditionSet retrieves the condition set for this resource. Implements the KRShaped interface.
func (*BlockchainSource) GetConditionSet() apis.ConditionSet {
	return BlockchainSourceCondSet
}

// GetStatus retrieves the duck status for this resource. Implements the KRShaped interface.
func (g *BlockchainSource) GetStatus() *duckv1.Status {
	return &g.Status.Status
}

// GetCondition returns the condition currently associated with the given type, or nil.
func (s *BlockchainSourceStatus) GetCondition(t apis.ConditionType) *apis.Condition {
	return BlockchainSourceCondSet.Manage(s).GetCondition(t)
}

// IsReady returns true if the resource is ready overall.
func (s *BlockchainSourceStatus) IsReady() bool {
	return BlockchainSourceCondSet.Manage(s).IsHappy()
}

// InitializeConditions sets relevant unset conditions to Unknown state.
func (s *BlockchainSourceStatus) InitializeConditions() {
	BlockchainSourceCondSet.Manage(s).InitializeConditions()
}

// MarkSecrets sets the condition that the source has a valid spec
func (s *BlockchainSourceStatus) MarkSecrets() {
	BlockchainSourceCondSet.Manage(s).MarkTrue(BlockchainSourceConditionSecretsProvided)
}

// MarkNoSecrets sets the condition that the source does not have a valid spec
func (s *BlockchainSourceStatus) MarkNoSecrets(reason, messageFormat string, messageA ...interface{}) {
	BlockchainSourceCondSet.Manage(s).MarkFalse(BlockchainSourceConditionSecretsProvided, reason, messageFormat, messageA...)
}

// MarkSink sets the condition that the source has a sink configured.
func (s *BlockchainSourceStatus) MarkSink(uri *apis.URL) {
	s.SinkURI = uri
	if uri != nil {
		BlockchainSourceCondSet.Manage(s).MarkTrue(BlockchainSourceConditionSinkProvided)
	} else {
		BlockchainSourceCondSet.Manage(s).MarkUnknown(BlockchainSourceConditionSinkProvided,
			"SinkEmpty", "Sink has resolved to empty.")
	}
}

// MarkNoSink sets the condition that the source does not have a sink configured.
func (s *BlockchainSourceStatus) MarkNoSink(reason, messageFormat string, messageA ...interface{}) {
	BlockchainSourceCondSet.Manage(s).MarkFalse(BlockchainSourceConditionSinkProvided, reason, messageFormat, messageA...)
}

// MarkWebhookConfigured sets the condition that the source has set its webhook configured.
func (s *BlockchainSourceStatus) MarkWebhookConfigured() {
	BlockchainSourceCondSet.Manage(s).MarkTrue(BlockchainSourceConditionWebhookConfigured)
}

// MarkWebhookNotConfigured sets the condition that the source does not have its webhook configured.
func (s *BlockchainSourceStatus) MarkWebhookNotConfigured(reason, messageFormat string, messageA ...interface{}) {
	BlockchainSourceCondSet.Manage(s).MarkFalse(BlockchainSourceConditionWebhookConfigured, reason, messageFormat, messageA...)
}

// MarkDeployed sets the condition that the source has been deployed.
//func (s *BlockchainSourceStatus) MarkServiceDeployed(d *appsv1.Deployment) {
//	if duckv1.DeploymentIsAvailable(&d.Status, false) {
//		BlockchainSourceCondSet.Manage(s).MarkTrue(GitHubServiceConditionDeployed)
//	} else {
//		BlockchainSourceCondSet.Manage(s).MarkFalse(GitHubServiceConditionDeployed, "ServiceDeploymentUnavailable", "The Deployment '%s' is unavailable.", d.Name)
//	}
//}

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BlockchainSource is the Schema for the BlockchainSources API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:categories=all,knative,eventing,sources
type BlockchainSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BlockchainSourceSpec   `json:"spec,omitempty"`
	Status BlockchainSourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BlockchainSourceList contains a list of BlockchainSource
type BlockchainSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BlockchainSource `json:"items"`
}
