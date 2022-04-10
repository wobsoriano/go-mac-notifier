package notifier

import (
	"errors"
	"fmt"

	"github.com/wobsoriano/go-jxa"
	version "github.com/wobsoriano/go-macos-version"
)

type notifierOptions struct {
	title    string
	text     string
	subtitle string
	sound    string
}

func DisplayNotification(n notifierOptions) error {
	version.AssertMacOSVersion(">=10.9")

	if len(n.title) == 0 {
		return errors.New("title or text is required")
	}

	_, err := jxa.RunJXA(fmt.Sprintf(`
		const app = Application.currentApplication()

		app.includeStandardAdditions = true
		
		app.displayNotification("%s", {
				withTitle: "%s",
				subtitle: "%s",
				soundName: "%s"
		})
	`, n.text, n.title, n.subtitle, n.sound))

	return err
}
