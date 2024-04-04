package kong

import "encoding/json"

// Service represents a Service in Kong.
// Read https://docs.konghq.com/gateway/latest/admin-api/#service-object
// +k8s:deepcopy-gen=true
type Service struct {
	ClientCertificate *Certificate `json:"client_certificate,omitempty" yaml:"client_certificate,omitempty"`
	ConnectTimeout    *int         `json:"connect_timeout,omitempty" yaml:"connect_timeout,omitempty"`
	CreatedAt         *int         `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Enabled           *bool        `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Host              *string      `json:"host,omitempty" yaml:"host,omitempty"`
	ID                *string      `json:"id,omitempty" yaml:"id,omitempty"`
	Name              *string      `json:"name,omitempty" yaml:"name,omitempty"`
	Path              *string      `json:"path,omitempty" yaml:"path,omitempty"`
	Port              *int         `json:"port,omitempty" yaml:"port,omitempty"`
	Protocol          *string      `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	ReadTimeout       *int         `json:"read_timeout,omitempty" yaml:"read_timeout,omitempty"`
	Retries           *int         `json:"retries,omitempty" yaml:"retries,omitempty"`
	UpdatedAt         *int         `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	URL               *string      `json:"url,omitempty" yaml:"url,omitempty"`
	WriteTimeout      *int         `json:"write_timeout,omitempty" yaml:"write_timeout,omitempty"`
	Tags              []*string    `json:"tags" yaml:"tags"`
	TLSVerify         *bool        `json:"tls_verify,omitempty" yaml:"tls_verify,omitempty"`
	TLSVerifyDepth    *int         `json:"tls_verify_depth,omitempty" yaml:"tls_verify_depth,omitempty"`
	CACertificates    []*string    `json:"ca_certificates" yaml:"ca_certificates"`
}

// FriendlyName returns the endpoint key name or ID.
func (s *Service) FriendlyName() string {
	if s.Name != nil {
		return *s.Name
	}
	if s.ID != nil {
		return *s.ID
	}
	return ""
}

func (s *Service) MarshalJSON() ([]byte, error) {
	// Create an alias to avoid infinite recursion
	type Alias Service
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
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

	// Remove null values recursively from the map
	s.removeNullValues(data)

	// Marshal the modified map back to JSON bytes
	return json.Marshal(data)
}

func (s *Service) removeNullValues(data map[string]any) {
	for key, value := range data {
		if value == nil {
			delete(data, key)
		} else if nestedMap, ok := value.(map[string]any); ok {
			s.removeNullValues(nestedMap)
		}
	}
}
