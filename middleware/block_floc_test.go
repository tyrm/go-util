package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

const TestHeaderPermissionsPolicy = "Permissions-Policy"
const TestExpectedPermissionsPolicy = "interest-cohort=()"

func TestBlockFlocGin(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	BlockFlocGin(c)
	permissionsPolicy := c.Writer.Header().Get(TestHeaderPermissionsPolicy)
	if permissionsPolicy != TestExpectedPermissionsPolicy {
		t.Errorf("%s header was incorrect, got: %v, want: %v.", TestHeaderPermissionsPolicy, permissionsPolicy, TestExpectedPermissionsPolicy)
	}
}

func TestBlockFlocMux(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		permissionsPolicy := w.Header().Get(TestHeaderPermissionsPolicy)
		if permissionsPolicy != TestExpectedPermissionsPolicy {
			t.Errorf("%s header was incorrect, got: %v, want: %v.", TestHeaderPermissionsPolicy, permissionsPolicy, TestExpectedPermissionsPolicy)
		}
	})

	req := httptest.NewRequest("GET", "http://localhost", nil)
	blockFloc := BlockFlocMux(testHandler)
	blockFloc.ServeHTTP(httptest.NewRecorder(), req)
}
