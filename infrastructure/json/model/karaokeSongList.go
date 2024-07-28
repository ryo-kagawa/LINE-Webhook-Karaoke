package model

import (
	"encoding/json"
	"io"
	"os"

	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/constants"
)

type KaraokeSongList []struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	DamId      *string `json:"damId"`
	JoysoundId *string `json:"joysoundId"`
	Songs      []struct {
		Id         int64   `json:"id"`
		Name       string  `json:"name"`
		Lyrics     string  `json:"lyrics"`
		DamId      *string `json:"damId"`
		JoysoundId *string `json:"joysoundId"`
	} `json:"songs"`
}

func Load() (KaraokeSongList, error) {
	file, err := os.Open(constants.KaraokeSongListFile)
	if err != nil {
		return KaraokeSongList{}, err
	}
	defer file.Close()
	binary, err := io.ReadAll(file)
	if err != nil {
		return KaraokeSongList{}, err
	}
	jsonData := KaraokeSongList{}
	err = json.Unmarshal(binary, &jsonData)
	if err != nil {
		return KaraokeSongList{}, err
	}
	return jsonData, nil
}
