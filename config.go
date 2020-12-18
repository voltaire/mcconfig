//go:generate go run github.com/go-bindata/go-bindata/go-bindata -pkg mcconfig templates
package mcconfig

import (
	"io"
	"text/template"
)

type Gamemode string

var (
	AdventureGamemode Gamemode = "adventure"
	CreativeGamemode  Gamemode = "creative"
	SpectatorGamemode Gamemode = "spectator"
	SurvivalGamemode  Gamemode = "survival"
)

type Difficulty string

var (
	PeacefulDifficulty Difficulty = "peaceful"
	EasyDifficulty     Difficulty = "easy"
	NormalDifficulty   Difficulty = "normal"
	HardDifficulty     Difficulty = "hard"
)

type LevelType string

var (
	DefaultLevelType     LevelType = "default"
	FlatLevelType        LevelType = "flat"
	LargeBiomesLevelType LevelType = "largeBiomes"
	AmplifiedLevelType   LevelType = "amplified"
	BuffetLevelType      LevelType = "buffet"
)

func Bool(b bool) *bool {
	return &b
}

func Int(i int) *int {
	return &i
}

func Int64(i int64) *int64 {
	return &i
}

func String(s string) *string {
	return &s
}

type ServerProperties struct {
	// spawn protection radius
	SpawnProtection             *int
	QueryPort                   *int
	GeneratorSettings           *string
	EnforceWhitelist            *bool
	ForceGamemode               *bool
	Gamemode                    *Gamemode
	EnableQuery                 *bool
	PlayerIdleTimeout           *int64
	Difficulty                  *Difficulty
	SpawnMonsters               *bool
	Pvp                         *bool
	LevelType                   *LevelType
	EnableStatus                *bool
	Hardcore                    *bool
	EnableCommandBlock          *bool
	NetworkCompressionThreshold *int64
	MaxPlayers                  *int64
	MaxWorldSize                *int64
	ResourcePackSha1            *string
	RconPort                    *int
	ServerPort                  *int
	Debug                       *bool
	ServerIp                    *string
	SpawnNpcs                   *bool
	AllowFlight                 *bool
	LevelName                   *string
	ViewDistance                *int64
	ResourcePack                *string
	SpawnAnimals                *bool
	WhiteList                   *bool
	RconPassword                *string
	GenerateStructures          *bool
	OnlineMode                  *bool
	MaxBuildHeight              *int64
	LevelSeed                   *string
	PreventProxyConnections     *bool
	Motd                        *string
	RateLimit                   *int64
	EnableRcon                  *bool
}

func (cfg *ServerProperties) WriteTemplate(w io.Writer) error {
	data, err := Asset("templates/server.properties")
	if err != nil {
		return err
	}
	tmpl, err := template.New("server.properties").Parse(string(data))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, cfg)
}

type RunScript struct {
	// memory limit in MB
	MemoryLimit int64
}

func (cfg *RunScript) WriteTemplate(w io.Writer) error {
	data, err := Asset("templates/run")
	if err != nil {
		return err
	}
	tmpl, err := template.New("run").Parse(string(data))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, cfg)
}
