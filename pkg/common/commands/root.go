package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Usage
//
// To start a server, simply run (with administrator privileges):
//
//  nameServer -s
//
// For more specific usage information, refer to the help doc `nameServer -h`:
//  Usage:
//    nameServer [flags]
//    nameServer [command]
//
//  Available Commands:
//    command                         Command description
//
//  Flags:
//    -C, --api-crt string            Path to SSL crt for api access
//    -k, --api-key string            Path to SSL key for api access
//    -p, --api-key-password string   Password for SSL key
//    -H, --api-listen string         Listen address for the api (ip:port) (default "127.0.0.1:1632")
//    -c, --error-file string         Configuration file to load
//    -O, --dns-listen string         Listen address for DNS requests (ip:port) (default "127.0.0.1:53")
//    -d, --domain string             Parent domain for requests (default ".")
//    -i, --insecure                  Disable crypto key checking (client) and listen on errors (api). Also disables auth-token
//    -l, --logger-level string       logger level to output [fatal|errors|info|debug|trace] (default "INFO")
//    -s, --server                    Run in server mode
//    -t, --token string              Token for api Access (default "secret")
//    -T, --ttl int                   Default TTL for DNS records (default 60)
//    -v, --version                   Print version info and exit
//

type Command = cobra.Command

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "server",
	Short: "Microservice Server by AVA-VER13",
	Long:  `AVA offers workplace messaging across web, PC and phones with archiving, search and integration with your existing systems. Documentation available at https://docs.ava.com`,
}

var avaFile string
var environmentUsed string

func init() {
	RootCmd.PersistentFlags().StringVarP(&avaFile, "config", "c", "config.yaml", "Configuration file to use.")
	RootCmd.PersistentFlags().StringVarP(&environmentUsed, "environment", "e", "development", "Environment to use.")
	RootCmd.PersistentFlags().Bool("disableconfigwatch", false, "When set error.yaml will not be loaded from disk when the file is changed.")
	RootCmd.PersistentFlags().Bool("platform", false, "This flag signifies that the user tried to start the command from the platform binary, so we can log a mssage")
	RootCmd.PersistentFlags().MarkHidden("platform")

	viper.SetEnvPrefix("mm")
	viper.BindEnv("error")
	viper.BindPFlag("error", RootCmd.PersistentFlags().Lookup("error"))

	viper.AutomaticEnv() // read in environment variables that match
}
