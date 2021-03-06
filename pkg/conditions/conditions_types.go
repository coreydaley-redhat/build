// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package conditions

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Conditions defines a list of Condition
type Conditions []Condition

// Type used for defining the conditiont Type field flavour
type Type string

const (
	// Succeeded specifies that the resource has finished.
	// For resources that run to completion.
	Succeeded Type = "Succeeded"
)

// Condition defines the required fields for populating
// Build controllers Conditions
type Condition struct {
	// Type of condition
	// +required
	Type Type `json:"type" description:"type of status condition"`

	// Status of the condition, one of True, False, Unknown.
	// +required
	Status corev1.ConditionStatus `json:"status" description:"status of the condition, one of True, False, Unknown"`

	// LastTransitionTime last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" description:"last time the condition transit from one status to another"`

	// The reason for the condition last transition.
	// +optional
	Reason string `json:"reason,omitempty" description:"one-word CamelCase reason for the condition's last transition"`

	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message,omitempty" description:"human-readable message indicating details about last transition"`
}

// GetReason returns the condition Reason, it ensures that by getting the Reason
// the call will not panic if the Condition is not present
func (c *Condition) GetReason() string {
	if c == nil {
		return ""
	}
	return c.Reason
}

// GetMessage returns the condition Message, it ensures that by getting the Message
// the call will not panic if the Condition is not present
func (c *Condition) GetMessage() string {
	if c == nil {
		return ""
	}
	return c.Message
}

// GetStatus returns the condition Status, it ensures that by getting the Status
// the call will not panic if the Condition is not present
func (c *Condition) GetStatus() corev1.ConditionStatus {
	if c == nil {
		return ""
	}
	return c.Status
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
// TODO: Added manually, it was not autogenereated.
func (c *Condition) DeepCopyInto(out *Condition) {
	*out = *c
	c.DeepCopyInto(out)
	return
}
