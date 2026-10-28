// Package config общая загрузка конфигурации сервисов. Сами поля у каждого сервиса свои:
// сервис описывает структуру с тегами env (github.com/caarlos0/env) и передаёт её в Load.
package config

import (
	"errors"
	"io/fs"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Load читает .env (если он есть), затем конфигурацию из окружения в cfg
// и проверяет обязательные поля. cfg указатель на структуру сервиса.
func Load(cfg any) error {
	// .env нужен только локально, в проде переменные задаются окружением.
	// Отсутствие файла не ошибка, а вот битый файл это ошибка.
	if err := godotenv.Load(); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}

	return env.Parse(cfg)
}
