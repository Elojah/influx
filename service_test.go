package influx_test

import (
	"testing"

	"github.com/elojah/influx"
	"github.com/elojah/services"
)

func TestUp(t *testing.T) {
	s := &influx.Service{}
	l := s.NewLauncher(influx.Namespaces{
		Influx: "influx",
	}, "influx")

	ls := services.Launchers{}
	ls = append(ls, l)
	if err := ls.Up("config_test.json"); err != nil {
		t.Error(err)
	}
}
