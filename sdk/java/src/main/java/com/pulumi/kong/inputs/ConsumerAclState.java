// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kong.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConsumerAclState extends com.pulumi.resources.ResourceArgs {

    public static final ConsumerAclState Empty = new ConsumerAclState();

    /**
     * the id of the consumer to be configured
     * 
     */
    @Import(name="consumerId")
    private @Nullable Output<String> consumerId;

    /**
     * @return the id of the consumer to be configured
     * 
     */
    public Optional<Output<String>> consumerId() {
        return Optional.ofNullable(this.consumerId);
    }

    /**
     * the acl group
     * 
     */
    @Import(name="group")
    private @Nullable Output<String> group;

    /**
     * @return the acl group
     * 
     */
    public Optional<Output<String>> group() {
        return Optional.ofNullable(this.group);
    }

    /**
     * A list of strings associated with the consumer acl for grouping and filtering
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A list of strings associated with the consumer acl for grouping and filtering
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private ConsumerAclState() {}

    private ConsumerAclState(ConsumerAclState $) {
        this.consumerId = $.consumerId;
        this.group = $.group;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConsumerAclState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConsumerAclState $;

        public Builder() {
            $ = new ConsumerAclState();
        }

        public Builder(ConsumerAclState defaults) {
            $ = new ConsumerAclState(Objects.requireNonNull(defaults));
        }

        /**
         * @param consumerId the id of the consumer to be configured
         * 
         * @return builder
         * 
         */
        public Builder consumerId(@Nullable Output<String> consumerId) {
            $.consumerId = consumerId;
            return this;
        }

        /**
         * @param consumerId the id of the consumer to be configured
         * 
         * @return builder
         * 
         */
        public Builder consumerId(String consumerId) {
            return consumerId(Output.of(consumerId));
        }

        /**
         * @param group the acl group
         * 
         * @return builder
         * 
         */
        public Builder group(@Nullable Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group the acl group
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param tags A list of strings associated with the consumer acl for grouping and filtering
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A list of strings associated with the consumer acl for grouping and filtering
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A list of strings associated with the consumer acl for grouping and filtering
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        public ConsumerAclState build() {
            return $;
        }
    }

}