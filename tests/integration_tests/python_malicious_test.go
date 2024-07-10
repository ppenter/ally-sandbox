package integrationtests_test

import (
	"strings"
	"testing"

	"github.com/ppenter/ally-sandbox/internal/core/runner/types"
	"github.com/ppenter/ally-sandbox/internal/service"
)

func TestSysFork(t *testing.T) {
	// Test case for sys_fork
	resp := service.RunPython3Code(`
import os
os.fork()
	`, "", &types.RunnerOptions{})

	if resp.Code != 0 {
		t.Error(resp)
	}

	if !strings.Contains(resp.Data.(*service.RunCodeResponse).Stderr, "operation not permitted") {
		t.Error(resp.Data.(*service.RunCodeResponse).Stderr)
	}
}

func TestExec(t *testing.T) {
	// Test case for exec
	resp := service.RunPython3Code(`
import os
os.execl("/bin/ls", "ls")
	`, "", &types.RunnerOptions{})
	if resp.Code != 0 {
		t.Error(resp)
	}

	if !strings.Contains(resp.Data.(*service.RunCodeResponse).Stderr, "operation not permitted") {
		t.Error(resp.Data.(*service.RunCodeResponse).Stderr)
	}
}

func TestRunCommand(t *testing.T) {
	// Test case for run_command
	resp := service.RunPython3Code(`
import subprocess
subprocess.run(["ls", "-l"])
	`, "", &types.RunnerOptions{})
	if resp.Code != 0 {
		t.Error(resp)
	}

	if !strings.Contains(resp.Data.(*service.RunCodeResponse).Stderr, "operation not permitted") {
		t.Error(resp.Data.(*service.RunCodeResponse).Stderr)
	}
}
