package test

import (
	"fmt"
	"testing"

	dm "github.com/nfjBill/gorm-driver-dm"
	"gorm.io/gorm"
)

type GormTable struct {
	gorm.Model
	Name string
	Age  int
}

func TestGormDM(t *testing.T) {

	dsn := "dm://SYSDBA:SYSDBA@10.0.93.38:5236?autoCommit=false"
	db, err := gorm.Open(dm.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	dm.TB(db)

	// if dm.Table().HasTable(&GormTable{}) {
	// 	err = dm.Table().DropTable(&GormTable{})
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }

	err = db.AutoMigrate(&GormTable{})
	if err != nil {
		fmt.Println(err)
	}

	// Create
	db.Create(&GormTable{Name: "S1", Age: 20})
	db.Create(&GormTable{Name: "S2", Age: 20})

	// Read
	var gt GormTable
	tx := db.First(&gt, 1)
	fmt.Printf("%+v, affected: %v\n", gt, tx.RowsAffected)
	tx = db.First(&gt, "age = ?", "20")
	fmt.Printf("%+v, affected: %v\n", gt, tx.RowsAffected)

	// Update
	tx = db.Model(&gt).Update("Age", 50)
	fmt.Printf("%+v, affected: %v\n", gt, tx.RowsAffected)
	tx = db.Model(&gt).Updates(GormTable{Name: "S3", Age: 30})
	fmt.Printf("%+v, affected: %v\n", gt, tx.RowsAffected)

	tx = db.Model(&gt).Updates(map[string]interface{}{"Name": "S4", "Age": 40})
	fmt.Printf("%+v, affected: %v\n", gt, tx.RowsAffected)

	// Delete
	tx = db.Delete(&gt, 1)
	fmt.Printf("%+v, affected: %v\n", gt, tx.RowsAffected)
}
