package graphbox

type Spacer struct {
	Margin Point
}

func (sp *Spacer) Constraint(r, c int, applier ConstraintApplier) {
	applier.Apply(SizeConstraint{r, c, sp.Margin.X / 2, sp.Margin.X / 2, sp.Margin.Y / 2, sp.Margin.Y / 2})
}

func (sp *Spacer) Draw(ctx DrawContext, point Point) {
}
