package formatter

import (
    "fmt"
    "time"
    "strconv"
    
    "github.com/wsxiaoys/terminal/color"
)

type Formatter interface {
    IndexStatus() string
    Master() string
    NodesSeen() string
    Uptime() string
    Version() string
}

var linePrefix string = "|- "
var headerLinePrefix string = "# "

type PlainFormatter struct {
    
}


func (p PlainFormatter) HostName(s string) string {
    return fmt.Sprint(headerLinePrefix,s)
}

func (p PlainFormatter) IndexStatus(s string) string {
    return fmt.Sprint(linePrefix,"Index: ",s)
}

func (p PlainFormatter) Master(s string) string {
    return fmt.Sprint(linePrefix,"Master: ",s)
}

func (p PlainFormatter) NodesSeen(s float64) string {
    return fmt.Sprint(linePrefix,"Nodes seen: ",s)
}

func (p PlainFormatter) Uptime(s float64) string {
	return fmt.Sprint(linePrefix,"Uptime: ",s)
}

func (p PlainFormatter) Version(s string) string {
    return fmt.Sprint(linePrefix,"Version: ",s)
}

type TerminalFormatter struct {
    
}

func (p TerminalFormatter) HostName(s string) string {
    return fmt.Sprint(headerLinePrefix,s)
}

func (p TerminalFormatter) IndexStatus(s string) string {
    formatting := ""
	if s == "GREEN" {
	    
		formatting = "@g"
	}
	if s == "YELLOW" {
		formatting = "@y"
	}
	if s == "RED" {
		formatting = "@r"
	}
    return color.Sprint(linePrefix,"Index:", formatting, s)
}

func (p TerminalFormatter) Master(s string) string {
    formatting := ""
	if s == "true" {
		formatting = "@g"
	}
	return color.Sprint(linePrefix,"Master:", formatting, s)
}

func (p TerminalFormatter) NodesSeen(s float64) string {
    return fmt.Sprint(linePrefix,"Nodes seen: ",s)
}

func (p TerminalFormatter) Uptime(s float64) string {
    uptimeString := strconv.FormatFloat(s, 'f', -1, 64)
	duration := fmt.Sprintf("%sms", uptimeString)
	formattedUptime, _ := time.ParseDuration(duration)
	formatting := "@b"
	return color.Sprint(linePrefix, "Uptime:", formatting, formattedUptime)
}

func (p TerminalFormatter) Version(s string) string {
    return fmt.Sprint(linePrefix,"Version: ",s)
}