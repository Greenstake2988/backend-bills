package utils

import (
	"crypto/ecdsa"
    "github.com/ethereum/go-ethereum/crypto"
	"sync"
)


var jwtPrivateKey *ecdsa.PrivateKey
var once sync.Once

// generateECDSAPrivateKey genera la clave privada ECDSA utilizando el paquete github.com/ethereum/go-ethereum/crypto.
func generateECDSAPrivateKey() (*ecdsa.PrivateKey, error) {
    // Generar la clave privada ECDSA
    privateKeyECDSA, err := crypto.GenerateKey()
    if err != nil {
        return nil, err
    }

    return privateKeyECDSA, nil
}

// GetJWTPrivateKey devuelve la clave privada ECDSA para firmar los tokens JWT.
// Si la clave aún no se ha generado, se generará una vez y se almacenará en la variable global jwtPrivateKey.
func GetJWTPrivateKey() (*ecdsa.PrivateKey, error) {
    once.Do(func() {
        // Generar la clave privada una sola vez
        privateKey, err := generateECDSAPrivateKey()
        if err != nil {
            panic("Error generando la clave privada ECDSA: " + err.Error())
        }

        jwtPrivateKey = privateKey
    })

    return jwtPrivateKey, nil
}