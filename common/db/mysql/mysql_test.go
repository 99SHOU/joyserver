package mysql_test

import (
	"github.com/99SHOU/joyserver/common/db/mysql"
)

func Example() {
	mysql.Open("root:123456@/test_mysql")
}
