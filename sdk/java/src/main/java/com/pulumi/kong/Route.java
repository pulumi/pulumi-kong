// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kong;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kong.RouteArgs;
import com.pulumi.kong.Utilities;
import com.pulumi.kong.inputs.RouteState;
import com.pulumi.kong.outputs.RouteDestination;
import com.pulumi.kong.outputs.RouteHeader;
import com.pulumi.kong.outputs.RouteSource;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # kong.Route
 * 
 * The route resource maps directly onto the json for the route endpoint in Kong. For more information on the parameters [see the Kong Route create documentation](https://docs.konghq.com/gateway-oss/2.5.x/admin-api/#route-object).
 * 
 * To create a tcp/tls route you set `sources` and `destinations` by repeating the corresponding element (`source` or `destination`) for each source or destination you want.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.kong.Route;
 * import com.pulumi.kong.RouteArgs;
 * import com.pulumi.kong.inputs.RouteHeaderArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var route = new Route("route", RouteArgs.builder()
 *             .name("MyRoute")
 *             .protocols(            
 *                 "http",
 *                 "https")
 *             .methods(            
 *                 "GET",
 *                 "POST")
 *             .hosts("example2.com")
 *             .paths("/test")
 *             .stripPath(false)
 *             .preserveHost(true)
 *             .regexPriority(1)
 *             .serviceId(service.id())
 *             .headers(RouteHeaderArgs.builder()
 *                 .name("x-test-1")
 *                 .values(                
 *                     "a",
 *                     "b")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * To create a tcp/tls route you set `sources` and `destinations` by repeating the corresponding element (`source` or `destination`) for each source or destination you want, for example:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.kong.Route;
 * import com.pulumi.kong.RouteArgs;
 * import com.pulumi.kong.inputs.RouteSourceArgs;
 * import com.pulumi.kong.inputs.RouteDestinationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var route = new Route("route", RouteArgs.builder()
 *             .protocols("tcp")
 *             .stripPath(true)
 *             .preserveHost(false)
 *             .sources(            
 *                 RouteSourceArgs.builder()
 *                     .ip("192.168.1.1")
 *                     .port(80)
 *                     .build(),
 *                 RouteSourceArgs.builder()
 *                     .ip("192.168.1.2")
 *                     .build())
 *             .destinations(RouteDestinationArgs.builder()
 *                 .ip("172.10.1.1")
 *                 .port(81)
 *                 .build())
 *             .snis("foo.com")
 *             .serviceId(service.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * To import a route:
 * 
 * ```sh
 * $ pulumi import kong:index/route:Route &lt;route_identifier&gt; &lt;route_id&gt;
 * ```
 * 
 */
@ResourceType(type="kong:index/route:Route")
public class Route extends com.pulumi.resources.CustomResource {
    /**
     * A list of destination `ip` and `port`
     * 
     */
    @Export(name="destinations", refs={List.class,RouteDestination.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RouteDestination>> destinations;

    /**
     * @return A list of destination `ip` and `port`
     * 
     */
    public Output<Optional<List<RouteDestination>>> destinations() {
        return Codegen.optional(this.destinations);
    }
    /**
     * One or more blocks of `name` to set name of header and `values` which is a list of `string` for the header values to match on.  See above example of how to set.  These headers will cause this Route to match if present in the request. The Host header cannot be used with this attribute: hosts should be specified using the hosts attribute.
     * 
     */
    @Export(name="headers", refs={List.class,RouteHeader.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RouteHeader>> headers;

    /**
     * @return One or more blocks of `name` to set name of header and `values` which is a list of `string` for the header values to match on.  See above example of how to set.  These headers will cause this Route to match if present in the request. The Host header cannot be used with this attribute: hosts should be specified using the hosts attribute.
     * 
     */
    public Output<Optional<List<RouteHeader>>> headers() {
        return Codegen.optional(this.headers);
    }
    /**
     * A list of domain names that match this Route
     * 
     */
    @Export(name="hosts", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> hosts;

    /**
     * @return A list of domain names that match this Route
     * 
     */
    public Output<Optional<List<String>>> hosts() {
        return Codegen.optional(this.hosts);
    }
    /**
     * The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is HTTP instead of HTTPS. Location header is injected by Kong if the field is set to `301`, `302`, `307` or `308`. Accepted values are: `426`, `301`, `302`, `307`, `308`. Default: `426`.
     * 
     */
    @Export(name="httpsRedirectStatusCode", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> httpsRedirectStatusCode;

    /**
     * @return The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is HTTP instead of HTTPS. Location header is injected by Kong if the field is set to `301`, `302`, `307` or `308`. Accepted values are: `426`, `301`, `302`, `307`, `308`. Default: `426`.
     * 
     */
    public Output<Optional<Integer>> httpsRedirectStatusCode() {
        return Codegen.optional(this.httpsRedirectStatusCode);
    }
    /**
     * A list of HTTP methods that match this Route
     * 
     */
    @Export(name="methods", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> methods;

    /**
     * @return A list of HTTP methods that match this Route
     * 
     */
    public Output<Optional<List<String>>> methods() {
        return Codegen.optional(this.methods);
    }
    /**
     * The name of the route
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the route
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Controls how the Service path, Route path and requested path are combined when sending a request to the upstream.
     * 
     */
    @Export(name="pathHandling", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pathHandling;

    /**
     * @return Controls how the Service path, Route path and requested path are combined when sending a request to the upstream.
     * 
     */
    public Output<Optional<String>> pathHandling() {
        return Codegen.optional(this.pathHandling);
    }
    /**
     * A list of paths that match this Route
     * 
     */
    @Export(name="paths", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> paths;

    /**
     * @return A list of paths that match this Route
     * 
     */
    public Output<Optional<List<String>>> paths() {
        return Codegen.optional(this.paths);
    }
    /**
     * When matching a Route via one of the hosts domain names, use the request Host header in the upstream request headers. If set to false, the upstream Host header will be that of the Service’s host.
     * 
     */
    @Export(name="preserveHost", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> preserveHost;

    /**
     * @return When matching a Route via one of the hosts domain names, use the request Host header in the upstream request headers. If set to false, the upstream Host header will be that of the Service’s host.
     * 
     */
    public Output<Optional<Boolean>> preserveHost() {
        return Codegen.optional(this.preserveHost);
    }
    /**
     * The list of protocols to use
     * 
     */
    @Export(name="protocols", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> protocols;

    /**
     * @return The list of protocols to use
     * 
     */
    public Output<List<String>> protocols() {
        return this.protocols;
    }
    /**
     * A number used to choose which route resolves a given request when several routes match it using regexes simultaneously.
     * 
     */
    @Export(name="regexPriority", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> regexPriority;

    /**
     * @return A number used to choose which route resolves a given request when several routes match it using regexes simultaneously.
     * 
     */
    public Output<Optional<Integer>> regexPriority() {
        return Codegen.optional(this.regexPriority);
    }
    /**
     * Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding. Default: true.
     * 
     */
    @Export(name="requestBuffering", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> requestBuffering;

    /**
     * @return Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding. Default: true.
     * 
     */
    public Output<Optional<Boolean>> requestBuffering() {
        return Codegen.optional(this.requestBuffering);
    }
    /**
     * Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding. Default: true.
     * 
     */
    @Export(name="responseBuffering", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> responseBuffering;

    /**
     * @return Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding. Default: true.
     * 
     */
    public Output<Optional<Boolean>> responseBuffering() {
        return Codegen.optional(this.responseBuffering);
    }
    /**
     * Service ID to map to
     * 
     */
    @Export(name="serviceId", refs={String.class}, tree="[0]")
    private Output<String> serviceId;

    /**
     * @return Service ID to map to
     * 
     */
    public Output<String> serviceId() {
        return this.serviceId;
    }
    /**
     * A list of SNIs that match this Route when using stream routing.
     * 
     */
    @Export(name="snis", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> snis;

    /**
     * @return A list of SNIs that match this Route when using stream routing.
     * 
     */
    public Output<Optional<List<String>>> snis() {
        return Codegen.optional(this.snis);
    }
    /**
     * A list of source `ip` and `port`
     * 
     */
    @Export(name="sources", refs={List.class,RouteSource.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RouteSource>> sources;

    /**
     * @return A list of source `ip` and `port`
     * 
     */
    public Output<Optional<List<RouteSource>>> sources() {
        return Codegen.optional(this.sources);
    }
    /**
     * When matching a Route via one of the paths, strip the matching prefix from the upstream request URL. Default: true.
     * 
     */
    @Export(name="stripPath", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stripPath;

    /**
     * @return When matching a Route via one of the paths, strip the matching prefix from the upstream request URL. Default: true.
     * 
     */
    public Output<Optional<Boolean>> stripPath() {
        return Codegen.optional(this.stripPath);
    }
    /**
     * A list of strings associated with the Route for grouping and filtering.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A list of strings associated with the Route for grouping and filtering.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Route(java.lang.String name) {
        this(name, RouteArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Route(java.lang.String name, RouteArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Route(java.lang.String name, RouteArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kong:index/route:Route", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Route(java.lang.String name, Output<java.lang.String> id, @Nullable RouteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kong:index/route:Route", name, state, makeResourceOptions(options, id), false);
    }

    private static RouteArgs makeArgs(RouteArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RouteArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Route get(java.lang.String name, Output<java.lang.String> id, @Nullable RouteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Route(name, id, state, options);
    }
}
