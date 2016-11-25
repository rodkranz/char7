package settings

import (
	"testing"
)


func TestWorkDirShouldReturnCurrentUserFolder(t *testing.T) {
	s := WorkDir()
	if len(s) == 0 {
		t.Errorf("Expected to return a $HOME but got %v", s)
	}
}
