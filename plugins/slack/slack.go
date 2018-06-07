package slack

import (
	"encoding/json"
	"flag"
	"image/color"

	"github.com/ieee0824/getenv"
	"github.com/ieee0824/sakuya"
	"github.com/ieee0824/skyclad/notifer"
)

func init() {
	notifer.Register("slack", New())
}

type Slack struct {
	APIURL   *string
	Username string
}

func New() *Slack {
	return &Slack{
		APIURL: flag.String("slack-api", getenv.String("SLACK_INCOMING_API_URL"), "slack incoming url. \"SLACK_INCOMING_API_URL\" may be specified as an environment variable."),
	}
}

func (s *Slack) SetEncoder(e notifer.Encoder) {
}

func (s *Slack) Notice(v interface{}) error {
	w := sakuya.NewIncomingWriter(*s.APIURL, s.Username)
	w.SetBaseColor(color.RGBA{
		0xff,
		0x00,
		0x00,
		0xff,
	})
	bin, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if _, err := w.Write(bin); err != nil {
		return err
	}
	return nil
}
