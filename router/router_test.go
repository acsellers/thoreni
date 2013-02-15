package router

import "testing"

func TestBasicGet(t *testing.T) {
	r := NewRouter()
	r.Get("/", func(ctx *Contextable) {
		ctx.Render("testing")
	})

	testReq := &testRequestish{path: "/", method: "GET"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "testing" {
			t.Fail()
		}
	}
}

func TestBasicGetInvalid(t *testing.T) {
	r := NewRouter()
	r.Get("/asdf", func(ctx *Contextable) {
		ctx.Render("testing")
	})

	testReq := &testRequestish{path: "/", method: "GET"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "Not Found" {
			t.Fail()
		}
	}
}

func TestTwoGets(t *testing.T) {
	r := NewRouter()
	r.Get("/asdf", func(ctx *Contextable) {
		ctx.Render("testing")
	})
	r.Get("/", func(ctx *Contextable) {
		ctx.Render("error")
	})

	testReq := &testRequestish{path: "/asdf", method: "GET"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "testing" {
			t.Fail()
		}
	}
}

func TestTwoGetsReversed(t *testing.T) {
	r := NewRouter()
	r.Get("/", func(ctx *Contextable) {
		ctx.Render("error")
	})
	r.Get("/asdf", func(ctx *Contextable) {
		ctx.Render("testing")
	})

	testReq := &testRequestish{path: "/asdf", method: "GET"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "testing" {
			t.Fail()
		}
	}
}

func TestBasicPost(t *testing.T) {
	r := NewRouter()
	r.Post("/", func(ctx *Contextable) {
		ctx.Render("testing")
	})

	testReq := &testRequestish{path: "/", method: "POST"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "testing" {
			t.Fail()
		}
	}
}

func TestGetAndPost(t *testing.T) {
	r := NewRouter()
	r.Get("/", func(ctx *Contextable) {
		ctx.Render("get")
	})

	r.Post("/", func(ctx *Contextable) {
		ctx.Render("post")
	})

	testReq := &testRequestish{path: "/", method: "POST"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "post" {
			t.Fail()
		}
	}

	testReq2 := &testRequestish{path: "/", method: "GET"}
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "get" {
			t.Fail()
		}
	}
}

func TestBasicDelete(t *testing.T) {
	r := NewRouter()
	r.Delete("/", func(ctx *Contextable) {
		ctx.Render("delete testing")
	})

	testReq := &testRequestish{path: "/", method: "DELETE"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "delete testing" {
			t.Fail()
		}
	}
}

func TestBasicPut(t *testing.T) {
	r := NewRouter()
	r.Put("/", func(ctx *Contextable) {
		ctx.Render("put testing")
	})

	testReq := &testRequestish{path: "/", method: "PUT"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "put testing" {
			t.Fail()
		}
	}
}

func TestBasicAny(t *testing.T) {
	r := NewRouter()
	r.Any("/", func(ctx *Contextable) {
		ctx.Render("any testing")
	})

	testReq := &testRequestish{path: "/", method: "POST"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "any testing" {
			t.Fail()
		}
	}

	testReq2 := &testRequestish{path: "/", method: "GET"}
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "any testing" {
			t.Fail()
		}
	}

	testReq3 := &testRequestish{path: "/", method: "PUT"}
	tc3 := newTestContext(testReq3)
	r.Match(testReq3)(tc3)
	if atc, ok := tc3.Renderable.(*testContext); ok {
		if atc.rendered != "any testing" {
			t.Fail()
		}
	}
}
