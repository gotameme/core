package internal

import (
	"fmt"
	"github.com/gotameme/core/internal/resources"
	"github.com/gotameme/core/internal/simulation"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"log"
)

var (
	startAnimation *resources.TimedAnimation
	Quit           = fmt.Errorf("quit game")
)

type GameState int

// define the game states
const (
	// GameStateStart is the initial state of the game
	GameStateStart GameState = iota
	// GameStateStartSimulation is the state when the game is starting
	GameStateStartSimulation
	// GameStateRunning is the state when the game is running
	GameStateRunning
	// GameStatePaused is the state when the game is paused
	GameStatePaused
	// GameStateEnd is the state when the game is over
	GameStateEnd
)

type Game struct {
	screenWidth, screenHeight int
	GameConfiguration
	s *simulation.Simulation
}

func NewGame(screenWidth, screenHeight int, cnf GameConfiguration) *Game {
	startAnimation = resources.NewGoTameMeAnimation(float64(screenWidth), float64(screenHeight))
	return &Game{
		screenWidth:       screenWidth,
		screenHeight:      screenHeight,
		GameConfiguration: cnf,
		s:                 simulation.NewSimulation(screenWidth, screenHeight, cnf.SimulationConfig),
	}
}

// region GameState

func (g *Game) togglePause() {
	if g.state == GameStateRunning {
		g.state = GameStatePaused
	} else if g.state == GameStatePaused {
		g.state = GameStateRunning
	}
}

func (g *Game) pauseGame() {
	g.state = GameStatePaused
}

func (g *Game) startSimulation() {
	g.state = GameStateStartSimulation
}

func (g *Game) startGame() {
	g.state = GameStateStart
}

func (g *Game) endGame() {
	g.state = GameStateEnd
}

// endregion

func (g *Game) handleKeyEvents() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.state == GameStateStart {
			g.startSimulation()
		} else if g.state == GameStateEnd {
			// start new game if game is over
			g.startGame()
		} else {
			g.togglePause()
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if g.state == GameStateRunning {
			g.pauseGame()
		} else if g.state == GameStatePaused {
			g.endGame()
		} else if g.state == GameStateEnd {
			return Quit
		}
	}
	return nil
}

func (g *Game) Update() error {
	if err := g.handleKeyEvents(); err != nil {
		return err
	}
	switch g.state {
	case GameStateStartSimulation:
		if !startAnimation.Update() {
			startAnimation.Reset()
			g.state = GameStateRunning
		}
		// fallthrough to update the game content while the game is starting
		fallthrough
	case GameStateRunning:
		// run simulation logic when the game is running
		g.updateRunning()
	case GameStatePaused:
		// placeholder for paused game logic
	case GameStateStart:
		// placeholder for start screen logic
	case GameStateEnd:
		// placeholder for end screen logic
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x80, G: 0xc0, B: 0xa0, A: 0xff})
	switch g.state {
	case GameStateRunning:
		g.drawRunning(screen)
	case GameStatePaused:
		// draw the game content when the game is paused
		ebitenutil.DebugPrint(screen, "Game is paused - press escape to end the game or space to continue")
		g.drawRunning(screen)
	case GameStateStart:
		// draw the start screen
		ebitenutil.DebugPrint(screen, "Start screen - press space to start the game")
	case GameStateStartSimulation:
		// draw the current game content while the game is starting
		g.drawRunning(screen)
		// draw the start animation
		startAnimation.Draw(screen)
	case GameStateEnd:
		// draw the end screen
		ebitenutil.DebugPrint(screen, "End screen - press escape to end the game or space to restart the game")
	}
}

func (g *Game) updateRunning() {
	// run simulation logic when the game is running
	if err := g.s.Update(); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) drawRunning(screen *ebiten.Image) {
	// draw the game content when the game is running
	g.s.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return g.screenWidth, g.screenHeight
}
