// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface RouteDestination {
    ip?: string;
    port?: number;
}

export interface RouteHeader {
    /**
     * The name of the route
     */
    name: string;
    values: string[];
}

export interface RouteSource {
    ip?: string;
    port?: number;
}

export interface UpstreamHealthchecks {
    active: outputs.UpstreamHealthchecksActive;
    passive: outputs.UpstreamHealthchecksPassive;
}

export interface UpstreamHealthchecksActive {
    concurrency?: number;
    healthy: outputs.UpstreamHealthchecksActiveHealthy;
    httpPath?: string;
    httpsSni?: string;
    httpsVerifyCertificate?: boolean;
    timeout?: number;
    type?: string;
    unhealthy: outputs.UpstreamHealthchecksActiveUnhealthy;
}

export interface UpstreamHealthchecksActiveHealthy {
    httpStatuses: number[];
    interval: number;
    successes: number;
}

export interface UpstreamHealthchecksActiveUnhealthy {
    httpFailures: number;
    httpStatuses: number[];
    interval: number;
    tcpFailures: number;
    timeouts: number;
}

export interface UpstreamHealthchecksPassive {
    healthy: outputs.UpstreamHealthchecksPassiveHealthy;
    type?: string;
    unhealthy: outputs.UpstreamHealthchecksPassiveUnhealthy;
}

export interface UpstreamHealthchecksPassiveHealthy {
    httpStatuses?: number[];
    successes?: number;
}

export interface UpstreamHealthchecksPassiveUnhealthy {
    httpFailures?: number;
    httpStatuses?: number[];
    tcpFailures?: number;
    timeouts?: number;
}

