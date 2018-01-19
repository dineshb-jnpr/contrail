package models

// BGPRouter

import "encoding/json"

// BGPRouter
type BGPRouter struct {
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
}

// String returns json representation of the object
func (model *BGPRouter) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeBGPRouter makes BGPRouter
func MakeBGPRouter() *BGPRouter {
	return &BGPRouter{
		//TODO(nati): Apply default
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
	}
}

// MakeBGPRouterSlice() makes a slice of BGPRouter
func MakeBGPRouterSlice() []*BGPRouter {
	return []*BGPRouter{}
}
