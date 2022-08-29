# tcp output

## How To Use

1. Clone this code to `elastic/beats/libbeat/output/`

2. Modify `elastic/beats/libbeat/publisher/includes/includes.go` :
   ```go
   // add import
   import _ "github.com/elastic/beats/v7/libbeat/output/tcp"
   ```

3. Compile beats

## Configuration

### Example

```yaml
output.tcp:
  host: 127.0.0.1
  port: 8080
  ssl:
    enable: true
    cert_path: ...
    key_path: ...
  buffer_size: 1024
  writev: true
  line_delimiter: \n
  codec: ...
```

### Options

#### buffer_size

The buffer size of `net.Buffers`. Default 1<<15 (32768).

#### line_delimiter

Specify the characters used to split the output events. Default \n.

#### codec

Output codec configuration. If the codec section is missing, events will be json encoded using the pretty option.

See [Change the output codec](https://www.elastic.co/guide/en/beats/filebeat/master/configuration-output-codec.html) for
more information.
