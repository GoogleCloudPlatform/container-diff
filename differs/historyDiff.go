package differs

import (
	"strings"

	"github.com/GoogleCloudPlatform/container-diff/utils"
)

type HistoryAnalyzer struct {
}

func (a HistoryAnalyzer) Diff(image1, image2 utils.Image) (utils.DiffResult, error) {
	diff, err := getHistoryDiff(image1, image2)
	return &utils.HistDiffResult{
		Image1:   image1.Source,
		Image2:   image2.Source,
		DiffType: "History",
		Diff:     diff,
	}, err
}

func (a HistoryAnalyzer) Analyze(image utils.Image) (utils.AnalyzeResult, error) {
	history := getHistoryList(image.Config.History)
	result := utils.ListAnalyzeResult{
		Image:       image.Source,
		AnalyzeType: "History",
		Analysis:    history,
	}
	return &result, nil
}

func getHistoryDiff(image1, image2 utils.Image) (utils.HistDiff, error) {
	history1 := getHistoryList(image1.Config.History)
	history2 := getHistoryList(image2.Config.History)

	adds := utils.GetAdditions(history1, history2)
	dels := utils.GetDeletions(history1, history2)
	diff := utils.HistDiff{adds, dels}
	return diff, nil
}

func getHistoryList(historyItems []utils.ImageHistoryItem) []string {
	strhistory := make([]string, len(historyItems))
	for i, layer := range historyItems {
		strhistory[i] = strings.TrimSpace(layer.CreatedBy)
	}
	return strhistory
}
