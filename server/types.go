package server

import "github.com/gin-gonic/gin"

type H = gin.H
type Context = gin.Context
type Engine = gin.Engine

type Middleware = gin.HandlerFunc

var WrapH = gin.WrapH
var NewEngine = gin.Default
