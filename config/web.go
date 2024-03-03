package config

import "os"

var TrustedProxies = os.Getenv("TRUSTED_PROXIES")

var JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
