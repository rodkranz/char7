package charset

import (
	"os"
	"testing"
	"io/ioutil"
	"bytes"
)

func checkSize(t *testing.T, path string, size int64) {
	dir, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat %q (looking for size %d): %s", path, size, err)
	}
	if dir.Size() != size {
		t.Errorf("Stat %q: size %d want %d", path, dir.Size(), size)
	}
}

func Test_CharSet_Should_Convert_Temporary_File(t *testing.T) {
	f, err := ioutil.TempFile("", "charset-test")
	if err != nil {
		t.Fatal(err)
	}

	originalText := []byte("Acentue as palavras, se for necessário: cláusula, prêmio, aliás, câncer, arguem " +
		"toxicômano chinês vatapá, níquel, baú, útil, incrível, vintém, tórax, lágrima, compôs, jóquei, ruína, " +
		"periódico, através, lápis, assembleia, sensível, voo, leem, vírus, córrego, cônsul, amáveis, xérox.")

	expectedText := []byte("Acentue as palavras, se for necess&aacute;rio: cl&aacute;usula, pr&ecirc;mio, " +
		"ali&aacute;s, c&acirc;ncer, arguem toxic&ocirc;mano chin&ecirc;s vatap&aacute;, n&iacute;quel, " +
		"ba&uacute;, &uacute;til, incr&iacute;vel, vint&eacute;m, t&oacute;rax, l&aacute;grima, comp&ocirc;s, " +
		"j&oacute;quei, ru&iacute;na, peri&oacute;dico, atrav&eacute;s, l&aacute;pis, assembleia, sens&iacute;vel, " +
		"voo, leem, v&iacute;rus, c&oacute;rrego, c&ocirc;nsul, am&aacute;veis, x&eacute;rox.\n")

	filename := f.Name()
	if err := ioutil.WriteFile(filename, originalText, 0644); err != nil {
		t.Fatalf("WriteFile %s: %v", filename, err)
	}
	f.Close()

	if err := CharSet(filename); err != nil {
		t.Fatalf("Convert file %s: %v", filename, err)
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Read converted file %s: %v", filename, err)
	}

	checkSize(t, filename, int64(len(contents)))

	if !bytes.Equal(contents, expectedText) {
		t.Fatalf("The final result is not expected for file: %v", filename)
	}

	if !HasChange {
		t.Error("The variable HasChange must return true but got false!")
	}
}
func Test_CharSet_Should_Return_Error_Message(t *testing.T) {
	err := CharSet("/_not_exits_file")
	if err == nil {
		t.Errorf("It should return a error message but got %v", err)
	}

	if err.Error() != "open /_not_exits_file: no such file or directory" {
		t.Errorf("It should return a error message with file or fir no such but got: %v", err)
	}

	if HasChange {
		t.Error("The variable HasChange must return false but got true!")
	}
}