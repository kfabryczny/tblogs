package main

import (
	"github.com/ezeoleaf/tblogs/cfg"
	"github.com/ezeoleaf/tblogs/helpers"
	"github.com/gdamore/tcell"
	"github.com/pkg/browser"
	"github.com/rivo/tview"
)

func SavedPosts(nextSlide func()) (title string, content tview.Primitive) {
	listSavedPosts := tview.NewList()

	appCfg := cfg.GetAPPConfig()

	if len(appCfg.SavedPosts) == 0 {
		listSavedPosts.AddItem("You don't have saved posts", "Try save them using Ctrl+S", ' ', nil)
	} else {
		listSavedPosts.Clear()

		posts := appCfg.SavedPosts

		for _, post := range posts {
			r := ' '
			isIn, _ := helpers.IsHash(post.Hash, appCfg.SavedPosts)
			if isIn {
				r = 's'
			}
			listSavedPosts.AddItem(post.Title, post.Blog+" - "+post.Published, r, func() {
				return
			})
		}

		listSavedPosts.SetSelectedFunc(func(x int, s string, s1 string, r rune) {
			post := posts[x]
			browser.OpenURL(post.Link)
		})

		listSavedPosts.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyCtrlS {
				appCfg := cfg.GetAPPConfig()

				x := listSavedPosts.GetCurrentItem()

				appCfg.SavedPosts = append(appCfg.SavedPosts[:x], appCfg.SavedPosts[x+1:]...)
				cfg.UpdateAppConfig(appCfg)

				listSavedPosts.RemoveItem(x)
				// listSavedPosts.InsertItem(x, post.Title, post.Blog+" - "+post.Published, '-', func() {
				// 	return
				// })
				return nil
			}
			return event
		})
	}

	return "Saved Posts", tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(listSavedPosts, 0, 1, true), 0, 1, true)
}