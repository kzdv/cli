package config

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"text/template"
)

const KUBECONFIG_TMPL = `
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
        client-secret: {{.ClientSecret}}
        id-token: {{.IdToken}}
        idp-issuer-url: https://auth.denartcc.org
        refresh-token: {{.RefreshToken}}
      name: oidc
`

func GetKubeConfig(idToken, clientSecret, refreshToken string) (string, error) {
	tmpl := template.Must(template.New("kubeconfig").Parse(KUBECONFIG_TMPL))
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, map[string]string{
		"IdToken":      idToken,
		"ClientSecret": clientSecret,
		"RefreshToken": refreshToken,
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func WriteKubeConfig(data string) error {
	dir, err := GetConfigDir()
	if err != nil {
		return err
	}
	return WriteFile(filepath.Join(dir, "kubeconfig"), data)
}

func WriteFile(file, data string) error {
	return ioutil.WriteFile(file, []byte(data), 0600)
}
