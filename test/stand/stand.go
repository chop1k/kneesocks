package stand

import (
	"github.com/go-playground/validator/v10"
	"github.com/sarulabs/di"
	"github.com/stretchr/testify/require"
	"os"
	"socks/pkg/protocol/auth/password"
	v4Protocol "socks/pkg/protocol/v4"
	v4aProtocol "socks/pkg/protocol/v4a"
	v5Protocol "socks/pkg/protocol/v5"
	"socks/pkg/utils"
	"socks/test/stand/config"
	"socks/test/stand/picture"
	"socks/test/stand/server"
	"socks/test/stand/v4"
	"socks/test/stand/v4a"
	"socks/test/stand/v5"
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

	scopeDef := di.Def{
		Name:  "scope",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)

			return config.NewScope(t, cfg)
		},
	}

	err := builder.Add(
		configPathDef,
		validatorDef,
		configDef,
		scopeDef,
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
			parser := ctn.Get("v5_parser").(v5Protocol.Parser)

			return picture.NewPicture(cfg, _case, t, parser)
		},
	}

	err := builder.Add(
		pictureDef,
	)

	require.NoError(t, err)

	s.registerServer(builder, t)
}

func (s Stand) registerServer(builder di.Builder, t *testing.T) {
	serverDef := di.Def{
		Name:  "server",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			pic := ctn.Get("picture").(picture.Picture)

			return server.NewServer(t, cfg, pic)
		},
	}

	err := builder.Add(
		serverDef,
	)

	require.NoError(t, err)

	s.registerV4(builder, t)
}

func (s Stand) registerV4(builder di.Builder, t *testing.T) {
	builderDef := di.Def{
		Name:  "v4_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4Protocol.NewBuilder(), nil
		},
	}

	senderDef := di.Def{
		Name:  "v4_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			builder := ctn.Get("v4_builder").(v4Protocol.Builder)

			return v4.NewSender(t, cfg, builder)
		},
	}

	comparatorDef := di.Def{
		Name:  "v4_comparator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			builder := ctn.Get("v4_builder").(v4Protocol.Builder)

			return v4.NewComparator(t, cfg, builder)
		},
	}

	bindTesterDef := di.Def{
		Name:  "v4_bind_tester",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			pic := ctn.Get("picture").(picture.Picture)
			sender := ctn.Get("v4_sender").(v4.Sender)
			comparator := ctn.Get("v4_comparator").(v4.Comparator)
			scope := ctn.Get("scope").(config.Scope)
			srv := ctn.Get("server").(server.Server)

			return v4.NewBindTester(cfg, t, pic, sender, comparator, scope, srv)
		},
	}

	connectTesterDef := di.Def{
		Name:  "v4_connect_tester",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			srv := ctn.Get("server").(server.Server)
			sender := ctn.Get("v4_sender").(v4.Sender)
			comparator := ctn.Get("v4_comparator").(v4.Comparator)
			scope := ctn.Get("scope").(config.Scope)

			return v4.NewConnectTester(cfg, t, srv, sender, comparator, scope)
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
		senderDef,
		comparatorDef,
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
			return v4aProtocol.NewBuilder(), nil
		},
	}

	bindTesterDef := di.Def{
		Name:  "v4a_bind_tester",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			pic := ctn.Get("picture").(picture.Picture)
			srv := ctn.Get("server").(server.Server)
			sender := ctn.Get("v4a_sender").(v4a.Sender)
			comparator := ctn.Get("v4a_comparator").(v4a.Comparator)
			scope := ctn.Get("scope").(config.Scope)

			return v4a.NewBindTester(cfg, t, pic, srv, sender, comparator, scope)
		},
	}

	senderDef := di.Def{
		Name:  "v4a_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			builder := ctn.Get("v4a_builder").(v4aProtocol.Builder)

			return v4a.NewSender(t, cfg, builder)
		},
	}

	comparatorDef := di.Def{
		Name:  "v4a_comparator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			builder := ctn.Get("v4a_builder").(v4aProtocol.Builder)

			return v4a.NewComparator(t, cfg, builder)
		},
	}

	connectTesterDef := di.Def{
		Name:  "v4a_connect_tester",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			srv := ctn.Get("server").(server.Server)
			sender := ctn.Get("v4a_sender").(v4a.Sender)
			comparator := ctn.Get("v4a_comparator").(v4a.Comparator)
			scope := ctn.Get("scope").(config.Scope)

			return v4a.NewConnectTester(cfg, t, srv, sender, comparator, scope)
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
		senderDef,
		comparatorDef,
		bindTesterDef,
		connectTesterDef,
		testDef,
	)

	require.NoError(t, err)

	s.registerV5(builder, t)
}

