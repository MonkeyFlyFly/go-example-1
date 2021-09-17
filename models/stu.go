package models

type Stu struct {
	Model

	Id   string
	Age  int
	Name string
}

func GetStus(pageNum int, pageSize int, maps interface{}) (stu []Stu) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&stu)
	return
}

func GetStusTotal(maps interface{}) (count int) {
	db.Model(&Stu{}).Where(maps).Count(&count)
	return
}

func GetStu(id string) (stu Stu) {
	db.Where("id = ?", id).First(&stu)
	db.Model(&stu)
	return
}

func AddStu(data map[string]interface{}) bool {
	db.Create(&Stu{
		Id:   data["id"].(string),
		Age:  data["age"].(int),
		Name: data["name"].(string),
	})

	return true
}

func DeleteStu(id int) bool {
	db.Where("id = ?", id).Delete(Stu{})

	return true
}

func CleanAllStu() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Stu{})

	return true
}
