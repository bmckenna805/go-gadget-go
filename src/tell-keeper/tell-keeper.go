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

func createRecord(name string) *ksm.Record {
  record := ksm.NewRecordCreate(name, "created by tell-keeper")
  return record
}

func appendLogin(record *ksm.Record, login string) *ksm.Record{
  record.Fields = append(customLogin.Fields,
          ksm.NewLogin(login),
  )
  return record
}

func appendPassword(record *ksm.Record, password string) *ksm.Record{
  record.Fields = append(customLogin.Fields,
          ksm.NewPassword(password),
  )
  return record
}

func createSecret(folder string, record *ksm.Record,  client *ksm.SecretsManager) []interface{} {
  recordUid, err := sm.CreateSecretWithRecordData("", folder, record)
  if err != nil {
        log.Fatal(err)
    } else {
        return recordUid
    }
    return nil
}

func main() {
    // Set up command line arguments
    tokenPtr := flag.String("token", "", "One Time app Token, only required for initial run")
    deleteConfigPtr := flag.Bool("delete-config", false, "Delete an initialized configuration")
    namePtr := flag.String("name", "", "Secret Name that will be added to Keeper")
    passwordPtr := flag.String("password", "", "Secret Password that will be added to Keeper")
    loginPtr := flag.String("login", "", "Associated Secret Login that will be added to Keeper")
    folderPtr := flag.String("folder", "", "Shared folder UID to place secret into")
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
    if *namePtr != "" {
        record := createRecord(*namePtr)
        if *passwordPtr != "" {
            appendPassword(*record, *passwordPtr)
        }
        if *loginPtr != "" {
            appendLogin(*record, *loginPtr)
        }
        secretUid := createSecret(*folderPtr, *record, client)
        fmt.Println(secretUid)
    }
}
