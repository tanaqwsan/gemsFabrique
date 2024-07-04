package res

import (
	"app/model"
	"app/model/web"
	"time"
)

func ConvertIndexWorld(worlds []model.World) []web.GetWorldResponse {
	var results []web.GetWorldResponse
	for _, world := range worlds {
		worldResponse := web.GetWorldResponse{
			Id:                    int(world.ID),
			Name:                  world.Name,
			NameId:                world.NameId,
			Owner:                 world.Owner,
			Type:                  world.Type,
			IsSmallLock:           world.IsSmallLock,
			SLOwner:               world.SLOwner,
			IsNuked:               world.IsNuked,
			SmallLockAge:          world.SmallLockAge,
			FloatPepperBlockCount: world.FloatPepperBlockCount,
			FloatPepperSeedCount:  world.FloatPepperSeedCount,
			TilePepperSeedCount:   world.TilePepperSeedCount,
			TilePepperBlockCount:  world.TilePepperBlockCount,
			FossilCount:           world.FossilCount,
			BotHandlerId:          world.BotHandlerId,
			Gems:                  world.Gems,
			LastAccessed:          world.LastAccessed,
		}
		results = append(results, worldResponse)
	}

	return results

}
func ConvertIndexWorldNameIdOnly(worlds []model.World) []web.GetWorldNameIdOnlyResponse {
	var results []web.GetWorldNameIdOnlyResponse
	for _, world := range worlds {
		worldResponse := web.GetWorldNameIdOnlyResponse{
			Id:           int(world.ID),
			Name:         world.Name,
			NameId:       world.NameId,
			BotHandlerId: world.BotHandlerId,
		}
		results = append(results, worldResponse)
	}

	return results

}

func ConvertIndexBot(bots []model.Bot) []web.GetBotResponse {
	var results []web.GetBotResponse
	for _, bot := range bots {
		botResponse := web.GetBotResponse{
			Id:          int(bot.ID),
			Growid:      bot.Growid,
			Age:         bot.Age,
			Gems:        bot.Gems,
			Level:       bot.Level,
			IsSuspended: bot.IsSuspended,
			Whatever:    bot.Whatever,
		}
		results = append(results, botResponse)
	}

	return results

}

func ConvertIndexWorldOnlyName(worlds []model.World) []web.GetWorldResponse {
	var results []web.GetWorldResponse
	for _, world := range worlds {
		worldResponse := web.GetWorldResponse{
			Name: world.Name,
		}
		results = append(results, worldResponse)
	}

	return results
}

func ConvertIndexWord(words []model.Word) []web.GetWordResponse {
	var results []web.GetWordResponse
	for _, word := range words {
		wordResponse := web.GetWordResponse{
			Id:     int(word.ID),
			Growid: word.Growid,
			Word:   word.Word,
			Target: word.Target,
		}
		results = append(results, wordResponse)
	}

	return results

}

func currentEpochTime() int64 {
	return time.Now().Unix()
}

func humanReadableToEpoch(date time.Time) int64 {
	return date.Unix()
}

func epochToHumanReadable(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}