func (s Stand) registerV5(builder di.Builder, t *testing.T) {
	addressUtilsDef := di.Def{
		Name:  "address_utils",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return utils.NewUtils()
		},
	}

	passwordBuilderDef := di.Def{
		Name:  "v5_password_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return password.NewBuilder()
		},
	}

	builderDef := di.Def{
		Name:  "v5_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v5Protocol.NewBuilder()
		},
	}

	parserDef := di.Def{
		Name:  "v5_parser",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)

			return v5Protocol.NewParser(addressUtils), nil
		},
	}

	senderDef := di.Def{
		Name:  "v5_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			builder := ctn.Get("v5_builder").(v5Protocol.Builder)
			passwordBuilder := ctn.Get("v5_password_builder").(password.Builder)

			return v5.NewSender(t, cfg, builder, passwordBuilder)
		},
	}

	comparatorDef := di.Def{
		Name:  "v5_comparator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			builder := ctn.Get("v5_builder").(v5Protocol.Builder)
			passwordBuilder := ctn.Get("v5_password_builder").(password.Builder)

			return v5.NewComparator(t, cfg, builder, passwordBuilder)
		},
	}

	bindTesterDef := di.Def{
		Name:  "v5_bind_tester",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			srv := ctn.Get("server").(server.Server)
			sender := ctn.Get("v5_sender").(v5.Sender)
			compare := ctn.Get("v5_comparator").(v5.Comparator)
			scope := ctn.Get("scope").(config.Scope)
			pic := ctn.Get("picture").(picture.Picture)

			return v5.NewBindTester(t, cfg, sender, compare, srv, scope, pic)
		},
	}

	associationTesterDef := di.Def{
		Name:  "v5_association_tester",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			srv := ctn.Get("server").(server.Server)
			sender := ctn.Get("v5_sender").(v5.Sender)
			compare := ctn.Get("v5_comparator").(v5.Comparator)
			scope := ctn.Get("scope").(config.Scope)
			pic := ctn.Get("picture").(picture.Picture)

			return v5.NewAssociationTester(t, cfg, sender, compare, srv, scope, pic)
		},
	}

	authTesterDef := di.Def{
		Name:  "v5_auth_tester",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			srv := ctn.Get("server").(server.Server)
			sender := ctn.Get("v5_sender").(v5.Sender)
			compare := ctn.Get("v5_comparator").(v5.Comparator)
			scope := ctn.Get("scope").(config.Scope)

			return v5.NewAuthTester(t, cfg, srv, sender, compare, scope)
		},
	}

	connectTesterDef := di.Def{
		Name:  "v5_connect_tester",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			cfg := ctn.Get("config").(config.Config)
			srv := ctn.Get("server").(server.Server)
			sender := ctn.Get("v5_sender").(v5.Sender)
			compare := ctn.Get("v5_comparator").(v5.Comparator)
			scope := ctn.Get("scope").(config.Scope)

			return v5.NewConnectTester(t, cfg, sender, compare, srv, scope)
		},
	}

	testDef := di.Def{
		Name:  "v5_test",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			t := ctn.Get("t").(*testing.T)
			_case := ctn.Get("case").(config.Case)
			auth := ctn.Get("v5_auth_tester").(v5.AuthTester)
			connect := ctn.Get("v5_connect_tester").(v5.ConnectTester)
			bind := ctn.Get("v5_bind_tester").(v5.BindTester)
			associate := ctn.Get("v5_association_tester").(v5.AssociationTester)

			return v5.NewTest(_case, t, auth, connect, bind, associate)
		},
	}

	err := builder.Add(
		addressUtilsDef,
		passwordBuilderDef,
		parserDef,
		builderDef,
		senderDef,
		comparatorDef,
		bindTesterDef,
		associationTesterDef,
		authTesterDef,
		connectTesterDef,
		testDef,
	)

	require.NoError(t, err)

	s.start(builder.Build())
}

func (s Stand) start(ctn di.Container) {
	ctn.Get("test").(Test).Start()
}
