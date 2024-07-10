package integrationtests_test

import (
	"strings"
	"testing"

	"github.com/ppenter/ally-sandbox/internal/core/runner/types"
	"github.com/ppenter/ally-sandbox/internal/service"
)

func TestNodejsBase64(t *testing.T) {
	// Test case for base64
	runMultipleTestings(t, 30, func(t *testing.T) {
		resp := service.RunNodeJsCode(`
const base64 = Buffer.from("hello world").toString("base64");
console.log(Buffer.from(base64, "base64").toString());
		`, "", &types.RunnerOptions{
			EnableNetwork: true,
		})
		if resp.Code != 0 {
			t.Fatal(resp)
		}

		if !strings.Contains(resp.Data.(*service.RunCodeResponse).Stdout, "hello world") {
			t.Fatalf("unexpected output: %s\n", resp.Data.(*service.RunCodeResponse).Stdout)
		}

		if resp.Data.(*service.RunCodeResponse).Stderr != "" {
			t.Fatalf("unexpected error: %s\n", resp.Data.(*service.RunCodeResponse).Stderr)
		}
	})
}

func TestNodejsJSON(t *testing.T) {
	// Test case for json
	runMultipleTestings(t, 30, func(t *testing.T) {
		resp := service.RunNodeJsCode(`
console.log(JSON.stringify({"hello": "world"}));
		`, "", &types.RunnerOptions{
			EnableNetwork: true,
		})
		if resp.Code != 0 {
			t.Error(resp)
		}

		if !strings.Contains(resp.Data.(*service.RunCodeResponse).Stdout, `{"hello":"world"}`) {
			t.Fatalf("unexpected output: %s\n", resp.Data.(*service.RunCodeResponse).Stdout)
		}

		if resp.Data.(*service.RunCodeResponse).Stderr != "" {
			t.Fatalf("unexpected error: %s\n", resp.Data.(*service.RunCodeResponse).Stderr)
		}
	})
}
