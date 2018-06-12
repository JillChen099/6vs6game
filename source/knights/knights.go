/*
Created on 2018/6/11 10:33

author: ChenJinLong

Content: 
*/
package knights

import "6vs6game/source/skills"

type IKnight interface {
	ReleaseAttack(beattacker ...IKnight) //发动攻击
	IsDead() bool      //是否死亡
	IncreaseDamage()  //增加技能伤害
	SufferDamage(attacker IKnight)  //遭受攻击
}



type Knight struct {
	Name    string           //骑士名字
	HP   int                //血量
	SwordsmanshipNum  int   // 剑术
	HorsemanshipNum   int   //骑术
	ShieldHitNum     int  //盾术
	ThrowingNum      int  //投掷
	Skills []*skills.ActiveSkill

}


func (k *Knight) IsDead() bool{
	return k.HP <= 0

}

func NewKnight(name string,hp,sw,ho,sh,th int) *Knight {
	k := new(Knight)
	k.Name = name
	k.HP = hp
	k.SwordsmanshipNum = sw
	k.HorsemanshipNum = ho
	k.ShieldHitNum = sh
	k.ThrowingNum = th





}


