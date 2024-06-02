package helm_test

import (
	"shuvava/helm-decrypt-aes"
	"testing"
)

func TestDecryptAES(t *testing.T) {
	t.Run("Give encrypted data", func(t *testing.T) {
		encryptedData := "30tEfhuJSVRhpG97XCuWgz2okj7L8vQ1s6V9zVUPeDQ="
		encryptionKey := "secretkey"
		t.Run("When app decrypt data", func(t *testing.T) {
			data, err := helm.DecryptAES(encryptedData, encryptionKey)
			t.Run("Then no error should be returned", func(t *testing.T) {
				if err != nil {
					t.Errorf("Then err should be nil, but got %v", err)
				}
			})
			t.Run("Then encryption key should be correct", func(t *testing.T) {
				if string(data) != "plaintext" {
					t.Errorf("Then encryption key should be correct, but got %v", string(data))
				}
			})
		})
	})
}
