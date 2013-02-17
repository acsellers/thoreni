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

func TestBasicUpdate(t *testing.T) {
	r := NewRouter()
	r.Update("/", func(ctx *Contextable) {
		ctx.Render("update testing")
	})

	testReq := &testRequestish{path: "/", method: "UPDATE"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "update testing" {
			t.Fail()
		}
	}
}

func TestBasicHead(t *testing.T) {
	r := NewRouter()
	r.Head("/", func(ctx *Contextable) {
		ctx.Render("head testing")
	})

	testReq := &testRequestish{path: "/", method: "HEAD"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "head testing" {
			t.Fail()
		}
	}
}

func TestBasicOptions(t *testing.T) {
	r := NewRouter()
	r.Options("/", func(ctx *Contextable) {
		ctx.Render("options testing")
	})

	testReq := &testRequestish{path: "/", method: "OPTIONS"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "options testing" {
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

func TestIDColorOperator(t *testing.T) {
	r := NewRouter()
	r.Get("/a/:id", func(ctx *Contextable) {
		ctx.Render("id testing")
	})

	testReq1 := &testRequestish{path: "/", method: "GET"}
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered == "id testing" {
			t.Fatal("ID Testing, invalid match")
		}
	}

	testReq2 := &testRequestish{path: "/a/123", method: "GET"}
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "id testing" {
			t.Fatal("ID Testing, valid match")
		}
	}
	testReq3 := &testRequestish{path: "/a/arel", method: "GET"}
	tc3 := newTestContext(testReq3)
	r.Match(testReq3)(tc3)
	if atc, ok := tc3.Renderable.(*testContext); ok {
		if atc.rendered == "id testing" {
			t.Fatal("ID Testing, invalid match (second)")
		}
	}
	testReq4 := &testRequestish{path: "/a/123arel", method: "GET"}
	tc4 := newTestContext(testReq4)
	r.Match(testReq4)(tc4)
	if atc, ok := tc4.Renderable.(*testContext); ok {
		if atc.rendered != "id testing" {
			t.Fatal("ID Testing, valid match (second)")
		}
	}

}
func TestIDColorOperator2(t *testing.T) {
	r := NewRouter()
	r.Get("/a/:id$", func(ctx *Contextable) {
		ctx.Render("id testing")
	})

	testReq1 := &testRequestish{path: "/", method: "GET"}
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered == "id testing" {
			t.Fatal("ID Testing, invalid match")
		}
	}

	testReq2 := &testRequestish{path: "/a/123", method: "GET"}
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "id testing" {
			t.Fatal("ID Testing, valid match")
		}
	}
	testReq3 := &testRequestish{path: "/a/arel", method: "GET"}
	tc3 := newTestContext(testReq3)
	r.Match(testReq3)(tc3)
	if atc, ok := tc3.Renderable.(*testContext); ok {
		if atc.rendered == "id testing" {
			t.Fatal("ID Testing, invalid match (second)")
		}
	}
	testReq4 := &testRequestish{path: "/a/123arel", method: "GET"}
	tc4 := newTestContext(testReq4)
	r.Match(testReq4)(tc4)
	if atc, ok := tc4.Renderable.(*testContext); ok {
		if atc.rendered == "id testing" {
			t.Fatal("ID Testing, invalid match (third)")
		}
	}
}

func TestkeyColorOperator(t *testing.T) {
	r := NewRouter()
	r.Get("/", func(ctx *Contextable) {
		ctx.Render("root testing")
	})

	r.Get("/a/:id", func(ctx *Contextable) {
		ctx.Render("id testing")
	})

	r.Get("/a/:key", func(ctx *Contextable) {
		ctx.Render("key testing")
	})

	testReq1 := &testRequestish{path: "/", method: "GET"}
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered != "root testing" {
			t.Fatal("key Testing, root match")
		}
	}

	testReq2 := &testRequestish{path: "/a/123", method: "GET"}
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "id testing" || atc.rendered != "key testing" {
			t.Fatal("key Testing, id match")
		}
	}
	testReq3 := &testRequestish{path: "/a/arel", method: "GET"}
	tc3 := newTestContext(testReq3)
	r.Match(testReq3)(tc3)
	if atc, ok := tc3.Renderable.(*testContext); ok {
		if atc.rendered != "key testing" {
			t.Fatal("key Testing, invalkey match (second)")
		}
	}

	testReq4 := &testRequestish{path: "/a/123arel", method: "GET"}
	tc4 := newTestContext(testReq4)
	r.Match(testReq4)(tc4)
	if atc, ok := tc4.Renderable.(*testContext); ok {
		if atc.rendered != "key testing" {
			t.Fatal("key Testing, valkey match (second)")
		}
	}

}
func TestkeyColorOperator2(t *testing.T) {
	r := NewRouter()
	r.Get("/a/:key$", func(ctx *Contextable) {
		ctx.Render("key testing")
	})

	testReq1 := &testRequestish{path: "/", method: "GET"}
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered == "key testing" {
			t.Fatal("key Testing, invalkey match")
		}
	}

	testReq2 := &testRequestish{path: "/a/123abcef", method: "GET"}
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "key testing" {
			t.Fatal("key Testing, valkey match")
		}
	}
	testReq3 := &testRequestish{path: "/a/arel", method: "GET"}
	tc3 := newTestContext(testReq3)
	r.Match(testReq3)(tc3)
	if atc, ok := tc3.Renderable.(*testContext); ok {
		if atc.rendered == "key testing" {
			t.Fatal("key Testing, invalkey match (second)")
		}
	}
	testReq4 := &testRequestish{path: "/a/123arel", method: "GET"}
	tc4 := newTestContext(testReq4)
	r.Match(testReq4)(tc4)
	if atc, ok := tc4.Renderable.(*testContext); ok {
		if atc.rendered == "key testing" {
			t.Fatal("key Testing, invalkey match (third)")
		}
	}
}

func TestslugColorOperator2(t *testing.T) {
	r := NewRouter()
	r.Get("/a/:slug$", func(ctx *Contextable) {
		ctx.Render("slug testing")
	})

	testReq1 := &testRequestish{path: "/", method: "GET"}
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered == "slug testing" {
			t.Fatal("slug Testing, invalslug match")
		}
	}

	testReq2 := &testRequestish{path: "/a/123abcef", method: "GET"}
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "slug testing" {
			t.Fatal("slug Testing, valslug match")
		}
	}
	testReq3 := &testRequestish{path: "/a/arel", method: "GET"}
	tc3 := newTestContext(testReq3)
	r.Match(testReq3)(tc3)
	if atc, ok := tc3.Renderable.(*testContext); ok {
		if atc.rendered != "slug testing" {
			t.Fatal("slug Testing, invalslug match (second)")
		}
	}
	testReq4 := &testRequestish{path: "/a/123arel", method: "GET"}
	tc4 := newTestContext(testReq4)
	r.Match(testReq4)(tc4)
	if atc, ok := tc4.Renderable.(*testContext); ok {
		if atc.rendered != "slug testing" {
			t.Fatal("slug Testing, invalslug match (third)")
		}
	}
	testReq5 := &testRequestish{path: "/a/123arel/t/asdf", method: "GET"}
	tc5 := newTestContext(testReq5)
	r.Match(testReq5)(tc5)
	if atc, ok := tc5.Renderable.(*testContext); ok {
		if atc.rendered != "slug testing" {
			t.Fatal("slug Testing, invalslug match (third)")
		}
	}
}
