# go-png-chunks-chara
Reads and writes embedded data in PNGs for AI Character Card, V1 &amp; V2 spec

Leverages [go-png-chunks](https://github.com/chrisbward/go-png-chunks) to read tEXt chunks from PNG, and then marshalls the b64 encoded data of the [Character Card V1 & V2 spec](https://github.com/malfoyslastname/character-card-spec-v2/blob/main/spec_v2.md) to a `CharacterCardV1V2` struct.

Example usage;

```go
package main

import (
	gopngchunkschara "github.com/chrisbward/go-png-chunks-chara"
)

func CreateCharacterCard(characterCard gopngchunkschara.CharacterCardV1V2, inputFilePath string, outputFilePath string) error {

    err := characterCard.WriteToPng(inputFilePath, outputFilePath)
	if err != nil {
		return err
	}
	return nil

}

func main() {

    characterCard := gopngchunkschara.CharacterCardV1V2{
        Name:        "Chris Ward",
        Description: "",
        ... etc
    }

    CreateCharacterCard(characterCard, "test.png", "out.png")
}

```
 