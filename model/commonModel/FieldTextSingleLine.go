package commonModel

type FieldTextSingleLine struct {
	Required     bool   `json:"required"`
	DefaultValue string `json:"defaultValue,omitempty"`
}
