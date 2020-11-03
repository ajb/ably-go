package proto_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ably/ably-go/ably/internal/ablyutil"
	"github.com/ably/ably-go/ably/proto"
)

func TestPresenceMessage(t *testing.T) {
	actions := []proto.PresenceAction{
		proto.PresenceAbsent,
		proto.PresencePresent,
		proto.PresenceEnter,
		proto.PresenceLeave,
	}

	for _, a := range actions {
		// pin
		a := a
		id := fmt.Sprint(a)
		m := proto.PresenceMessage{
			Message: proto.Message{
				ID: id,
			},
			Action: a,
		}

		t.Run("json", func(ts *testing.T) {
			b, err := json.Marshal(m)
			if err != nil {
				ts.Fatal(err)
			}
			msg := proto.PresenceMessage{}
			err = json.Unmarshal(b, &msg)
			if err != nil {
				ts.Fatal(err)
			}
			if msg.ID != id {
				ts.Errorf("expected id to be %s got %s", id, msg.ID)
			}
			if msg.Action != a {
				ts.Errorf("expected action to be %d got %d", a, msg.Action)
			}
		})
		t.Run("msgpack", func(ts *testing.T) {
			b, err := ablyutil.MarshalMsgpack(m)
			if err != nil {
				ts.Fatal(err)
			}
			msg := proto.PresenceMessage{}
			err = ablyutil.UnmarshalMsgpack(b, &msg)
			if err != nil {
				ts.Fatal(err)
			}
			if msg.ID != id {
				ts.Errorf("expected id to be %s got %s", id, msg.ID)
			}
			if msg.Action != a {
				ts.Errorf("expected action to be %d got %d", a, msg.Action)
			}
		})
	}
}
