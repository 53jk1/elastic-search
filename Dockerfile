FROM docker.elastic.co/elasticsearch/elasticsearch:6.2.4
RUN bin/elasticsearch-plugin install analysis-icu
RUN bin/elasticsearch-plugin install analysis-kuromoji
RUN bin/elasticsearch-plugin install analysis-smartcn
RUN bin/elasticsearch-plugin install analysis-stempel
