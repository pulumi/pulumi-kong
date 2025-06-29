// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kong
{
    /// <summary>
    /// ## # kong.ConsumerBasicAuth
    /// 
    /// Consumer basic auth is a resource that allows you to configure the basic auth plugin for a consumer.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Kong = Pulumi.Kong;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myConsumer = new Kong.Consumer("my_consumer", new()
    ///     {
    ///         Username = "User1",
    ///         CustomId = "123",
    ///     });
    /// 
    ///     var basicAuthPlugin = new Kong.Plugin("basic_auth_plugin", new()
    ///     {
    ///         Name = "basic-auth",
    ///     });
    /// 
    ///     var consumerBasicAuth = new Kong.ConsumerBasicAuth("consumer_basic_auth", new()
    ///     {
    ///         ConsumerId = myConsumer.Id,
    ///         Username = "foo_updated",
    ///         Password = "bar_updated",
    ///         Tags = new[]
    ///         {
    ///             "myTag",
    ///             "anotherTag",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [KongResourceType("kong:index/consumerBasicAuth:ConsumerBasicAuth")]
    public partial class ConsumerBasicAuth : global::Pulumi.CustomResource
    {
        /// <summary>
        /// the id of the consumer to be configured with basic auth
        /// </summary>
        [Output("consumerId")]
        public Output<string> ConsumerId { get; private set; } = null!;

        /// <summary>
        /// password to be used for basic auth
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// A list of strings associated with the consumer basic auth for grouping and filtering
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// username to be used for basic auth
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a ConsumerBasicAuth resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsumerBasicAuth(string name, ConsumerBasicAuthArgs args, CustomResourceOptions? options = null)
            : base("kong:index/consumerBasicAuth:ConsumerBasicAuth", name, args ?? new ConsumerBasicAuthArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsumerBasicAuth(string name, Input<string> id, ConsumerBasicAuthState? state = null, CustomResourceOptions? options = null)
            : base("kong:index/consumerBasicAuth:ConsumerBasicAuth", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConsumerBasicAuth resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConsumerBasicAuth Get(string name, Input<string> id, ConsumerBasicAuthState? state = null, CustomResourceOptions? options = null)
        {
            return new ConsumerBasicAuth(name, id, state, options);
        }
    }

    public sealed class ConsumerBasicAuthArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the id of the consumer to be configured with basic auth
        /// </summary>
        [Input("consumerId", required: true)]
        public Input<string> ConsumerId { get; set; } = null!;

        /// <summary>
        /// password to be used for basic auth
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of strings associated with the consumer basic auth for grouping and filtering
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// username to be used for basic auth
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public ConsumerBasicAuthArgs()
        {
        }
        public static new ConsumerBasicAuthArgs Empty => new ConsumerBasicAuthArgs();
    }

    public sealed class ConsumerBasicAuthState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the id of the consumer to be configured with basic auth
        /// </summary>
        [Input("consumerId")]
        public Input<string>? ConsumerId { get; set; }

        /// <summary>
        /// password to be used for basic auth
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of strings associated with the consumer basic auth for grouping and filtering
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// username to be used for basic auth
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ConsumerBasicAuthState()
        {
        }
        public static new ConsumerBasicAuthState Empty => new ConsumerBasicAuthState();
    }
}
