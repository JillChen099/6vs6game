/*
Created on 2018/6/12 9:23

author: ChenJinLong

Content: 
*/
package formation

import (
	"6vs6game/source/knights"
	"errors"
)

type IFormation interface {
	generateSingleRow() []knights.IKnight  //生成一行骑士对象
	generateSingleColumn() []knights.IKnight  //生成一列骑士对象。
	generateDoubleColumn() []knights.IKnight  //生成两列骑士对象
	generateSingle() knights.IKnight           //生成单个骑士对象，从左到右,从上到下
}





type Formation struct {
	position [][]knights.IKnight
	rows        int     //行数量
	columns      int     //列数量
}

//生成一个新阵型布局。
func NewFormation(rows,columns int,knights ...knights.IKnight) error {
	if len(knights) != rows*columns {
		return errors.New("骑士数量不符合要求")
	}
	f := new(Formation)
	f.rows = rows
	f.columns = columns
	var position [][]knights.IKnight
	for i:=0;i<rows;i++ {
		singleRow := make([]knights.IKnight,columns)
		for j:=i*columns;j<(i+1)*columns;j++{
			singleRow = append(singleRow,knights[j])
		}
		position = append(position,singleRow)
	}
	f.position = position
	return nil

}

//生成一行骑士对象
func (f *Formation) generateSingleRow() []knights.IKnight {
	singleRow := make([]knights.IKnight,f.columns)
	for i:=0;i<f.rows ; i++ {
		for j:=0;j<f.columns ;j++  {
			if !f.position[i][j].IsDead() {
				singleRow = append(singleRow,f.position[i][j])
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
	singleColumn := make([]knights.IKnight,f.rows)
	for i:=0;i<f.rows ;i++  {
		for j:=0;j<f.columns;j++ {
			if !f.position[i][j].IsDead() {
				singleColumn = append(singleColumn,f.position[i][j])
				break
			}

		}
	}
	return singleColumn


}

//生成两列骑士对象
func (f *Formation) generateDoubleColumn() []knights.IKnight {
	doubleColumn := make([]knights.IKnight,f.rows * 2)
	for i:=0; i<f.rows ;i++  {
		for j:=0;j<f.columns ;j++  {
			if !f.position[i][j].IsDead() {
				doubleColumn =append(doubleColumn,f.position[i][j])
				if j+1 < f.columns && !f.position[i][j+1].IsDead() {
					doubleColumn = append(doubleColumn,f.position[i][j+1])
				}
				break
			}
		}
	}
	return doubleColumn
}


//生成单个骑士对象
func (f *Formation)generateSingle() knights.IKnight {
	for i:=0;i<f.rows ;i++  {
		for j:=0;j<f.columns;j++ {
			if !f.position[i][j].IsDead() {
				return f.position[i][j]
			}
		}

	}
	return nil
}


