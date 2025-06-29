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
    /// ## # kong.ConsumerAcl
    /// 
    /// Consumer ACL is a resource that allows you to configure the acl plugin for a consumer.
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
    ///     var aclPlugin = new Kong.Plugin("acl_plugin", new()
    ///     {
    ///         Name = "acl",
    ///         ConfigJson = @"	{
    /// 		""allow"": [""group1"", ""group2""]
    /// 	}
    /// ",
    ///     });
    /// 
    ///     var consumerAcl = new Kong.ConsumerAcl("consumer_acl", new()
    ///     {
    ///         ConsumerId = myConsumer.Id,
    ///         Group = "group2",
    ///         Tags = new[]
    ///         {
    ///             "myTag",
    ///             "otherTag",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [KongResourceType("kong:index/consumerAcl:ConsumerAcl")]
    public partial class ConsumerAcl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// the id of the consumer to be configured
        /// </summary>
        [Output("consumerId")]
        public Output<string> ConsumerId { get; private set; } = null!;

        /// <summary>
        /// the acl group
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// A list of strings associated with the consumer acl for grouping and filtering
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ConsumerAcl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsumerAcl(string name, ConsumerAclArgs args, CustomResourceOptions? options = null)
            : base("kong:index/consumerAcl:ConsumerAcl", name, args ?? new ConsumerAclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsumerAcl(string name, Input<string> id, ConsumerAclState? state = null, CustomResourceOptions? options = null)
            : base("kong:index/consumerAcl:ConsumerAcl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConsumerAcl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConsumerAcl Get(string name, Input<string> id, ConsumerAclState? state = null, CustomResourceOptions? options = null)
        {
            return new ConsumerAcl(name, id, state, options);
        }
    }

    public sealed class ConsumerAclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the id of the consumer to be configured
        /// </summary>
        [Input("consumerId", required: true)]
        public Input<string> ConsumerId { get; set; } = null!;

        /// <summary>
        /// the acl group
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of strings associated with the consumer acl for grouping and filtering
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public ConsumerAclArgs()
        {
        }
        public static new ConsumerAclArgs Empty => new ConsumerAclArgs();
    }

    public sealed class ConsumerAclState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the id of the consumer to be configured
        /// </summary>
        [Input("consumerId")]
        public Input<string>? ConsumerId { get; set; }

        /// <summary>
        /// the acl group
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of strings associated with the consumer acl for grouping and filtering
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public ConsumerAclState()
        {
        }
        public static new ConsumerAclState Empty => new ConsumerAclState();
    }
}
