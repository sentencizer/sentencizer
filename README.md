# **Sentencizer: Sentence Splitting (Sentence Boundary Disambiguation) Library for Go**

<img align="right" width="320" src="/artifacts/sbd-gopher.png" alt="sentencizer-logo" title="dsbd-logo" />

[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/sentencizer/sentencizer)

Sentencizer is a library for segmenting text into sentences, designed to make it easier to build Retrieval Augmented Generation (RAG) systems in Go. It is inspired by [pySBD](https://github.com/nipunsadvilkar/pySBD) and [pragmatic_segmenter](https://github.com/diasks2/pragmatic_segmenter), and works out-of-the-box with a rule-based approach.

## Playground

Try out Sentencizer in our [online playground](https://sentencizer.pages.dev/).

## Features

- **Sentence Splitting**: Efficiently breaks down a block of text into individual sentences.
- **Lightweight and Easy Integration**: Designed to be lightweight and easy to integrate into existing Go projects.
- **High Accuracy**: Offers high accuracy in sentence segmentation. For more details, see [pySBD](https://github.com/nipunsadvilkar/pySBD).
- **Fast Sentence Splitting**: Sentencizer aims to provide high-performance sentence splitting by leveraging Go's efficiency.
- **Non-Destructive Splitting**: Segments text into sentences without altering the original content.
- **Language-Specific Configuration**: Adaptable to handle punctuation rules specific to different languages.
- **Text Cleaning**: Equipped with features to manage and clean noisy text, including:
  - Handling irregular newline characters and spacing
  - Processing Tables of Contents
  - Recognizing and managing URLs and HTML tags
  - Dealing with sentences that are delimited without any space

_Note: Text Cleaning feature is to be implemented. Contributions are greatly welcomed._

## Installation

To install sentencizer, you can use `go get`:

```sh
go get github.com/sentencizer/sentencizer
```

## Usage

Here's a basic example of how to use sentencizer:

```go
package main

import (
    "fmt"
    "github.com/sentencizer/sentencizer"
)

// This example segments a text string into individual sentences.
func main() {
    segmenter := sentencizer.NewSegmenter("en")
    text := "This is a sentence. And this is another one."
    sentences := segmenter.Segment(text)
    for _, sentence := range sentences {
        fmt.Println(sentence)
    }
}
```

## Roadmap

- [x] Add Online Playground.
- [ ] Add chuking feature with overlapping option.
- [ ] Setup Codecov for monitoring test coverage.
- [ ] Implement text cleaner.
- [ ] Add support for more languages.
- [ ] Add benchmark test.
- [ ] Setup GitHub Action for testing.

## Language Support Roadmap

The following table outlines our current language support. We're actively seeking contributions to expand this list. If you're interested in contributing, consider helping us add support for a language, whether it's listed below or not. Your expertise in a language not listed here could be a valuable addition to our project.

| Language  | ISO Code | Supported |
| --------- | -------- | --------- |
| Amharic   | am       | Planned   |
| Arabic    | ar       | Planned   |
| Armenian  | hy       | Planned   |
| Bulgarian | bg       | Planned   |
| Burmese   | my       | Planned   |
| Chinese   | zh       | Yes       |
| Danish    | da       | Planned   |
| Deutsch   | de       | Planned   |
| Dutch     | nl       | Planned   |
| English   | en       | Yes       |
| French    | fr       | Planned   |
| Greek     | el       | Planned   |
| Hindi     | hi       | Planned   |
| Italian   | it       | Planned   |
| Japanese  | ja       | Yes       |
| Kazakh    | kk       | Planned   |
| Marathi   | mr       | Planned   |
| Persian   | fa       | Planned   |
| Polish    | pl       | Planned   |
| Russian   | ru       | Yes       |
| Slovak    | sk       | Planned   |
| Spanish   | es       | Planned   |
| Urdu      | ur       | Planned   |

We welcome contributions that help us add support for these languages. Please feel free to submit a Pull Request with your contributions.

## Motivation

Sentence splitting is a crucial step in the preprocessing pipeline of Natural Language Processing (NLP) tasks, especially for building Retrieval Augmented Generation (RAG) systems. RAG systems rely on accurately segmented sentences to retrieve relevant information and generate coherent responses.

While libraries like pragmatic_segmenter and pySBD are known for their high accuracy and efficiency in sentence splitting, there are no equivalent libraries available in Go. This poses a challenge for developers building RAG systems in Go, as they need to rely on external libraries or implement their own sentence splitting logic.

Sentencizer aims to bridge this gap by providing a reliable and efficient sentence splitting solution in Go. By offering a native Go library for sentence splitting, Sentencizer simplifies the process of building RAG systems and other NLP applications entirely within the Go ecosystem. This not only streamlines the development workflow but also enables faster execution times by leveraging Go's performance characteristics.

## Acknowledgement

This library builds upon the excellent foundations laid by [pySBD](https://github.com/nipunsadvilkar/pySBD) and [pragmatic_segmenter](https://github.com/diasks2/pragmatic_segmenter).

## Contributing

Contributions are greatly appreciated and crucial for this project! Here are a few ways you can contribute:

- **Add new tests and rules**: Improve the accuracy of sentence segmentation by adding new tests and rules.
- **Add support for a new language**: Help expand the reach of this library by adding support for new languages.
- **Port features**: Help improve this library by porting features that are supported in pySBD and pragmatic_segmenter.

Please feel free to submit a Pull Request with your contributions.

## License

This project is licensed under the MIT License.
