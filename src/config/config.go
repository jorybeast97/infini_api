package config

import (
    "os"
)

type Config struct {
    Port       string
    Secret     string
    UploadDir  string
}

func LoadConfig() Config {
    port := os.Getenv("PORT")
    if port == "" { port = "8080" }
    secret := os.Getenv("INFINI_SECRET")
    if secret == "" { secret = "infini-dev-secret" }
    upload := os.Getenv("UPLOAD_DIR")
    if upload == "" { upload = "uploads" }
    return Config{Port: port, Secret: secret, UploadDir: upload}
}
