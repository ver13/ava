# logger

## Levels
The following table defines the log levels and messages in logger, in decreasing order of severity. The left column lists the log level designation in and the right column provides a brief description of each log level.

|  Level   |  Description                                                                                                                                                                                                    |
| -------- | -----------------------------------------------------------------------------------------------------------------------------------------------------|
| Fatal    | Severe errors that cause premature termination.                                                                                                      |                                                           |
|          | Expect these to be immediately visible on a status console.                                                                                          |
| Error    | Other runtime errors or unexpected conditions.                                                                                                       |
|          | Expect these to be immediately visible on a status console.                                                                                          |
| Critical |                                                                                                                                                      |
| Warn     | Use of deprecated APIs, poor use of APIConfig, 'almost' errors, other runtime situations that are undesirable or unexpected, but not necessarily "wrong".  |
|          | Expect these to be immediately visible on a status console.                                                                                          |
| Info     | Interesting runtime events (startup/shutdown).                                                                                                       |
|          | Expect these to be immediately visible on a console, so be conservative and keep to a minimum.                                                       |
| Debug    | Detailed information on the flow through the system.                                                                                                 |
|          | Expect these to be written to logs only.                                                                                                             |
| Trace    | Most detailed information.                                                                                                                           |
|          | Expect these to be written to logs only.                                                                                                             |

## Colors
The supported keys in the :colors keyword list are:

|  Level   |  Color                |
| -------- | --------------------- |
| Fatal    | Defaults to: Red      |
| Error    | Defaults to: Red      |
| Critical | Defaults to: LightRed |
| Warn     | Defaults to: Yellow   |
| Info     | Defaults to: Green    |
| Debug    | Defaults to: Blue     |
| Trace    | Defaults to: Magenta  |

See the IO.ANSI module for a list of colors and attributes.

For example, info takes precedence over debug. If your log level is set to info, info, warn, and error will be printed to the console. If your log level is set to warn, only warn and error will be printed.

## Types

|  Type                  |  Format                                                                                                                                         |
| ---------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| Text                   | %[levelName]s:%[name]s:%[message]s%[fields]s\n                                                                                                  |
| Apache Common          | {host} {userIdentifier} {authUserId} [{datetime}] "{method} {request} HTTP/1.0" {responseCode} {bytes}                                          |
| Apache Combined        | %[host]s - %[user]s %[authUserId]d [%[ASC_TIME]s] \"%[method]s %[request]s HTTP/1.0\" %[responseCode]d %[bytes]d \"%[referrer]s\" \"%[agent]s\" |
| Apache Error           | [{timestamp}] [{module}:{severity}] [pid {pid}:tid {threadID}] [client: %{client}] %{message}                                                   |
| RFC3164                | <priority>{timestamp} {hostname} {application}[{pid}]: {message}                                                                                |
| RFC5424                | <priority>{version} {iso-timestamp} {hostname} {application} {pid} {message-id} {structured-data} {message}                                     |
| Common Log File Format | %[host]s - %[user]s %[authUserId]d [%[ASC_TIME]s] \"%[method]s %[request]s HTTP/1.0\" %[responseCode]d %[bytes]d                                |

## Metadata
In addition to the keys provided by the user via logger.metadata/1, the following extra keys are available to the :metadata list:

|  Name                  |  Description                   |
| ---------------------- | ------------------------------ |
| host                   |                                |
| application            | The current application        |
| module                 | The current module             |
| function               | The current function           |
| file                   | The current file               |
| line                   | The current line               |
| pid                    | The current process identifier |
| threadID               |                                |
| user                   |                                |
| authUserId             |                                |
| method                 |                                |
| request                |                                |
| responseCode           |                                |
| bytes                  |                                |

