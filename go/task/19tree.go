package task

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/tree"
)

type treeTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewTreeTask(pathPrefix string) Task {
	return &treeTask{
		dataPath: pathPrefix + "/19tree/Tree%03d/%03d/%03d",
		name:     "Tree",
		count:    100,
	}
}

func (tt *treeTask) GetCount() int { return tt.count }

func (tt *treeTask) GetName() string { return tt.name }

func (tt *treeTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(tt.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	tt.scanner = scanner
	tt.checker = checker
	return nil
}

func (tt *treeTask) ScannerEOF() bool { return tt.scanner.EOF() }

func (tt *treeTask) CheckerEOF() bool { return tt.checker.EOF() }

func (tt *treeTask) CleanData() {
	tt.scanner.RemoveFiles()
	tt.checker.RemoveFiles()
}

func (tt *treeTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return tt.tree1
	case 2:
		return tt.tree2
	case 3:
		return tt.tree3
	case 4:
		return tt.tree4
	case 5:
		return tt.tree5
	case 6:
		return tt.tree6
	case 7:
		return tt.tree7
	case 8:
		return tt.tree8
	case 9:
		return tt.tree9
	case 10:
		return tt.tree10
	case 11:
		return tt.tree11
	case 12:
		return tt.tree12
	case 13:
		return tt.tree13
	case 14:
		return tt.tree14
	case 15:
		return tt.tree15
	case 16:
		return tt.tree16
	case 17:
		return tt.tree17
	case 18:
		return tt.tree18
	case 19:
		return tt.tree19
	case 20:
		return tt.tree20
	case 21:
		return tt.tree21
	case 22:
		return tt.tree22
	case 23:
		return tt.tree23
	case 24:
		return tt.tree24
	case 25:
		return tt.tree25
	case 26:
		return tt.tree26
	case 27:
		return tt.tree27
	case 28:
		return tt.tree28
	case 29:
		return tt.tree29
	case 30:
		return tt.tree30
	case 31:
		return tt.tree31
	case 32:
		return tt.tree32
	case 33:
		return tt.tree33
	case 34:
		return tt.tree34
	case 35:
		return tt.tree35
	case 36:
		return tt.tree36
	case 37:
		return tt.tree37
	case 38:
		return tt.tree38
	case 39:
		return tt.tree39
	case 40:
		return tt.tree40
	case 41:
		return tt.tree41
	case 42:
		return tt.tree42
	case 43:
		return tt.tree43
	case 44:
		return tt.tree44
	case 45:
		return tt.tree45
	case 46:
		return tt.tree46
	case 47:
		return tt.tree47
	case 48:
		return tt.tree48
	case 49:
		return tt.tree49
	case 50:
		return tt.tree50
	case 51:
		return tt.tree51
	case 52:
		return tt.tree52
	case 53:
		return tt.tree53
	case 54:
		return tt.tree54
	case 55:
		return tt.tree55
	case 56:
		return tt.tree56
	case 57:
		return tt.tree57
	case 58:
		return tt.tree58
	case 59:
		return tt.tree59
	case 60:
		return tt.tree60
	case 61:
		return tt.tree61
	case 62:
		return tt.tree62
	case 63:
		return tt.tree63
	case 64:
		return tt.tree64
	case 65:
		return tt.tree65
	case 66:
		return tt.tree66
	case 67:
		return tt.tree67
	case 68:
		return tt.tree68
	case 69:
		return tt.tree69
	case 70:
		return tt.tree70
	case 71:
		return tt.tree71
	case 72:
		return tt.tree72
	case 73:
		return tt.tree73
	case 74:
		return tt.tree74
	case 75:
		return tt.tree75
	case 76:
		return tt.tree76
	case 77:
		return tt.tree77
	case 78:
		return tt.tree78
	case 79:
		return tt.tree79
	case 80:
		return tt.tree80
	case 81:
		return tt.tree81
	case 82:
		return tt.tree82
	case 83:
		return tt.tree83
	case 84:
		return tt.tree84
	case 85:
		return tt.tree85
	case 86:
		return tt.tree86
	case 87:
		return tt.tree87
	case 88:
		return tt.tree88
	case 89:
		return tt.tree89
	case 90:
		return tt.tree90
	case 91:
		return tt.tree91
	case 92:
		return tt.tree92
	case 93:
		return tt.tree93
	case 94:
		return tt.tree94
	case 95:
		return tt.tree95
	case 96:
		return tt.tree96
	case 97:
		return tt.tree97
	case 98:
		return tt.tree98
	case 99:
		return tt.tree99
	case 100:
		return tt.tree100
	default:
		return nil
	}
}

