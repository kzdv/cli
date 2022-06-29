package config

import (
	"os"
	"testing"
)

func TestGetKubeConfig(t *testing.T) {
	tests := []struct {
		name         string
		idToken      string
		clientSecret string
		refreshToken string
		expected     string
	}{
		{
			name:         "test1",
			idToken:      "idToken",
			clientSecret: "clientSecret",
			refreshToken: "refreshToken",
			expected: `
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJlRENDQVIyZ0F3SUJBZ0lCQURBS0JnZ3Foa2pPUFFRREFqQWpNU0V3SHdZRFZRUUREQmhyTTNNdGMyVnkKZG1WeUxXTmhRREUyTlRVMU1qTXpOamd3SGhjTk1qSXdOakU0TURNek5qQTRXaGNOTXpJd05qRTFNRE16TmpBNApXakFqTVNFd0h3WURWUVFEREJock0zTXRjMlZ5ZG1WeUxXTmhRREUyTlRVMU1qTXpOamd3V1RBVEJnY3Foa2pPClBRSUJCZ2dxaGtqT1BRTUJCd05DQUFRWWZYbU8xeUpUUmJMd3ZuaURyTE5WZjBZYUZibHRCUGNtQmxOcHh5NDUKREhmNHg1NlBUVjZTTHFiYWdEWllPbXVVbmRGRnJaNjVKaWE2aTJKcWhwRzlvMEl3UURBT0JnTlZIUThCQWY4RQpCQU1DQXFRd0R3WURWUjBUQVFIL0JBVXdBd0VCL3pBZEJnTlZIUTRFRmdRVUtNL1VDNnk3eGdNNFllVDhsRE43CkpvdVR0NkV3Q2dZSUtvWkl6ajBFQXdJRFNRQXdSZ0loQUsvRFRPbFVoUE1UOVNidUxiVGdhQjJrTVBvaVlaTGkKRmNzUnNONlVmTnhjQWlFQWtpQml4akdUUXNSd1NkVm1hMFY2RDZQS05GK05sTWQrVFZadXRjOXhEV0k9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    server: https://server1.denartcc.org:6443
  name: kzdv
contexts:
- context:
    cluster: kzdv
    user: kzdv
  name: kzdv
current-context: kzdv
kind: Config
preferences: {}
users:
- name: kzdv
  user:
    auth-provider:
      config:
        client-id: kubernetes
        client-secret: clientSecret
        id-token: idToken
        idp-issuer-url: https://auth.denartcc.org
        refresh-token: refreshToken
      name: oidc
`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := GetKubeConfig(test.idToken, test.clientSecret, test.refreshToken)
			if err != nil {
				t.Errorf("GetKubeConfig() error = %v", err)
				return
			}
			if actual != test.expected {
				t.Errorf("expected %s, got %s", test.expected, actual)
			}
		})
	}
}

func TestWriteFile(t *testing.T) {
	cases := []struct {
		name       string
		kubeConfig string
		filename   string
	}{
		{
			name:       "test1",
			kubeConfig: "test12345",
			filename:   "test.yaml",
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			err := WriteFile(test.filename, test.kubeConfig)
			if err != nil {
				t.Errorf("WriteFile() error = %v", err)
				return
			}
			if _, err := os.Stat(test.filename); os.IsNotExist(err) {
				t.Errorf("WriteFile() error = %v", err)
				return
			}
			os.Remove(test.filename)
		})
	}
}
