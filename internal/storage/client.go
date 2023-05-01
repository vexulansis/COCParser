package storage

import "fmt"

type DBClient struct {
	Storage     *Storage
	APIAccounts []*APIAccount
}

func NewClient() (*DBClient, error) {
	var err error
	client := new(DBClient)
	client.Storage, err = NewStorage()
	if err != nil {
		return nil, err
	}
	err = client.getAccounts()
	if err != nil {
		return nil, err
	}
	return client, nil
}
func (c *DBClient) getAccounts() error {
	accounts := []*APIAccount{}
	// Rows: email, password
	rows, err := c.Storage.DB.Query(getAccountsQuery)
	if err != nil {
		return err
	}
	defer rows.Close()
	// Iterating through rows, creating new accounts using credentials
	for rows.Next() {
		acc := &APIAccount{}
		err := rows.Scan(&acc.ID, &acc.Credentials.Email, &acc.Credentials.Password)
		if err != nil {
			return err
		}
		accounts = append(accounts, acc)
	}
	c.APIAccounts = accounts
	return nil
}
func (c *DBClient) AddAccount(cred *Credentials) error {
	_, err := c.Storage.DB.Exec(addAccountQuery, cred.Email, cred.Password)
	if err != nil {
		return err
	}
	return nil
}
func (c *DBClient) GenerateAccounts(n int) error {
	for i := 1; i <= n; i++ {
		cred := &Credentials{}
		cred.Email = fmt.Sprintf(c.Storage.Config.APIemail, i)
		cred.Password = c.Storage.Config.APIpassword
		err := c.AddAccount(cred)
		if err != nil {
			return err
		}
	}
	return c.getAccounts()
}
