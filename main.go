/*
Created on 2018/6/11 10:25

author: ChenJinLong

Content: 
*/
package main

import (
	"fmt"
	"6vs6game/source/knights"
	"6vs6game/source/formation"
	"time"
)

func main()  {
	fmt.Println("游戏开始")
	fmt.Println("生成我方骑士")
	selfk1 := knights.NewKnight("Aaron","我方",100,20,65,10,20)
	selfk2 := knights.NewKnight("Bart","我方",80,50,15,30,25)
	selfK3 := knights.NewKnight("Colin","我方",115,10,15,45,35)
	selfk4 := knights.NewKnight("Duke","我方",85,30,15,15,10)
	selfk5 := knights.NewKnight("Evan","我方",90,10,25,15,45)
	selfk6 := knights.NewKnight("Frank","我方",70,15,55,25,10)
	fmt.Println("生成我方阵容")
	selfFormation,err := formation.NewFormation(2,3,selfk1,selfk2,selfK3,selfk4,selfk5,selfk6)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("生成敌方骑士")
	enemyk1 := knights.NewKnight("Aaron","敌方",100,20,65,10,20)
	enemyk2 := knights.NewKnight("Bart","敌方方",140,50,15,30,25)
	enemyK3 := knights.NewKnight("Colin","敌方",95,10,15,45,35)
	enemyk4 := knights.NewKnight("Duke","敌方",104,30,15,15,10)
	enemyk5 := knights.NewKnight("Evan","敌方",75,10,25,15,45)
	enemyk6 := knights.NewKnight("Frank","敌方",85,15,55,25,10)
	fmt.Println("生成敌方阵容")
	enemyFormation,err := formation.NewFormation(2,3,enemyk1,enemyk2,enemyK3,enemyk4,enemyk5,enemyk6)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("开始战斗")
	for {


			fmt.Println("我方阵容：")
			selfFormation.Display()
			fmt.Println("敌方阵容：")
			enemyFormation.Display()
			selfKnight := selfFormation.GenerateOneKnight()
			enemyBeAttackedKnights := formation.NewBeattackedKnights(enemyFormation, selfKnight)
			selfKnight.ReleaseAttack(enemyBeAttackedKnights)
			time.Sleep(2 * time.Second)

			if enemyFormation.IsAllDead() {
				fmt.Println("游戏结束,敌方全部阵亡!")
				return
			}

			enemyKnight := enemyFormation.GenerateOneKnight()
			selfBeAttackedKnights := formation.NewBeattackedKnights(selfFormation, enemyKnight)
			enemyKnight.ReleaseAttack(selfBeAttackedKnights)
			if selfFormation.IsAllDead() {
				fmt.Println("游戏结束我方全部阵亡！")
				return
			}



	}



}