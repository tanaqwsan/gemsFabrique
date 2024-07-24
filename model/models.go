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
	GroupType   string `json:"group_type" form:"group_type"`
	GroupOwner  string `json:"group_owner" form:"group_owner"`
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
	LastAccessed          int    `json:"last_accessed" form:"last_accessed"`
	Problem               string `json:"problem" form:"problem"`
	TileCCTVCount         int    `json:"tile_cctv_count" form:"tile_cctv_count"`
	TileSLCount           int    `json:"tile_sl_count" form:"tile_sl_count"`
	TileBLCount           int    `json:"tile_bl_count" form:"tile_bl_count"`
	TileHLCount           int    `json:"tile_hl_count" form:"tile_hl_count"`
	TileDoorCount         int    `json:"tile_door_count" form:"tile_door_count"`
	IsDoorAccessible      int    `json:"is_door_accessible" form:"is_door_accessible"`
	IsPathFindAble        int    `json:"is_path_find_able" form:"is_path_find_able"`
	LastEvent             string `json:"last_event" form:"last_event"`
	*gorm.Model
}

type Word struct {
	Growid string `json:"growid" form:"growid"`
	Word   string `json:"word" form:"word"`
	Target string `json:"target" form:"target"`
	*gorm.Model
}

type CustomWhere struct {
	Where     string `json:"where" form:"where"`
	FieldSort string `json:"field_sort" form:"field_sort"`
	TypeSort  string `json:"type_sort" form:"type_sort"`
}
