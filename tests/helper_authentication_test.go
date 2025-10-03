package imagekit_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/stainless-sdks/imagekit-go"
	"github.com/stainless-sdks/imagekit-go/option"
)

func TestGetAuthenticationParametersDetailed(t *testing.T) {
	privateKey := "private_key_test"
	client := imagekit.NewClient(option.WithPrivateKey(privateKey))

	t.Run("should return correct authentication parameters with provided token and expire", func(t *testing.T) {
		token := "your_token"
		expire := int64(1582269249)

		params := client.Helper.GetAuthenticationParameters(token, expire)

		// Expected exact match with Node.js output
		expectedSignature := "e71bcd6031016b060d349d212e23e85c791decdd"

		if params["token"] != token {
			t.Errorf("Expected token %s, got %v", token, params["token"])
		}
		if params["expire"] != expire {
			t.Errorf("Expected expire %d, got %v", expire, params["expire"])
		}
		if params["signature"] != expectedSignature {
			t.Errorf("Expected signature %s, got %v", expectedSignature, params["signature"])
		}
	})

	t.Run("should return authentication parameters with required properties when no params provided", func(t *testing.T) {
		params := client.Helper.GetAuthenticationParameters("", 0)

		// Check that all required properties exist
		if _, exists := params["token"]; !exists {
			t.Errorf("Expected token parameter")
		}
		if _, exists := params["expire"]; !exists {
			t.Errorf("Expected expire parameter")
		}
		if _, exists := params["signature"]; !exists {
			t.Errorf("Expected signature parameter")
		}

		// Token should be a UUID (36 characters with dashes)
		token, ok := params["token"].(string)
		if !ok {
			t.Errorf("Expected token to be string, got %T", params["token"])
		} else {
			uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
			if !uuidRegex.MatchString(token) {
				t.Errorf("Expected token to be UUID format, got %s", token)
			}
		}

		// Expire should be a number greater than current time
		expire, ok := params["expire"].(int64)
		if !ok {
			t.Errorf("Expected expire to be int64, got %T", params["expire"])
		} else {
			currentTime := time.Now().Unix()
			if expire <= currentTime {
				t.Errorf("Expected expire %d to be greater than current time %d", expire, currentTime)
			}
		}

		// Signature should be a hex string (40 characters for HMAC-SHA1)
		signature, ok := params["signature"].(string)
		if !ok {
			t.Errorf("Expected signature to be string, got %T", params["signature"])
		} else {
			signatureRegex := regexp.MustCompile(`^[a-f0-9]{40}$`)
			if !signatureRegex.MatchString(signature) {
				t.Errorf("Expected signature to be 40 character hex string, got %s", signature)
			}
		}
	})

	t.Run("should handle edge case with expire time 0", func(t *testing.T) {
		token := "test-token"
		expire := int64(0)

		params := client.Helper.GetAuthenticationParameters(token, expire)

		if params["token"] != token {
			t.Errorf("Expected token %s, got %v", token, params["token"])
		}

		// When expire is 0 (falsy), it should use default expire time (30 minutes from now)
		expireResult, ok := params["expire"].(int64)
		if !ok {
			t.Errorf("Expected expire to be int64, got %T", params["expire"])
		} else {
			expectedExpire := time.Now().Unix() + 60*30
			// Allow a 10 second tolerance for test execution time
			if expireResult < expectedExpire-10 || expireResult > expectedExpire+10 {
				t.Errorf("Expected expire to be close to %d (30 minutes from now), got %d", expectedExpire, expireResult)
			}
		}

		// Signature should be a hex string (40 characters for HMAC-SHA1)
		signature, ok := params["signature"].(string)
		if !ok {
			t.Errorf("Expected signature to be string, got %T", params["signature"])
		} else {
			signatureRegex := regexp.MustCompile(`^[a-f0-9]{40}$`)
			if !signatureRegex.MatchString(signature) {
				t.Errorf("Expected signature to be 40 character hex string, got %s", signature)
			}
		}
	})

	t.Run("should handle empty string token", func(t *testing.T) {
		token := "" // Empty string is falsy
		expire := int64(1582269249)

		params := client.Helper.GetAuthenticationParameters(token, expire)

		// Since empty string is falsy, it should generate a UUID
		tokenResult, ok := params["token"].(string)
		if !ok {
			t.Errorf("Expected token to be string, got %T", params["token"])
		} else {
			if tokenResult == "" {
				t.Errorf("Expected token to be generated when empty string is provided")
			}
			uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
			if !uuidRegex.MatchString(tokenResult) {
				t.Errorf("Expected generated token to be UUID format, got %s", tokenResult)
			}
		}

		if params["expire"] != expire {
			t.Errorf("Expected expire %d, got %v", expire, params["expire"])
		}

		// Signature should be a hex string (40 characters for HMAC-SHA1)
		signature, ok := params["signature"].(string)
		if !ok {
			t.Errorf("Expected signature to be string, got %T", params["signature"])
		} else {
			signatureRegex := regexp.MustCompile(`^[a-f0-9]{40}$`)
			if !signatureRegex.MatchString(signature) {
				t.Errorf("Expected signature to be 40 character hex string, got %s", signature)
			}
		}
	})

	// Additional test to match Node.js behavior - panic when private key is not provided
	t.Run("should panic when private key is not provided", func(t *testing.T) {
		client := imagekit.NewClient(option.WithPrivateKey(""))

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when private key is empty")
			}
		}()

		client.Helper.GetAuthenticationParameters("test", 123)
	})
}
