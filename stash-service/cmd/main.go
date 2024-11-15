package main

import (
	"joicejoseph.dev/clutter-cache/stash-service"
)


func main(){
	app, err := app.NewApp()
    if err != nil {
        // Handle the error
        return
    }

    app.Run();

	
}