package formatter

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/haraldringvold/enonicstatus/jsonstruct"
)

var green string = "#36a64f"
var yellow string = "#f5f625"
var red string = "#df0000"

type SlackFormatter struct {
}

func (s SlackFormatter) HostName(name string) string {
	return name
}

func (s SlackFormatter) IndexStatus(index string) string {
	return index
}

func (s SlackFormatter) Master(master string) string {
	return master
}

func (s SlackFormatter) NodesSeen(nodesSeen float64) string {
	return strconv.FormatFloat(nodesSeen, 'f', -1, 64)
}

func (s SlackFormatter) Uptime(uptime float64) string {
	uptimeString := strconv.FormatFloat(uptime, 'f', -1, 64)
	duration := fmt.Sprintf("%sms", uptimeString)
	formattedUptime, _ := time.ParseDuration(duration)

	return formattedUptime.String()
}

func (s SlackFormatter) Version(version string) string {
	return version
}

func (s SlackFormatter) String(jsonData []jsonstruct.Status) string {
	slackmessage := SlackMessage{Attachments: []SlackAttachment{}}
	for _,element := range jsonData {
		slackmessage.AddAttachment(SlackAttachment{
					Fallback: s.HostName(element.Cluster.LocalNode.HostName) + "s index is " + s.IndexStatus(element.Index.Status),
					Color:   s.SlackAttachmentColor(element.Index.Status),
					Title: s.HostName(element.Cluster.LocalNode.HostName),
					Fields: []SlackAttachmentField{
						SlackAttachmentField{
							Title: "Index",
							Value: s.IndexStatus(element.Index.Status),
							Inline: true},
						SlackAttachmentField{
							Title: "Master",
							Value: s.Master(element.Cluster.LocalNode.Master),
							Inline: true},
						SlackAttachmentField{
							Title: "Nodes seen",
							Value: s.NodesSeen(element.Cluster.LocalNode.NumberOfNodesSeen),
							Inline: true},
						SlackAttachmentField{
							Title: "Uptime",
							Value: s.Uptime(element.Jvm.UpTime),
							Inline: true},
						SlackAttachmentField{
							Title: "Version",
							Value: s.Version(element.Product.Version),
							Inline: true}}})
	}
	slackmessageAsJson, _ := json.Marshal(slackmessage)

	return string(slackmessageAsJson)
}

func (s SlackFormatter) SlackAttachmentColor(index string) string {
	var color string
	if "GREEN" == index {
		color = green
	} else if "YELLOW" == index {
		color = yellow
	} else {
		color = red
	}
	return color
}

type SlackMessage struct {
	Attachments []SlackAttachment `json:"attachments"`
}

func (sm *SlackMessage) AddAttachment(attachment SlackAttachment) []SlackAttachment {
	sm.Attachments = append(sm.Attachments, attachment)
	return sm.Attachments
}

type SlackAttachment struct {
	Fallback string `json:"fallback"`
	Color string `json:"color"`
	Title string `json:"title"`
	Fields []SlackAttachmentField `json:"fields"`
}

type SlackAttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Inline bool `json:"short"`
}
