package main

import (
	"fmt"
	"task_list/config"
	"task_list/driver"
	"task_list/pkg/seeds"

	"gorm.io/gorm"
)

func main() {
	config.InitConfig()
	orm := driver.InitGorm()

	// Create Tasks ----------------------------------------------------------------
	taskSeeds := seeds.AllTask()
	run(orm, taskSeeds)
}

func run(orm *gorm.DB, channelSeeds []seeds.Seed) {
	for _, seed := range channelSeeds {
		fmt.Println(seed.Name)
		err := seed.Run(orm)
		if err != nil {
			fmt.Println(seed.Name + " Failed")
			fmt.Println(err.Error())
		}
	}
}
