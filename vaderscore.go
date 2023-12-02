package main

type VaderScoredSentence struct {
	Sentence           string             `json:"sentence"`
	VaderEmotionScores VaderEmotionScores `json:"vader_emotion_scores"`
}

type VaderEmotionScores struct {
	Neg      float64 `json:"neg"`
	Neu      float64 `json:"neu"`
	Pos      float64 `json:"pos"`
	Compound float64 `json:"compound"`
}
