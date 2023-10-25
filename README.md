Run app using ```docker-compose up --build```

failed logs
```
 => CANCELED [api 6/6] RUN go build -o /usr/bin/letterpress github.com/sonichigo/letterpress  28.9s
 => CACHED [logstash builder 1/2] FROM docker.elastic.co/logstash/logstash:8.10.4@sha256:8239  0.0s
 => ERROR [logstash builder 2/2] RUN /usr/share/logstash/bin/logstash-plugin install logstas  26.9s
------
 > [logstash builder 2/2] RUN /usr/share/logstash/bin/logstash-plugin install logstash-integration-jdbc:
#0 0.403 Using bundled JDK: /usr/share/logstash/jdk
#0 14.08 Validating logstash-integration-jdbc
#0 18.22 Resolving mixin dependencies
#0 26.87 Killed
------
failed to solve: process "/bin/sh -c /usr/share/logstash/bin/logstash-plugin install logstash-integration-jdbc" did not complete successfully: exit code: 137
```