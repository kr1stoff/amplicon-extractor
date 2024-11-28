package check

import "testing"

func TestIsWrongPrimer(t *testing.T) {
	// 测试正常情况
	primer1 := "ATGCGGCTTA"
	if IsWrongPrimer(primer1) {
		t.Errorf("Expected false, got true for primer1")
	}

	// 测试包含不允许字符的情况
	primer2 := "ATGX1CGGCTTA"
	if !IsWrongPrimer(primer2) {
		t.Errorf("Expected true, got false for primer2")
	}

	// 测试空字符串
	primer3 := ""
	if IsWrongPrimer(primer3) {
		t.Errorf("Expected false, got true for primer3")
	}
}
