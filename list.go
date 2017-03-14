package hosts

import (
	"bufio"
	"github.com/google/gxui"
	"log"
	"os"
	"time"
)

func getListHosts() []string {
	file, err := os.Open("/etc/hosts")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		item := Host{}
		item.fromString(scanner.Text())
		lstHosts.append(item)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return lstHosts.getIps()
}

func ListHosts(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {
	adapter := gxui.CreateDefaultAdapter()
	driver := theme.Driver()
	go func() {
		for {
			driver.CallSync(func() {
				adapter.SetItems(getListHosts())
			})
			time.Sleep(1500 * time.Millisecond)
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
