package main

import (
	"time"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
)

const gridCentre = 128
const outsideGrid = gridCentre + 1
const maxUndo = 10000
const clipboardSize = 64
const clipboardFile = "cache/clipboards.dat"

const superWidth = 2048
const superHeight = 1280
const superTiles = (superWidth/128)*(superHeight/128)

var (
	windowTitlePrefix    = "OrthoEngine"
	frames                                              = 0
	undoFrame                                  = 1
	second                                              = time.Tick(time.Second)
	win               *pixelgl.Window
	textRenderer      *text.Text
	luaRenderer      *text.Text
	textLine          int
	luaLine			  int
	tilePic           pixel.Picture
	tileTexture       [superTiles]*pixel.Sprite
	tileOverlay       *pixel.Batch
	tileOverlayWidth  uint16
	tileOverlayHeight uint16
	imd1              *imdraw.IMDraw
	imd2              *imdraw.IMDraw
	batch             *pixel.Batch
	grid              [2*gridCentre][2*gridCentre][16][2]uint16
	clipboard         [10][clipboardSize][clipboardSize][16][2]uint16
	clipboardWidth    [10]int
	clipboardHeight   [10]int
	clipboardShift                 = 0
	currentClipboard         = 1
	previewClipboard         = -1
	clobber                                                               		 = true
	undo            [maxUndo][6]int //0 frame,  1 x,  2 y,  3 z,  4 base,  5 front
	undoCounter            = 0
	scale                  = 0.5
	aspect                 = 0.5
	viewDirection          = 0
	compass                = [4]string{"North", "East", "South", "West"}
	hScale                 = 64.0
	vScale                 = hScale * aspect
	lastTileX              = outsideGrid
	lastTileY              = 0
	selectedTile1   uint16 = 4
	tileRow1        uint16 = 0
	selectedTile2   uint16 = 0
	tileRow2        uint16 = 0
	cameraX                = 0.0 //128.0*gridCentre
	cameraY                = 0.0 //128.0*gridCentre
	mouseX                 = 0.0
	mouseY                 = 0.0
	tileX                  = 0
	tileY                  = 0
	tileZ                  = 0
	hideTile               = false
	showShadows            = true
	showGrid               = 1
	xRay                   = false
	zRay                   = false
	selectionStartX        = 0
	selectionStartY        = 0
	selectionEndX          = 0
	selectionEndY          = 0
	selectionLive          = false
	leftAltPressed         = false
	rightAltPressed        = false
	quit                   = 0
	help                   = false
)

