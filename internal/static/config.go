package static

import (
	"os"
	"strconv"
	"strings"

	"github.com/ppenter/ally-sandbox/internal/types"
	"github.com/ppenter/ally-sandbox/internal/utils/log"
	"gopkg.in/yaml.v3"
)

var allySandboxGlobalConfigurations types.DifySandboxGlobalConfigurations

func InitConfig(path string) error {
	allySandboxGlobalConfigurations = types.DifySandboxGlobalConfigurations{}

	// read config file
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer configFile.Close()

	// parse config file
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&allySandboxGlobalConfigurations)
	if err != nil {
		return err
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err == nil {
		allySandboxGlobalConfigurations.App.Debug = debug
	}

	max_workers := os.Getenv("MAX_WORKERS")
	if max_workers != "" {
		allySandboxGlobalConfigurations.MaxWorkers, _ = strconv.Atoi(max_workers)
	}

	max_requests := os.Getenv("MAX_REQUESTS")
	if max_requests != "" {
		allySandboxGlobalConfigurations.MaxRequests, _ = strconv.Atoi(max_requests)
	}

	port := os.Getenv("SANDBOX_PORT")
	if port != "" {
		allySandboxGlobalConfigurations.App.Port, _ = strconv.Atoi(port)
	}

	timeout := os.Getenv("WORKER_TIMEOUT")
	if timeout != "" {
		allySandboxGlobalConfigurations.WorkerTimeout, _ = strconv.Atoi(timeout)
	}

	api_key := os.Getenv("API_KEY")
	if api_key != "" {
		allySandboxGlobalConfigurations.App.Key = api_key
	}

	python_path := os.Getenv("PYTHON_PATH")
	if python_path != "" {
		allySandboxGlobalConfigurations.PythonPath = python_path
	}

	if allySandboxGlobalConfigurations.PythonPath == "" {
		allySandboxGlobalConfigurations.PythonPath = "/usr/local/bin/python3"
	}

	python_lib_path := os.Getenv("PYTHON_LIB_PATH")
	if python_lib_path != "" {
		allySandboxGlobalConfigurations.PythonLibPaths = strings.Split(python_lib_path, ",")
	}

	if len(allySandboxGlobalConfigurations.PythonLibPaths) == 0 {
		allySandboxGlobalConfigurations.PythonLibPaths = DEFAULT_PYTHON_LIB_REQUIREMENTS
	}

	nodejs_path := os.Getenv("NODEJS_PATH")
	if nodejs_path != "" {
		allySandboxGlobalConfigurations.NodejsPath = nodejs_path
	}

	if allySandboxGlobalConfigurations.NodejsPath == "" {
		allySandboxGlobalConfigurations.NodejsPath = "/usr/local/bin/node"
	}

	enable_network := os.Getenv("ENABLE_NETWORK")
	if enable_network != "" {
		allySandboxGlobalConfigurations.EnableNetwork, _ = strconv.ParseBool(enable_network)
	}

	if allySandboxGlobalConfigurations.EnableNetwork {
		log.Info("network has been enabled")
		socks5_proxy := os.Getenv("SOCKS5_PROXY")
		if socks5_proxy != "" {
			allySandboxGlobalConfigurations.Proxy.Socks5 = socks5_proxy
		}

		if allySandboxGlobalConfigurations.Proxy.Socks5 != "" {
			log.Info("using socks5 proxy: %s", allySandboxGlobalConfigurations.Proxy.Socks5)
		}

		https_proxy := os.Getenv("HTTPS_PROXY")
		if https_proxy != "" {
			allySandboxGlobalConfigurations.Proxy.Https = https_proxy
		}

		if allySandboxGlobalConfigurations.Proxy.Https != "" {
			log.Info("using https proxy: %s", allySandboxGlobalConfigurations.Proxy.Https)
		}

		http_proxy := os.Getenv("HTTP_PROXY")
		if http_proxy != "" {
			allySandboxGlobalConfigurations.Proxy.Http = http_proxy
		}

		if allySandboxGlobalConfigurations.Proxy.Http != "" {
			log.Info("using http proxy: %s", allySandboxGlobalConfigurations.Proxy.Http)
		}
	}
	return nil
}

// avoid global modification, use value copy instead
func GetDifySandboxGlobalConfigurations() types.DifySandboxGlobalConfigurations {
	return allySandboxGlobalConfigurations
}

type RunnerDependencies struct {
	PythonRequirements string
}

var runnerDependencies RunnerDependencies

func GetRunnerDependencies() RunnerDependencies {
	return runnerDependencies
}

func SetupRunnerDependencies() error {
	file, err := os.ReadFile("dependencies/python-requirements.txt")
	if err != nil {
		if err == os.ErrNotExist {
			return nil
		}
		return err
	}

	runnerDependencies.PythonRequirements = string(file)

	return nil
}
