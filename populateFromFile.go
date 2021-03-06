package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Users struct {
	Publishers  []Publisher  `json:"publisher"`
	Subscribers []Subscriber `json:"subscriber"`
}

type Publisher struct {
	PubID     float64 `json:"pub_id"`
	NodeID    int     `json:"node_id"`
	TopicList []int   `json:"topic_list"`
}

type Subscriber struct {
	SubID     float64 `json:"sub_id"`
	NodeID    int     `json:"node_id"`
	TopicList []int   `json:"topic_list"`
}

func populateFromFile(fileName string, nodeport int) (Users, []map[string]byte, map[int]string) {

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	var user Users
	err = json.Unmarshal(file, &user)
	if err != nil {
		fmt.Println(err)
	}

	nodeIDs := make(map[int]string)

	nodePort := strconv.Itoa(nodeport)
	nodeIDs[0] = "tcp://localhost:1883"
	// insert here the list of broker nodes
	nodeIDs[1] = "tcp://192.168.3.5:" + nodePort

	arraySubTopics := make([]map[string]byte, len(user.Subscribers))
	var str string

	for indexSub, sub := range user.Subscribers {
		subTopics := make(map[string]byte)

		for _, top := range sub.TopicList {
			str = strconv.Itoa(top)
			subTopics[str] = 0
		}
		arraySubTopics[indexSub] = subTopics
	}

	return user, arraySubTopics, nodeIDs
}
