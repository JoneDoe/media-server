module bitbucket.org/vadimtitov/istorage

go 1.13

require (
	github.com/boltdb/bolt v1.3.1
	github.com/gin-gonic/gin v1.5.0
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/google/uuid v1.1.1
)

replace bitbucket.org/vadimtitov/istorage/attachment => ./attachment

replace bitbucket.org/vadimtitov/istorage/config => ./config

replace bitbucket.org/vadimtitov/istorage/controllers => ./controllers

replace bitbucket.org/vadimtitov/istorage/upload => ./upload

replace bitbucket.org/vadimtitov/istorage/utils => ./utils
