module github.com/JoneDoe/istorage

go 1.13

require (
	github.com/JoneDoe/istorage/attachment v0.0.0 // indirect
	github.com/JoneDoe/istorage/config v0.0.0
	github.com/JoneDoe/istorage/controllers v0.0.0
	github.com/JoneDoe/istorage/upload v0.0.0 // indirect
	github.com/JoneDoe/istorage/utils v0.0.0 // indirect
	github.com/boltdb/bolt v1.3.1
	github.com/gin-gonic/gin v1.5.0
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/google/uuid v1.1.1
)

replace github.com/JoneDoe/istorage/attachment => ./attachment

replace github.com/JoneDoe/istorage/config => ./config

replace github.com/JoneDoe/istorage/controllers => ./controllers

replace github.com/JoneDoe/istorage/upload => ./upload

replace github.com/JoneDoe/istorage/utils => ./utils
