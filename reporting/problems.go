package reporting

import "fmt"

import (
	"github.com/smartystreets/goconvey/gotest"
	"github.com/smartystreets/goconvey/printing"
)

func (self *problem) BeginStory(test gotest.T) {}

func (self *problem) Enter(title, id string) {}

func (self *problem) Report(r *AssertionReport) {
	if r.Error != nil {
		self.errors = append(self.errors, r)
	} else if r.Failure != "" {
		self.failures = append(self.failures, r)
	}
}

func (self *problem) Exit() {}

func (self *problem) EndStory() {
	self.out.Println("")
	self.show(self.showErrors, redColor)
	self.show(self.showFailures, yellowColor)
	self.prepareForNextStory()
}
func (self *problem) show(display func(), color string) {
	fmt.Print(color)
	display()
	fmt.Print(resetColor)
	self.out.Dedent()
}
func (self *problem) showErrors() {
	for i, e := range self.errors {
		if i == 0 {
			self.out.Println("\nErrors:\n")
			self.out.Indent()
		}
		self.out.Println(errorTemplate, e.File, e.Line, e.Error, e.stackTrace)
	}
}
func (self *problem) showFailures() {
	for i, f := range self.failures {
		if i == 0 {
			self.out.Println("\nFailures:\n")
			self.out.Indent()
		}
		self.out.Println(failureTemplate, f.File, f.Line, f.Failure)
	}
}

func NewProblemReporter(out *printing.Printer) *problem {
	self := problem{}
	self.out = out
	self.prepareForNextStory()
	return &self
}
func (self *problem) prepareForNextStory() {
	self.errors = []*AssertionReport{}
	self.failures = []*AssertionReport{}
}

type problem struct {
	out      *printing.Printer
	errors   []*AssertionReport
	failures []*AssertionReport
}
