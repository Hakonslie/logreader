package events

import (
	"encoding/json"
	"errors"
	"regexp"
)

func parseAttributeWithName(name, s string) string {
	pattern := regexp.MustCompile(" " + name + ":(-?[0-9]+) ")
	return pattern.FindStringSubmatch(s)[1]
}
func parseSingleCharacterName(s string) string {
	pattern := regexp.MustCompile("(\\S+)<[0-9a-z]{8}>")
	return pattern.FindStringSubmatch(s)[1]
}
func parseSingleCharacterID(s string) string {
	pattern := regexp.MustCompile("<([0-9a-z]{8})>")
	return pattern.FindStringSubmatch(s)[1]
}

func ReadEvent(s []string) (Event, error) {
	var e Event
	switch s[0] {
	case "ActionAddPlayer":
		e.Name = s[0]
		j := struct{ Health, Player, Id, Place string }{}
		j.Health = parseAttributeWithName("Health", s[1])
		j.Place = parseAttributeWithName("Place", s[1])
		j.Player = parseSingleCharacterName(s[1])
		j.Id = parseSingleCharacterID(s[1])
		m, err := json.Marshal(j)
		if err != nil {
			return Event{}, err
		}
		e.EventJ = string(m)
		return e, nil
	case "NewGameStarted":
		e.Name = "New Game"
		return e, nil
	default:
		return e, errors.New("no event")
	}
}
