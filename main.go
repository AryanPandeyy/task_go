package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type ListTask struct {
	Tasks []Task
}

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide argument")
	}
	if os.Args[1] == "add" {
		f, err := os.OpenFile("/home/ap/code/fun/tp/go/sample", os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			log.Fatal(err)
		}

		b, e := os.ReadFile("/home/ap/code/fun/tp/go/sample")
		if e != nil {
			log.Fatal(e)
		}

		tmp := ListTask{Tasks: []Task{}}
		json.Unmarshal(b, &tmp)
		fmt.Println("EXIST CONTENT: ", tmp.Tasks)

		var tsk string
		fmt.Print("Enter the task: ")
		fmt.Scan(&tsk)
		id := len(tmp.Tasks) + 1;
		t := Task{id, tsk, "NOTDONE", time.Now().Format(time.DateTime), time.Now().Format(time.DateTime)}
		if err != nil {
			log.Fatal(err)
		}

		if (len(tmp.Tasks) >= 1) {
			tmp.Tasks = append(tmp.Tasks, t)
			b, e := json.MarshalIndent(tmp, "", "  ")
			// b, e := json.Marshal(tmp)
			if e != nil {
				log.Fatal(e)
			}
			fmt.Println("AFTER ADDING CONTENT: ", string(b))

			if _,err := f.Write(b); err != nil {
				log.Fatal(err)
			}
		} else {
			t1 := ListTask{Tasks: []Task{t}}
			b, e := json.MarshalIndent(t1, "", "  ")
			// b, e := json.Marshal(t1)
			if e != nil {
				log.Fatal(e)
			}
			fmt.Println("ADDING CONTENT FIRST: ", string(b))

			if _,err := f.Write(b); err != nil {
				log.Fatal(err)
			}
		}


		f.Close()
	} else if os.Args[1] == "list" {
	}
}
