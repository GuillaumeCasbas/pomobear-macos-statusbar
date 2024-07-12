package adapters

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/objc"
)

func Start() {
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		item := appkit.StatusBar_SystemStatusBar().StatusItemWithLength(-1)
		objc.Retain(&item)
		item.Button().SetTitle("Pomobear UI")

		go func() {
			for {
				cmd := exec.Command("pomobear", "status")

				out, err := cmd.Output()

				trimmed := bytes.TrimSuffix(out, []byte{10})

				dispatch.MainQueue().DispatchAsync(func() {
					item.Button().SetTitle(string(trimmed))
				})

				if err != nil {
					fmt.Println(err)
				}

				time.Sleep(1 * time.Second)
			}
		}()

		itemStart := appkit.NewMenuItemWithAction("Start", "start", func(sender objc.Object) {
			exec.Command("pomobear", "start").Run()
		})

		itemStop := appkit.NewMenuItemWithAction("Stop", "st", func(sender objc.Object) {
			exec.Command("pomobear", "stop").Run()
		})

		itemQuit := appkit.NewMenuItem()
		itemQuit.SetTitle("Quit")
		itemQuit.SetAction(objc.Sel("terminate:"))

		menu := appkit.NewMenu()
		menu.AddItem(itemStart)
		menu.AddItem(itemStop)
		menu.AddItem(itemQuit)
		item.SetMenu(menu)
	})
}
