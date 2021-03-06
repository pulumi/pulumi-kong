// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kong
{
    [KongResourceType("kong:index/upstream:Upstream")]
    public partial class Upstream : Pulumi.CustomResource
    {
        [Output("hashFallback")]
        public Output<string?> HashFallback { get; private set; } = null!;

        [Output("hashFallbackHeader")]
        public Output<string?> HashFallbackHeader { get; private set; } = null!;

        [Output("hashOn")]
        public Output<string?> HashOn { get; private set; } = null!;

        [Output("hashOnCookie")]
        public Output<string?> HashOnCookie { get; private set; } = null!;

        [Output("hashOnCookiePath")]
        public Output<string?> HashOnCookiePath { get; private set; } = null!;

        [Output("hashOnHeader")]
        public Output<string?> HashOnHeader { get; private set; } = null!;

        [Output("healthchecks")]
        public Output<Outputs.UpstreamHealthchecks> Healthchecks { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("slots")]
        public Output<int?> Slots { get; private set; } = null!;


        /// <summary>
        /// Create a Upstream resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Upstream(string name, UpstreamArgs? args = null, CustomResourceOptions? options = null)
            : base("kong:index/upstream:Upstream", name, args ?? new UpstreamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Upstream(string name, Input<string> id, UpstreamState? state = null, CustomResourceOptions? options = null)
            : base("kong:index/upstream:Upstream", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Upstream resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Upstream Get(string name, Input<string> id, UpstreamState? state = null, CustomResourceOptions? options = null)
        {
            return new Upstream(name, id, state, options);
        }
    }

    public sealed class UpstreamArgs : Pulumi.ResourceArgs
    {
        [Input("hashFallback")]
        public Input<string>? HashFallback { get; set; }

        [Input("hashFallbackHeader")]
        public Input<string>? HashFallbackHeader { get; set; }

        [Input("hashOn")]
        public Input<string>? HashOn { get; set; }

        [Input("hashOnCookie")]
        public Input<string>? HashOnCookie { get; set; }

        [Input("hashOnCookiePath")]
        public Input<string>? HashOnCookiePath { get; set; }

        [Input("hashOnHeader")]
        public Input<string>? HashOnHeader { get; set; }

        [Input("healthchecks")]
        public Input<Inputs.UpstreamHealthchecksArgs>? Healthchecks { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("slots")]
        public Input<int>? Slots { get; set; }

        public UpstreamArgs()
        {
        }
    }

    public sealed class UpstreamState : Pulumi.ResourceArgs
    {
        [Input("hashFallback")]
        public Input<string>? HashFallback { get; set; }

        [Input("hashFallbackHeader")]
        public Input<string>? HashFallbackHeader { get; set; }

        [Input("hashOn")]
        public Input<string>? HashOn { get; set; }

        [Input("hashOnCookie")]
        public Input<string>? HashOnCookie { get; set; }

        [Input("hashOnCookiePath")]
        public Input<string>? HashOnCookiePath { get; set; }

        [Input("hashOnHeader")]
        public Input<string>? HashOnHeader { get; set; }

        [Input("healthchecks")]
        public Input<Inputs.UpstreamHealthchecksGetArgs>? Healthchecks { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("slots")]
        public Input<int>? Slots { get; set; }

        public UpstreamState()
        {
        }
    }
}
