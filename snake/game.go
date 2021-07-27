package snake

type Game struct {
	stage *Stage
}

func NewGame(stg *Stage) *Game {
	return &Game{stage: stg}
}

func (g *Game) IsOver() bool {
	// スネークの頭の位置が壁か、スネークの胴体ならゲームオーバー
	h := g.stage.snake.GetCurrentHeadPosition()
	return g.stage.IsWall(h) || g.stage.IsSnakeBody(h)
}
