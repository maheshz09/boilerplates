version: '2'
 
services:
  mongodb:
    image: mongo:6.0
    volumes:
    - /data/graylog/mongodb/data/db:/data/db
    ports:
    - 27017:27017
    networks:
      - graylog_network # Ensure it's on the same network as other services
  # Elasticsearch: https://www.elastic.co/guide/en/elasticsearch/reference/6.x/docker.html
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:7.10.2
    volumes:
    - /data/graylog/elasticsearch/usr/share/elasticsearch/data:/usr/share/elasticsearch/data
    environment:
      - http.host=0.0.0.0
      - transport.host=localhost
      - network.host=0.0.0.0
      # Disable X-Pack security: https://www.elastic.co/guide/en/elasticsearch/reference/6.x/security-settings.html#general-security-settings
      - xpack.security.enabled=false
      - xpack.watcher.enabled=false
      - xpack.monitoring.enabled=false
      - xpack.security.audit.enabled=false
      - xpack.ml.enabled=false
      - xpack.graph.enabled=false
      - "ES_JAVA_OPTS=-Xms512m -Xmx8500m"
      # log4j Vulnerability Solution
      - ES_JAVA_OPTS=-Dlog4j2.formatMsgNoLookups=true
    ports:
      # Port exposed for direct access from Grafana
      - 9200:9200
    ulimits:
      memlock:
        soft: -1
        hard: -1
    mem_limit: 10g
    networks:
      - graylog_network # Ensure it's on the same network as other services

  opensearch:
    image: opensearchproject/opensearch:2.11.0 # Use a specific, tested version
    container_name: graylog_opensearch_1
    volumes:
      - /data/graylog/opensearch/data:/usr/share/opensearch/data # Persist data
      - /data/graylog/snapshots:/usr/share/opensearch/backup # For backups (optional)
    environment:
      - discovery.type=single-node # For single-node setup. Adjust for clusters.
      - network.host=0.0.0.0 # Listen on all interfaces
      - OPENSEARCH_JAVA_OPTS=-Xms512m -Xmx8500m # Adjust JVM heap size as needed
      - bootstrap.memory_lock=true # Prevent swapping
      - DISABLE_SECURITY_PLUGIN=true # Disable security for simplicity (CONSIDER SECURITY!)
    ports:
      - 9201:9200 # Host:Container port mapping (change host port if needed)
      - 9301:9300 # Transport port (used for internal communication)
    ulimits:
      memlock:
        soft: -1 # Unlimited memory lock
        hard: -1
    mem_limit: 10g # Adjust memory limit as needed
    networks:
      - graylog_network # Ensure it's on the same network as other services
  #kibana:
  #  image: kibana:6.8.15
  #  container_name: kibana
  #  environment:
  #   ELASTICSEARCH_HOST: http://10.70.88.10:9200
  #  kibana:
  #  image: kibana:6.8.15
  #  container_name: kibana
  #  environment:
  #   ELASTICSEARCH_HOST: http://10.70.88.10
  #   ELASTICSEARCH_PORT: 9200
  # depends_on:
  #   - elasticsearch
  # ports:
  #   - 5601:5601
 
# Graylog: https://hub.docker.com/r/graylog/graylog/
  graylog:
    image: graylog/graylog:6.1.12
    volumes:
    - /data/graylog/graylog/plugin:/usr/share/graylog/plugin
    - /data/graylog/graylog/config:/usr/share/graylog/data/config
    environment:
      - http.host=0.0.0.0
      - transport.host=localhost
      - network.host=0.0.0.0
      # CHANGE ME!
      - GRAYLOG_PASSWORD_SECRET=C@pGray20192019!
      # Password: admin
      - GRAYLOG_ROOT_PASSWORD_SHA2=4da2ed7bbc9f8df61ea7eb606c4189a63883d6a11e6de1554d909f714a094e8e
      #- GRAYLOG_HTTP_EXTERNAL_URI=
      # log4j Vulnerability Solution
      - GRAYLOG_SERVER_JAVA_OPTS=-Dlog4j2.formatMsgNoLookups=true
      # the below parameter is to test
      - TRUSTED_PROXIES=10.24.221.150/32, 10.24.221.251/32
    links:
      - mongodb:mongo
      - elasticsearch
    depends_on:
      - mongodb
      - elasticsearch
    ports:
      # Graylog web interface and REST API
      - 30000:9000
      # Syslog TCP
      - 514:514
      # Syslog UDP
      - 514:514/udp
      # GELF TCP
      - 12201:12201
      # GELF UDP
      - 30001:30001/udp
      # GELF TCP
      - 30001:30001
    networks:
      - graylog_network # Ensure it's on the same network as other services

networks:
  graylog_network:
    driver: bridge