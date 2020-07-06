package startinstance

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestStartInstance(t *testing.T) {
	cases := []struct {
		desc string
		data string
		want string
	}{
		{
			desc: "sucess no data",
			data: "",
			want: "Hello, Daigo!\n",
		},
		{
			desc: "sucess  data",
			data: "Go",
			want: "Hello, Go!\n",
		},
	}

	for _, c := range cases {
		r, w, _ := os.Pipe()
		log.SetOutput(w)
		originalFlags := log.Flags()
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

		m := PubSubMessage{
			Data: []byte(c.data),
		}
		StartInstance(context.Background(), m)

		w.Close()
		log.SetOutput(os.Stderr)
		log.SetFlags(originalFlags)

		out, err := ioutil.ReadAll(r)
		if err != nil {
			t.Fatalf("ReadAll: %v", err)
		}
		if got := string(out); got != c.want {
			t.Errorf("StartInstance(%q) = %q, want %q", c.data, got, c.want)
		}
	}
}
