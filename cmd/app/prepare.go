package main

import (
	"fmt"
	"time"

	"github.com/vexulansis/COCParser/internal/db"
	"github.com/vexulansis/COCParser/pkg/api"
)

var accountsNum = 100

func Prepare(AC *api.APIClient, DC *db.DBClient, logger *MainLogger) error {
	// Generating credentials for reserved accounts
	dbPool := db.NewPool(10, DC)
	dbPool.Start()
	dbPool.WG.Add(accountsNum)
	go db.GenerateCredentials(accountsNum, DC)
	dbPool.WG.Wait()
	f := MainLoggerFields{
		Source:  "CREDENTIALS",
		Method:  "GENERATION",
		Subject: "SUCCESS",
	}
	dbPool.Stop()
	logger.Print(f, 0)
	// Getting credentials data from DataBase
	credentials, err := DC.GetCredentials()
	if err != nil {
		return err
	}
	for _, cred := range credentials {
		account := &api.Account{
			Email:    cred.Email,
			Password: cred.Password,
		}
		AC.AccountPool = append(AC.AccountPool, account)
	}
	f = MainLoggerFields{
		Source:  "ACCOUNTPOOL",
		Method:  "GENERATION",
		Subject: "SUCCESS",
	}
	logger.Print(f, fmt.Sprintf("Total accounts created: %d", len(AC.AccountPool)))
	time.Sleep(time.Second * 5)
	// Creating KeyPool
	apiPool := api.NewPool(100, AC)
	apiPool.Start()
	apiPool.WG.Add(len(AC.AccountPool))
	for i, a := range AC.AccountPool {
		task := &api.Task{
			ID:   i,
			Data: a,
		}
		AC.TaskChan <- task
	}
	apiPool.WG.Wait()
	apiPool.Stop()
	f = MainLoggerFields{
		Source:  "KEYPOOL",
		Method:  "GENERATION",
		Subject: "SUCCESS",
	}
	AC.CreateKeyPool()
	logger.Print(f, fmt.Sprintf("Total keys collected: %d", len(AC.KeyPool)))
	time.Sleep(time.Second * 5)
	// Start parsing clans
	apiPool = api.NewPoolWithKeys(AC)
	go api.GenerateTags(AC)
	apiPool.WG.Add(1024000)
	apiPool.Start()
	start := time.Now()
	apiPool.WG.Wait()
	dur := time.Since(start)
	fmt.Printf("TIME TO PARSE 1.024.000 clans: %.2f sec\n", dur.Seconds())
	fmt.Printf("AVERAGE SPEED: %.2f clans/sec", 1024000/dur.Seconds())
	return nil
}
