package asset

import "embed"

//go:embed all:swagger
var Dist embed.FS

//go:embed swagger/index.html
var IndexHTML embed.FS
