/*
Created on 2018/6/12 9:23

author: ChenJinLong

Content: 
*/
package formation

import (
	"6vs6game/source/knights"
	"errors"
	"6vs6game/source/skills"
	"fmt"
)

type IFormation interface {
	generateSingleRow() []knights.IKnight  //生成一行骑士对象
	generateSingleColumn() []knights.IKnight  //生成一列骑士对象。
	generateDoubleColumn() []knights.IKnight  //生成两列骑士对象
	generateSingle() []knights.IKnight           //生成被攻击单个骑士对象，从左到右,从上到下
	IsAllDead() bool
}




//判断骑士是否全部死亡
func (f *Formation)IsAllDead() bool {
	for i:=0; i<f.Rows ;i++ {
		for j := 0; j < f.Columns; j++ {
			if !f.Position[i][j].IsDead() {
				return false
			}
		}
	}
	return true

}

type Formation struct {
	Position [][]knights.IKnight
	Rows        int     //行数量
	Columns      int     //列数量
	enableAttackKnight []knights.IKnight
}






//生成一个新阵型布局。
func NewFormation(rows,columns int,k ...knights.IKnight) (*Formation,error) {
	if len(k) != rows*columns {
		return nil,errors.New("骑士数量不符合要求")
	}
	f := new(Formation)
	f.Rows = rows
	f.Columns = columns
	var position  [][]knights.IKnight
	for i:=0;i<rows;i++ {
		var singleRow []knights.IKnight
		for j:=i*columns;j<(i+1)*columns;j++{
			singleRow = append(singleRow,k[j])
		}
		position = append(position,singleRow)
	}
	f.Position = position
	f.enableAttackKnight = k
	return f,nil


}

// 生成一个待攻击的骑士
func (f *Formation) GenerateOneKnight() knights.IKnight {
	var enableAttackKnight []knights.IKnight
	for _,v := range f.enableAttackKnight {
		if !v.IsDead() {
			enableAttackKnight = append(enableAttackKnight,v)
		}
	}
	if len(enableAttackKnight) > 2 {
		f.enableAttackKnight = append(enableAttackKnight[1:len(enableAttackKnight)],enableAttackKnight[0])
		return enableAttackKnight[0]
	}else if len(enableAttackKnight) >= 1 {
		return enableAttackKnight[0]
	}else {
		return nil
	}


}

//生成一行骑士对象
func (f *Formation) generateSingleRow() []knights.IKnight {
	var singleRow []knights.IKnight
	for i:=0;i<f.Rows ; i++ {
		for j:=0;j<f.Columns ;j++  {
			if !f.Position[i][j].IsDead() {
				singleRow = append(singleRow,f.Position[i][j])
			}

		}
		if len(singleRow) !=0 {
			return singleRow
		}
	}
	return singleRow
}

//生成一列骑士对象
func (f *Formation) generateSingleColumn() []knights.IKnight  {
	var singleColumn  []knights.IKnight
	for i:=0;i<f.Rows ;i++  {
		for j:=0;j<f.Columns;j++ {
			if !f.Position[i][j].IsDead() {
				singleColumn = append(singleColumn,f.Position[i][j])
				break
			}

		}
	}
	return singleColumn


}

//生成两列骑士对象
func (f *Formation) generateDoubleColumn() []knights.IKnight {
	var doubleColumn []knights.IKnight
	for i:=0; i<f.Rows ;i++  {
		for j:=0;j<f.Columns ;j++  {
			if !f.Position[i][j].IsDead() {
				doubleColumn =append(doubleColumn,f.Position[i][j])
				if j+1 < f.Columns && !f.Position[i][j+1].IsDead() {
					doubleColumn = append(doubleColumn,f.Position[i][j+1])
				}
				break
			}
		}
	}
	return doubleColumn
}


//生成单个骑士对象
func (f *Formation)generateSingle() []knights.IKnight {
	for i:=0;i<f.Rows ;i++  {
		for j:=0;j<f.Columns;j++ {
			if !f.Position[i][j].IsDead() {
				return []knights.IKnight{f.Position[i][j]}
			}
		}

	}
	return nil
}


//显示阵容
func (f *Formation) Display() {
	var showString string
	for i:=0;i<f.Rows ;i++  {
		showString = showString + "\n------------------------------\n"
		for j:=0;j<f.Columns;j++ {
				showString =showString + fmt.Sprintf(" %s:%d |",f.Position[i][j].GetName(),f.Position[i][j].GetCurrrentHp())
			}
	}
	fmt.Println(showString)

}

//根据敌方postion和我方骑士技能攻击范围生成被攻击的骑士对象列表
func NewBeattackedKnights(formation *Formation,knight knights.IKnight ) []knights.IKnight {
	s := knight.RandomOneActiveSkill()
	switch s.AttackRange {
	case skills.HorsemanshipAttack:
		return formation.generateDoubleColumn()
	case skills.SwordsmanshipAttack:
		return formation.generateSingle()
	case skills.ShieldHitAttack:
		return formation.generateSingleColumn()
	default:
		return formation.generateSingleRow()
	}

}

