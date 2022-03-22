package build

import (
	"errors"
	"os"
	"socks/internal/kneesocks/config/tree"

	"github.com/go-playground/validator/v10"
	"github.com/sarulabs/di"
)

func TcpConfigPath(ctn di.Container) (interface{}, error) {
	path, ok := os.LookupEnv("tcp_config_path")

	if !ok {
		return "/etc/kneesocks/proxy/tcp.json", nil
	}

	return path, nil
}

func UdpConfigPath(ctn di.Container) (interface{}, error) {
	path, ok := os.LookupEnv("udp_config_path")

	if !ok {
		return "/etc/kneesocks/proxy/udp.json", nil
	}

	return path, nil
}

func HttpConfigPath(ctn di.Container) (interface{}, error) {
	path, ok := os.LookupEnv("http_config_path")

	if !ok {
		return "/etc/kneesocks/proxy/http.json", nil
	}

	return path, nil
}

func V4ConfigPath(ctn di.Container) (interface{}, error) {
	path, ok := os.LookupEnv("v4_config_path")

	if !ok {
		return "/etc/kneesocks/proxy/v4.json", nil
	}

	return path, nil
}

func V4aConfigPath(ctn di.Container) (interface{}, error) {
	path, ok := os.LookupEnv("v4a_config_path")

	if !ok {
		return "/etc/kneesocks/proxy/v4a.json", nil
	}

	return path, nil
}

func V5ConfigPath(ctn di.Container) (interface{}, error) {
	path, ok := os.LookupEnv("v5_config_path")

	if !ok {
		return "/etc/kneesocks/proxy/v5.json", nil
	}

	return path, nil
}

func LogConfigPath(ctn di.Container) (interface{}, error) {
	path, ok := os.LookupEnv("log_config_path")

	if !ok {
		return "/etc/kneesocks/proxy/log.json", nil
	}

	return path, nil
}

func TcpConfigFile(ctn di.Container) (interface{}, error) {
	path := ctn.Get("tcp_config_path").(string)

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func UdpConfigFile(ctn di.Container) (interface{}, error) {
	path := ctn.Get("udp_config_path").(string)

	file, err := os.Open(path)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return file, nil
}

func HttpConfigFile(ctn di.Container) (interface{}, error) {
	path := ctn.Get("http_config_path").(string)

	file, err := os.Open(path)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return file, nil
}

func V4ConfigFile(ctn di.Container) (interface{}, error) {
	path := ctn.Get("v4_config_path").(string)

	file, err := os.Open(path)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return file, nil
}

func V4aConfigFile(ctn di.Container) (interface{}, error) {
	path := ctn.Get("v4a_config_path").(string)

	file, err := os.Open(path)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return file, nil
}

func V5ConfigFile(ctn di.Container) (interface{}, error) {
	path := ctn.Get("v5_config_path").(string)

	file, err := os.Open(path)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return file, nil
}

func LogConfigFile(ctn di.Container) (interface{}, error) {
	path := ctn.Get("log_config_path").(string)

	file, err := os.Open(path)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return file, nil
}

func Validator(ctn di.Container) (interface{}, error) {
	return *validator.New(), nil
}

func TcpTree(ctn di.Container) (interface{}, error) {
	file := ctn.Get("tcp_config_file").(*os.File)
	builder := ctn.Get("tcp_tree_builder").(tree.TcpBuilder)

	defer file.Close()

	return builder.Build(file)
}

func UdpTree(ctn di.Container) (interface{}, error) {
	_file := ctn.Get("udp_config_file")
	builder := ctn.Get("udp_tree_builder").(tree.UdpBuilder)

	if _file == nil {
		return nil, nil
	}

	file := _file.(*os.File)

	defer file.Close()

	return builder.Build(file)
}

func HttpTree(ctn di.Container) (interface{}, error) {
	return nil, nil
}

func V4Tree(ctn di.Container) (interface{}, error) {
	_file := ctn.Get("v4_config_file")
	builder := ctn.Get("v4_tree_builder").(tree.SocksV4Builder)

	if _file == nil {
		return nil, nil
	}

	file := _file.(*os.File)

	defer file.Close()

	return builder.Build(file)
}

func V4aTree(ctn di.Container) (interface{}, error) {
	_file := ctn.Get("v4a_config_file")
	builder := ctn.Get("v4a_tree_builder").(tree.SocksV4aBuilder)

	if _file == nil {
		return nil, nil
	}

	file := _file.(*os.File)

	defer file.Close()

	return builder.Build(file)
}

func V5Tree(ctn di.Container) (interface{}, error) {
	_file := ctn.Get("v5_config_file")
	builder := ctn.Get("v5_tree_builder").(tree.SocksV5Builder)

	if _file == nil {
		return nil, nil
	}

	file := _file.(*os.File)

	defer file.Close()

	return builder.Build(file)
}

func LogTree(ctn di.Container) (interface{}, error) {
	_file := ctn.Get("log_config_file")
	builder := ctn.Get("log_tree_builder").(tree.LogBuilder)

	if _file == nil {
		return nil, nil
	}

	file := _file.(*os.File)

	defer file.Close()

	return builder.Build(file)
}

func TcpTreeBuilder(ctn di.Container) (interface{}, error) {
	validator := ctn.Get("validator").(validator.Validate)

	return tree.NewTcpBuilder(validator)
}

func UdpTreeBuilder(ctn di.Container) (interface{}, error) {
	validator := ctn.Get("validator").(validator.Validate)

	return tree.NewUdpBuilder(validator)
}

func HttpTreeBuilder(ctn di.Container) (interface{}, error) {
	return nil, nil
}

func V4TreeBuilder(ctn di.Container) (interface{}, error) {
	validator := ctn.Get("validator").(validator.Validate)

	return tree.NewSocksV4Builder(validator)
}

func V4aTreeBuilder(ctn di.Container) (interface{}, error) {
	validator := ctn.Get("validator").(validator.Validate)

	return tree.NewSocksV4aBuilder(validator)
}

func V5TreeBuilder(ctn di.Container) (interface{}, error) {
	validator := ctn.Get("validator").(validator.Validate)

	return tree.NewSocksV5Builder(validator)
}

func LogTreeBuilder(ctn di.Container) (interface{}, error) {
	validator := ctn.Get("validator").(validator.Validate)

	return tree.NewLogBuilder(validator)
}
