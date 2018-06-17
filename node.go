package kongo

import (
	"context"
	"net/http"
	"net/url"
)

const (
	nodeInfoResourcePath   = "/"
	nodeStatusResourcePath = "/status"
)

type (
	// Node retrieves the info about the server nodes.
	Node interface {
		// Info retrieves the information about the server node.
		Info() (*NodeInfo, *http.Response, error)

		// InfoWithContext retrieves the information about the server node.
		InfoWithContext(ctx context.Context) (*NodeInfo, *http.Response, error)

		// Status retrieves the status of the server node.
		Status() (*NodeStatus, *http.Response, error)

		// StatusWithContext retrieves the status of the server node.
		StatusWithContext(ctx context.Context) (*NodeStatus, *http.Response, error)
	}

	// NodeService it's a concrete instance of node.
	NodeService struct {
		// Kongo client manages communication throught API.
		client *Kongo
	}

	// NodeInfo it's a structure of API result.
	NodeInfo struct {
		Configuration *NodeInfoConfiguration `json:"configuration"`
		Hostname      string                 `json:"hostname"`
		LuaVersion    string                 `json:"lua_version"`
		Plugins       *NodeInfoPlugins       `json:"plugins"`
		PrngSeeds     map[string]int         `json:"prng_seeds"`
		Tagline       string                 `json:"tagline"`
		Timers        *NodeInfoTimers        `json:"timers"`
		Version       string                 `json:"version"`
	}

	// NodeInfoConfiguration it's a structure of API result.
	NodeInfoConfiguration struct {
		AdminAcessLog                 string              `json:"admin_access_log"`
		AdminErrorLog                 string              `json:"admin_error_log"`
		AdminSSLCertificateDefault    string              `json:"admin_ssl_cert_default"`
		AdminSSLCertificateCsrDefault string              `json:"admin_ssl_cert_csr_default"`
		AdminSSLCertificateKeyDefault string              `json:"admin_ssl_cert_key_default"`
		AdminListen                   []string            `json:"admin_listen"`
		AdminListeners                []*NodeInfoListener `json:"admin_listeners"`

		CassandraConsistency            string   `json:"cassandra_consistency"`
		CassandraKeyspace               string   `json:"cassandra_keyspace"`
		CassandraLBPolicy               string   `json:"cassandra_lb_policy"`
		CassandraReplicationStrategy    string   `json:"cassandra_repl_strategy"`
		CassandraUsername               string   `json:"cassandra_username"`
		CassandraContactPoints          []string `json:"cassandra_contact_points"`
		CassandraDataCenters            []string `json:"cassandra_data_centers"`
		CassandraPort                   int      `json:"cassandra_port"`
		CassandraReplicationFactor      int      `json:"cassandra_repl_factor"`
		CassandraSchemaConsensusTimeout int      `json:"cassandra_schema_consensus_timeout"`
		CassandraTimeout                int      `json:"cassandra_timeout"`

		ClientBodyBufferSize           string `json:"client_body_buffer_size"`
		ClientMaxBodySize              string `json:"client_max_body_size"`
		ClientSSLCertificateCsrDefault string `json:"client_ssl_cert_csr_default"`
		ClientSSLCertificateDefault    string `json:"client_ssl_cert_default"`
		ClientSSLCertificateKeyDefault string `json:"client_ssl_cert_key_default"`

		CustomPlugins interface{} `json:"custom_plugins"`

		Database                  string `json:"database"`
		DatabaseCacheTTL          int    `json:"db_cache_ttl"`
		DatabaseUpdateFrequency   int    `json:"db_update_frequency"`
		DatabaseUpdatePropagation int    `json:"db_update_propagation"`

		DNSHostsFile   string      `json:"dns_hostsfile"`
		DNSOrder       []string    `json:"dns_order"`
		DNSResolver    interface{} `json:"dns_resolver"`
		DNSNotFoundTTL int         `json:"dns_not_found_ttl"`
		DNSErrorTTL    int         `json:"dns_error_ttl"`
		DNSStaleTTL    int         `json:"dns_stale_ttl"`

		ErrorDefaultType string `json:"error_default_type"`

		KongEnv string `json:"kong_env"`

		LuaPackageCPath   string `json:"lua_package_cpath"`
		LuaPackagePath    string `json:"lua_package_path"`
		LuaSocketPoolSize int    `json:"lua_socket_pool_size"`
		LuaSSLVerifyDepth int    `json:"lua_ssl_verify_depth"`

		LogLevel string `json:"log_level"`

		MemoryCacheSize string `json:"mem_cache_size"`

		NginxAccessLogs      string `json:"nginx_acc_logs"`
		NginxAdminAccessLog  string `json:"nginx_admin_acc_logs"`
		NginxConf            string `json:"nginx_conf"`
		NginxDaemon          string `json:"nginx_daemon"`
		NginxErrorLogs       string `json:"nginx_err_logs"`
		NginxKongConf        string `json:"nginx_kong_conf"`
		NginxPID             string `json:"nginx_pid"`
		NginxWorkerProcesses string `json:"nginx_worker_processes"`

		Plugins map[string]bool `json:"plugins"`

		PostgresDatabase string `json:"pg_database"`
		PostgresHost     string `json:"pg_host"`
		PostgresUsername string `json:"pg_user"`
		PostgresPort     int    `json:"pg_port"`

		Prefix string `json:"prefix"`

		ProxyAccessLog string              `json:"proxy_access_log"`
		ProxyErrorLog  string              `json:"proxy_error_log"`
		ProxyListen    []string            `json:"proxy_listen"`
		ProxyListeners []*NodeInfoListener `json:"proxy_listeners"`

		RealIPHeader    string `json:"real_ip_header"`
		RealIPRecursive string `json:"real_ip_recursive"`

		SSLCertificate           string `json:"ssl_cert"`
		SSLCertificateDefault    string `json:"ssl_cert_default"`
		SSLCertificateKey        string `json:"ssl_cert_key"`
		SSLCertificateDefaultKey string `json:"ssl_cert_key_default"`
		SSLCertificateCsrDefault string `json:"ssl_cert_csr_default"`
		SSLCiphers               string `json:"ssl_ciphers"`
		SSLCipherSuite           string `json:"ssl_cipher_suite"`

		TrustedIPs interface{} `json:"trusted_ips"`

		UpstreamKeepAlive int `json:"upstream_keepalive"`

		AdminSSLEnabled bool `json:"admin_ssl_enabled"`

		AnonymousReports bool `json:"anonymous_reports"`

		ClientSSL bool `json:"client_ssl"`

		CassandraSSL       bool `json:"cassandra_ssl"`
		CassandraSSLVerify bool `json:"cassandra_ssl_verify"`

		DNSNoSync bool `json:"dns_no_sync"`

		LatencyTokens bool `json:"latency_tokens"`

		NginxOptimizations bool `json:"nginx_optimizations"`

		PostgresSSL       bool `json:"pg_ssl"`
		PostgresSSLVerify bool `json:"pg_ssl_verify"`

		ProxySSLEnabled bool `json:"proxy_ssl_enabled"`

		ServerTokens bool `json:"server_tokens"`
	}

	// NodeInfoListener it's a structure of API result.
	NodeInfoListener struct {
		IP       string `json:"ip"`
		Listener string `json:"listener"`
		Port     int    `json:"port"`
		Protocol bool   `json:"protocol"`
		SSL      bool   `json:"ssl"`
		HTTP2    bool   `json:"http2"`
	}

	// NodeInfoPlugins it's a structure of API result.
	NodeInfoPlugins struct {
		AvailableOnServer map[string]bool `json:"available_on_server"`
		EnabledInCluster  []string        `json:"enabled_in_cluster"`
	}

	// NodeInfoTimers it's a structure of API result.
	NodeInfoTimers struct {
		Pending int `json:"pending"`
		Running int `json:"running"`
	}

	// NodeStatus it's a structure of API result.
	NodeStatus struct {
		Database *NodeStatusDatabase `json:"database"`
		Server   *NodeStatusServer   `json:"server"`
	}

	// NodeStatusDatabase it's a structure of API result.
	NodeStatusDatabase struct {
		Reachable bool `json:"reachable, omitempty"`
	}

	// NodeStatusServer it's a structure of API result.
	NodeStatusServer struct {
		ConnectionsAccepted int `json:"connections_accepted, omitempty"`
		ConnectionsActive   int `json:"connections_active, omitempty"`
		ConnectionsHandled  int `json:"connections_handled, omitempty"`
		ConnectionsReading  int `json:"connections_reading, omitempty"`
		ConnectionsWaiting  int `json:"connections_waiting, omitempty"`
		ConnectionsWriting  int `json:"connections_writing, omitempty"`
		TotalRequests       int `json:"total_requests, omitempty"`
	}
)

