//go:generate go-enum -f=$GOFILE --marshal --lower

package model

// ConfigFileType x ENUM(
// Json
// Toml
// Yaml
// Yml
// Xml
// Properties
// Props
// Prop
// Env
// Dotenv
// Unknown
// )
type ConfigFileType int32
