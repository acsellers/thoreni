package router

func Default404(ctx *Contextable) {
	ctx.Render("Not Found")
}
