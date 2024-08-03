package config

import (
	"errors"
	"os"
	"strconv"
)

// COMMON
var NATS_URL = getenvStr("NATS_URL", "nats://localhost:4222")
var TG_API_BASE_URL = getenvStr("TG_API_BASE_URL", "https://api.telegram.org")
var POSTGRES_CONN_STRING = getenvStr("POSTGRES_CONN_STRING", "")

var USER_BOT_TOKEN = getenvStr("USER_BOT_TOKEN", "")

var ADMIN_BOT_TOKEN = getenvStr("ADMIN_BOT_TOKEN", "")
var ADMIN_AUTH_CODE = getenvStr("ADMIN_AUTH_CODE", "")

// MESSAGES
var NATS_MESSAGES_STREAM = getenvStr("NATS_MESSAGES_STREAM", "")
var NATS_MESSAGES_SUBJECT = getenvStr("NATS_MESSAGES_SUBJECT", "")

var RATE_LIMIT_GLOBAL = getenvInt("RATE_LIMIT_GLOBAL", 30)
var RATE_LIMIT_USER = getenvInt("RATE_LIMIT_USER", 1)

// IMGPROXY
var IMGPROXY_PUBLIC_URL = getenvStr("IMGPROXY_PUBLIC_URL", "")

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

func getenvStr(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func getenvInt(key string, defaultValue int) int {
	s := getenvStr(key, "")
	if s == "" {
		return defaultValue
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}

	return v
}
