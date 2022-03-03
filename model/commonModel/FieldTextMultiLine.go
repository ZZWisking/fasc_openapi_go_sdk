package commonModel

type FieldTextMultiLine struct {
	Required     bool   `json:"required"`
	DefaultValue string `json:"defaultValue,omitempty"`
}
