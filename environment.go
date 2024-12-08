package metaorm

import "github.com/metadiv-io/env"

var (
	METAORM_HOST        = env.New("METAORM_HOST", false)
	METAORM_PORT        = env.New("METAORM_PORT", false)
	METAORM_USERNAME    = env.New("METAORM_USERNAME", false)
	METAORM_PASSWORD    = env.New("METAORM_PASSWORD", false)
	METAORM_DATABASE    = env.New("METAORM_DATABASE", false)
	METAORM_ENCRYPT_KEY = env.New("METAORM_ENCRYPT_KEY", false)
)
