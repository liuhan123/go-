package testcase

import (
	"gl.zego.im/goserver/han/unittest/cmd"
	"testing"
)

var sum  = []int{2, 1, 3}
var max = []int{2, 1, 2}

var a , b = 2, 1

func Test_Add (t *testing.T){
	value := cmd.Add(sum[0], sum[1])
	if (value != sum[2]){
		t.Error("test add error")
	}
}

func Test_Max(t *testing.T){
	value := cmd.Max(max[0], max[1])
	if (value != max[2]){
		t.Error("test max error")
	}
}