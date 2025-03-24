package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	ld "github.com/launchdarkly/go-server-sdk/v7"
)

var LDClient *ld.LDClient


// Initiate LaunchDarkly Client
func InitiateLDClient() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or failed to load")
	}

	// Ambil SDK Key dari environment variable
	sdkKey := os.Getenv("SDK_LD")
	if sdkKey == "" {
		log.Fatal("Error: SDK_LD environment variable is missing")
	}

	// Inisialisasi LaunchDarkly Client dengan timeout 10 detik
	client, err := ld.MakeClient(sdkKey, 10*time.Second)
	if err != nil {
		log.Fatalf("Failed to initialize LaunchDarkly client: %v", err)
	}

	// Simpan client ke variabel global
	LDClient = client
	log.Println("LaunchDarkly client successfully initialized")
}
// Fungsi untuk menutup LaunchDarkly Client saat aplikasi berhenti
func CloseLDClient() {
	if LDClient != nil {
		LDClient.Close()
	}
}
