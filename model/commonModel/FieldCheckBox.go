package commonModel

type FieldCheckBox struct {
	Required     bool   `json:"required"`
	DefaultValue string `json:"defaultValue,omitempty"`
}
