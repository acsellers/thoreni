package router

import "github.com/acsellers/thoreni"
import "testing"

func TestBasicGet(t *testing.T) {
	r := NewRouter()
	r.Get("/", func(ctx *thoreni.Contextable) {
		ctx.Render("testing", nil)
	})

	testReq := requestish("/", "GET")
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
	r.Get("/asdf", func(ctx *thoreni.Contextable) {
		ctx.Render("testing", nil)
	})

	testReq := requestish("/", "GET")
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
	r.Get("/asdf", func(ctx *thoreni.Contextable) {
		ctx.Render("testing", nil)
	})
	r.Get("/", func(ctx *thoreni.Contextable) {
		ctx.Render("error", nil)
	})

	testReq := requestish("/asdf", "GET")
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
	r.Get("/", func(ctx *thoreni.Contextable) {
		ctx.Render("error", nil)
	})
	r.Get("/asdf", func(ctx *thoreni.Contextable) {
		ctx.Render("testing", nil)
	})

	testReq := requestish("/asdf", "GET")
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
	r.Post("/", func(ctx *thoreni.Contextable) {
		ctx.Render("testing", nil)
	})

	testReq := requestish("/", "POST")
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
	r.Get("/", func(ctx *thoreni.Contextable) {
		ctx.Render("get", nil)
	})

	r.Post("/", func(ctx *thoreni.Contextable) {
		ctx.Render("post", nil)
	})

	testReq := requestish("/", "POST")
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "post" {
			t.Fail()
		}
	}

	testReq2 := requestish("/", "GET")
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
	r.Update("/", func(ctx *thoreni.Contextable) {
		ctx.Render("update testing", nil)
	})

	testReq := requestish("/", "UPDATE")
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
	r.Head("/", func(ctx *thoreni.Contextable) {
		ctx.Render("head testing", nil)
	})

	testReq := requestish("/", "HEAD")
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
	r.Options("/", func(ctx *thoreni.Contextable) {
		ctx.Render("options testing", nil)
	})

	testReq := requestish("/", "OPTIONS")
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
	r.Delete("/", func(ctx *thoreni.Contextable) {
		ctx.Render("delete testing", nil)
	})

	testReq := requestish("/", "DELETE")
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
	r.Put("/", func(ctx *thoreni.Contextable) {
		ctx.Render("put testing", nil)
	})

	testReq := requestish("/", "PUT")
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
	r.Any("/", func(ctx *thoreni.Contextable) {
		ctx.Render("any testing", nil)
	})

	testReq := requestish("/", "POST")
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "any testing" {
			t.Fail()
		}
	}

	testReq2 := requestish("/", "GET")
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "any testing" {
			t.Fail()
		}
	}

	testReq3 := requestish("/", "PUT")
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
	r.Get("/a/:id", func(ctx *thoreni.Contextable) {
		ctx.Render("id testing", nil)
	})

	testReq1 := requestish("/", "GET")
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered == "id testing" {
			t.Fatal("ID Testing, invalid match")
		}
	}

	testReq2 := requestish("/a/123", "GET")
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "id testing" {
			t.Fatal("ID Testing, valid match")
		}
	}
	testReq3 := requestish("/a/arel", "GET")
	tc3 := newTestContext(testReq3)
	r.Match(testReq3)(tc3)
	if atc, ok := tc3.Renderable.(*testContext); ok {
		if atc.rendered == "id testing" {
			t.Fatal("ID Testing, invalid match (second)")
		}
	}
	testReq4 := requestish("/a/123arel", "GET")
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
	r.Get("/a/:id$", func(ctx *thoreni.Contextable) {
		ctx.Render("id testing", nil)
	})

	testReq1 := requestish("/", "GET")
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered == "id testing" {
			t.Fatal("ID Testing, invalid match")
		}
	}

	testReq2 := requestish("/a/123", "GET")
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "id testing" {
			t.Fatal("ID Testing, valid match")
		}
	}
	testReq3 := requestish("/a/arel", "GET")
	tc3 := newTestContext(testReq3)
	r.Match(testReq3)(tc3)
	if atc, ok := tc3.Renderable.(*testContext); ok {
		if atc.rendered == "id testing" {
			t.Fatal("ID Testing, invalid match (second)")
		}
	}
	testReq4 := requestish("/a/123arel", "GET")
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
	r.Get("/", func(ctx *thoreni.Contextable) {
		ctx.Render("root testing", nil)
	})

	r.Get("/a/:id", func(ctx *thoreni.Contextable) {
		ctx.Render("id testing", nil)
	})

	r.Get("/a/:key", func(ctx *thoreni.Contextable) {
		ctx.Render("key testing", nil)
	})

	testReq1 := requestish("/", "GET")
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered != "root testing" {
			t.Fatal("key Testing, root match")
		}
	}

	testReq2 := requestish("/a/123", "GET")
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "id testing" || atc.rendered != "key testing" {
			t.Fatal("key Testing, id match")
		}
	}
	testReq3 := requestish("/a/arel", "GET")
	tc3 := newTestContext(testReq3)
	r.Match(testReq3)(tc3)
	if atc, ok := tc3.Renderable.(*testContext); ok {
		if atc.rendered != "key testing" {
			t.Fatal("key Testing, invalkey match (second)")
		}
	}

	testReq4 := requestish("/a/123arel", "GET")
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
	r.Get("/a/:key$", func(ctx *thoreni.Contextable) {
		ctx.Render("key testing", nil)
	})

	testReq1 := requestish("/", "GET")
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered == "key testing" {
			t.Fatal("key Testing, invalkey match")
		}
	}

	testReq2 := requestish("/a/123abcef", "GET")
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "key testing" {
			t.Fatal("key Testing, valkey match")
		}
	}
	testReq3 := requestish("/a/arel", "GET")
	tc3 := newTestContext(testReq3)
	r.Match(testReq3)(tc3)
	if atc, ok := tc3.Renderable.(*testContext); ok {
		if atc.rendered == "key testing" {
			t.Fatal("key Testing, invalkey match (second)")
		}
	}
	testReq4 := requestish("/a/123arel", "GET")
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
	r.Get("/a/:slug$", func(ctx *thoreni.Contextable) {
		ctx.Render("slug testing", nil)
	})

	testReq1 := requestish("/", "GET")
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered == "slug testing" {
			t.Fatal("slug Testing, invalslug match")
		}
	}

	testReq2 := requestish("/a/123abcef", "GET")
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "slug testing" {
			t.Fatal("slug Testing, valslug match")
		}
	}
	testReq3 := requestish("/a/arel", "GET")
	tc3 := newTestContext(testReq3)
	r.Match(testReq3)(tc3)
	if atc, ok := tc3.Renderable.(*testContext); ok {
		if atc.rendered != "slug testing" {
			t.Fatal("slug Testing, invalslug match (second)")
		}
	}
	testReq4 := requestish("/a/123arel", "GET")
	tc4 := newTestContext(testReq4)
	r.Match(testReq4)(tc4)
	if atc, ok := tc4.Renderable.(*testContext); ok {
		if atc.rendered != "slug testing" {
			t.Fatal("slug Testing, invalslug match (third)")
		}
	}
	testReq5 := requestish("/a/123arel/t/asdf", "GET")
	tc5 := newTestContext(testReq5)
	r.Match(testReq5)(tc5)
	if atc, ok := tc5.Renderable.(*testContext); ok {
		if atc.rendered != "slug testing" {
			t.Fatal("slug Testing, invalslug match (third)")
		}
	}
}

