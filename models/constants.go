package models

const (
	ContentTypeHeaderKey        = "Content-Type"
	AppJsonContentTypeVal       = "application/json"
	TextPlainTypeVal            = "text/plain"
	CharsetCodeVal              = "charset=utf-8"
	AppJSONContentTypeHeaderVal = AppJsonContentTypeVal + "; " + CharsetCodeVal
	AuthorizationHeaderKey      = "Authorization"
	AppUrlencodedHeaderVal      = "application/x-www-form-urlencoded"
	TokenTypeBearer             = "Bearer"

	// for Cache
	CacheTokenClientId = "mbr_accts:clientId:"
	CacheEmailCode     = "mbr_accts:emailCode:"
)
