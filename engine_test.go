package xorm

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func TestEngine_CloseGraceful(t *testing.T) {

	dsn := "go_pk:3VaX7GMI7UPLWjuk@tcp(127.0.0.1:9566)/go_pk?charset=utf8mb4"
	engine, err := NewEngine("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}

	engine.DB().SetMaxIdleConns(10)
	//go func() {
	//	rows, _ := engine.DB().Query("SELECT * FROM audit_report_history WHERE audit_status = 4 LIMIT 10 ")
	//	rows.Close()
	//
	//	//engine.DB().Query("SELECT * FROM audit_report_history WHERE audit_status = 4 LIMIT 10 ")
	//	//engine.DB().Query("SELECT * FROM audit_report_history WHERE audit_status = 4 LIMIT 10 ")
	//}()
	//go func() {
	//	sess := engine.NewSession()
	//	sess.Query("SELECT * FROM audit_report_history WHERE audit_status = 4 LIMIT 10 ")
	//	_, err := sess.Query("SELECT * FROM audit_report_history WHERE audit_status = 4 LIMIT 10 ")
	//	t.Logf("err %v", err)
	//
	//}()
	//go func() {
	//	engine.NewSession().Query("SELECT * FROM audit_report_history WHERE audit_status = 4 LIMIT 10 ")
	//}()
	//
	////go func() {
	////	engine.DB().Query("SELECT * FROM audit_report_history WHERE extra_info  LIKE '%https://oss-pk-arab-zl.badambiz.com/icon_9965834_774be0331dfbdee5d56aa2d7b396e9d9.jpeg%' ")
	////}()
	//go func() {
	//	session := engine.NewSession()
	//	session.Query("select last_insert_id(1000)")
	//	r, _ := session.Query("select last_insert_id()")
	//	t.Logf("r %v", r)
	//}()
	engine.CloseGraceful()

	for i := 0; i < 20; i++ {

		//engine.DB().SetMaxIdleConns(10)
		session := engine.NewSession()
		//session.isAutoCommit = false
		go func() {
			//time.Sleep(time.Second * time.Duration(rand.Intn(3)))

			_, _ = session.Prepare().Query("select last_insert_id(?)", 1000)
			r, _ := session.Query("select last_insert_id()")
			//li, _ := rs.LastInsertId()
			t.Logf("r1 %v", r)
		}()
		//go func() {
		//	//time.Sleep(time.Second * time.Duration(rand.Intn(3)))
		//	session.Exec("select last_insert_id(10000)")
		//	r, _ := session.Query("select last_insert_id()")
		//	t.Logf("r2 %v", r)
		//
		//}()
		time.Sleep(time.Second * time.Duration(3))

		//r, _ := session.Query("select last_insert_id()")
		//t.Logf("r %v", r)

		//t.Logf("stats %#v", engine.DB().Stats())
		//t.Log("使用中链接", engine.DB().Stats().InUse)
		//t.Log("使用中的空闲链接", engine.DB().Stats().Idle)
		//t.Log("总链接数", engine.DB().Stats().OpenConnections)
		time.Sleep(time.Second)
	}

}
