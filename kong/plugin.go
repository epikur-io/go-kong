package kong

import "encoding/json"

// Plugin represents a Plugin in Kong.
// Read https://docs.konghq.com/gateway/latest/admin-api/#plugin-object
// +k8s:deepcopy-gen=true
type Plugin struct {
	CreatedAt     *int            `json:"created_at" yaml:"created_at"`
	ID            *string         `json:"id" yaml:"id"`
	Name          *string         `json:"name" yaml:"name"`
	InstanceName  *string         `json:"instance_name" yaml:"instance_name"`
	Route         *Route          `json:"route" yaml:"route"`
	Service       *Service        `json:"service" yaml:"service"`
	Consumer      *Consumer       `json:"consumer" yaml:"consumer"`
	ConsumerGroup *ConsumerGroup  `json:"consumer_group" yaml:"consumer_group"`
	Config        Configuration   `json:"config" yaml:"config"`
	Enabled       *bool           `json:"enabled" yaml:"enabled"`
	RunOn         *string         `json:"run_on" yaml:"run_on"`
	Ordering      *PluginOrdering `json:"ordering" yaml:"ordering"`
	Protocols     []*string       `json:"protocols" yaml:"protocols"`
	Tags          []*string       `json:"tags" yaml:"tags"`
}

// PluginOrdering contains before or after instructions for plugin execution order
// +k8s:deepcopy-gen=true
type PluginOrdering struct {
	Before PluginOrderingPhase `json:"before,omitempty"`
	After  PluginOrderingPhase `json:"after,omitempty"`
}

// TODO this explanation is bad, but the organization of the overall struct defies a good explanation at this level
// beyond "they're the things used in PluginOrdering. This is a map from a phase name (which can only be "access"
// in the initial 3.0 release) to a list of plugins that the plugin containing the PluginOrdering should run before
// or after

// PluginOrderingPhase indicates which plugins in a phase should affect the target plugin's order
// +k8s:deepcopy-gen=true
type PluginOrderingPhase map[string][]string

// FriendlyName returns the endpoint key name or ID.
func (p *Plugin) FriendlyName() string {
	if p.Name != nil {
		return *p.Name
	}
	if p.ID != nil {
		return *p.ID
	}
	return ""
}

func (p *Plugin) MarshalJSON() ([]byte, error) {
	// Create an alias to avoid infinite recursion
	type Alias Plugin
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	// Marshal the struct to JSON bytes
	jsonBytes, err := json.Marshal(aux)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON bytes into a map
	var data map[string]any
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return nil, err
	}

	// Remove null values in first layer from the map
	for key, value := range data {
		if value == nil {
			delete(data, key)
		}
	}

	// Marshal the modified map back to JSON bytes
	return json.Marshal(data)
}