func (tt *treeTask) tree1() bool {
	var p1 = tt.scanner.NextTree()
	if !tt.checker.CompareInt(p1.Data, p1.Left.Data, p1.Right.Data) {
		return false
	}
	return tt.checker.CompareTree(p1, p1.Left, p1.Right)
}

func (tt *treeTask) tree2() bool {
	var p1 = tt.scanner.NextTree()
	var nodesCount = tree.GetCount(p1, 0, 0, "all")
	return tt.checker.CompareInt(nodesCount)
}

func (tt *treeTask) tree3() bool {
	var p1 = tt.scanner.NextTree()
	var k = tt.scanner.NextInt()
	var kCounts = tree.GetCount(p1, k, 0, "data")
	return tt.checker.CompareInt(kCounts)
}

func (tt *treeTask) tree4() bool {
	var p1 = tt.scanner.NextTree()
	var sum = tree.GetSum(p1, 0, 0, "all")
	return tt.checker.CompareInt(sum)
}

func (tt *treeTask) tree5() bool {
	var p1 = tt.scanner.NextTree()
	var leftNodesCount = tree.GetCount(p1, 0, 0, "left")
	return tt.checker.CompareInt(leftNodesCount)
}

func (tt *treeTask) tree6() bool {
	var p1 = tt.scanner.NextTree()
	var leavesCount = tree.GetCount(p1, 0, 0, "leaf")
	return tt.checker.CompareInt(leavesCount)
}

func (tt *treeTask) tree7() bool {
	var p1 = tt.scanner.NextTree()
	var leavesSum = tree.GetSum(p1, 0, 0, "leaf")
	return tt.checker.CompareInt(leavesSum)
}

func (tt *treeTask) tree8() bool {
	var p1 = tt.scanner.NextTree()
	var rightLeavesCount = tree.GetCount(p1, 0, 0, "right", "leaf")
	return tt.checker.CompareInt(rightLeavesCount)
}

func (tt *treeTask) tree9() bool {
	return false
}

func (tt *treeTask) tree10() bool {
	return false
}

func (tt *treeTask) tree11() bool {
	return false
}

func (tt *treeTask) tree12() bool {
	return false
}

func (tt *treeTask) tree13() bool {
	return false
}

func (tt *treeTask) tree14() bool {
	return false
}

func (tt *treeTask) tree15() bool {
	return false
}

func (tt *treeTask) tree16() bool {
	return false
}

func (tt *treeTask) tree17() bool {
	return false
}

func (tt *treeTask) tree18() bool {
	return false
}

func (tt *treeTask) tree19() bool {
	return false
}

func (tt *treeTask) tree20() bool {
	return false
}

func (tt *treeTask) tree21() bool {
	return false
}

func (tt *treeTask) tree22() bool {
	return false
}

func (tt *treeTask) tree23() bool {
	return false
}

func (tt *treeTask) tree24() bool {
	return false
}

func (tt *treeTask) tree25() bool {
	return false
}

func (tt *treeTask) tree26() bool {
	return false
}

func (tt *treeTask) tree27() bool {
	return false
}

func (tt *treeTask) tree28() bool {
	return false
}

func (tt *treeTask) tree29() bool {
	return false
}

func (tt *treeTask) tree30() bool {
	return false
}

func (tt *treeTask) tree31() bool {
	return false
}

func (tt *treeTask) tree32() bool {
	return false
}

func (tt *treeTask) tree33() bool {
	return false
}

