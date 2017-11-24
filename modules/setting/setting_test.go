package setting

import (
	"testing"
)

func Test_WorkDir_Should_Return_Current_User_Folder(t *testing.T) {
	s := WorkDir()
	if len(s) == 0 {
		t.Errorf("Expected to return a $HOME but got %v", s)
	}
}
