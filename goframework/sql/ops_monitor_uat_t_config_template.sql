INSERT INTO ops_monitor_uat.t_config_template (type, name, field, comment, `default`, required)
VALUES
#     ('rabbitmq', '前缀', 'prefix', '前缀', 'rabbitmq', 'false'),
#     ('rabbitmq', '类型', 'type', '类型', 'rabbitmq', 'false'),
#     ('rabbitmq', '别名', 'asName', '别名', 'rabbitmq', 'false'),
#     ('rabbitmq', '连接地址', 'RABBIT_URL', '地址', 'http://127.0.0.1:15672', 'true'),
#     ('rabbitmq', '用户名', 'RABBIT_USER', '用户名', 'admin', 'true'),
#     ('rabbitmq', '密码', 'RABBIT_PASSWORD', '密码', '********', 'true'),
#     ('rabbitmq', '密码', 'RABBIT_USER_FILE', '密码', '********', 'true'),

('solr', '前缀', 'prefix', '前缀', 'solr', 'false'),
('solr', '类型', 'type', '类型', 'solr', 'false'),
('solr', '别名', 'asName', '别名', 'solr', 'false'),
('solr', '账号', 'user', '账号', 'root', 'true'),
('solr', '密码', 'pass', '密码', '123456', 'true'),
('solr', '地址', 'url', '地址', 'http://localhost/sdk', 'true'),

('vmware', '前缀', 'prefix', '前缀', 'vmware', 'false'),
('vmware', '类型', 'type', '类型', 'vmware', 'false'),
('vmware', '别名', 'asName', '别名', 'asName', 'false'),
('vmware', '账号', 'username', '账号', 'root', 'true'),
('vmware', '密码', 'password', '密码', '123456', 'true'),
('vmware', '地址', 'url', '地址', 'http://localhost/sdk', 'true'),

('tidb', '前缀', 'prefix', '前缀', 'tidb', 'false'),
('tidb', '类型', 'type', '类型', 'tidb', 'false'),
('tidb', '别名', 'asName', '别名', 'tidb', 'false'),
('tidb', '地址', 'url', '地址', 'http://localhost/:8080', 'true'),

('tomcat', '前缀', 'prefix', '前缀', 'tomcat', 'false'),
('tomcat', '类型', 'type', '类型', 'tomcat', 'false'),
('tomcat', '别名', 'asName', '别名', 'tomcat', 'false'),
('tomcat', '账号', 'user', '账号', 'root', 'true'),
('tomcat', '密码', 'pass', '密码', '123456', 'true'),
('tomcat', '地址', 'url', '地址', 'http://localhost/:8080', 'true'),

('zookeeper', '前缀', 'prefix', '前缀', 'zookeeper', 'false'),
('zookeeper', '类型', 'type', '类型', 'zookeeper', 'false'),
('zookeeper', '别名', 'asName', '别名', 'zookeeper', 'false'),
('zookeeper', '地址', 'url', '地址', 'http://localhost/:8080', 'true'),

('apache', '前缀', 'prefix', '前缀', 'apache', 'false'),
('apache', '类型', 'type', '类型', 'apache', 'false'),
('apache', '别名', 'asName', '别名', 'apache', 'false'),
('apache', '地址', 'url', '地址', 'http://localhost/:8080', 'true'),

('ceph', '前缀', 'prefix', '前缀', 'ceph', 'false'),
('ceph', '类型', 'type', '类型', 'ceph', 'false'),
('ceph', '别名', 'asName', '别名', 'ceph', 'false'),
('ceph', 'CephBinary', 'CephBinary', 'CephBinary', '-', 'true'),
('ceph', 'MonPrefix', 'MonPrefix', 'MonPrefix', '-', 'true'),
('ceph', 'MdsPrefix', 'MdsPrefix', 'MdsPrefix', '-', 'true'),
('ceph', 'RgwPrefix', 'RgwPrefix', 'RgwPrefix', '-', 'true'),
('ceph', 'SocketDir', 'SocketDir', 'SocketDir', '-', 'true'),
('ceph', 'SocketSuffix', 'SocketSuffix', 'SocketSuffix', '-', 'true'),
('ceph', 'CephUser', 'CephUser', 'CephUser', '-', 'true'),
('ceph', 'CephConfig', 'CephConfig', 'CephConfig', '-', 'true'),

('consul', '前缀', 'prefix', '前缀', 'consul', 'false'),
('consul', '类型', 'type', '类型', 'consul', 'false'),
('consul', '别名', 'asName', '别名', 'consul', 'false'),
('consul', '地址', 'url', '地址', 'http://localhost:8500', 'true'),

