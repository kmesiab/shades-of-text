package main

import (
	"encoding/json"
	"image/color"
	"log"
	"os/exec"
	"path/filepath"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/kmesiab/shades-of-text/ui"
)

func main() {
	application := app.New()
	window, appWindow := ui.Initialize(application)

	application.Settings().SetTheme(theme.DarkTheme())

	appWindow.InputBox.OnChanged = func(text string) {
		processed, err := processInputString(text)

		if err != nil {
			log.Println(err)
			return
		}

		for i, sentence := range processed {
			var style *widget.CustomTextGridStyle

			c := sentimentToColor(sentence.VaderEmotionScores)
			style = &widget.CustomTextGridStyle{
				BGColor: c,
			}

			textGridRow := widget.TextGridRow{}

			// Style the sentence, rune by rune
			for _, r := range sentence.Sentence {
				textGridRow.Cells = append(
					textGridRow.Cells,
					widget.TextGridCell{Rune: r, Style: style},
				)
			}

			appWindow.TextGrid.SetRow(i, textGridRow)

			logMsgf := "Sentence: %s\nVaderEmotionScores: %s\n"
			logBody, err := json.MarshalIndent(sentence.VaderEmotionScores, "", "  ")

			if err != nil {
				continue
			}

			log.Printf(logMsgf, sentence.Sentence, string(logBody))
		}
	}

	window.ShowAndRun()
}

func processInputString(input string) ([]VaderScoredSentence, error) {
	// Invoke Python script
	interpreter := "python3"
	scriptPath := "./bin/emote-score.py"
	scriptAbsPath, err := filepath.Abs(scriptPath)

	if err != nil {
		return nil, err
	}

	cmd := exec.Command(interpreter, scriptAbsPath, input)
	stdout, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	// Unmarshal output
	var sentences []VaderScoredSentence
	if err := json.Unmarshal(stdout, &sentences); err != nil {
		return nil, err
	}

	return sentences, nil
}

func sentimentToColor(scores VaderEmotionScores) color.RGBA {

	green := uint8((scores.Compound + 1) / 2 * 255)
	red := uint8((1 - (scores.Compound+1)/2) * 255)

	if red < 0 {
		red = 0
	} else if red > 255 {
		red = 255
	}

	if green < 0 {
		green = 0
	} else if green > 255 {
		green = 255
	}

	return color.RGBA{
		R: red,
		G: green,
		B: 0,
		A: 255,
	}
}
