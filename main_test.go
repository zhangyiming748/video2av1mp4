package video2av1mp4

import (
	"github.com/zhangyiming748/log"
	"os/exec"
	"strconv"
	"testing"
)

func TestUnit(t *testing.T) {
	src := "/Users/zen/Movies/shiny/spandex"
	dst := "/Users/zen/Movies/shiny"
	pattern := "mp4;webm"
	threads := "4"
	save, total := ConvToH265(src, dst, pattern, threads)
	log.Debug.Printf("节省的空间:%v\n", save)
	log.Debug.Printf("共处理的文件数:%v\n", total)

}

func BenchmarkBeep(b *testing.B) {
	var cmd *exec.Cmd
	cmd = exec.Command("echo", "-e", "\\a")
	for i := 0; i < b.N; i++ {
		cmd.Run()
	}
}

func TestFakeThreads(t *testing.T) {
	thread := "13"
	ret := fakeThreads(thread)
	t.Log(ret)

}
func fakeThreads(threads string) bool {
	maxThreads := 12
	if t, err := strconv.Atoi(threads); err != nil {
		return false
	} else if t >= maxThreads {
		return false
	} else {
		return true
	}
}

func TestGetSize(t *testing.T) {

}
