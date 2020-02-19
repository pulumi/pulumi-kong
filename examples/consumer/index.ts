import * as kong from "@pulumi/kong";

const consumer = new kong.Consumer("my-consumer", {
    username: "my-username-1",
    customId: "123"
});

export const consumerId = consumer.id;

