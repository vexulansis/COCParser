package api

import (
	"fmt"
	"sync"
	"time"

	"github.com/vexulansis/COCParser/internal/algo"
	db "github.com/vexulansis/COCParser/internal/storage"
)

type Client struct {
	Keys    []string
	IP      string
	Auth    Auth
	Storage *db.Storage
}
type Auth struct {
	Login    string
	Password string
}

func (c *Client) UpdateClans(start int, end int) error {
	var wg sync.WaitGroup
	timestart := time.Now()
	wg.Add(end - start)
	worksize := (end - start) / len(c.Keys)
	index := 0
	for i := start; i < end; i += worksize {
		go c.UpdateClansInfo(i, i+worksize, index, &wg)
		index++
	}
	wg.Wait()
	fmt.Printf("Time to execute: %v", time.Since(timestart))
	fmt.Printf("Avg. clans per minute: %.1f", 2950/time.Since(timestart).Minutes())
	return nil
}
func (c *Client) UpdateClansInfo(start int, end int, keyindex int, wg *sync.WaitGroup) error {
	for i := start; i < end; i++ {
		tag := algo.TagFromId(i)
		clan, err := GetClanByTag(tag, c.Keys[keyindex])
		if clan != nil {
			if err != nil {
				return err
			}
			clan.Tag = tag
			//fmt.Printf("Worker: %d, Clan: %v\n", keyindex, clan)
			err = UpdateClanInfo(clan, *c.Storage)
			if err != nil {
				return err
			}
		}
		wg.Done()
	}
	return nil
}
