package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

const rawInit = `
name: users
type: mysql_gorm
fields:
  - name: fullname
    type: text
    validate: 'required' #https://github.com/go-playground/validator 
    comment: 'fullname of user'
    sql_hint: 'type=BIGINT(20),key=candidate,is_not_null=true, xxx'
  - name: age
    type: int16
    validate: 'gte=1,lte=200' #https://github.com/go-playground/validator 
    comment: 'age'
    sql_hint: 'type=INT(10),key=candidate,is_not_null=true,default=18'
  - name: email
    type: text
    validate: 'required,email' #https://github.com/go-playground/validator 
    comment: 'email'
    sql_hint: 'type=VARCHAR(64),key=unique,is_not_null=true'
  - name: sex
    type: bool
    comment: 'true:male,false:female'    
  - name: account_balance
    type: int64
    comment: 'total money'    
`

func initCmd(ctx *cli.Context) error {
	fmt.Println("this init func")
	return os.WriteFile("collections.yaml", []byte(rawInit), 0644)
}
