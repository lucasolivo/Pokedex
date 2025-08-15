package main

import (
	"fmt"
	"net/http"       
    "encoding/json"
	"io"
)

type Movestruct struct {
	Power       *int   `json:"power"`       // use *int because null is possible
    Accuracy    *int   `json:"accuracy"`    // use *int because null is possible
    Type        NamedResource `json:"type"` // e.g. { "name": "normal", "url": "..." }
    DamageClass NamedResource `json:"damage_class"`
    Priority    int           `json:"priority"`
}

type NamedResource struct {
    Name string `json:"name"`
}

func addMoveData(mon Pokemon, move string) (Pokemon, error) {
    if _, exists := mon.Movedata[move]; exists {
        return mon, nil
    }

    url := "https://pokeapi.co/api/v2/move/" + move
    res, err := http.Get(url)
    if err != nil {
        return mon, err
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return mon, fmt.Errorf("Move %s not found", move)
    }

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return mon, err
    }

    var moveRaw Movestruct
    err = json.Unmarshal(body, &moveRaw)
    if err != nil {
        return mon, err
    }

    if mon.Movedata == nil {
        mon.Movedata = make(map[string]Move)
    }

    // Convert *int to string safely
    var powerStr, accuracyStr string
    if moveRaw.Power != nil {
        powerStr = fmt.Sprintf("%d", *moveRaw.Power)
    } else {
        powerStr = "N/A"
    }



    if moveRaw.Accuracy != nil {
        accuracyStr = fmt.Sprintf("%d", *moveRaw.Accuracy)
    } else {
        accuracyStr = "N/A"
    }

    newMove := Move{
        Power:       powerStr,
        Accuracy:    accuracyStr,
        Poketype:    moveRaw.Type.Name,
        Damagetype:  moveRaw.DamageClass.Name,
        Priority:    moveRaw.Priority,
    }

    mon.Movedata[move] = newMove

    return mon, nil
}
