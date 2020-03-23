//go:generate avaEnum -f=$GOFILE --marshal --lower

package logger

// LogLevelType x ENUM(
// Panic    // Panic level, highest level of severity. Logs and then calls panic with the message passed to debug, info, ...
// Fatal    // Fatal level. Logs and then calls `logger.Exit(1)`. It will exit even if the logging level is set to Panic.
// Error    // Error level. Logs. Used for errors that should definitely be noted.
// Warn     // Warn level. Non-critical entries that deserve eyes.
// Info     // Info level. General operational entries about what's going on inside the application.
// Debug    // Debug level. Usually only enabled when debugging. Very verbose logging.
// Trace    // Trace level. Designates finer-grained informational events than the debug.
// Unknown
// )
type LogLevelType uint32
