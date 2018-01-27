package log

import (
	"fmt"
	"math"

	"gopkg.in/cheggaaa/pb.v1"
)

const (
	NOCOLOR = 0
	RED     = 31
	GREEN   = 32
	YELLOW  = 33
	BLUE    = 36
	GRAY    = 37
	UNICODE_FULL_BLOCK = "█"
	UNICODE_HALF_BLOCK = "▌"
)

type Report struct {
	O           float64
	Chunk		float64
	ProgressBar *pb.ProgressBar
}

func NewReport() Report {
	return Report{}
}

func (r *Report) ErrorMessage(err error) {
	msg := fmt.Sprintf("\x1b[%dmERROR : %s\x1b[%dm", RED, err, NOCOLOR)
	fmt.Println(msg)
}

func (r *Report) PrintMessage(msg string) {
	fmt.Println(msg)
}

func (r *Report) FormatMessage(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	r.PrintMessage(msg)
}

func (r *Report) CMessage(msg string) {
	fmt.Print(msg, " ... ")
}

func (r *Report) FormatCMessage(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	r.CMessage(msg)
}

func (r *Report) NewProgressBar(prefix string, o int) {
	bar := pb.New(o).Prefix(prefix).SetMaxWidth(100)
	bar.SetUnits(pb.U_BYTES_DEC)
	bar.Format("|" + UNICODE_FULL_BLOCK + UNICODE_HALF_BLOCK + " |")
	bar.ShowCounters = false
	bar.ShowSpeed = true
	r.ProgressBar = bar.Start()
	r.O = float64(o)
}

func (r *Report) ProgressTick(c float64) {
	r.Chunk = r.Chunk + c
	chunk := math.Trunc(r.Chunk)
	if c > 0.0 {
		r.ProgressBar.Add(int(chunk))
		r.Chunk = r.Chunk - chunk
	}
	r.O = r.O - c
}

func (r *Report) ProgressDone() {
	r.ProgressBar.Finish()
}

func (r *Report) CreateChunk(m int) float64 {
	return r.O / float64(m)
}