func TestnameColorOperator2(t *testing.T) {
	r := NewRouter()
	r.Get("/a/:name$", func(ctx *thoreni.Contextable) {
		ctx.Render("name testing", nil)
	})

	testReq1 := requestish("/", "GET")
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered == "name testing" {
			t.Fatal("name Testing, invalname match")
		}
	}

	testReq2 := requestish("/a/123abcef", "GET")
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "name testing" {
			t.Fatal("name Testing, valname match")
		}
	}
	testReq3 := requestish("/a/arel", "GET")
	tc3 := newTestContext(testReq3)
	r.Match(testReq3)(tc3)
	if atc, ok := tc3.Renderable.(*testContext); ok {
		if atc.rendered != "name testing" {
			t.Fatal("name Testing, invalname match (second)")
		}
	}
	testReq4 := requestish("/a/123arel", "GET")
	tc4 := newTestContext(testReq4)
	r.Match(testReq4)(tc4)
	if atc, ok := tc4.Renderable.(*testContext); ok {
		if atc.rendered != "name testing" {
			t.Fatal("name Testing, invalname match (third)")
		}
	}
	testReq5 := requestish("/a/123arel/t/asdf", "GET")
	tc5 := newTestContext(testReq5)
	r.Match(testReq5)(tc5)
	if atc, ok := tc5.Renderable.(*testContext); ok {
		if atc.rendered != "name testing" {
			t.Fatal("name Testing, invalname match (third)")
		}
	}
}

func TestRoot(t *testing.T) {
	r := NewRouter()
	r.Root(func(ctx *thoreni.Contextable) {
		ctx.Render("root testing", nil)
	})

	testReq1 := requestish("/", "GET")
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered != "root testing" {
			t.Fatal("root Testing, valid match")
		}
	}

	testReq2 := requestish("/a/123abcef", "GET")
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered == "root testing" {
			t.Fatal("root Testing, invalid match")
		}
	}

}

func BasicNamespaceTest(t *testing.T) {
	r := NewRouter()
	nm := r.Namespace("what")
	nm.Root(func(ctx *thoreni.Contextable) {
		ctx.Render("namespace testing", nil)
	})

	testReq1 := requestish("/", "GET")
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered == "namespace testing" {
			t.Fatal("namespace Testing, invalid match")
		}
	}

	testReq2 := requestish("/what/", "GET")
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "namespace testing" {
			t.Fatal("namespace Testing, valid match")
		}
	}
}

func BasicNamespaceGetTest(t *testing.T) {
	r := NewRouter()
	nm := r.Namespace("what")
	nm.Get("now", func(ctx *thoreni.Contextable) {
		ctx.Render("namespace get testing", nil)
	})

	testReq1 := requestish("/", "GET")
	tc1 := newTestContext(testReq1)
	r.Match(testReq1)(tc1)
	if atc, ok := tc1.Renderable.(*testContext); ok {
		if atc.rendered == "namespace get testing" {
			t.Fatal("namespace Testing, invalid match")
		}
	}

	testReq2 := requestish("/what/now", "GET")
	tc2 := newTestContext(testReq2)
	r.Match(testReq2)(tc2)
	if atc, ok := tc2.Renderable.(*testContext); ok {
		if atc.rendered != "namespace get testing" {
			t.Fatal("namespace Testing, valid match")
		}
	}
}
