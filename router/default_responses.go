package router

// Default404 is the default 404 handler. If you want to change the content of the page, 
// simply put a valid template file in templates/builtins/ with the name 404. This function
// will then use that template. If you wish to implement a custom function for this purpose, 
// set the NotFound RoutingFunc on the Router for you application with your new RoutingFunc.
func Default404(ctx *Contextable) {
	ctx.Render("Not Found")
}
