// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kong.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class UpstreamHealthchecksPassiveHealthy {
    private @Nullable List<Integer> httpStatuses;
    private @Nullable Integer successes;

    private UpstreamHealthchecksPassiveHealthy() {}
    public List<Integer> httpStatuses() {
        return this.httpStatuses == null ? List.of() : this.httpStatuses;
    }
    public Optional<Integer> successes() {
        return Optional.ofNullable(this.successes);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(UpstreamHealthchecksPassiveHealthy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<Integer> httpStatuses;
        private @Nullable Integer successes;
        public Builder() {}
        public Builder(UpstreamHealthchecksPassiveHealthy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.httpStatuses = defaults.httpStatuses;
    	      this.successes = defaults.successes;
        }

        @CustomType.Setter
        public Builder httpStatuses(@Nullable List<Integer> httpStatuses) {

            this.httpStatuses = httpStatuses;
            return this;
        }
        public Builder httpStatuses(Integer... httpStatuses) {
            return httpStatuses(List.of(httpStatuses));
        }
        @CustomType.Setter
        public Builder successes(@Nullable Integer successes) {

            this.successes = successes;
            return this;
        }
        public UpstreamHealthchecksPassiveHealthy build() {
            final var _resultValue = new UpstreamHealthchecksPassiveHealthy();
            _resultValue.httpStatuses = httpStatuses;
            _resultValue.successes = successes;
            return _resultValue;
        }
    }
}
