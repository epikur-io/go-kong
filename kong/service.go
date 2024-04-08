package kong

import "encoding/json"

// Service represents a Service in Kong.
// Read https://docs.konghq.com/gateway/latest/admin-api/#service-object
// +k8s:deepcopy-gen=true
type Service struct {
	ClientCertificate *Certificate `json:"client_certificate" yaml:"client_certificate"`
	ConnectTimeout    *int         `json:"connect_timeout" yaml:"connect_timeout"`
	CreatedAt         *int         `json:"created_at" yaml:"created_at"`
	Enabled           *bool        `json:"enabled" yaml:"enabled"`
	Host              *string      `json:"host" yaml:"host"`
	ID                *string      `json:"id" yaml:"id"`
	Name              *string      `json:"name" yaml:"name"`
	Path              *string      `json:"path" yaml:"path"`
	Port              *int         `json:"port" yaml:"port"`
	Protocol          *string      `json:"protocol" yaml:"protocol"`
	ReadTimeout       *int         `json:"read_timeout" yaml:"read_timeout"`
	Retries           *int         `json:"retries" yaml:"retries"`
	UpdatedAt         *int         `json:"updated_at" yaml:"updated_at"`
	URL               *string      `json:"url" yaml:"url"`
	WriteTimeout      *int         `json:"write_timeout" yaml:"write_timeout"`
	Tags              []*string    `json:"tags" yaml:"tags"`
	TLSVerify         *bool        `json:"tls_verify" yaml:"tls_verify"`
	TLSVerifyDepth    *int         `json:"tls_verify_depth" yaml:"tls_verify_depth"`
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

	// Remove null values in first layer from the map
	for key, value := range data {
		if value == nil {
			delete(data, key)
		}
	}

	// Marshal the modified map back to JSON bytes
	return json.Marshal(data)
}
