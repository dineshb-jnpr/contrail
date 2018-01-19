package models

// ServiceInstance

import "encoding/json"

// ServiceInstance
type ServiceInstance struct {
	ServiceInstanceBindings   *KeyValuePairs       `json:"service_instance_bindings,omitempty"`
	Perms2                    *PermType2           `json:"perms2,omitempty"`
	ParentUUID                string               `json:"parent_uuid,omitempty"`
	IDPerms                   *IdPermsType         `json:"id_perms,omitempty"`
	ServiceInstanceProperties *ServiceInstanceType `json:"service_instance_properties,omitempty"`
	Annotations               *KeyValuePairs       `json:"annotations,omitempty"`
	UUID                      string               `json:"uuid,omitempty"`
	ParentType                string               `json:"parent_type,omitempty"`
	FQName                    []string             `json:"fq_name,omitempty"`
	DisplayName               string               `json:"display_name,omitempty"`

	ServiceTemplateRefs []*ServiceInstanceServiceTemplateRef `json:"service_template_refs,omitempty"`
	InstanceIPRefs      []*ServiceInstanceInstanceIPRef      `json:"instance_ip_refs,omitempty"`

	PortTuples []*PortTuple `json:"port_tuples,omitempty"`
}

// ServiceInstanceServiceTemplateRef references each other
type ServiceInstanceServiceTemplateRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// ServiceInstanceInstanceIPRef references each other
type ServiceInstanceInstanceIPRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *ServiceInterfaceTag
}

// String returns json representation of the object
func (model *ServiceInstance) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceInstance makes ServiceInstance
func MakeServiceInstance() *ServiceInstance {
	return &ServiceInstance{
		//TODO(nati): Apply default
		FQName:                    []string{},
		DisplayName:               "",
		ServiceInstanceProperties: MakeServiceInstanceType(),
		Annotations:               MakeKeyValuePairs(),
		UUID:                      "",
		ParentType:                "",
		ServiceInstanceBindings:   MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		ParentUUID:                "",
		IDPerms:                   MakeIdPermsType(),
	}
}

// MakeServiceInstanceSlice() makes a slice of ServiceInstance
func MakeServiceInstanceSlice() []*ServiceInstance {
	return []*ServiceInstance{}
}
