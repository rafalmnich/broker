package broker

import (
	"strconv"
	"strings"

	"github.com/dghubble/sling"
	"golang.org/x/xerrors"
)

type Performer interface {
	MakeActions(actions Actions) error
}

type httpPerformer struct {
	client *sling.Sling
	path   string
}

func NewPerformer(client *sling.Sling, path string) *httpPerformer {
	return &httpPerformer{client: client, path: path}
}

func (p httpPerformer) MakeActions(actions Actions) error {
	for _, action := range actions {
		resp, err := p.client.
			Set("Content-Type", "application/x-www-form-urlencoded").
			Set("Accept", "*/*").
			Set("Access-Control-Allow-Origin", "*").
			Post(p.path).
			Body(strings.NewReader(action.Name + "=" + strconv.Itoa(action.Value))).
			ReceiveSuccess(nil)

		_ = resp
		if err != nil {
			return xerrors.Errorf("mass client request failed: %w", err)
		}
	}

	return nil
}
