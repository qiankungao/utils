package search

import (
	"testing"
)

// 测试最后一个小于等于
func TestSearchLastLessElement(t *testing.T) {
	arrs := []int{1, 5, 10, 15, 20}
	if res := SearchLastLessElement(arrs, 1); res != 0 {
		t.Errorf("res is 0,not %d", res)
	}
	if res := SearchLastLessElement(arrs, 5); res != 1 {
		t.Errorf("res is 1,not %d", res)
	}
	if res := SearchLastLessElement(arrs, 100); res != 4 {
		t.Errorf("res is 4,not %d", res)
	}
	if res := SearchLastLessElement(arrs, 0); res != -1 {
		t.Errorf("res is -1,not %d", res)
	}
	arrs = []int{}
	if res := SearchLastLessElement(arrs, 100); res != -1 {
		t.Errorf("res is -1,not %d", res)
	}
}

//查找最后一个与 target 相等的元素
func TestSearchLastEqualElement(t *testing.T) {
	arrs := []int{1, 5, 5, 10, 15, 20}
	if res := SearchLastEqualElement(arrs, 1); res != 0 {
		t.Errorf("res is 0,not %d", res)
	}
	if res := SearchLastEqualElement(arrs, 5); res != 1 {
		t.Errorf("res is 1,not %d", res)
	}
	if res := SearchLastEqualElement(arrs, 20); res != 4 {
		t.Errorf("res is 4,not %d", res)
	}
}

//查找第一个大于等于target的元素
func TestSearchFirstGreaterElement(t *testing.T) {
	arrs := []int{1, 5, 5, 10, 15, 20}
	if res := SearchFirstGreaterElement(arrs, 4); res != 1 {
		t.Errorf("res is 1,not %d", res)
	}

	if res := SearchFirstGreaterElement(arrs, 100); res != -1 {
		t.Errorf("res is -1,not %d", res)
	}
}

//查找第一个与 target 相等的元素
func TestSearchFirstEqualElement(t *testing.T) {
	arrs := []int{1, 5, 5, 10, 15, 20}
	if res := SearchFirstEqualElement(arrs, 4); res != -1 {
		t.Errorf("res is -1,not %d", res)
	}

	if res := SearchFirstEqualElement(arrs, 20); res != 5 {
		t.Errorf("res is 5,not %d", res)
	}

	if res := SearchFirstEqualElement(arrs, 200); res != -1 {
		t.Errorf("res is -1,not %d", res)
	}

}
