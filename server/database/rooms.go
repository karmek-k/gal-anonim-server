package database

import (
	"errors"
	"log"
	"time"
)

func InsertRoom(uniqueRoomID string, roomName string, passwordHash string, ownerID int) error {
	stmt, err := db.Prepare("insert into rooms (id, uniqueRoomID, roomName, passwordHash, ownerID, created, updated) values (default, $1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(uniqueRoomID, roomName, passwordHash, ownerID, time.Now(), time.Now())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func CheckIfUniqueRoomIDExists(uniqueRoomID string) error {
	var id string
	err := db.QueryRow("select uniqueRoomID from rooms where uniqueRoomID = $1", uniqueRoomID).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func ChceckIfUserIsOwnerOfTheRoom(uniqueRoomID string, userID int) error {
	var ownerID int
	err := db.QueryRow("select ownerID from rooms where uniqueRoomID=$1", uniqueRoomID).Scan(&ownerID)
	if err != nil {
		log.Println(err)
		return err
	}

	if ownerID != userID {
		return errors.New("Not the owner")
	}

	return nil

}

func DeleteRoom(uniqueRoomID string) error {
	stmt, err := db.Prepare("delete from rooms where uniqueRoomID=$1")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(uniqueRoomID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
