# Emote-Score Script

## Overview

Emote-Score is a Python script designed to analyze the
emotional sentiment of text. Utilizing the VADER
(Valence Aware Dictionary and sEntiment Reasoner)
SentimentIntensityAnalyzer, it breaks down the input
text into sentences and calculates sentiment scores for
each. The script is ideal for evaluating the emotional
tone of any textual content, such as social media posts,
reviews, or articles.

## Requirements

- Python 3.x
- VADERSentiment Python package

To install VADERSentiment, run:

```bash
pip install vaderSentiment
```

## Usage

The script can be run from the command line with an input
string and an optional output file.

### Command Line Syntax

```bash
python emote-score.py [-o <output_file>] "<input_string>"
```

- -o <output_file>: Optional. Specify the output file
to save the results in JSON format. If not provided,
- `output.json` will be used.
- <input_string>: Required. The string of text you want
to analyze.

### Examples

Analyzing a direct input string:

```bash
python emote-score.py -o results.json "This is a sample text to analyze."
```

Reading input from a file:

```bash
cat textfile.txt | python emote-score.py -o results.json
```

## Output

The script outputs a JSON file containing the analyzed sentences
along with their VADER sentiment scores. The sentiment scores
include `compound`, `neg`, `neu`, and `pos` values indicating
the overall sentiment and the mix of negative, neutral, and
positive tones.

## Functionality

The Emote-Score script is designed with the following
functionalities:

**Sentence Splitting**: The script processes the input
text and splits it into individual sentences based
on punctuation marks (periods, exclamation marks, and
question marks). This feature ensures that the sentiment
analysis is done on a per-sentence basis, providing a
more granular understanding of the text's emotional tone.

**Sentiment Analysis**: Each sentence is analyzed using the
VADER SentimentIntensityAnalyzer. This tool is specifically
attuned to sentiments expressed in social media and can
handle a variety of textual formats. It provides a composite
score that quantifies the overall sentiment of the sentence
on a scale from highly negative to highly positive.

**Output Generation**: The results of the sentiment analysis
are compiled into a structured JSON format. Each sentence is
listed with its corresponding sentiment scores, including the
composite score as well as separate scores for negative,
neutral, and positive sentiment.

**Flexible Input Handling**: The script can accept input text
either directly through the command line or piped from another
command. This flexibility allows for easy integration with other
tools and workflows.

**Customizable Output**: Users have the option to specify the
output file's name and location, allowing for better organization
of the results.

This functionality makes Emote-Score an effective tool for
analyzing the sentiment of various types of text, especially
useful in fields like social media analysis, market research,
and customer feedback assessment.
