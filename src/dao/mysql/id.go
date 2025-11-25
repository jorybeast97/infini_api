package mysql

import (
    "crypto/rand"
)

func genID() string {
    b := make([]byte, 8)
    _, err := rand.Read(b)
    if err != nil { return "id" }
    const hex = "0123456789abcdef"
    out := make([]byte, 16)
    for i, v := range b { out[i*2] = hex[v>>4]; out[i*2+1] = hex[v&0x0f] }
    return string(out)
}

