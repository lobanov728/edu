package main

import (
	"reflect"
	"testing"
)

//func TestSimple(t *testing.T) {
//	var assert int = 42
//
//	res := MyAtoi("42")
//	if !reflect.DeepEqual(assert, res) {
//		t.Error(res, " not equal ", assert)
//	}
//}
//
//func TestSimpleWithPlus(t *testing.T) {
//	var assert int = 42
//
//	res := MyAtoi("+42")
//	if !reflect.DeepEqual(assert, res) {
//		t.Error(res, " not equal ", assert)
//	}
//}

//func TestMinusSimple(t *testing.T) {
//	var assert int = -42
//	res := MyAtoi("   -42")
//	if !reflect.DeepEqual(assert, res) {
//		t.Error(res, " not equal ", assert)
//	}
//}

func TestWithWords(t *testing.T) {
	var assert int = 4193
	res := MyAtoi("4193 with words что то еще")
	if !reflect.DeepEqual(assert, res) {
		t.Error(res, " not equal ", assert)
	}
}

//func TestWithWordsInBegin(t *testing.T) {
//	var assert int = 0
//	res := MyAtoi("with words что то еще 4193")
//	if !reflect.DeepEqual(assert, res) {
//		t.Error(res, " not equal ", assert)
//	}
//}
//
//func TestOverFlow(t *testing.T) {
//	var assert int = math.MinInt32
//	res := MyAtoi("-91283472332")
//	if !reflect.DeepEqual(assert, res) {
//		t.Error(res, " not equal ", assert)
//	}
//}