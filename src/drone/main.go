package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strconv"

	_ "go.uber.org/automaxprocs"
	tele "gopkg.in/telebot.v3"
)

func must(err error) {
	if err == nil {
		return
	}

	slog.Error("err in main", slog.Any("err", err))
	os.Exit(1)
}

func parseEnvInt(key string) (int, error) {
	val := os.Getenv(key)
	if len(val) == 0 {
		return 0, nil
	}
	return strconv.Atoi(val)
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	boarDWhiteChatID, err := parseEnvInt("DRONE_BOAR_D_WHITE_CHAT_ID")
	must(err)
	boarDWhiteLeetCodeThreadID, err := parseEnvInt("DRONE_BOAR_D_WHITE_LEET_CODE_THREAD_ID")
	must(err)
	cfg := Config{
		TgKey:                      os.Getenv("DRONE_TG_BOT_API_KEY"),
		LCDailyCron:                os.Getenv("DRONE_LC_DAILY_CRON"),
		LCDailyStickerID:           os.Getenv("DRONE_LC_DAILY_STICKER_ID"),
		BoarDWhiteChatID:           tele.ChatID(boarDWhiteChatID),
		BoarDWhiteLeetCodeThreadID: boarDWhiteLeetCodeThreadID,
	}
	if len(cfg.LCDailyCron) == 0 {
		cfg.LCDailyCron = "0 1 * * *" // every day at 01:00 UTC
	}
	if len(cfg.LCDailyStickerID) == 0 {
		cfg.LCDailyStickerID = "CAACAgIAAxkBAAELpiFl7G4Rn8WQBK3AaDiAMn6ixTUR7gACzzkAAr-TAAFK91qMnVpp9TQ0BA"
	}
	if cfg.BoarDWhiteChatID == 0 {
		cfg.BoarDWhiteChatID = tele.ChatID(-1001640461540)
	}
	if cfg.BoarDWhiteLeetCodeThreadID == 0 {
		cfg.BoarDWhiteLeetCodeThreadID = 10095
	}

	must(StartDrone(ctx, cfg))
}
