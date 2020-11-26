/**
  @author:tyj
  @date:2020/11/26
  @note:1670171244@qq.com
  @TODO:
  @Param:
  @return:
**/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
)
func checkErr(err error) {
	if err != nil {
		//panic(err)
	}
}
func main() {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=tyj sslmode=disable")
	//查询数据
	rows, err := db.Query("SELECT * FROM company;")
	//打印数据
	for rows.Next() {
		var id int
		var name string
		var age int
		var address string
		var salary int
		var join_date string
		err = rows.Scan(&id, &name, &age, &address,&salary,&join_date)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(age)
		fmt.Println(address)
		fmt.Println(salary)
		fmt.Println(join_date)
	}
	db.Close()
}

