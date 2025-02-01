package main

import (
	"fmt"
	"encoding/json"
	"log"
	// "io"
	// "io/fs"
	"os"
)

type Task struct {
	Id			int
	Description	string
	Status		string
	CreatedAt	int
	UpdatedAt	int
}

func main() {
	var cmd string
	fmt.Scan(&cmd)
	if cmd == "add" {
		var tsk string
		fmt.Print("Enter the task: ")
		fmt.Scan(&tsk)
		t1 := Task{1, tsk, "NOTDONE", 1, 1}
		b, e := json.Marshal(t1)
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println(string(b))

		f,err := os.OpenFile("/home/ap/code/fun/tp/go/sample", os.O_RDWR|os.O_CREATE, 0644);
		if  err != nil {
			log.Fatal(err)
		}

		if _,err := f.WriteString(string(b)); err != nil {
			log.Fatal(err)
		}

		// err := os.WriteFile("/home/ap/code/fun/tp/go/sample", b, 0644);
		// if  err != nil {
		// 	log.Fatal(err)
		// }

		f.Close()
	} else if cmd == "list" {
	}
}
