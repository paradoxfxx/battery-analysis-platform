package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/model"
	"battery-anlysis-platform/pkg/checker"
	"battery-anlysis-platform/pkg/mysqlx"
	"errors"
	"strings"
)

type MiningBaseService struct {
	DataComeFrom string `form:"dataComeFrom" binding:"required"`
	StartDate    string `form:"startDate" binding:"required"`
	NeedParams   string `form:"needParams" binding:"required"`
	DataLimit    int    `form:"dataLimit" binding:"required"`
}

func (s *MiningBaseService) Query() ([]map[string]interface{}, error) {
	// 校验字段合法性
	table, ok := model.BatteryMysqlNameToTable[s.DataComeFrom]
	if !ok {
		return nil, errors.New("参数 dataComeFrom 不合法")
	}
	fields := strings.Split(s.NeedParams, ",")
	for _, field := range fields {
		_, ok = table.FieldToName[field]
		if !ok {
			return nil, errors.New("参数 needParams 不合法")
		}
	}
	if !checker.ReDatetime.MatchString(s.StartDate) {
		return nil, errors.New("参数 startDate 不合法")
	}
	if s.DataLimit > 10000 {
		return nil, errors.New("参数 dataLimit 不合法")
	}

	rows, err := dao.MysqlDB.Table(table.Name).
		Where("timestamp >= ?", s.StartDate).
		Select("timestamp," + strings.Join(fields, ",")).
		Limit(s.DataLimit).
		Rows()
	if err != nil {
		return nil, errors.New("查询失败")
	}

	records, err := mysqlx.GetRecords(rows)
	_ = rows.Close()
	if err != nil {
		panic(err)
	}
	if len(records) == 0 {
		return nil, errors.New("未查询到相关数据")
	}
	return records, nil
}
