package common

// ExampleSpec defines the desired state of Example
type ExampleSpec struct {
	// Foo is an example field of Example. Edit example_types.go to remove/update
	Foo string `json:"foo,omitempty"`

	// Bar is another exemple field
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:hidden"}
	Bar string `json:"bar,omitempty"`

	Options CommonOptions `json:"options"`
}

// CommonOptions defines some common options of our apis
type CommonOptions struct {
	// Common1 is an option
	Common1 string `json:"common1,omitempty"`

	// Hidden1 is an internal options that should be hidden in UI
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:hidden"}
	Hidden1 string `json:"hidden1,omitempty"`
}
