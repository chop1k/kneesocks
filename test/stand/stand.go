package stand

import (
	"github.com/go-playground/validator/v10"
	"github.com/sarulabs/di"
	"github.com/stretchr/testify/require"
	"os"
	"socks/cmd/e2e_test_server/protocol"
	v42 "socks/protocol/v4"
	v4a2 "socks/protocol/v4a"
	"socks/test/stand/config"
	"socks/test/stand/picture"
	v4 "socks/test/stand/v4"
	"socks/test/stand/v4a"
	v5 "socks/test/stand/v5"
	"testing"
)

type Stand struct {
}

func New() Stand {
	return Stand{}
}

func (s Stand) Execute(protocol string, command string, number int, t *testing.T) {
	builder, err := di.NewBuilder()

	require.NoError(t, err)

	s.register(protocol, command, number, *builder, t)
}

func (s Stand) register(protocol string, command string, number int, builder di.Builder, t *testing.T) {
	testingDef := di.Def{
		Name:  "t",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return t, nil
		},
	}

	testDef := di.Def{
		Name:  "test",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			_case := ctn.Get("case").(config.Case)
			v4Test := ctn.Get("v4_test").(v4.Test)
			v4aTest := ctn.Get("v4a_test").(v4a.Test)
			v5Test := ctn.Get("v5_test").(v5.Test)

			return NewTest(t, _case, v4Test, v4aTest, v5Test)
		},
	}

	err := builder.Add(
		testingDef,
		testDef,
	)

	require.NoError(t, err)

	s.registerConfig(protocol, command, number, builder, t)
}

func (s Stand) registerConfig(protocol string, command string, number int, builder di.Builder, t *testing.T) {
	configPathDef := di.Def{
		Name:  "config_path",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)

			path, ok := os.LookupEnv("config_path")

			require.True(t, ok, "Config path is not specified. ")

			return path, nil
		},
	}

	validatorDef := di.Def{
		Name:  "validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return *validator.New(), nil
		},
	}

	configDef := di.Def{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			validate := ctn.Get("validator").(validator.Validate)
			path := ctn.Get("config_path").(string)

			return config.NewConfig(validate, path, t), nil
		},
	}

	caseDef := di.Def{
		Name:  "case",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return config.NewCase(protocol, command, number), nil
		},
	}

	err := builder.Add(
		configPathDef,
		validatorDef,
		configDef,
		caseDef,
	)

	require.NoError(t, err)

	s.registerPicture(builder, t)
}

func (s Stand) registerPicture(builder di.Builder, t *testing.T) {
	pictureDef := di.Def{
		Name:  "picture",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			_case := ctn.Get("case").(config.Case)
			cfg := ctn.Get("config").(config.Config)

			return picture.NewPicture(cfg, _case, t)
		},
	}

	err := builder.Add(
		pictureDef,
	)

	require.NoError(t, err)

	s.registerServer(builder, t)
}

func (s Stand) registerServer(builder di.Builder, t *testing.T) {
	builderDef := di.Def{
		Name:  "server_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return protocol.NewBuilder()
		},
	}

	err := builder.Add(
		builderDef,
	)

	require.NoError(t, err)

	s.registerV4(builder, t)
}

func (s Stand) registerV4(builder di.Builder, t *testing.T) {
	builderDef := di.Def{
		Name:  "v4_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v42.NewBaseBuilder(), nil
		},
	}

	bindTesterDef := di.Def{
		Name:  "v4_bind_tester",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			builder := ctn.Get("v4_builder").(v42.Builder)

			pic := ctn.Get("picture").(picture.Picture)

			return v4.NewBindTester(cfg, t, builder, pic)
		},
	}

	connectTesterDef := di.Def{
		Name:  "v4_connect_tester",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			builder := ctn.Get("v4_builder").(v42.Builder)
			serverBuilder := ctn.Get("server_builder").(protocol.Builder)
			pic := ctn.Get("picture").(picture.Picture)

			return v4.NewConnectTester(cfg, t, builder, serverBuilder, pic)
		},
	}

	testDef := di.Def{
		Name:  "v4_test",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			_case := ctn.Get("case").(config.Case)
			connectTester := ctn.Get("v4_connect_tester").(v4.ConnectTester)
			bindTester := ctn.Get("v4_bind_tester").(v4.BindTester)

			return v4.NewTest(_case, t, connectTester, bindTester)
		},
	}

	err := builder.Add(
		builderDef,
		bindTesterDef,
		connectTesterDef,
		testDef,
	)

	require.NoError(t, err)

	s.registerV4a(builder, t)
}

func (s Stand) registerV4a(builder di.Builder, t *testing.T) {
	builderDef := di.Def{
		Name:  "v4a_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4a2.NewBaseBuilder(), nil
		},
	}

	bindTesterDef := di.Def{
		Name:  "v4a_bind_tester",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			builder := ctn.Get("v4a_builder").(v4a2.Builder)
			pic := ctn.Get("picture").(picture.Picture)

			return v4a.NewBindTester(cfg, t, builder, pic)
		},
	}

	connectTesterDef := di.Def{
		Name:  "v4a_connect_tester",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			builder := ctn.Get("v4a_builder").(v4a2.Builder)
			serverBuilder := ctn.Get("server_builder").(protocol.Builder)
			pic := ctn.Get("picture").(picture.Picture)

			return v4a.NewConnectTester(cfg, t, builder, serverBuilder, pic)
		},
	}

	testDef := di.Def{
		Name:  "v4a_test",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			_case := ctn.Get("case").(config.Case)
			connectTester := ctn.Get("v4a_connect_tester").(v4a.ConnectTester)
			bindTester := ctn.Get("v4a_bind_tester").(v4a.BindTester)

			return v4a.NewTest(_case, t, connectTester, bindTester)
		},
	}

	err := builder.Add(
		builderDef,
		bindTesterDef,
		connectTesterDef,
		testDef,
	)

	require.NoError(t, err)

	s.registerV5(builder, t)
}

func (s Stand) registerV5(builder di.Builder, t *testing.T) {
	testDef := di.Def{
		Name:  "v5_test",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			_case := ctn.Get("case").(config.Case)

			return v5.NewTest(_case, t)
		},
	}

	err := builder.Add(
		testDef,
	)

	require.NoError(t, err)
	s.start(builder.Build())
}

func (s Stand) start(ctn di.Container) {
	ctn.Get("test").(Test).Start()
}
