package entity

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota //iota là đếm tự động tăng dần
	ItemStatusDone
	ItemStatusDeleted
)

// Thực hiện parse từ int thành  String cho json data
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

// Từ json về dữ liệu
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "") // \" tương đương với dấu " " trong JSON
	itemVal, err := parseItemStatus(str)
	if err != nil {
		return err
	}
	*item = itemVal
	return nil
}

//Do trc đó mình đã set Status là string giờ là int nên mysql không the scan
// -> để scan cần có

var allItemStatus = [3]string{"Doing", "Done", "Deleted"}

func (item *ItemStatus) String() string {
	return allItemStatus[*item]
}

func parseItemStatus(status string) (ItemStatus, error) {
	for i := range allItemStatus {
		if allItemStatus[i] == status {
			return ItemStatus(i), nil
		}
	}
	return ItemStatus(0), errors.New("Invalid Item Status")
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	v, err := parseItemStatus(string(bytes))

	if err != nil {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	*item = v
	return nil

}

// Lấy từ itemStauts chuyển thành dữ liệu mysql
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}