// InfoWithContext retrieves the server node information.
func (n *NodeService) InfoWithContext(ctx context.Context) (*NodeInfo, *http.Response, error) {
	resource, _ := url.Parse(nodeInfoResourcePath)

	req, err := n.client.NewRequest(ctx, http.MethodGet, resource, nil)

	if err != nil {
		return nil, nil, err
	}

	nodeInfo := new(NodeInfo)

	res, err := n.client.Do(req, nodeInfo)

	if err != nil {
		return nil, res, err
	}

	return nodeInfo, res, nil
}

// Info retrieves the server node information.
func (n *NodeService) Info() (*NodeInfo, *http.Response, error) {
	return n.InfoWithContext(context.TODO())
}

// StatusWithContext retrieves the server node status.
func (n *NodeService) StatusWithContext(ctx context.Context) (*NodeStatus, *http.Response, error) {
	resource, _ := url.Parse(nodeStatusResourcePath)

	req, err := n.client.NewRequest(ctx, http.MethodGet, resource, nil)

	if err != nil {
		return nil, nil, err
	}

	nodeStatus := new(NodeStatus)

	res, err := n.client.Do(req, nodeStatus)

	if err != nil {
		return nil, res, err
	}

	return nodeStatus, res, nil
}

// Status retrieves the server node status.
func (n *NodeService) Status() (*NodeStatus, *http.Response, error) {
	return n.StatusWithContext(context.TODO())
}
