package config

import (
	"log"
	"time"

	ld "github.com/launchdarkly/go-server-sdk/v7"
)

var LDClient *ld.LDClient

// Initiate LaunchDarkly Client
func InitiateLDClient() {
	// Ganti dengan SDK Key dari LaunchDarkly
	sdkKey := "sdk-d4a84fa2-23f0-486b-af59-8985243ede13"
	
	// Inisialisasi LaunchDarkly Client
	client, err := ld.MakeClient(sdkKey, 10*time.Second)
	if err != nil {
		log.Fatalf("Failed to initialize LaunchDarkly client: %v", err)
	}

	// Simpan client ke variabel global
	LDClient = client
}

// Fungsi untuk menutup LaunchDarkly Client saat aplikasi berhenti
func CloseLDClient() {
	if LDClient != nil {
		LDClient.Close()
	}
}
