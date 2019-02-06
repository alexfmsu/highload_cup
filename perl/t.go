package main

import (
	"encoding/json"
)

type Account3 struct {
	Id      uint32            `json:id`
	Email   string            `json:email`
	Fname   string            `json:fname`
	Sname   string            `json:sname`
	Phone   string            `json:phone`
	Sex     string            `json:sex`
	Birth   uint64            `json:birth`
	Country string            `json:country`
	City    string            `json:city`
	Joined  uint32            `json:joined`
	Status  string            `json:status`
	Premium map[string]uint32 `json:premium`
}

func main() {
	var account Account3

	body := `{"likes":[{"ts":1458532958,"id":2068},{"ts":1530543279,"id":8192},{"ts":1504195330,"id":14222},{"ts":1531801936,"id":13052},{"ts":1483571968,"id":436},{"ts":1477798464,"id":13798},{"ts":1540794608,"id":27366},{"ts":1501938063,"id":2640},{"ts":1475711148,"id":18730},{"ts":1483182988,"id":1438},{"ts":1502868081,"id":20254},{"ts":1457154892,"id":5898},{"ts":1499109873,"id":27434},{"ts":1492498978,"id":3906},{"ts":1492727181,"id":148},{"ts":1455009416,"id":5130},{"ts":1499787738,"id":20974},{"ts":1457697787,"id":18186},{"ts":1468773125,"id":27452},{"ts":1514133181,"id":23554},{"ts":1482863630,"id":5190},{"ts":1459079122,"id":4274},{"ts":1529525143,"id":28170},{"ts":1501810072,"id":20680},{"ts":1454398908,"id":17898},{"ts":1499489038,"id":19682},{"ts":1507146242,"id":5252},{"ts":1520412653,"id":18584},{"ts":1507117971,"id":7142},{"ts":1475659744,"id":12788},{"ts":1527026490,"id":10682},{"ts":1455581885,"id":5518},{"ts":1501760787,"id":21064}],"email":"tifiminisugirwedy@ymail.com","birth":642681139,"status":"\u0432\u0441\u0451 \u0441\u043b\u043e\u0436\u043d\u043e","joined":1344124800,"sex":"m","interests":["\u041d\u0430 \u043e\u0442\u043a\u0440\u044b\u0442\u043e\u043c \u0432\u043e\u0437\u0434\u0443\u0445\u0435","\u0412\u0435\u0447\u0435\u0440 \u0441 \u0434\u0440\u0443\u0437\u044c\u044f\u043c\u0438"],"sname":"\u0424\u0430\u0430\u0442\u043e\u0432\u0438\u0447","phone":"8(918)5543777","country":"\u0420\u043e\u0441\u043b\u044f\u043d\u0434\u0438\u044f","id":30001,"premium":{"start":1519546195,"finish":1527408595},"fname":"\u041d\u0438\u043a\u0438\u0442\u0430"}`

	if err := json.Unmarshal([]byte(body), &account); err != nil {
		print("errorr\n")
	} else {
		print("ok\n")
	}
}
