//go:generate avaEnum -f=$GOFILE --marshal --lower

package serializer

// SerializerType x ENUM(
// json
// xml
// yaml
// toml
// hcl
// proto
// unknown
// )
type SerializerType int32
