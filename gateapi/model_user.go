/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type User struct {
	Authorities []GrantedAuthority `json:"authorities,omitempty"`
	AccountNonLocked bool `json:"accountNonLocked,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	AllowedAccounts []string `json:"allowedAccounts,omitempty"`
	Roles []string `json:"roles,omitempty"`
	CredentialsNonExpired bool `json:"credentialsNonExpired,omitempty"`
	Email string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	AccountNonExpired bool `json:"accountNonExpired,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	LastName string `json:"lastName,omitempty"`
}