('docker', '前缀', 'prefix', '前缀', 'consul', 'false'),
('docker', '类型', 'type', '类型', 'consul', 'false'),
('docker', '别名', 'asName', '别名', 'consul', 'false'),
('docker', '地址', 'url', '地址', 'unix://localhost:8500', 'true'),

('etcd', '前缀', 'prefix', '前缀', 'etcd', 'false'),
('etcd', '类型', 'type', '类型', 'etcd', 'false'),
('etcd', '别名', 'asName', '别名', 'etcd', 'false'),
('etcd', '地址', 'url', '地址', 'http://localhost:8500', 'true'),

('eventstoredb', '前缀', 'prefix', '前缀', 'eventstoredb', 'false'),
('eventstoredb', '类型', 'type', '类型', 'eventstoredb', 'false'),
('eventstoredb', '别名', 'asName', '别名', 'eventstoredb', 'false'),
('eventstoredb', '账号', 'user', '账号', 'root', 'true'),
('eventstoredb', '密码', 'pass', '密码', '123456', 'true'),
('eventstoredb', '模式', 'mode', '模式', '-', 'true'),
('eventstoredb', '地址', 'url', '地址', 'http://localhost/:8080', 'true'),

('graphite', '前缀', 'prefix', '前缀', 'graphite', 'false'),
('graphite', '类型', 'type', '类型', 'graphite', 'false'),
('graphite', '别名', 'asName', '别名', 'graphite', 'false'),
('graphite', '地址', 'url', '地址', 'http://localhost:8500', 'true'),

('influxdb', '前缀', 'prefix', '前缀', 'influxdb', 'false'),
('influxdb', '类型', 'type', '类型', 'influxdb', 'false'),
('influxdb', '别名', 'asName', '别名', 'influxdb', 'false'),
('influxdb', '地址', 'url', '地址', 'http://localhost:8500', 'true'),

('jenkins', '前缀', 'prefix', '前缀', 'jenkins', 'false'),
('jenkins', '类型', 'type', '类型', 'jenkins', 'false'),
('jenkins', '别名', 'asName', '别名', 'jenkins', 'false'),
('jenkins', '账号', 'user', '账号', 'root', 'true'),
('jenkins', '密码', 'pass', '密码', '123456', 'true'),
('jenkins', '地址', 'url', '地址', 'http://localhost/:8080', 'true'),

('kafka', '前缀', 'prefix', '前缀', 'kafka', 'false'),
('kafka', '类型', 'type', '类型', 'kafka', 'false'),
('kafka', '别名', 'asName', '别名', 'kafka', 'false'),
('kafka', '地址', 'url', '地址', 'http://localhost/:8080', 'true'),

('clickhouse', '前缀', 'prefix', '前缀', 'clickhouse', 'false'),
('clickhouse', '类型', 'type', '类型', 'clickhouse', 'false'),
('clickhouse', '别名', 'asName', '别名', 'clickhouse', 'false'),
('clickhouse', '账号', 'user', '账号', 'root', 'true'),
('clickhouse', '密码', 'pass', '密码', '123456', 'true'),
('clickhouse', '地址', 'url', '地址', 'http://localhost/:8080', 'true'),

('kubernetes', '前缀', 'prefix', '前缀', 'kubernetes', 'false'),
('kubernetes', '类型', 'type', '类型', 'kubernetes', 'false'),
('kubernetes', '别名', 'asName', '别名', 'kubernetes', 'false'),
('kubernetes', '配置文件', 'config_file_path', '配置文件目录', './admin.conf', 'true'),
('kubernetes', '地址[ApiServer]', 'api_server_url', 'ApiServer地址', 'https://localhost:6443', 'true'),
('kubernetes', '地址[KubelatApi]', 'metrics_url', 'KubelatApi地址', 'https://localhost:10250', 'true'),

('mongodb', '前缀', 'prefix', '前缀', 'mongodb', 'false'),
('mongodb', '类型', 'type', '类型', 'mongodb', 'false'),
('mongodb', '别名', 'asName', '别名', 'mongodb', 'false'),
('mongodb', '地址', 'url', '地址', 'http://localhost/:8080', 'true'),

