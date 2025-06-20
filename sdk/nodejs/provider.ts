// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the kong package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'kong';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * An basic auth password for kong admin
     */
    public readonly kongAdminPassword!: pulumi.Output<string | undefined>;
    /**
     * API key for the kong api (Enterprise Edition)
     */
    public readonly kongAdminToken!: pulumi.Output<string | undefined>;
    /**
     * The address of the kong admin url e.g. http://localhost:8001
     */
    public readonly kongAdminUri!: pulumi.Output<string | undefined>;
    /**
     * An basic auth user for kong admin
     */
    public readonly kongAdminUsername!: pulumi.Output<string | undefined>;
    /**
     * API key for the kong api (if you have locked it down)
     */
    public readonly kongApiKey!: pulumi.Output<string | undefined>;
    /**
     * Workspace context (Enterprise Edition)
     */
    public readonly kongWorkspace!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["kongAdminPassword"] = args ? args.kongAdminPassword : undefined;
            resourceInputs["kongAdminToken"] = args ? args.kongAdminToken : undefined;
            resourceInputs["kongAdminUri"] = args ? args.kongAdminUri : undefined;
            resourceInputs["kongAdminUsername"] = args ? args.kongAdminUsername : undefined;
            resourceInputs["kongApiKey"] = args ? args.kongApiKey : undefined;
            resourceInputs["kongWorkspace"] = args ? args.kongWorkspace : undefined;
            resourceInputs["strictPluginsMatch"] = pulumi.output((args ? args.strictPluginsMatch : undefined) ?? utilities.getEnvBoolean("STRICT_PLUGINS_MATCH")).apply(JSON.stringify);
            resourceInputs["tlsSkipVerify"] = pulumi.output((args ? args.tlsSkipVerify : undefined) ?? (utilities.getEnvBoolean("TLS_SKIP_VERIFY") || false)).apply(JSON.stringify);
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }

    /**
     * This function returns a Terraform config object with terraform-namecased keys,to be used with the Terraform Module Provider.
     */
    terraformConfig(): pulumi.Output<Provider.TerraformConfigResult> {
        return pulumi.runtime.call("pulumi:providers:kong/terraformConfig", {
            "__self__": this,
        }, this);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * An basic auth password for kong admin
     */
    kongAdminPassword?: pulumi.Input<string>;
    /**
     * API key for the kong api (Enterprise Edition)
     */
    kongAdminToken?: pulumi.Input<string>;
    /**
     * The address of the kong admin url e.g. http://localhost:8001
     */
    kongAdminUri?: pulumi.Input<string>;
    /**
     * An basic auth user for kong admin
     */
    kongAdminUsername?: pulumi.Input<string>;
    /**
     * API key for the kong api (if you have locked it down)
     */
    kongApiKey?: pulumi.Input<string>;
    /**
     * Workspace context (Enterprise Edition)
     */
    kongWorkspace?: pulumi.Input<string>;
    /**
     * Should plugins `configJson` field strictly match plugin configuration
     */
    strictPluginsMatch?: pulumi.Input<boolean>;
    /**
     * Whether to skip tls verify for https kong api endpoint using self signed or untrusted certs
     */
    tlsSkipVerify?: pulumi.Input<boolean>;
}

export namespace Provider {
    /**
     * The results of the Provider.terraformConfig method.
     */
    export interface TerraformConfigResult {
        readonly result: {[key: string]: any};
    }

}
