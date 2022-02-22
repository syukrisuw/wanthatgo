package wtdtomdl

type WtDtoModel struct {
	ID          uint64  `json:"ID"`
	ObjTypeName string  `json:"obj_type_name"`
	ObjTypeDesc *string `json:"obj_type_desc,omitempty"`
}
