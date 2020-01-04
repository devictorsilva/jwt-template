package dtos

// DataToEncodeDto get the method we need to encode and other params
type DataToEncodeDto struct {

	// Method refers to method used to encode the claim.
	Method string `json:"method"`

	// ExpiresIn refers to the number of seconds after the token is generated to expire.
	ExpiresIn *int64 `json:"exp_in,omitempty"`

	// The "aud" (audience) claim identifies the recipients that the JWT is intended for.
	Audience *string `json:"aud,omitempty"`

	// The "iss" (issuer) claim identifies the principal that issued the JWT.
	Issuer *string `json:"iss,omitempty"`

	// The "nbf" (not before) claim identifies the time before which
	// the JWT MUST NOT be accepted for processing.
	NotBefore *int64 `json:"nbf_in,omitempty"`

	// The "sub" (subject) claim identifies the principal that is the subject of the JWT.
	Subject *string `json:"sub,omitempty"`
}
