package main

import (
	"flag"
	"io/ioutil"
	"os"
	"gyg/middleware/app"
	"gyg/middleware/config"
	"gyg/middleware/services"
)

func main() {

	serviceLocator := services.NewLocator()
	logger := serviceLocator.Logger()
	logger.Info([]string{`
        GGGGGGGGGGGGGYYYYYYY       YYYYYYY        GGGGGGGGGGGGG
     GGG::::::::::::GY:::::Y       Y:::::Y     GGG::::::::::::G
   GG:::::::::::::::GY:::::Y       Y:::::Y   GG:::::::::::::::G
  G:::::GGGGGGGG::::GY::::::Y     Y::::::Y  G:::::GGGGGGGG::::G
 G:::::G       GGGGGGYYY:::::Y   Y:::::YYY G:::::G       GGGGGG
G:::::G                 Y:::::Y Y:::::Y   G:::::G
G:::::G                  Y:::::Y:::::Y    G:::::G
G:::::G    GGGGGGGGGG     Y:::::::::Y     G:::::G    GGGGGGGGGG
G:::::G    G::::::::G      Y:::::::Y      G:::::G    G::::::::G
G:::::G    GGGGG::::G       Y:::::Y       G:::::G    GGGGG::::G
G:::::G        G::::G       Y:::::Y       G:::::G        G::::G
 G:::::G       G::::G       Y:::::Y        G:::::G       G::::G
  G:::::GGGGGGGG::::G       Y:::::Y         G:::::GGGGGGGG::::G
   GG:::::::::::::::G    YYYY:::::YYYY       GG:::::::::::::::G
     GGG::::::GGG:::G    Y:::::::::::Y         GGG::::::GGG:::G
        GGGGGG   GGGG    YYYYYYYYYYYYY            GGGGGG   GGGG
`}[0])
	logger.Warn([]string{`
████████╗██╗  ██╗███████╗    ███████╗████████╗ ██████╗ ██████╗ ███╗   ███╗    ██╗███████╗     ██████╗ ██████╗ ███╗   ███╗██╗███╗   ██╗ ██████╗
╚══██╔══╝██║  ██║██╔════╝    ██╔════╝╚══██╔══╝██╔═══██╗██╔══██╗████╗ ████║    ██║██╔════╝    ██╔════╝██╔═══██╗████╗ ████║██║████╗  ██║██╔════╝
   ██║   ███████║█████╗      ███████╗   ██║   ██║   ██║██████╔╝██╔████╔██║    ██║███████╗    ██║     ██║   ██║██╔████╔██║██║██╔██╗ ██║██║  ███╗
   ██║   ██╔══██║██╔══╝      ╚════██║   ██║   ██║   ██║██╔══██╗██║╚██╔╝██║    ██║╚════██║    ██║     ██║   ██║██║╚██╔╝██║██║██║╚██╗██║██║   ██║
   ██║   ██║  ██║███████╗    ███████║   ██║   ╚██████╔╝██║  ██║██║ ╚═╝ ██║    ██║███████║    ╚██████╗╚██████╔╝██║ ╚═╝ ██║██║██║ ╚████║╚██████╔╝██╗██╗██╗
   ╚═╝   ╚═╝  ╚═╝╚══════╝    ╚══════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝    ╚═╝╚══════╝     ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝╚═╝╚═╝
	`}[0])

	configPath := flag.String("config.path", "config/main.json", "Base config.")
	flag.Parse()

	data, err := ioutil.ReadFile(*configPath)
	if nil != err {
		logger.Error(err.Error())
		os.Exit(0)
	}

	config.Configuration, err = serviceLocator.LoadConfig(data)
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(0)
	}

	server := app.NewServer(app.NewDispatcher(logger).Run())
	server.Start()

	logger.Success([]string{`
 _     _                         _                _                     _           _           _
| |   | |           _           | |              (_)     _             | |         | |         | |
| |__ | | ____  ___| |_  ____   | | ____    _   _ _  ___| |_  ____     | | _   ____| | _  _   _| |
|  __)| |/ _  |/___)  _)/ _  |  | |/ _  |  | | | | |/___)  _)/ _  |    | || \ / _  | || \| | | |_|
| |   | ( ( | |___ | |_( ( | |  | ( ( | |   \ V /| |___ | |_( ( | |_   | |_) | ( | | |_) ) |_| |_
|_|   |_|\_||_(___/ \___)_||_|  |_|\_||_|    \_/ |_(___/ \___)_||_( )  |____/ \_||_|____/ \__  |_|
                                                                  |/                     (____/
	`}[0])
}
