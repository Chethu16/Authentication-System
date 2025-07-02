package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_"github.com/lib/pq"
)


type Register struct{
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
		}
type Login struct{
	Email string `json:"email"`
	Password string `json:"password"`
}


func main(){
	db,err:=sql.Open("postgres","postgresql://free_d1iq_user:KHOULvBoVAL5elQcbXqE6ZwcRad8nCDr@dpg-d1iecrali9vc73fu5ueg-a.oregon-postgres.render.com/free_d1iq")
	if err!=nil{
		fmt.Println("error")
		return
	}
	fmt.Println("connected")
	defer db.Close()


	http.HandleFunc("/Register", func(w http.ResponseWriter, r*http.Request){
		var reg Register
		err := json.NewDecoder(r.Body).Decode(&reg)
		if err !=nil{
			fmt.Println("error in json")
			return
		}
		query:=`INSERT INTO users VALUES($1,$2,$3)`
			_,err = db.Exec(query,reg.Name,reg.Email,reg.Password)
			if err!=nil{
				fmt.Println("error in sql query")
				return
			}
			fmt.Println("Register Successfuly")
	})


	http.HandleFunc("/Login", func(w http.ResponseWriter, r*http.Request){
		var log Login
		err:=json.NewDecoder(r.Body).Decode(&log)
		if err !=nil{
		fmt.Println("Error in login json")
		return
	}

		query2:=`SELECT name,email FROM users WHERE email=$1`
		_,err = db.Exec(query2,log.Email)
		if err!=nil{
		fmt.Println("error in sql query2")
		return
		}
		fmt.Println("Login Successfuly")
})

	http.ListenAndServe(":8000",nil)

}