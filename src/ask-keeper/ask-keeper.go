package main

import (
    ksm "github.com/keeper-security/secrets-manager-go/core"
    "fmt"
    "log"
    "flag"
    "os"
)

func deleteConfig(configLoc string) {
    fmt.Println("Deleting KSM Configuration")
    err := os.Remove(configLoc)
    if err != nil {
        log.Fatal(err)
    }
}

func getClient(token string, configLoc string) *ksm.SecretsManager {
	clientOptions := &ksm.ClientOptions{
		Token:  token,
		Config: ksm.NewFileKeyValueStorage(configLoc)}
	return ksm.NewSecretsManager(clientOptions)
}

func getPassword(secretUID string, client *ksm.SecretsManager) []interface{} {
    notation := secretUID + "/field/password"
    value, err := client.GetNotation(notation)
    if err != nil {
        log.Fatal(err)
    } else {
        return value
    }
    return nil
}

func main() {
    // Set up command line arguments
    tokenPtr := flag.String("token", "", "One Time app Token, only required for initial run")
    deleteConfigPtr := flag.Bool("delete-config", false, "Delete an initialized configuration")
    secretPtr := flag.String("secret", "", "Secret UID to look up")
    flag.Parse()

    // Establish ask-keeper config location
    configLoc := os.Getenv("HOME") + "/.ksm/ask-keeper-config.json"

    // If requested by flag, destroy the KSM config
    if *deleteConfigPtr == true {
        deleteConfig(configLoc)
    }

    // Establish a client
    client := getClient(*tokenPtr, configLoc)

    // Get and print the password, if not empty
    if *secretPtr != "" {
        secret := getPassword(*secretPtr, client)
        fmt.Println(secret)
    }
}
