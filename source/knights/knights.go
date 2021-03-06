/*
Created on 2018/6/11 10:33

author: ChenJinLong

Content: 
*/
package knights

import (
	"6vs6game/source/skills"
	"6vs6game/tools/randomutils"
	"fmt"
)

type IKnight interface {
	ReleaseAttack(beattacker []IKnight) //发动攻击
	IsDead() bool      //是否死亡
	//IncreaseDamage()  //增加技能伤害
	SufferDamage(damage int)  //遭受攻击
	RandomOneActiveSkill() *skills.ActiveSkill  //随机发动一个技能
	GetCurrrentHp()  int //获取当前骑士血量
	GetName() string
}



type Knight struct {
	Name    string           //骑士名字
	HP   int                //血量
	SwordsmanshipNum  int   // 剑术
	HorsemanshipNum   int   //骑术
	ShieldHitNum     int  //盾术
	ThrowingNum      int  //投掷
	CurrentActiveSkill *skills.ActiveSkill
	Role            string          //角色立场:我方或敌方

}

//获取当前骑士血量
func (k *Knight) GetCurrrentHp() int {
	return k.HP
}

func (k *Knight)GetName()string {
	return k.Name
}


func (k *Knight) ReleaseAttack(beattacker []IKnight) {
	fmt.Println("----------------------")
	fmt.Printf("%s骑士%s发动技能%s\n",k.Role,k.Name,k.CurrentActiveSkill.Name)
	for _,v := range beattacker {
		v.SufferDamage(k.CurrentActiveSkill.Damage)
	}


}



func (k *Knight) SufferDamage(damage int) {
	k.HP = k.HP - damage
	fmt.Printf("%s骑士%s遭受%d伤害\n",k.Role,k.Name,damage)
}

func (k *Knight) IsDead() bool{
	return k.HP <= 0

}


//随机发动一个主动技能
func (k *Knight) RandomOneActiveSkill() *skills.ActiveSkill {

	newActiveSkill := new(skills.ActiveSkill)
	allAttributes := float64(k.ThrowingNum+k.SwordsmanshipNum+k.ShieldHitNum+k.HorsemanshipNum)
	tp := float64(k.ThrowingNum)/ allAttributes
	swp := float64(k.SwordsmanshipNum) / allAttributes
	shp := float64(k.ShieldHitNum) / allAttributes
	hp := float64(k.HorsemanshipNum) / allAttributes
	percentChoiceList := [][]int{{skills.SwordsmanshipAttack,int(swp*100)},{skills.HorsemanshipAttack,int(hp*100)},{skills.ShieldHitAttack,int(shp*100)},{skills.ThrowingAttack,int(tp*100)}}
	resultType := randomutils.GetAwardByWeight(percentChoiceList)
	switch resultType[0] {
	case skills.SwordsmanshipAttack:
		newActiveSkill.AttackRange = skills.SwordsmanshipAttack
		newActiveSkill.Name = "剑刺"
		newActiveSkill.Damage = int(float64(k.SwordsmanshipNum) * swp)


	case skills.ThrowingAttack:
		newActiveSkill.AttackRange = skills.ThrowingAttack
		newActiveSkill.Name = "长矛攻击"
		newActiveSkill.Damage = int(float64(k.ThrowingNum)*tp)

	case skills.ShieldHitAttack:
		newActiveSkill.AttackRange = skills.ShieldHitAttack
		newActiveSkill.Name = "盾击"
		newActiveSkill.Damage = int(float64(k.ShieldHitNum)*shp)

	default :
		newActiveSkill.AttackRange = skills.HorsemanshipAttack
		newActiveSkill.Name = "骑兵冲锋"
		newActiveSkill.Damage = int(float64(k.HorsemanshipNum)*hp)
	}
	k.CurrentActiveSkill = newActiveSkill
	return newActiveSkill

}

func NewKnight(name ,role string,hp,sw,ho,sh,th int) *Knight {
	k := new(Knight)
	k.Name = name
	k.HP = hp
	k.SwordsmanshipNum = sw
	k.HorsemanshipNum = ho
	k.ShieldHitNum = sh
	k.ThrowingNum = th
	k.Role = role
	return k

}


