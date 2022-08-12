package config

import "os"

var TrustedProxies = os.Getenv("TRUSTED_PROXIES")
