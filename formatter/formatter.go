package formatter

import (
	"fmt"
	"strconv"
	"time"
	"bytes"

	"github.com/wsxiaoys/terminal/color"

	"github.com/haraldringvold/enonicstatus/jsonstruct"
)

type Formatter interface {
	HostName(string) string
	IndexStatus(string) string
	Master(string) string
	NodesSeen(float64) string
	Uptime(float64) string
	Version(string) string
	String(jsonstruct.Status) string
}

var linePrefix string = "|- "
var headerLinePrefix string = "# "

type PlainFormatter struct {
}

func (p PlainFormatter) HostName(name string) string {
	return fmt.Sprint(headerLinePrefix, name)
}

func (p PlainFormatter) IndexStatus(index string) string {
	return fmt.Sprint(linePrefix, "Index: ", index)
}

func (p PlainFormatter) Master(master string) string {
	return fmt.Sprint(linePrefix, "Master: ", master)
}

func (p PlainFormatter) NodesSeen(nodesSeen float64) string {
	return fmt.Sprint(linePrefix, "Nodes seen: ", nodesSeen)
}

func (p PlainFormatter) Uptime(uptime float64) string {
	uptimeString := strconv.FormatFloat(uptime, 'f', -1, 64)
	duration := fmt.Sprintf("%sms", uptimeString)
	formattedUptime, _ := time.ParseDuration(duration)
	return fmt.Sprint(linePrefix, "Uptime: ", formattedUptime)
}

func (p PlainFormatter) Version(version string) string {
	return fmt.Sprint(linePrefix, "Version: ", version)
}

func (p PlainFormatter) String(json jsonstruct.Status) string {
	var buffer bytes.Buffer
	buffer.WriteString("\n")
	buffer.WriteString(p.HostName(json.Cluster.LocalNode.HostName) + "\n")
	buffer.WriteString(p.IndexStatus(json.Index.Status) + "\n")
	buffer.WriteString(p.Master(json.Cluster.LocalNode.Master) + "\n")
	buffer.WriteString(p.NodesSeen(json.Cluster.LocalNode.NumberOfNodesSeen) + "\n")
	buffer.WriteString(p.Uptime(json.Jvm.UpTime) + "\n")
	buffer.WriteString(p.Version(json.Product.Version))
	return buffer.String()
}

type TerminalFormatter struct {
}

func (t TerminalFormatter) HostName(name string) string {
	return fmt.Sprint(headerLinePrefix, name)
}

func (t TerminalFormatter) IndexStatus(index string) string {
	formatting := ""
	if index == "GREEN" {

		formatting = "@g"
	}
	if index == "YELLOW" {
		formatting = "@y"
	}
	if index == "RED" {
		formatting = "@r"
	}
	return color.Sprint(linePrefix, "Index:", formatting, index)
}

func (t TerminalFormatter) Master(master string) string {
	formatting := ""
	if master == "true" {
		formatting = "@g"
	}
	return color.Sprint(linePrefix, "Master:", formatting, master)
}

func (t TerminalFormatter) NodesSeen(nodesSeen float64) string {
	return fmt.Sprint(linePrefix, "Nodes seen: ", nodesSeen)
}

func (t TerminalFormatter) Uptime(uptime float64) string {
	uptimeString := strconv.FormatFloat(uptime, 'f', -1, 64)
	duration := fmt.Sprintf("%sms", uptimeString)
	formattedUptime, _ := time.ParseDuration(duration)
	formatting := "@b"
	return color.Sprint(linePrefix, "Uptime:", formatting, formattedUptime)
}

func (t TerminalFormatter) Version(version string) string {
	return fmt.Sprint(linePrefix, "Version: ", version)
}

func (t TerminalFormatter) String(json jsonstruct.Status) string {
	var buffer bytes.Buffer
	buffer.WriteString("\n")
	buffer.WriteString(t.HostName(json.Cluster.LocalNode.HostName) + "\n")
	buffer.WriteString(t.IndexStatus(json.Index.Status) + "\n")
	buffer.WriteString(t.Master(json.Cluster.LocalNode.Master) + "\n")
	buffer.WriteString(t.NodesSeen(json.Cluster.LocalNode.NumberOfNodesSeen) + "\n")
	buffer.WriteString(t.Uptime(json.Jvm.UpTime) + "\n")
	buffer.WriteString(t.Version(json.Product.Version))
	return buffer.String()
}
