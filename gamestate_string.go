// Code generated by "stringer -type=GameState"; DO NOT EDIT.

package main

import "strconv"

const _GameState_name = "InitialLoginCreatePlayerInGame"

var _GameState_index = [...]uint8{0, 7, 12, 24, 30}

func (i GameState) String() string {
	if i < 0 || i >= GameState(len(_GameState_index)-1) {
		return "GameState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GameState_name[_GameState_index[i]:_GameState_index[i+1]]
}
