package main

import (
    ksm "github.com/keeper-security/secrets-manager-go/core"
    "fmt"
    "log"
    "flag"
    "os"
)

func get_client(token string, config_loc string) *ksm.SecretsManager {
	clientOptions := &ksm.ClientOptions{
		Token:  token,
		Config: ksm.NewFileKeyValueStorage(config_loc)}
	return ksm.NewSecretsManager(clientOptions)
}

func get_password(secretUID string, client *ksm.SecretsManager) []interface{} {
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
    //deleteConfigPtr := flag.Bool("delete-config", false, "Delete an initialized configuration")
    secretPtr := flag.String("secret", "", "Secret UID to look up")
    flag.Parse()

    // Establish ask-keeper config location
    config_loc := os.Getenv("HOME") + "/.ksm/ask-keeper-config.json"

    // Establish a client
	client := get_client(*tokenPtr, config_loc)

    // Get and print the password
    secret := get_password(*secretPtr, client)
    fmt.Println(secret)
}
