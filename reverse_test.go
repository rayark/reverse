package reverse

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	showError := func(info string) {
		t.Error(fmt.Sprintf("Error: %s. urlStore: %s", info, Urls))
	}

	if Urls.MustAdd("firstUrl", "/first") != "/first" {
		showError("0")
	}

	if Urls.MustAdd("helloUrl", "/hello/:p1:p2", "1", "2") != "/hello/:p1:p2" {
		showError("0-1")
	}

	Urls.MustAdd("secondUrl", "/second/:param/:param2", ":param", ":param2")

	// re := regexp.MustCompile("^/comment/(?P<id>\d+)$")
	Urls.MustAdd("thirdUrl", "/comment/:p1", ":p1")

	if Urls.getParam("helloUrl", 1) != "2" {
		showError("1")
	}

	if Urls.Get("helloUrl") != "/hello/:p1:p2" {
		showError("2")
	}

	if Urls.getParam("secondUrl", 0) != ":param" {
		showError("3")
	}

	if Urls.MustReverse("firstUrl") != "/first" {
		showError("4")
	}

	if Urls.MustReverse("secondUrl", "123", "ABC") != "/second/123/ABC" {
		showError("5")
	}

	if Urls.MustReverse("thirdUrl", "123") != "/comment/123" {
		t.Error(Urls.Reverse("thirdUrl", "123"))
		showError("6")
	}

}

func TestAppend(t *testing.T) {

	root := NewURLStore()
	root.MustAdd("APIIndex", "/")

	admin := NewURLStore()
	admin.MustAdd("AdminIndex", "/")
	admin.MustAdd("AdminPost", "/:post", ":post")
	admin.MustAdd("AdminEditPost", "/:post/edit", ":post")

	err := root.Append("/admin", admin)
	if err != nil {
		panic(err)
	}

	if root.Rev("APIIndex") != "/" {
		t.Fatal("APIIndex")
	}

	if root.Rev("AdminIndex") != "/admin/" {
		t.Fatal("AdminIndex")
	}

	if root.Rev("AdminPost", "test") != "/admin/test" {
		t.Fatal("AdminPost")
	}

	if root.Rev("AdminEditPost", "test") != "/admin/test/edit" {
		t.Fatal("AdminEditPost")
	}

}
