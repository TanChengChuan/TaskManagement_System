package tests

import (
	"TaskManagement_System/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_task(t *testing.T) {
	var err error
	models.DB, err = gorm.Open(mysql.Open(models.DSN))
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	//可加循环，然后退出。-》加屁，错了
	models.DB.AutoMigrate(&models.Task{}) //每一次都迁移 ->可以优化为 创建，删除，更新，导入时 再迁移。 （后面再说）

	// 测试用例1：创建新任务
	t.Run("CreateTask", testCreateTask)

	// 测试用例2：获取所有任务
	t.Run("GetAllTasks", testGetAllTasks)
}

func testCreateTask(t *testing.T) {
	task := models.Task{Title: "Test Task", Description: "This is a test task"}
	result := models.DB.Create(&task)
	if result.Error != nil {
		t.Fatalf("failed to create task: %v", result.Error)
	}
	fmt.Printf("Task created with ID: %d\n", task.ID)
	// 可以添加更多的断言来验证任务是否正确创建
}

func testGetAllTasks(t *testing.T) {
	var tasks []models.Task
	result := models.DB.Find(&tasks)
	if result.Error != nil {
		t.Fatalf("failed to get tasks: %v", result.Error)
	}
	for _, task := range tasks {
		fmt.Printf("Task: %s, Description: %s\n", task.Title, task.Description)
	}
	// 可以添加更多的断言来验证返回的任务列表是否正确
}
