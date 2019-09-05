package disk

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/HouzuoGuo/tiedot/db"
)

const (
	DatabasePath       = "signaller.db" // TODO: make changing the path of database
	UsersColName       = "users"
	RoomsColName       = "rooms"
	EventsColName      = "events"
	RoomAliasesColName = "roomAliases"
)

func createCollections(database *db.DB) {
	if err := database.Create(UsersColName); err != nil {
		panic(err)
	}
	if err := database.Create(RoomsColName); err != nil {
		panic(err)
	}
	if err := database.Create(EventsColName); err != nil {
		panic(err)
	}
	if err := database.Create(RoomAliasesColName); err != nil {
		panic(err)
	}
}

func defaultValidationUsernameFunc(userName string) error {
	const re = `^\w{5,}$`

	if !regexp.MustCompile(re).MatchString(userName) {
		return fmt.Errorf("username does not match %s", re)
	}

	return nil
}

type DiskBackend struct {
	currentDatabase      *db.DB
	users                *db.Col
	rooms                *db.Col
	events               *db.Col
	roomAliases          *db.Col
	hostname             string
	validateUsernameFunc func(string) error // TODO: create ability to redefine validation func
	mutex                sync.RWMutex
}

func NewDiskBackend(hostname string) *DiskBackend {
	database, err := db.OpenDB(DatabasePath)
	if err != nil {
		panic(err)
	}
	createCollections(database)

	return &DiskBackend{
		currentDatabase:      database,
		users:                database.Use(UsersColName),
		rooms:                database.Use(RoomsColName),
		events:               database.Use(EventsColName),
		roomAliases:          database.Use(RoomAliasesColName),
		hostname:             hostname,
		validateUsernameFunc: defaultValidationUsernameFunc}
}