func (tt *treeTask) tree34() bool {
	return false
}

func (tt *treeTask) tree35() bool {
	return false
}

func (tt *treeTask) tree36() bool {
	return false
}

func (tt *treeTask) tree37() bool {
	return false
}

func (tt *treeTask) tree38() bool {
	return false
}

func (tt *treeTask) tree39() bool {
	return false
}

func (tt *treeTask) tree40() bool {
	return false
}

func (tt *treeTask) tree41() bool {
	return false
}

func (tt *treeTask) tree42() bool {
	return false
}

func (tt *treeTask) tree43() bool {
	return false
}

func (tt *treeTask) tree44() bool {
	return false
}

func (tt *treeTask) tree45() bool {
	return false
}

func (tt *treeTask) tree46() bool {
	return false
}

func (tt *treeTask) tree47() bool {
	return false
}

func (tt *treeTask) tree48() bool {
	return false
}

func (tt *treeTask) tree49() bool {
	return false
}

func (tt *treeTask) tree50() bool {
	return false
}

func (tt *treeTask) tree51() bool {
	return false
}

func (tt *treeTask) tree52() bool {
	return false
}

func (tt *treeTask) tree53() bool {
	return false
}

func (tt *treeTask) tree54() bool {
	return false
}

func (tt *treeTask) tree55() bool {
	return false
}

func (tt *treeTask) tree56() bool {
	return false
}

func (tt *treeTask) tree57() bool {
	return false
}

func (tt *treeTask) tree58() bool {
	return false
}

func (tt *treeTask) tree59() bool {
	return false
}

func (tt *treeTask) tree60() bool {
	return false
}

func (tt *treeTask) tree61() bool {
	return false
}

func (tt *treeTask) tree62() bool {
	return false
}

func (tt *treeTask) tree63() bool {
	return false
}

func (tt *treeTask) tree64() bool {
	return false
}

func (tt *treeTask) tree65() bool {
	return false
}

func (tt *treeTask) tree66() bool {
	return false
}

func (tt *treeTask) tree67() bool {
	return false
}

func (tt *treeTask) tree68() bool {
	return false
}

func (tt *treeTask) tree69() bool {
	return false
}

func (tt *treeTask) tree70() bool {
	return false
}

func (tt *treeTask) tree71() bool {
	return false
}

func (tt *treeTask) tree72() bool {
	return false
}

func (tt *treeTask) tree73() bool {
	return false
}

func (tt *treeTask) tree74() bool {
	return false
}

func (tt *treeTask) tree75() bool {
	return false
}

func (tt *treeTask) tree76() bool {
	return false
}

func (tt *treeTask) tree77() bool {
	return false
}

func (tt *treeTask) tree78() bool {
	return false
}

func (tt *treeTask) tree79() bool {
	return false
}

func (tt *treeTask) tree80() bool {
	return false
}

func (tt *treeTask) tree81() bool {
	return false
}

func (tt *treeTask) tree82() bool {
	return false
}

func (tt *treeTask) tree83() bool {
	return false
}

func (tt *treeTask) tree84() bool {
	return false
}

func (tt *treeTask) tree85() bool {
	return false
}

func (tt *treeTask) tree86() bool {
	return false
}

func (tt *treeTask) tree87() bool {
	return false
}

func (tt *treeTask) tree88() bool {
	return false
}

func (tt *treeTask) tree89() bool {
	return false
}

func (tt *treeTask) tree90() bool {
	return false
}

func (tt *treeTask) tree91() bool {
	return false
}

func (tt *treeTask) tree92() bool {
	return false
}

func (tt *treeTask) tree93() bool {
	return false
}

func (tt *treeTask) tree94() bool {
	return false
}

func (tt *treeTask) tree95() bool {
	return false
}

func (tt *treeTask) tree96() bool {
	return false
}

func (tt *treeTask) tree97() bool {
	return false
}

func (tt *treeTask) tree98() bool {
	return false
}

func (tt *treeTask) tree99() bool {
	return false
}

func (tt *treeTask) tree100() bool {
	return false
}
