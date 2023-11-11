package options

import (
	cliflag "k8s.io/component-base/cli/flag"
)

// OperatorOptions runs operator.
type OperatorOptions struct {
	MetricsAddr             string
	ProbeAddr               string
	EnableLeaderElection    bool
	MaxConcurrentReconciles int
}

// NewOperatorOptions creates a new OperatorOptions object with default parameters
func NewOperatorOptions() *OperatorOptions {
	o := &OperatorOptions{}
	return o
}

// Flags returns flags for a specific NewOperatorCommand by section name
func (o *OperatorOptions) Flags() (fss cliflag.NamedFlagSets) {
	fs := fss.FlagSet("operator")
	fs.StringVar(&o.MetricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	fs.StringVar(&o.ProbeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	fs.BoolVar(&o.EnableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	fs.IntVar(&o.MaxConcurrentReconciles, "max-concurrent-reconciles", 1, "max concurrent reconciles")
	return fss
}

// ValidateOptions validates OperatorOptions
func (o *OperatorOptions) ValidateOptions() error {
	return nil
}

// Complete set default OperatorOptions.
// Should be called after flags parsed.
func (o *OperatorOptions) Complete() error {
	return nil
}
