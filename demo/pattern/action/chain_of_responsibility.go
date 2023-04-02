package action

type ResponsibilityLibrary interface {
	Render(*UI)
	SetNext(library ResponsibilityLibrary)
}

type AResponsibility struct {
	next ResponsibilityLibrary
}

func (r *AResponsibility) Render(ui *UI) {
	// do logic
	if ui.Done {
		return
	}
	if r.next == nil {
		return
	}
	r.next.Render(ui)
}

func (r *AResponsibility) SetNext(next ResponsibilityLibrary) { r.next = next }

type UI struct {
	Done bool
	Data string
}