('nginx', '前缀', 'prefix', '前缀', 'nginx', 'false'),
('nginx', '类型', 'type', '类型', 'nginx', 'false'),
('nginx', '别名', 'asName', '别名', 'nginx', 'false'),
('nginx', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('nginx_plus', '前缀', 'prefix', '前缀', 'nginx_plus', 'false'),
('nginx_plus', '类型', 'type', '类型', 'nginx_plus', 'false'),
('nginx_plus', '别名', 'asName', '别名', 'nginx_plus', 'false'),
('nginx_plus', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('openstack', '前缀', 'prefix', '前缀', 'nginx_plus', 'false'),
('openstack', '类型', 'type', '类型', 'nginx_plus', 'false'),
('openstack', '别名', 'asName', '别名', 'nginx_plus', 'false'),
('openstack', '鉴权服务Api版本', 'identity_api_version', '鉴权服务Api版本', '-', 'true'),
('openstack', '存储Api版本', 'volume_api_version', '存储Api版本', '-', 'true'),
('openstack', '认证地址', 'auth_url', '认证地址', '-', 'true'),
('openstack', '区域', 'user_domain_name', '区域', '-', 'true'),
('openstack', '项目名称', 'project_name', '项目名称', '-', 'true'),
('openstack', '用户名', 'username', '用户名', '-', 'true'),
('openstack', '密码', 'password', '密码', '-', 'true'),

('oracle', '前缀', 'prefix', '前缀', 'oracle', 'false'),
('oracle', '类型', 'type', '类型', 'oracle', 'false'),
('oracle', '别名', 'asName', '别名', 'oracle', 'false'),
('oracle', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('postgres', '前缀', 'prefix', '前缀', 'postgres', 'false'),
('postgres', '类型', 'type', '类型', 'postgres', 'false'),
('postgres', '别名', 'asName', '别名', 'postgres', 'false'),
('postgres', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('activemq', '前缀', 'prefix', '前缀', 'activemq', 'false'),
('activemq', '类型', 'type', '类型', 'activemq', 'false'),
('activemq', '别名', 'asName', '别名', 'activemq', 'false'),
('activemq', '端口', 'port', '端口', 'activemq', 'false'),
('activemq', '用户名', 'user', '用户名', '用户名', 'false'),
('activemq', '密码', 'pass', '密码', '密码', 'false'),
('activemq', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('lustre', '前缀', 'prefix', '前缀', 'lustre', 'false'),
('lustre', '类型', 'type', '类型', 'lustre', 'false'),
('lustre', '别名', 'asName', '别名', 'lustre', 'false'),
('lustre', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('elasticsearch', '前缀', 'prefix', '前缀', 'elasticsearch', 'false'),
('elasticsearch', '类型', 'type', '类型', 'elasticsearch', 'false'),
('elasticsearch', '别名', 'asName', '别名', 'elasticsearch', 'false'),
('elasticsearch', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('gpfs', '前缀', 'prefix', '前缀', 'gpfs', 'false'),
('gpfs', '类型', 'type', '类型', 'gpfs', 'false'),
('gpfs', '别名', 'asName', '别名', 'gpfs', 'false'),
('gpfs', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('gluster', '前缀', 'prefix', '前缀', 'gluster', 'false'),
('gluster', '类型', 'type', '类型', 'gluster', 'false'),
('gluster', '别名', 'asName', '别名', 'gluster', 'false'),
('gluster', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('opentsdb', '前缀', 'prefix', '前缀', 'opentsdb', 'false'),
('opentsdb', '类型', 'type', '类型', 'opentsdb', 'false'),
('opentsdb', '别名', 'asName', '别名', 'opentsdb', 'false'),
('opentsdb', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('gearman', '前缀', 'prefix', '前缀', 'gearman', 'false'),
('gearman', '类型', 'type', '类型', 'gearman', 'false'),
('gearman', '别名', 'asName', '别名', 'gearman', 'false'),
('gearman', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('druid', '前缀', 'prefix', '前缀', 'druid', 'false'),
('druid', '类型', 'type', '类型', 'druid', 'false'),
('druid', '别名', 'asName', '别名', 'druid', 'false'),
('druid', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('aurora', '前缀', 'prefix', '前缀', 'aurora', 'false'),
('aurora', '类型', 'type', '类型', 'aurora', 'false'),
('aurora', '别名', 'asName', '别名', 'aurora', 'false'),
('aurora', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('logstash', '前缀', 'prefix', '前缀', 'logstash', 'false'),
('logstash', '类型', 'type', '类型', 'logstash', 'false'),
('logstash', '别名', 'asName', '别名', 'logstash', 'false'),
('logstash', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('kibana', '前缀', 'prefix', '前缀', 'kibana', 'false'),
('kibana', '类型', 'type', '类型', 'kibana', 'false'),
('kibana', '别名', 'asName', '别名', 'kibana', 'false'),
('kibana', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('nsq', '前缀', 'prefix', '前缀', 'nsq', 'false'),
('nsq', '类型', 'type', '类型', 'nsq', 'false'),
('nsq', '别名', 'asName', '别名', 'nsq', 'false'),
('nsq', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('nats', '前缀', 'prefix', '前缀', 'nats', 'false'),
('nats', '类型', 'type', '类型', 'nats', 'false'),
('nats', '别名', 'asName', '别名', 'nats', 'false'),
('nats', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('twemproxy', '前缀', 'prefix', '前缀', 'twemproxy', 'false'),
('twemproxy', '类型', 'type', '类型', 'twemproxy', 'false'),
('twemproxy', '别名', 'asName', '别名', 'twemproxy', 'false'),
('twemproxy', 'pools', 'pools', 'pools', '-', 'true'),
('twemproxy', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('memcached', '前缀', 'prefix', '前缀', 'memcached', 'false'),
('memcached', '类型', 'type', '类型', 'memcached', 'false'),
('memcached', '别名', 'asName', '别名', 'memcached', 'false'),
('memcached', '地址', 'uri', '地址', 'http://localhost:9080', 'true'),

('uwsgi', '前缀', 'prefix', '前缀', 'uwsgi', 'false'),
('uwsgi', '类型', 'type', '类型', 'uwsgi', 'false'),
('uwsgi', '别名', 'asName', '别名', 'uwsgi', 'false'),
('uwsgi', '地址', 'uri', '地址', 'http://localhost:9080', 'true'),

('udp_listener', '前缀', 'prefix', '前缀', 'udp_listener', 'false'),
('udp_listener', '类型', 'type', '类型', 'udp_listener', 'false'),
('udp_listener', '别名', 'asName', '别名', 'udp_listener', 'false'),
('udp_listener', '地址', 'host', '地址', 'localhost', 'true'),
('udp_listener', '端口', 'port', '端口', '8080', 'true'),
('udp_listener', '地址', 'uri', '地址', 'http://localhost:9080', 'true'),

('sqlserver', '前缀', 'prefix', '前缀', 'sqlserver', 'false'),
('sqlserver', '类型', 'type', '类型', 'sqlserver', 'false'),
('sqlserver', '别名', 'asName', '别名', 'sqlserver', 'false'),
('sqlserver', '用户名', 'username', '用户名', '用户名', 'false'),
('sqlserver', '密码', 'password', '密码', '密码', 'false'),
('sqlserver', '地址', 'url', '地址', 'http://127.0.0.1:8080/stub_status', 'true'),

('zfs', '前缀', 'prefix', '前缀', 'zfs', 'false'),
('zfs', '类型', 'type', '类型', 'zfs', 'false'),
('zfs', '别名', 'asName', '别名', 'zfs', 'false'),
('zfs', '地址', 'uri', '地址', 'http://localhost:9080', 'true'),

('apisix', '前缀', 'prefix', '前缀', 'apisix', 'false'),
('apisix', '类型', 'type', '类型', 'apisix', 'false'),
('apisix', '别名', 'asName', '别名', 'apisix', 'false'),
('apisix', '地址', 'uri', '地址', 'http://localhost:9080', 'true'),

('mysql', '前缀', 'prefix', '前缀', 'mysql', 'false'),
('mysql', '类型', 'type', '类型', 'mysql', 'false'),
('mysql', '别名', 'asName', '别名', 'mysql', 'false'),
('mysql', '地址', 'uri', '地址', 'user:password@(localhost:3306)/', 'true'),

('redis', '前缀', 'prefix', '前缀', 'redis', 'false'),
('redis', '类型', 'type', '类型', 'redis', 'false'),
('redis', '别名', 'asName', '别名', 'redis', 'false'),
('redis', '地址', 'uri', '地址', 'redis://localhost:7000', 'true'),

('couchbase', '前缀', 'prefix', '前缀', 'couchbase', 'false'),
('couchbase', '类型', 'type', '类型', 'couchbase', 'false'),
('couchbase', '别名', 'asName', '别名', 'couchbase', 'false'),
('couchbase', '地址', 'uri', '地址', 'http://localhost:8091', 'true'),

('beat', '前缀', 'prefix', '前缀', 'beat', 'false'),
('beat', '类型', 'type', '类型', 'beat', 'false'),
('beat', '别名', 'asName', '别名', 'beat', 'false'),
('beat', '地址', 'uri', '地址', 'http://localhost:5066', 'true'),

('dm', '前缀', 'prefix', '前缀', 'dm', 'false'),
('dm', '类型', 'type', '类型', 'dm', 'false'),
('dm', '别名', 'asName', '别名', 'dm', 'false'),
('dm', '地址', 'uri', '地址', 'http://localhost:5066', 'true');