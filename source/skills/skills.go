/*
Created on 2018/6/11 10:39

author: ChenJinLong

Content: 
*/
package skills

import "6vs6game/source/knights"

const  (SwordsmanshipAttack = 1 << iota
		HorsemanshipAttack
		ShieldHitAttack
		ThrowingAttack
)





//主动技能结构体
type ActiveSkill struct {
	name           string             //技能名
	AttackRange     int             //攻击范围
	triggerProbability    int      //触发概率
	Damage      int                  //技能伤害值，通过骑士四维计算得出
}

//通过攻击范围和骑士对象生成主动技能
func NewActiveSkill(attackRange int,k *knights.Knight) *ActiveSkill {
	if attackRange == SwordsmanshipAttack {
		a := new(ActiveSkill)
		a.name = "骑兵冲锋"
		a.AttackRange = attackRange




	}



}


//被动技能结构体

type PassiveSkill struct {
	name string        //技能名
	effectType int     //效果类型
	effectValue int   //效果值
}


