package worldserver

type Phase struct {
	Obj map[int64]*Object
}

type Object struct {
	X, Y, Z     float32
	Orientation float32
}

func (ws *WorldServer) GetPhase(i int64) *Phase {
	ws.PhaseL.Lock()
	ph := ws.Phases[i]
	if ph == nil {
		ph = &Phase{
			Obj: make(map[int64]*Object),
		}

		ws.Phases[i] = ph
	}
	ws.PhaseL.Unlock()
	return ph
}
