# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ConsumerKeyAuthArgs', 'ConsumerKeyAuth']

@pulumi.input_type
class ConsumerKeyAuthArgs:
    def __init__(__self__, *,
                 consumer_id: pulumi.Input[str],
                 key: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ConsumerKeyAuth resource.
        :param pulumi.Input[str] consumer_id: the id of the consumer to associate the credentials to
        :param pulumi.Input[str] key: Unique key to authenticate the client; if omitted the plugin will generate one
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of strings associated with the consumer key auth for grouping and filtering
        """
        pulumi.set(__self__, "consumer_id", consumer_id)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="consumerId")
    def consumer_id(self) -> pulumi.Input[str]:
        """
        the id of the consumer to associate the credentials to
        """
        return pulumi.get(self, "consumer_id")

    @consumer_id.setter
    def consumer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "consumer_id", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Unique key to authenticate the client; if omitted the plugin will generate one
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of strings associated with the consumer key auth for grouping and filtering
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ConsumerKeyAuthState:
    def __init__(__self__, *,
                 consumer_id: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ConsumerKeyAuth resources.
        :param pulumi.Input[str] consumer_id: the id of the consumer to associate the credentials to
        :param pulumi.Input[str] key: Unique key to authenticate the client; if omitted the plugin will generate one
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of strings associated with the consumer key auth for grouping and filtering
        """
        if consumer_id is not None:
            pulumi.set(__self__, "consumer_id", consumer_id)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="consumerId")
    def consumer_id(self) -> Optional[pulumi.Input[str]]:
        """
        the id of the consumer to associate the credentials to
        """
        return pulumi.get(self, "consumer_id")

    @consumer_id.setter
    def consumer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consumer_id", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Unique key to authenticate the client; if omitted the plugin will generate one
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of strings associated with the consumer key auth for grouping and filtering
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class ConsumerKeyAuth(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_id: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## # ConsumerKeyAuth

        Resource that allows you to configure the [Key Authentication](https://docs.konghq.com/hub/kong-inc/key-auth/) plugin for a consumer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_kong as kong

        my_consumer = kong.Consumer("myConsumer",
            username="User1",
            custom_id="123")
        key_auth_plugin = kong.Plugin("keyAuthPlugin")
        consumer_key_auth = kong.ConsumerKeyAuth("consumerKeyAuth",
            consumer_id=my_consumer.id,
            key="secret",
            tags=[
                "myTag",
                "anotherTag",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consumer_id: the id of the consumer to associate the credentials to
        :param pulumi.Input[str] key: Unique key to authenticate the client; if omitted the plugin will generate one
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of strings associated with the consumer key auth for grouping and filtering
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConsumerKeyAuthArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # ConsumerKeyAuth

        Resource that allows you to configure the [Key Authentication](https://docs.konghq.com/hub/kong-inc/key-auth/) plugin for a consumer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_kong as kong

        my_consumer = kong.Consumer("myConsumer",
            username="User1",
            custom_id="123")
        key_auth_plugin = kong.Plugin("keyAuthPlugin")
        consumer_key_auth = kong.ConsumerKeyAuth("consumerKeyAuth",
            consumer_id=my_consumer.id,
            key="secret",
            tags=[
                "myTag",
                "anotherTag",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param ConsumerKeyAuthArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConsumerKeyAuthArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_id: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConsumerKeyAuthArgs.__new__(ConsumerKeyAuthArgs)

            if consumer_id is None and not opts.urn:
                raise TypeError("Missing required property 'consumer_id'")
            __props__.__dict__["consumer_id"] = consumer_id
            __props__.__dict__["key"] = key
            __props__.__dict__["tags"] = tags
        super(ConsumerKeyAuth, __self__).__init__(
            'kong:index/consumerKeyAuth:ConsumerKeyAuth',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            consumer_id: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ConsumerKeyAuth':
        """
        Get an existing ConsumerKeyAuth resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consumer_id: the id of the consumer to associate the credentials to
        :param pulumi.Input[str] key: Unique key to authenticate the client; if omitted the plugin will generate one
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of strings associated with the consumer key auth for grouping and filtering
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConsumerKeyAuthState.__new__(_ConsumerKeyAuthState)

        __props__.__dict__["consumer_id"] = consumer_id
        __props__.__dict__["key"] = key
        __props__.__dict__["tags"] = tags
        return ConsumerKeyAuth(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="consumerId")
    def consumer_id(self) -> pulumi.Output[str]:
        """
        the id of the consumer to associate the credentials to
        """
        return pulumi.get(self, "consumer_id")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        Unique key to authenticate the client; if omitted the plugin will generate one
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of strings associated with the consumer key auth for grouping and filtering
        """
        return pulumi.get(self, "tags")
