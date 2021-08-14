package database

import (
	"log"
	"time"
)

func InsertMember(roomID, userID int) error {
	stmt, err := db.Prepare("insert into members (id, roomID, userID, joined) values (default, $1, $2, $3)")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(roomID, userID, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func InsertMewmberWithUniqueRoomID(uniqueRoomID string, userID int) error {
	//get normal room id
	id, err := GetRoomIDByUniqueRoomID(uniqueRoomID)
	if err != nil {
		log.Println(err)
		return err
	}

	err = InsertMember(id, userID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func CheckIfUserIsAMemberOfASpecificRoom(uniqueRoomID string, userID int) error {
	var roomid int
	var id int
	err := db.QueryRow("select distinct m.userid, r.id from members as m join rooms as r on r.id = m.roomid where r.uniqueroomid=$1 and m.userid=$2", uniqueRoomID, userID).Scan(&id, &roomid)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetRoomMembers(uniqueRoomID string) ([]string, error) {
	var usernames []string
	rows, err := db.Query("select u.username from rooms as r join members as m on m.roomid = r.id join users as u on m.userid = u.id where r.uniqueroomid=$1", uniqueRoomID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		usernames = append(usernames, username)
	}

	return usernames, nil
}

func DeleteMember(roomID, userID int) error {
	stmt, err := db.Prepare("delete from members where roomID=$1 and userID=$2")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(roomID, userID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteMemberWithUnqueRoomID(uniqueRoomID string, userID int) error {
	id, err := GetRoomIDByUniqueRoomID(uniqueRoomID)
	if err != nil {
		log.Println(err)
		return err
	}

	err = DeleteMember(id, userID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
