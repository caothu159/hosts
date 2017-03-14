package hosts

import (
	"github.com/google/gxui"
	"reflect"
	"sort"
	"strings"
)

type Host struct {
	ip     string
	domain []string
}

func (h *Host) fromString(str string) *Host {
	strs := strings.Fields(str)
	if len(strs) > 1 {
		h.ip = strs[0]
		h.domain = strs[1:]
	}
	return h
}

type Hosts struct {
	list map[string]Host
}

func (list *Hosts) append(item Host) *Hosts {
	if val, ok := list.list[item.ip]; !ok {
		list.list[item.ip] = item
	} else {
		list.list[item.ip] = val
	}

	return list
}

func (list *Hosts) getIps() []string {
	var ips []string
	rv := reflect.ValueOf(list.list)
	if rv.Kind() != reflect.Map {
		return ips
	}
	t := rv.Type()
	if t.Key().Kind() != reflect.String {
		return ips
	}
	for _, kv := range rv.MapKeys() {
		ips = append(ips, kv.String())
	}
	sort.Strings(ips)
	return ips
}

var lstHosts = Hosts{make(map[string]Host)}

func CreateHosts(theme gxui.Theme) gxui.LinearLayout {
	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)

	overlay := theme.CreateBubbleOverlay()

	layout.AddChild(ListHosts(theme, overlay))

	return layout
}
