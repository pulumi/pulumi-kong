// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kong
{
    /// <summary>
    /// ## # kong.ConsumerKeyAuth
    /// 
    /// Resource that allows you to configure the [Key Authentication](https://docs.konghq.com/hub/kong-inc/key-auth/) plugin for a consumer.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Kong = Pulumi.Kong;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myConsumer = new Kong.Consumer("myConsumer", new Kong.ConsumerArgs
    ///         {
    ///             Username = "User1",
    ///             CustomId = "123",
    ///         });
    ///         var keyAuthPlugin = new Kong.Plugin("keyAuthPlugin", new Kong.PluginArgs
    ///         {
    ///         });
    ///         var consumerKeyAuth = new Kong.ConsumerKeyAuth("consumerKeyAuth", new Kong.ConsumerKeyAuthArgs
    ///         {
    ///             ConsumerId = myConsumer.Id,
    ///             Key = "secret",
    ///             Tags = 
    ///             {
    ///                 "myTag",
    ///                 "anotherTag",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [KongResourceType("kong:index/consumerKeyAuth:ConsumerKeyAuth")]
    public partial class ConsumerKeyAuth : Pulumi.CustomResource
    {
        /// <summary>
        /// the id of the consumer to associate the credentials to
        /// </summary>
        [Output("consumerId")]
        public Output<string> ConsumerId { get; private set; } = null!;

        /// <summary>
        /// Unique key to authenticate the client; if omitted the plugin will generate one
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// A list of strings associated with the consumer key auth for grouping and filtering
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ConsumerKeyAuth resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsumerKeyAuth(string name, ConsumerKeyAuthArgs args, CustomResourceOptions? options = null)
            : base("kong:index/consumerKeyAuth:ConsumerKeyAuth", name, args ?? new ConsumerKeyAuthArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsumerKeyAuth(string name, Input<string> id, ConsumerKeyAuthState? state = null, CustomResourceOptions? options = null)
            : base("kong:index/consumerKeyAuth:ConsumerKeyAuth", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConsumerKeyAuth resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConsumerKeyAuth Get(string name, Input<string> id, ConsumerKeyAuthState? state = null, CustomResourceOptions? options = null)
        {
            return new ConsumerKeyAuth(name, id, state, options);
        }
    }

    public sealed class ConsumerKeyAuthArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// the id of the consumer to associate the credentials to
        /// </summary>
        [Input("consumerId", required: true)]
        public Input<string> ConsumerId { get; set; } = null!;

        /// <summary>
        /// Unique key to authenticate the client; if omitted the plugin will generate one
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of strings associated with the consumer key auth for grouping and filtering
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public ConsumerKeyAuthArgs()
        {
        }
    }

    public sealed class ConsumerKeyAuthState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// the id of the consumer to associate the credentials to
        /// </summary>
        [Input("consumerId")]
        public Input<string>? ConsumerId { get; set; }

        /// <summary>
        /// Unique key to authenticate the client; if omitted the plugin will generate one
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of strings associated with the consumer key auth for grouping and filtering
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public ConsumerKeyAuthState()
        {
        }
    }
}