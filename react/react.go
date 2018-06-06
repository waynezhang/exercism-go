package react

type reactorImpl struct {
	cellMap map[Cell][]*computeCellImpl
}

func (reactor *reactorImpl) addCell(source Cell, destination *computeCellImpl) {
	if list, ok := reactor.cellMap[source]; ok {
		reactor.cellMap[source] = append(list, destination)
	} else {
		reactor.cellMap[source] = []*computeCellImpl{destination}
	}
}

func (reactor *reactorImpl) getCells(source Cell) []*computeCellImpl {
	var cells = make([]*computeCellImpl, 0)
	for _, c := range reactor.cellMap[source] {
		cells = append(cells, c)
		cells = append(cells, reactor.getCells(c)...)
	}
	return cells
}

func (reactor *reactorImpl) notify(source Cell) {
	var toNotify = make(map[*computeCellImpl]int)
	for _, c := range reactor.getCells(source) {
		if _, ok := toNotify[c]; !ok {
			toNotify[c] = c.value
		}
		c.compute()
	}
	for cell, oldValue := range toNotify {
		if cell.value == oldValue {
			continue
		}
		for _, canceler := range cell.cancelers {
			canceler.callback(cell.value)
		}
	}
}

// Cell
type cellImpl struct {
	value int
}

func (cell *cellImpl) Value() int {
	return cell.value
}

// InputCell
type inputCellImpl struct {
	cellImpl
	reactor *reactorImpl
}

func (cell *inputCellImpl) SetValue(value int) {
	if cell.value == value {
		return
	}
	cell.value = value
	cell.reactor.notify(cell)
}

// CreateInput doc here
func (reactor *reactorImpl) CreateInput(value int) InputCell {
	c := inputCellImpl{reactor: reactor}
	c.value = value
	return &c
}

// ComputeCell
type computeCellImpl struct {
	cellImpl
	reactor   *reactorImpl
	cancelers []*cancelerImpl
	compute   func()
}

type cancelerImpl struct {
	cell     *computeCellImpl
	callback func(int)
}

func (canceler *cancelerImpl) Cancel() {
	cancelers := canceler.cell.cancelers
	for i, c := range cancelers {
		if c != canceler {
			continue
		}

		if i == 0 {
			cancelers = cancelers[i+1:]
		} else if i == len(cancelers)-1 {
			cancelers = cancelers[:i]
		} else if i == len(cancelers)-1 {
			cancelers = append(cancelers[:i-1], cancelers[i+1:]...)
		}
		break
	}
	canceler.cell.cancelers = cancelers
}

func (cell *computeCellImpl) AddCallback(callback func(int)) Canceler {
	canceler := cancelerImpl{cell: cell, callback: callback}
	cell.cancelers = append(cell.cancelers, &canceler)
	return &canceler
}

// CreateCompute1 doc here
func (reactor *reactorImpl) CreateCompute1(cell Cell, computeFunc func(value int) int) ComputeCell {
	c := computeCellImpl{reactor: reactor}
	c.compute = func() {
		c.value = computeFunc(cell.Value())
	}
	c.compute()

	reactor.addCell(cell, &c)
	return &c
}

// CreateCompute2 doc here
func (reactor *reactorImpl) CreateCompute2(cell1 Cell, cell2 Cell, computeFunc func(int, int) int) ComputeCell {
	c := computeCellImpl{reactor: reactor}
	c.compute = func() {
		c.value = computeFunc(cell1.Value(), cell2.Value())
	}
	c.compute()

	reactor.addCell(cell1, &c)
	reactor.addCell(cell2, &c)
	return &c
}

// New doc here
func New() Reactor {
	return &reactorImpl{cellMap: make(map[Cell][]*computeCellImpl)}
}
