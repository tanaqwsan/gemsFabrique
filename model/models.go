package model

import (
	"gorm.io/gorm"
)

type Bot struct {
	Growid      string `json:"growid" form:"growid"`
	Age         int    `json:"age" form:"age"`
	Gems        int    `json:"gems" form:"gems"`
	Level       int    `json:"level" form:"level"`
	IsSuspended int    `json:"is_suspended" form:"is_suspended"` //0 = No, 1 = Yes
	Whatever    string `json:"whatever" form:"whatever"`
	*gorm.Model
}

type World struct {
	Name                  string `json:"name" form:"name"`
	NameId                string `json:"name_id" form:"name_id"`
	Owner                 string `json:"owner" form:"owner"`
	Type                  string `json:"type" form:"type"`                   // Farm or Storage
	IsSmallLock           int    `json:"is_small_lock" form:"is_small_lock"` //0 = No, 1 = Yes
	SLOwner               string `json:"sl_owner" form:"sl_owner"`
	IsNuked               int    `json:"is_nuked" form:"is_nuked"` //0 = No, 1 = Yes
	SmallLockAge          int    `json:"small_lock_age" form:"small_lock_age"`
	FloatPepperBlockCount int    `json:"float_pepper_block_count" form:"float_pepper_block_count"`
	FloatPepperSeedCount  int    `json:"float_pepper_seed_count" form:"float_pepper_seed_count"`
	TilePepperSeedCount   int    `json:"tile_pepper_seed_count" form:"tile_pepper_seed_count"`
	TilePepperBlockCount  int    `json:"tile_pepper_block_count" form:"tile_pepper_block_count"`
	FossilCount           int    `json:"fossil_count" form:"fossil_count"`
	BotHandlerId          int    `json:"bot_handler_id" form:"bot_handler_id"`
	Gems                  int    `json:"gems" form:"gems"`
	*gorm.Model
}

type Word struct {
	Growid string `json:"growid" form:"growid"`
	Word   string `json:"word" form:"word"`
	Target string `json:"target" form:"target"`
	*gorm.Model
}
