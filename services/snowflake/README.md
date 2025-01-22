# Snowflake

Snowflake is a simple ID generator based on [Twitter's snowflakes](https://en.wikipedia.org/wiki/Snowflake_ID).

The benefits of using snowflakes over other popular unique identifier strategies (e.g. UUID) are well documented, but in this particular case
it has been chosen due to its size (64 bits vs UUID's 128 bits), along with other scaling benefits.

### Getting Started

To get a development instance of this service up and running, you may use the following command:

```
make run
```
