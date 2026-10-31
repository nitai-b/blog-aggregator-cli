package main

import (
	"github.com/nitai-b/blog-aggregator/internal/config"
	"github.com/nitai-b/blog-aggregator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}
