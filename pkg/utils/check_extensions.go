package utils

import "github.com/mdyssr/azmena/pkg/store"

func extensionSupported(element string) bool {
	for _, e := range store.Extensions {
		if e == element {
			return true
		}
	}
	return false
}

func ValidateExtensions(extensions []string) bool {
	for _, e := range extensions {
		if !extensionSupported(e) {
			return false
		}
	}
	return true
}
