// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kong;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConsumerJwtAuthArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConsumerJwtAuthArgs Empty = new ConsumerJwtAuthArgs();

    /**
     * The algorithm used to verify the token’s signature. Can be HS256, HS384, HS512, RS256, or ES256, Default is `HS256`
     * 
     */
    @Import(name="algorithm")
    private @Nullable Output<String> algorithm;

    /**
     * @return The algorithm used to verify the token’s signature. Can be HS256, HS384, HS512, RS256, or ES256, Default is `HS256`
     * 
     */
    public Optional<Output<String>> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    /**
     * the id of the consumer to be configured with jwt auth
     * 
     */
    @Import(name="consumerId", required=true)
    private Output<String> consumerId;

    /**
     * @return the id of the consumer to be configured with jwt auth
     * 
     */
    public Output<String> consumerId() {
        return this.consumerId;
    }

    /**
     * A unique string identifying the credential. If left out, it will be auto-generated.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return A unique string identifying the credential. If left out, it will be auto-generated.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * If algorithm is `RS256` or `ES256`, the public key (in PEM format) to use to verify the token’s signature
     * 
     */
    @Import(name="rsaPublicKey", required=true)
    private Output<String> rsaPublicKey;

    /**
     * @return If algorithm is `RS256` or `ES256`, the public key (in PEM format) to use to verify the token’s signature
     * 
     */
    public Output<String> rsaPublicKey() {
        return this.rsaPublicKey;
    }

    /**
     * If algorithm is `HS256` or `ES256`, the secret used to sign JWTs for this credential. If left out, will be auto-generated
     * 
     */
    @Import(name="secret")
    private @Nullable Output<String> secret;

    /**
     * @return If algorithm is `HS256` or `ES256`, the secret used to sign JWTs for this credential. If left out, will be auto-generated
     * 
     */
    public Optional<Output<String>> secret() {
        return Optional.ofNullable(this.secret);
    }

    /**
     * A list of strings associated with the consumer JWT auth for grouping and filtering
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A list of strings associated with the consumer JWT auth for grouping and filtering
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private ConsumerJwtAuthArgs() {}

    private ConsumerJwtAuthArgs(ConsumerJwtAuthArgs $) {
        this.algorithm = $.algorithm;
        this.consumerId = $.consumerId;
        this.key = $.key;
        this.rsaPublicKey = $.rsaPublicKey;
        this.secret = $.secret;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConsumerJwtAuthArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConsumerJwtAuthArgs $;

        public Builder() {
            $ = new ConsumerJwtAuthArgs();
        }

        public Builder(ConsumerJwtAuthArgs defaults) {
            $ = new ConsumerJwtAuthArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param algorithm The algorithm used to verify the token’s signature. Can be HS256, HS384, HS512, RS256, or ES256, Default is `HS256`
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm The algorithm used to verify the token’s signature. Can be HS256, HS384, HS512, RS256, or ES256, Default is `HS256`
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param consumerId the id of the consumer to be configured with jwt auth
         * 
         * @return builder
         * 
         */
        public Builder consumerId(Output<String> consumerId) {
            $.consumerId = consumerId;
            return this;
        }

        /**
         * @param consumerId the id of the consumer to be configured with jwt auth
         * 
         * @return builder
         * 
         */
        public Builder consumerId(String consumerId) {
            return consumerId(Output.of(consumerId));
        }

        /**
         * @param key A unique string identifying the credential. If left out, it will be auto-generated.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key A unique string identifying the credential. If left out, it will be auto-generated.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param rsaPublicKey If algorithm is `RS256` or `ES256`, the public key (in PEM format) to use to verify the token’s signature
         * 
         * @return builder
         * 
         */
        public Builder rsaPublicKey(Output<String> rsaPublicKey) {
            $.rsaPublicKey = rsaPublicKey;
            return this;
        }

        /**
         * @param rsaPublicKey If algorithm is `RS256` or `ES256`, the public key (in PEM format) to use to verify the token’s signature
         * 
         * @return builder
         * 
         */
        public Builder rsaPublicKey(String rsaPublicKey) {
            return rsaPublicKey(Output.of(rsaPublicKey));
        }

        /**
         * @param secret If algorithm is `HS256` or `ES256`, the secret used to sign JWTs for this credential. If left out, will be auto-generated
         * 
         * @return builder
         * 
         */
        public Builder secret(@Nullable Output<String> secret) {
            $.secret = secret;
            return this;
        }

        /**
         * @param secret If algorithm is `HS256` or `ES256`, the secret used to sign JWTs for this credential. If left out, will be auto-generated
         * 
         * @return builder
         * 
         */
        public Builder secret(String secret) {
            return secret(Output.of(secret));
        }

        /**
         * @param tags A list of strings associated with the consumer JWT auth for grouping and filtering
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A list of strings associated with the consumer JWT auth for grouping and filtering
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A list of strings associated with the consumer JWT auth for grouping and filtering
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        public ConsumerJwtAuthArgs build() {
            $.consumerId = Objects.requireNonNull($.consumerId, "expected parameter 'consumerId' to be non-null");
            $.rsaPublicKey = Objects.requireNonNull($.rsaPublicKey, "expected parameter 'rsaPublicKey' to be non-null");
            return $;
        }
    }

}