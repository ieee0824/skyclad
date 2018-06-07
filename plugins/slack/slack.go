package slack

import (
	"encoding/json"
	"flag"
	"fmt"
	"reflect"

	"github.com/ieee0824/getenv"
	"github.com/ieee0824/sakuya"
	"github.com/ieee0824/skyclad/client/io"
	"github.com/ieee0824/skyclad/notifer"
)

func init() {
	notifer.Register("slack", New())
}

type Slack struct {
	APIURL   *string
	Username *string
	Channel  *string
	Title    *string
}

func New() *Slack {
	return &Slack{
		APIURL:   flag.String("slack-api", getenv.String("SLACK_INCOMING_API_URL"), "slack incoming url. \"SLACK_INCOMING_API_URL\" may be specified as an environment variable."),
		Username: flag.String("slack-username", getenv.String("SLACK_USERNAME"), "slack username."),
		Channel:  flag.String("slack-channel", "", "slack channel"),
		Title:    flag.String("slack-title", "", "slack message title"),
	}
}

func (s *Slack) SetEncoder(e notifer.Encoder) {
}

func (s *Slack) Notice(v interface{}) error {
	w := sakuya.NewIncomingWriter(*s.APIURL, *s.Username)
	switch v := v.(type) {
	case []clientio.GetContainerOutput:
		w.AddTitle(*s.Title)
		for _, o := range v {
			bin, err := json.MarshalIndent(o, "", "  ")
			if err != nil {
				continue
			}
			w.AddInfo(string(bin))
		}

		return w.Flush()
	default:
		fmt.Println(reflect.TypeOf(v))
		bin, err := json.Marshal(v)
		if err != nil {
			return err
		}
		if _, err := w.Write(bin); err != nil {
			return err
		}
		return nil
	}
}
