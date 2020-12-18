package mcconfig

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TemplateTestSuite struct {
	suite.Suite

	output bytes.Buffer
}

func (suite *TemplateTestSuite) TestRunScript() {
	cfg := &RunScript{
		MemoryLimit: 512,
	}
	suite.Assert().NoError(cfg.WriteTemplate(&suite.output))
	expected := `#!/bin/sh

java -Xms512M -Xmx512M -XX:+UseG1GC -XX:+ParallelRefProcEnabled -XX:MaxGCPauseMillis=200 -XX:+UnlockExperimentalVMOptions -XX:+DisableExplicitGC -XX:+AlwaysPreTouch -XX:G1NewSizePercent=30 -XX:G1MaxNewSizePercent=40 -XX:G1HeapRegionSize=8M -XX:G1ReservePercent=20 -XX:G1HeapWastePercent=5 -XX:G1MixedGCCountTarget=4 -XX:InitiatingHeapOccupancyPercent=15 -XX:G1MixedGCLiveThresholdPercent=90 -XX:G1RSetUpdatingPauseTimePercent=5 -XX:SurvivorRatio=32 -XX:+PerfDisableSharedMem -XX:MaxTenuringThreshold=1 -Dusing.aikars.flags=https://mcflags.emc.gs -Daikars.new.flags=true -jar /spigot.jar nogui 2>&1`
	suite.Assert().Equal(expected, suite.output.String())
}

func (suite *TemplateTestSuite) TestServerProperties() {
	cfg := &ServerProperties{
		SpawnProtection:             16,
		QueryPort:                   25565,
		ForceGamemode:               false,
		EnforceWhitelist:            false,
		Gamemode:                    SurvivalGamemode,
		EnableQuery:                 false,
		PlayerIdleTimeout:           0,
		Difficulty:                  EasyDifficulty,
		SpawnMonsters:               true,
		Pvp:                         true,
		LevelType:                   DefaultLevelType,
		EnableStatus:                true,
		Hardcore:                    false,
		EnableCommandBlock:          false,
		NetworkCompressionThreshold: 256,
		MaxPlayers:                  20,
		MaxWorldSize:                29999984,
		RconPort:                    25575,
		ServerPort:                  25565,
		Debug:                       false,
		SpawnNpcs:                   true,
		AllowFlight:                 false,
		LevelName:                   "world",
		ViewDistance:                10,
		SpawnAnimals:                true,
		WhiteList:                   false,
		RconPassword:                "password",
		GenerateStructures:          true,
		OnlineMode:                  true,
		MaxBuildHeight:              256,
		PreventProxyConnections:     false,
		Motd:                        "A Minecraft Server",
		RateLimit:                   0,
		EnableRcon:                  true,
	}
	suite.Assert().NoError(cfg.WriteTemplate(&suite.output))
	expected := `spawn-protection=16
max-tick-time=60000
query.port=25565
generator-settings=
sync-chunk-writes=true
force-gamemode=false
allow-nether=true
enforce-whitelist=false
gamemode=survival
broadcast-console-to-ops=true
enable-query=false
player-idle-timeout=0
text-filtering-config=
difficulty=easy
broadcast-rcon-to-ops=true
spawn-monsters=true
op-permission-level=4
pvp=true
entity-broadcast-range-percentage=100
snooper-enabled=false
level-type=default
enable-status=true
hardcore=false
enable-command-block=false
network-compression-threshold=256
max-players=20
max-world-size=29999984
resource-pack-sha1=
function-permission-level=2
rcon.port=25575
server-port=25565
debug=false
server-ip=
spawn-npcs=true
allow-flight=false
level-name=world
view-distance=10
resource-pack=
spawn-animals=true
white-list=false
rcon.password=password
generate-structures=true
online-mode=true
max-build-height=256
level-seed=
prevent-proxy-connections=false
use-native-transport=true
enable-jmx-monitoring=false
motd=A Minecraft Server
rate-limit=0
enable-rcon=true
`
	suite.Assert().Equal(expected, suite.output.String())
}

func TestTemplate(t *testing.T) {
	suite.Run(t, &TemplateTestSuite{})
}
