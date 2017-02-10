package hosts

import (
	"github.com/google/gxui"
)

func CreateHosts(theme gxui.Theme) gxui.LinearLayout {
	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)

	overlay := theme.CreateBubbleOverlay()

	layout.AddChild(ListHosts(theme, overlay))

	return layout
}
