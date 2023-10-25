FROM docker.elastic.co/logstash/logstash:8.10.4 as builder
RUN /usr/share/logstash/bin/logstash-plugin install logstash-integration-jdbc
FROM docker.elastic.co/logstash/logstash:8.10.4
COPY --from=builder /usr/share/logstash /usr/share/logstash
ENTRYPOINT ["/usr/local/bin/docker-entrypoint"]