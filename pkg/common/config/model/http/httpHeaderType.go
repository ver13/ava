//go:generate go-enum -f=$GOFILE --marshal --lower

package http

// HTTP headers let the client and the server pass additional information with an HTTP request or response.
/*
ENUM(
All

// Authentication
WWW-Authenticate //Defines the authService method that should be used to access a resource.
Authorization // Contains the credentials to authenticate a user-agent with a server.
Proxy-Authenticate // Defines the authService method that should be used to access a resource behind a proxy server.
Proxy-Authorization //Contains the credentials to authenticate a user agent with a proxy server.

// Caching
Age // The time, in seconds, that the object has been in a proxy cache.
Cache-Control // Directives for caching mechanisms in both requests and responses.
Clear-Site-Data // Clears browsing data (e.g. cookies, storage, cache) associated with the requesting website.
Expires // The date/time after which the response is considered stale.
Pragma // Implementation-specific header that may have various effects anywhere along the request-response chain. Used for backwards compatibility with HTTP/1.0 caches where the Cache-Control header is not yet present.
Warning // General warning information about possible problems.

// ClientI hints. HTTP ClientI hints are a work in progress. Actual documentation can be found on the website of the HTTP working group.
Accept-CH // Servers can advertise support for ClientI Hints using the Accept-CH header field or an equivalent HTML <meta> element with http-equiv attribute ([HTML5]).
Accept-CH-Lifetime // Servers can ask the client to remember the set of ClientI Hints that the server supports for a specified period of time, to enable delivery of ClientI Hints on subsequent requests to the server’s origin ([RFC6454]).
Early-Data // Indicates that the request has been conveyed in early data.
Content-DPR // A number that indicates the ratio between physical pixels over CSS pixels of the selected image response.
DPR // A number that indicates the client’s current Device Pixel Ratio (DPR), which is the ratio of physical pixels over CSS pixels (Section 5.2 of [CSSVAL]) of the layout viewport (Section 9.1.1 of [CSS2]) on the device.
Device-Memory // Technically a part of Device Memory api, this header represents an approximate amount of RAM client has.
Save-Data // A boolean that indicates the user agent's preference for reduced data usage.
Viewport-Width // A number that indicates the layout viewport width in CSS pixels. The provided pixel value is a number rounded to the smallest following integer (i.e. ceiling value).
Width // The Width request header field is a number that indicates the desired resource width in physical pixels (i.e. intrinsic size of an image). The provided pixel value is a number rounded to the smallest following integer (i.e. ceiling value).

// Conditionals
Last-Modified // The last modification date of the resource, used to compare several versions of the same resource. It is less accurate than ETag, but easier to calculate in some environments. Conditional requests using If-Modified-Since and If-Unmodified-Since use this value to change the behavior of the request.
ETag // A unique string identifying the version of the resource. Conditional requests using If-Match and If-None-Match use this value to change the behavior of the request.
If-Match // Makes the request conditional, and applies the method only if the stored resource matches one of the given ETags.
If-None-Match // Makes the request conditional, and applies the method only if the stored resource doesn't match any of the given ETags. This is used to update caches (for safe requests), or to prevent to upload a new resource when one already exists.
If-Modified-Since // Makes the request conditional, and expects the entity to be transmitted only if it has been modified after the given date. This is used to transmit data only when the cache is out of date.
If-Unmodified-Since // Makes the request conditional, and expects the entity to be transmitted only if it has not been modified after the given date. This ensures the coherence of a new fragment of a specific range with previous ones, or to implement an optimistic concurrency control system when modifying existing documents.
Vary // Determines how to match request headers to decide whether a cached response can be used rather than requesting a fresh one from the origin server.

// Connection management
Connection // Controls whether the network connection stays open after the current transaction finishes.
Keep-Alive // Controls how long a persistent connection should stay open.

// Content negotiation
Accept // Informs the server about the types of data that can be sent back.
Accept-Charset // Which character encodings the client understands.
Accept-encoding // The error algorithm, usually a compression algorithm, that can be used on the resource sent back.
Accept-Language // Informs the server about the human language the server is expected to send back. This is a hint and is not necessarily under the full control of the user: the server should always pay attention not to override an explicit user choice (like selecting a language from a dropdown).

// Controls
Expect // Indicates expectations that need to be fulfilled by the server to properly handle the request.
Max-Forwards

// Cookies
Cookie // Contains stored HTTP cookies previously sent by the server with the Set-Cookie header.
Set-Cookie // Send cookies from the server to the user-agent.
Cookie2 // Contains an HTTP cookie previously sent by the server with the Set-Cookie2 header, but has been obsoleted. Use Cookie instead.
Set-Cookie2 // Sends cookies from the server to the user-agent, but has been obsoleted. Use Set-Cookie instead.

// cors. Learn more about cors here.
Access-Control-Allow-Origin // Indicates whether the response can be shared.
Access-Control-Allow-Credentials // Indicates whether the response to the request can be exposed when the credentials flag is true.
Access-Control-Allow-headers // Used in response to a preflight request to indicate which HTTP headers can be used when making the actual request.
Access-Control-Allow-Methods // Specifies the methods allowed when accessing the resource in response to a preflight request.
Access-Control-Expose-headers // Indicates which headers can be exposed as part of the response by listing their names.
Access-Control-Max-Age // Indicates how long the results of a preflight request can be cached.
Access-Control-Request-headers // Used when issuing a preflight request to let the server know which HTTP headers will be used when the actual request is made.
Access-Control-Request-method // Used when issuing a preflight request to let the server know which HTTP method will be used when the actual request is made.
Origin // Indicates where a fetch originates from.
ServiceI-Worker-Allowed // Used to remove the path restriction by including this header in the response of the ServiceI Worker script.
Timing-Allow-Origin // Specifies origins that are allowed to see values of attributes retrieved via features of the Resource Timing api, which would otherwise be reported as zero due to cross-origin restrictions.
X-Permitted-Cross-Domain-Policies // Specifies if a cross-domain policy file (crossdomain.xml) is allowed. The file may define a policy to grant clients, such as Adobe's Flash Player, Adobe Acrobat, Microsoft Silverlight, or Apache Flex, permission to handle data across domains that would otherwise be restricted due to the Same-Origin Policy. See the Cross-domain Policy file Specification for more information.

// Do Not Track
DNT // Expresses the user's tracking preference.
Tk // Indicates the tracking status of the corresponding response.

// Downloads
Content-Disposition // Indicates if the resource transmitted should be displayed inline (default behavior without the header), or if it should be handled like a download and the browser should present a “Save As” dialog.

// message body information
Content-Length // The size of the resource, in decimal number of bytes.
Content-Type // Indicates the media type of the resource.
Content-encoding // Used to specify the compression algorithm.
Content-Language // Describes the human language(s) intended for the audience, so that it allows a user to differentiate according to the users' own preferred language.
Content-Location // Indicates an alternate location for the returned data.

// Proxies
Forwarded // Contains information from the client-facing side of proxy servers that is altered or lost when a proxy is involved in the path of the request.
X-Forwarded-For // Identifies the originating IP addresses of a client connecting to a web server through an HTTP proxy or a load balancer.
X-Forwarded-host // Identifies the original host requested that a client used to connect to your proxy or load balancer.
X-Forwarded-Proto // Identifies the protocol (HTTP or HTTPS) that a client used to connect to your proxy or load balancer.
Via // Added by proxies, both forward and reverse proxies, and can appear in the request headers and the response headers.

// Redirects
Location // Indicates the url to redirect a page to.

// Request context
From // Contains an Internet email address for a human user who controls the requesting user agent.
host // Specifies the domain name of the server (for virtual hosting), and (optionally) the TCP port number on which the server is listening.
Referer // The address of the previous web page from which a link to the currently requested page was followed.
Referrer-Policy // Governs which referrer information sent in the Referer header should be included with requests made.
User-Agent // Contains a characteristic string that allows the network protocol peers to identify the application type, operating system, software vendor or software version of the requesting software user agent. See also the Firefox user agent string reference.

// Response context
Allow // Lists the set of HTTP request methods support by a resource.
Server // Contains information about the software used by the origin server to handle the request.

// Range requests
Accept-Ranges // Indicates if the server supports range requests, and if so in which unit the range can be expressed.
Range // Indicates the part of a document that the server should return.
If-Range // Creates a conditional range request that is only fulfilled if the given etag or date matches the remote resource. Used to prevent downloading two ranges from incompatible version of the resource.
Content-Range // Indicates where in a full body message a partial message belongs.

// Security
Cross-Origin-Opener-Policy // Prevents other domains from opening/controlling a window.
Cross-Origin-Resource-Policy // Prevents other domains from reading the response of the resources to which this header is applied.
Content-Security-Policy // Controls resources the user agent is allowed to load for a given page.
Content-Security-Policy-Report-Only // Allows web developers to experiment with policies by monitoring, but not enforcing, their effects. These violation reports consist of JSON documents sent via an HTTP POST request to the specified URI.
Expect-CT // Allows sites to opt in to reporting and/or enforcement of Certificate Transparency requirements, which prevents the use of misissued certificates for that site from going unnoticed. When a site enables the Expect-CT header, they are requesting that Chrome check that any certificate for that site appears in public CT logs.
Feature-Policy // Provides a mechanism to allow and deny the use of browser features in its own frame, and in iframes that it embeds.
Public-Key-Pins // Associates a specific cryptographic public key with a certain web server to decrease the risk of MITM attacks with forged certificates.
Public-Key-Pins-Report-Only // Sends reports to the report-uri specified in the header and does still allow clients to connect to the server even if the pinning is violated.
Strict-Transport-Security // Force communication using HTTPS instead of HTTP.
Upgrade-Insecure-Requests // Sends a signal to the server expressing the client’s preference for an encrypted and authenticated response, and that it can successfully handle the upgrade-insecure-requests directive.
X-Content-Type-Options // Disables MIME sniffing and forces browser to use the type given in Content-Type.
X-Download-Options // Indicates that the browser (Internet Explorer) should not display the option to "Open" a file that has been downloaded from an application, to prevent phishing attacks as the file otherwise would gain access to execute in the context of the application.
X-Frame-Options // Indicates whether a browser should be allowed to render a page in a <frame>, <iframe>, <embed> or <object>.
X-Powered-By // May be set by hosting environments or other frameworks and contains information about them while not providing any usefulness to the application or its visitors. Unset this header to avoid exposing potential vulnerabilities.
X-XSS-Protection // Enables cross-site scripting filtering.

// Server-sent events
Last-Event-ID
NEL // Defines a mechanism that enables developers to declare a network error reporting policy.
Ping-From
Ping-To
Report-To // Used to specify a server endpoint for the browser to send warning and error reports to.

// Transfer coding
Transfer-encoding // Specifies the form of error used to safely transfer the entity to the user.
TE // Specifies the transfer encodings the user agent is willing to accept.
Trailer // Allows the sender to include additional fields at the end of chunked message.

// WebSockets
Sec-WebSocket-Key
Sec-WebSocket-Extensions
Sec-WebSocket-Accept
Sec-WebSocket-Protocol
Sec-WebSocket-version

// Other
Accept-Push-Policy // A client can express the desired push policy for a request by sending an Accept-Push-Policy header field in the request.
Accept-Signature // A client can send the Accept-Signature header field to indicate intention to take advantage of any available signatures and to indicate what kinds of signatures it supports.
Alt-Svc // Used to list alternate ways to reach this stackMicroservices.
Date // Contains the date and time at which the message was originated.
Large-Allocation // Tells the browser that the page being loaded is going to want to perform a large allocation.
Link // The Link entity-header field provides a means for serialising one or more links in HTTP headers. It is semantically equivalent to the HTML <link> element.
Push-Policy // A Push-Policy defines the server behaviour regarding push when processing a request.
Retry-After // Indicates how long the user agent should wait before making a follow-up request.
Signature // The Signature header field conveys a list of signatures for an exchange, each one accompanied by information about how to determine the authority of and refresh that signature.
Signed-headers // The Signed-headers header field identifies an ordered list of response header fields to include in a signature.
Server-Timing // Communicates one or more metricsService and descriptions for the given request-response cycle.
SourceMap // Links generated code to a source map.
Upgrade // The relevant RFC document for the Upgrade header field is RFC 7230, section 6.7. The standard establishes rules for upgrading or changing to a different protocol on the current client, server, transport protocol connection. For example, this header standard allows a client to change from HTTP 1.1 to HTTP 2.0, assuming the server decides to acknowledge and implement the Upgrade header field. Neither party is required to accept the terms specified in the Upgrade header field. It can be used in both client and server headers. If the Upgrade header field is specified, then the sender MUST also send the Connection header field with the upgrade option specified. For details on the Connection header field please see section 6.1 of the aforementioned RFC.
X-DNS-Prefetch-Control // Controls DNS prefetching, a feature by which browsers proactively perform domain name resolution on both links that the user may choose to follow as well as URLs for items referenced by the document, including images, CSS, JavaScript, and so forth.
X-Firefox-Spdy
X-Pingback
X-Requested-With
X-Robots-Tag // Used to indicate how a web page is to be indexed within public search engine results. The header is effectively equivalent to <meta name="robots" content="...">.
X-UA-Compatible // Used by Internet Explorer to signal which document mode to use.

X-CSRF-Token // Used to prevent cross-site request forgery. Alternative header names are: X-CSRFToken and X-XSRF-TOKEN
X-UIDH // Server-side deep packet insertion of a unique ID identifying customers of Verizon Wireless; also known as "perma-cookie" or "supercookie"
X-Request-ID // Correlates HTTP requests between a client and server.
X-Correlation-ID // Correlates HTTP requests between a client and server.
Proxy-Connection // Implemented as a misunderstanding of the HTTP specifications. Common because of mistakes in implementations of early HTTP versions. Has exactly the same functionality as standard Connection field.
X-Wap-Profile // Links to an XML file on the Internet with a full description and details about the device currently connecting. In the example to the right is an XML file for an AT&T Samsung Galaxy S2.
X-ATT-DeviceId // Allows easier parsing of the MakeModel/Firmware that is usually found in the User-Agent String of AT&T Devices
X-Http-method-Override // Requests a web application to override the method specified in the request (typically POST) with the method given in the header field (typically PUT or DELETE). This can be used when a user agent or firewall prevents PUT or DELETE methods from being sent directly (note that this is either a bug in the software component, which ought to be fixed, or an intentional configurationServiceI, in which case bypassing it may be the wrong thing to do).
Front-End-Https // Non-standard header field used by Microsoft applications and load-balancers

Unknown
)
*/
type HTTPHeaderType int32
