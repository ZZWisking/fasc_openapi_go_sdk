package commonModel

// Field 文档控件
type Field struct {
	FieldId             string               `json:"fieldId,omitempty"`
	FieldName           string               `json:"fieldName,omitempty"`
	Position            *FieldPosition       `json:"position,omitempty"`
	FieldType           string               `json:"fieldType,omitempty"`
	FieldTextSingleLine *FieldTextSingleLine `json:"fieldTextSingleLine,omitempty"`
	FieldTextMultiLine  *FieldTextMultiLine  `json:"fieldTextMultiLine,omitempty"`
	FieldCheckBox       *FieldCheckBox       `json:"fieldCheckBox,omitempty"`
}
