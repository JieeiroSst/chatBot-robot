package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Robot struct {
	ID      string     `json:id`
	Name    string     `json:name`
	Gmail   string     `json:gmail`
	Old     int        `json:old`
	Product string     `json:product`
	Date    time.Month `json:date`
}

var robots map[string]*Robot

func handlerRobot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	switch r.Method {
	case "GET":
		fmt.Println("GET/robot/" + vars["ID"])

		robot := robots[vars["ID"]]

		rt, err := json.Marshal(robot)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(w, string(rt))
		return

	case "PUT":
		fmt.Println("put/robot/" + vars["ID"])

		robot := robots[vars["ID"]]

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&robot)

		if err != nil {
			log.Fatal(err)
		}

		robot.ID = vars["ID"]

		retRt, err := json.Marshal(robot)
		if err != nil {
			log.Fatal(err)
		}

	case "DELETE":
		fmt.Println("delete/robot/" + vars["ID"])

		delete(robots, vars["ID"])
		fmt.Fprint(w, "susses")
	}
}

func handlerRobots(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rt, err := json.Marshal(robots)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(w, string(rt))

	case "POST":
		robot := new(Robot)

		robot.ID = GENID()

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&robot)
		if err != nil {
			log.Fatal(err)
		}

		robots[robot.ID] = robot

		rtRt, err := json.Marshal(robot)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(rtRt))
	}
}

func GENID() string {
	id := make([]byte, 20)
	_, err := rand.Read(id[:])
	if err != nil {
		log.Fatal(err)
	}

	id[8] = (id[8] | 0x40) & 0x7F
	id[6] = (id[6] & 0xF) | (4 << 4)
	return fmt.Sprintf("%x-%x-%x-%x-%x", id[:4], id[4:8], id[8:12], id[12:20])
}

func main() {
	robot = map[string]*Robot{
		"12": {"12", "manh quan", "luumanhquan.91@gmail.com", 19, "alphaR", 7},
	}
	router := mux.NewRouter()
	router.HandleFunc("/robot/{id}", handlerRobot)
	router.HandleFunc("/robot", handlerRobots)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
