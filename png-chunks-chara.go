package gopngchunkschara

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	gopngchunks "github.com/chrisbward/go-png-chunks"
)

var charaKey = "chara"

func ReadtEXtChunksFromFile(inputFilePath string) (tEXtChunks []gopngchunks.TEXtChunk, err error) {
	f, err := os.Open(inputFilePath)
	if err != nil {
		return tEXtChunks, fmt.Errorf("os.Open(): %s", err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return tEXtChunks, fmt.Errorf("ioutil.ReadAll(): %s", err)
	}
	tEXtChunks, err = gopngchunks.GetAlltEXtChunks(data)
	if err != nil {
		panic(err)
	}
	return tEXtChunks, nil
}

func (character *CharacterCardV1V2) PopulateFromPng(inputFilePath string) error {
	tEXtChunks, err := ReadtEXtChunksFromFile(inputFilePath)
	if err != nil {
		fmt.Println("Error  :", err)
		return err
	}

	for _, chunk := range tEXtChunks {
		if chunk.Key == charaKey {
			sDec, err := base64.StdEncoding.DecodeString(chunk.Value)
			if err != nil {
				fmt.Println("Error  :", err)
				return err
			}
			fmt.Println("Updated string:", string(sDec))

			err = json.Unmarshal(sDec, &character)
			if err != nil {
				fmt.Println("Error  :", err)
				return err
			}

		}
	}

	return nil

}

func (character *CharacterCardV1V2) WriteToPng(inputFilePath string, outputFilePath string) error {
	f, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("os.Open(): %s", err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll(): %s", err)
	}

	jsonData, err := json.Marshal(character)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Base64 encode the JSON data
	base64Data := base64.StdEncoding.EncodeToString(jsonData)

	tEXtChunkToWrite := gopngchunks.TEXtChunk{
		Key:   charaKey,
		Value: base64Data,
	}
	w, err := gopngchunks.WritetEXtToPngBytes(data, tEXtChunkToWrite)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	out.Write(w.Bytes())

	return nil
}
