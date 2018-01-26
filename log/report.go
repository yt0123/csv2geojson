package log

import (
	"fmt"
	"math"

	"gopkg.in/cheggaaa/pb.v1"
)

type Report struct {
	O           float64
	ProgressBar *pb.ProgressBar
}

func NewReport() Report {
	return Report{}
}

func (r *Report) PrintMessage(msg string) {
	fmt.Println(msg)
}

func (r *Report) FormatMessage(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	r.PrintMessage(msg)
}

func (r *Report) CMessage(msg string) {
	fmt.Print(msg + " ... ")
}

func (r *Report) FormatCMessage(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	r.CMessage(msg)
}

func (r *Report) NewProgressBar(prefix string, total int) {
	p := pb.New(total)
	p.Prefix(prefix)
	p.SetMaxWidth(100)
	r.ProgressBar = p.Start()
	r.O = float64(total)
}

func (r *Report) ProgressTick(chunk float64) {
	c := math.Floor(chunk)
	r.ProgressBar.Add(int(c))
}

func (r *Report) ProgressDone() {
	r.ProgressBar.Finish()
}
