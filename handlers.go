package main

import (
	"fmt"
	"strings"
	"time"
)

func findCylinder(data *DataStore, id string) (*Cylinder, int) {
	for i := range data.Cylinders {
		if data.Cylinders[i].ID == id {
			return &data.Cylinders[i], i
		}
	}
	return nil, -1
}

func isValidType(t string) bool {
	for _, vt := range ValidCylinderTypes {
		if vt == t {
			return true
		}
	}
	return false
}

func isValidOwner(o string) bool {
	for _, vo := range ValidOwnerTypes {
		if vo == o {
			return true
		}
	}
	return false
}

func joinTypes(types []string) string {
	return strings.Join(types, "/")
}

func HandleAddCylinder(data *DataStore, id, cylType, owner string) error {
	if c, _ := findCylinder(data, id); c != nil {
		return fmt.Errorf("气瓶 %s 已存在", id)
	}
	if !isValidType(cylType) {
		return fmt.Errorf("气瓶类型无效，只能是: %s", joinTypes(ValidCylinderTypes))
	}
	if !isValidOwner(owner) {
		return fmt.Errorf("所有者类型无效，只能是: %s", joinTypes(ValidOwnerTypes))
	}
	cylinder := Cylinder{
		ID:        id,
		Type:      cylType,
		Owner:     owner,
		Status:    StatusPending,
		CreatedAt: time.Now().Format("2006-01-02"),
	}
	data.Cylinders = append(data.Cylinders, cylinder)
	fmt.Printf("气瓶 %s 登记成功，类型: %s，所有者: %s，状态: %s\n", id, cylType, owner, StatusPending)
	return nil
}

func HandleFill(data *DataStore, id, date, operator string) error {
	c, _ := findCylinder(data, id)
	if c == nil {
		return fmt.Errorf("气瓶 %s 不存在", id)
	}
	if c.Status != StatusPending {
		return fmt.Errorf("气瓶 %s 当前状态为 %s，只有待充状态才能充装", id, c.Status)
	}
	c.Status = StatusFull
	data.FillRecords = append(data.FillRecords, FillRecord{
		CylinderID: id,
		Date:       date,
		Operator:   operator,
	})
	fmt.Printf("气瓶 %s 充装完成，操作人: %s，状态: %s\n", id, operator, StatusFull)
	return nil
}

func HandleLend(data *DataStore, id, customer, phone, date string) error {
	c, _ := findCylinder(data, id)
	if c == nil {
		return fmt.Errorf("气瓶 %s 不存在", id)
	}
	if c.Status != StatusFull {
		return fmt.Errorf("气瓶 %s 当前状态为 %s，只有满瓶状态才能借出", id, c.Status)
	}
	c.Status = StatusInUse
	data.LendRecords = append(data.LendRecords, LendRecord{
		CylinderID: id,
		Customer:   customer,
		Phone:      phone,
		Date:       date,
	})
	fmt.Printf("气瓶 %s 已借出，客户: %s，电话: %s，状态: %s\n", id, customer, phone, StatusInUse)
	return nil
}

func HandleReturn(data *DataStore, id, date string) error {
	c, _ := findCylinder(data, id)
	if c == nil {
		return fmt.Errorf("气瓶 %s 不存在", id)
	}
	if c.Status != StatusInUse {
		return fmt.Errorf("气瓶 %s 当前状态为 %s，只有使用中状态才能归还", id, c.Status)
	}
	c.Status = StatusPending
	data.ReturnRecords = append(data.ReturnRecords, ReturnRecord{
		CylinderID: id,
		Date:       date,
	})
	fmt.Printf("气瓶 %s 已归还，状态: %s\n", id, StatusPending)
	return nil
}

func HandleMonthly(data *DataStore, month string) {
	var fillCount, lendCount int
	for _, r := range data.FillRecords {
		if len(r.Date) >= 7 && r.Date[:7] == month {
			fillCount++
		}
	}
	for _, r := range data.LendRecords {
		if len(r.Date) >= 7 && r.Date[:7] == month {
			lendCount++
		}
	}
	fmt.Printf("月份: %s\n充装次数: %d\n借出次数: %d\n", month, fillCount, lendCount)
}

func HandleStatus(data *DataStore) {
	counts := map[string]int{}
	for _, c := range data.Cylinders {
		counts[c.Status]++
	}
	fmt.Printf("气瓶总数: %d\n", len(data.Cylinders))
	fmt.Printf("待充: %d\n", counts[StatusPending])
	fmt.Printf("满瓶: %d\n", counts[StatusFull])
	fmt.Printf("使用中: %d\n", counts[StatusInUse])
}
