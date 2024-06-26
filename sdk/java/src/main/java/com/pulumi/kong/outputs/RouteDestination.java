// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kong.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RouteDestination {
    private @Nullable String ip;
    private @Nullable Integer port;

    private RouteDestination() {}
    public Optional<String> ip() {
        return Optional.ofNullable(this.ip);
    }
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RouteDestination defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String ip;
        private @Nullable Integer port;
        public Builder() {}
        public Builder(RouteDestination defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ip = defaults.ip;
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder ip(@Nullable String ip) {

            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {

            this.port = port;
            return this;
        }
        public RouteDestination build() {
            final var _resultValue = new RouteDestination();
            _resultValue.ip = ip;
            _resultValue.port = port;
            return _resultValue;
        }
    }
}
