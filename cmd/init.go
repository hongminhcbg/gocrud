package main

import (
	"os"

	"github.com/urfave/cli/v3"
)

const rawInit = `
# data_type
#   text => free text 
#   bool => true, false
#   json =>  https://github.com/go-gorm/datatypes/blob/master/json.go
#   int8 => uint8
#   int16 => int16
#   int32 => int32
#   int64 => int64
name: users
type: mysql_gorm
fields:
  - name: fullname # required, snake case
    type: text # required, in list ['text', 'bool', 'json', 'int8', 'int16', 'int32', 'int64']
    validate: 'required' # optional, https://github.com/go-playground/validator 
    comment: 'fullname of user' # optional
    sql_hint: 'type=BIGINT(20),key=candidate,is_not_null=true' # optional, hint for create table,  support flags [], if empty create default 
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
  - name: raw
    type: json
    comment: 'raw  json fo user'
`

func initCmd(ctx *cli.Context) error {
	return os.WriteFile("collections.yaml", []byte(rawInit), 0644)
}
