package hosts

import (
	"bufio"
	"github.com/google/gxui"
	"log"
	"os"
)

func getListHosts() []string {
	// file, err := os.Open("/etc/hosts.deny")
	file, err := os.Open("/etc/hosts")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	items := *[]string

	for scanner.Scan() {
		append(items, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return items
}

func ListHosts(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {
	adapter := gxui.CreateDefaultAdapter()
	listHostsChan := make()
	go func() {
		for {
			<-listHostsChan
			adapter.SetItems(getListHosts())
		}
	}()

	list := theme.CreateList()
	list.SetAdapter(adapter)
	list.SetOrientation(gxui.Vertical)

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)
	layout.AddChild(list)

	return layout
}
