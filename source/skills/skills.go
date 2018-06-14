/*
Created on 2018/6/11 10:39

author: ChenJinLong

Content: 
*/
package skills

const  (SwordsmanshipAttack = 1 << iota
		HorsemanshipAttack
		ShieldHitAttack
		ThrowingAttack
)



//主动技能结构体
type ActiveSkill struct {
	Name           string             //技能名
	AttackRange     int             //攻击范围
	Damage      int                  //技能伤害值，通过骑士四维计算得出
}



//被动技能结构体
type PassiveSkill struct {
	name string        //技能名
	effectType int     //效果类型
	effectValue int   //效果值
}


