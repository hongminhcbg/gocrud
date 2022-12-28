package YOUR_PACKAGE_HERE
type Users struct {
  Fullname 	string 	`gorm:"column:fullname" json:"fullname,omitempty"` // fullname of user
  Age 	int16 	`gorm:"column:age" json:"age,omitempty"` // age
  Sex 	bool 	`gorm:"column:sex" json:"sex,omitempty"` // true:male,false:female
  AccountBalance 	int64 	`gorm:"column:account_balance" json:"account_balance,omitempty"` // total money
}
func (Users) TableName() string {
  return "users"
}